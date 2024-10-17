// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package serviceusagev2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

const (
	dependentServiceWorkers int = 4
)

type ApiError = string

const (
	ApiErrorSuConflictingConcurrentModification   = "SU_CONFLICTING_CONCURRENT_MODIFICATION"
	ApiErrorCommonSuServicesHaveUsage             = "COMMON_SU_SERVICES_HAVE_USAGE"
	ApiErrorCommonSuServicesHaveHierarchicalUsage = "COMMON_SU_SERVICES_HAVE_HIERARCHICAL_USAGE"
	ApiErrorSuGroupNotFound                       = "SU_GROUP_NOT_FOUND"
)

type ApiCallFn = func(url string) (map[string]interface{}, error)

type EnableRules struct {
	Services []string
}

type DependentService struct {
	ServiceName string
}
type Dependencies struct {
	Services      []DependentService
	NextPageToken string
}

var urlRegexp = regexp.MustCompile(`(.*)/consumerPolicies/.*$`)
var servicesRegexp = regexp.MustCompile(`([^(\s|,)]+\.googleapis.com)`)

func parseObject(obj any, resp interface{}) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, resp); err != nil {
		return err
	}
	return nil
}
func removeDuplicateElementsWithDuplicates[T any](strSlice []T) ([]T, []T) {
	allKeys := make(map[interface{}]bool)
	duplicateKeys := make(map[interface{}]bool)

	var uniqueList, duplicates []T
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			uniqueList = append(uniqueList, item)
		} else if _, value := duplicateKeys[item]; !value {
			duplicateKeys[item] = true
			duplicates = append(duplicates, item)
		}
	}
	return uniqueList, duplicates
}

func removeDuplicateElements[T any](strSlice []T) []T {
	response, _ := removeDuplicateElementsWithDuplicates(strSlice)
	return response
}

func getServices(rules interface{}, output []string) ([]string, error) {
	var enableRules []EnableRules
	if err := parseObject(rules, &enableRules); err != nil {
		return nil, err
	}
	for _, rule := range enableRules {
		output = append(output, rule.Services...)
	}
	return output, nil
}

type DependentServicesResult struct {
	service     string
	descendants []string
	error       error
}

func getDescendantsForService(url, service string, apiDependenciesFunc ApiCallFn) ([]string, error) {
	var dependentServices []string
	var nextPageToken string

	dependenciesUrl := urlRegexp.ReplaceAllString(url, fmt.Sprintf(`${1}/%s/groups/dependencies/descendantServices`, service))

	for i := 0; i < 1000; i++ {
		if nextPageToken != "" {
			dependenciesUrl = fmt.Sprintf("%s?pageToken=%s", dependenciesUrl, nextPageToken)
		}
		resp, err := apiDependenciesFunc(dependenciesUrl)

		if err != nil {
			if isError(err, ApiErrorSuGroupNotFound) {
				return nil, nil
			}
			return nil, err
		}
		dependencies := Dependencies{}
		if err := parseObject(resp, &dependencies); err != nil {
			return nil, err
		}

		for _, responseService := range dependencies.Services {
			dependentServices = append(dependentServices, responseService.ServiceName)
		}

		if nextPageToken = dependencies.NextPageToken; nextPageToken == "" {
			break
		}
	}
	log.Printf("[DEBUG] Dependent services for service '%v': %v", service, dependentServices)
	return dependentServices, nil
}
func dependentServicesWorker(ctx context.Context, cancel context.CancelFunc, url string, apiDependenciesFunc ApiCallFn, services <-chan string, result chan<- DependentServicesResult) {

	for {
		select {
		case service, ok := <-services:
			if !ok {
				return
			}

			dependentServices, err := getDescendantsForService(url, service, apiDependenciesFunc)
			result <- DependentServicesResult{service, dependentServices, err}
			if err != nil {
				cancel()
			}

		case <-ctx.Done():
			return
		}
	}
}

func getDependentServicesMap(url string, services []string, apiDependenciesFunc ApiCallFn) (map[string][]string, error) {
	jobs := make(chan string, len(services))
	results := make(chan DependentServicesResult, len(services))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for w := 0; w < dependentServiceWorkers; w++ {
		go dependentServicesWorker(ctx, cancel, url, apiDependenciesFunc, jobs, results)
	}
	for _, service := range services {
		jobs <- service
	}
	close(jobs)

	resultMap := make(map[string][]string)
	for range services {
		result := <-results
		if result.error != nil {
			return nil, result.error
		}
		resultMap[result.service] = result.descendants
	}

	return resultMap, nil
}

func reverseDependencyMap(dependencies map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for svc, deps := range dependencies {
		for _, rsvc := range deps {
			rdeps := []string{svc}
			if rdepsExisting, ok := result[rsvc]; ok {
				rdeps = append(rdeps, rdepsExisting...)
			}
			result[rsvc] = rdeps
		}
	}
	return result
}

func getTransitiveDependents(dependents map[string][]string, service string) []string {
	var result, services []string
	initServices := []string{service}
	circular := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		services = slices.Clone(initServices)
		initServices = nil
		for _, svc := range services {

			if deps, ok := dependents[svc]; ok {
				for _, dep := range deps {
					if _, ok := circular[dep]; ok {
						continue
					}
					circular[svc] = true
					result = append(result, dep)
					initServices = append(initServices, dep)
				}
			}
		}
		if len(initServices) == 0 {
			break
		}

	}
	return removeDuplicateElements(result)
}

func validateDependencies(oldServices, newServices []string, dependentServicesMap map[string][]string) error {

	missingDependenciesMap := make(map[string][]string)

	for _, service := range newServices {
		dependencies, ok := dependentServicesMap[service]
		if !ok {
			return fmt.Errorf("Something went wrong, couldn't find dependencies for service: %v\n", service)
		}
		notDefinedDependencies := slices.DeleteFunc(dependencies, func(svc string) bool { return slices.Contains(newServices, svc) })

		if len(notDefinedDependencies) > 0 {
			missingDependenciesMap[service] = notDefinedDependencies
		}
	}

	// check dependencies for removal

	servicesToRemove := slices.DeleteFunc(slices.Clone(oldServices), func(svc string) bool { return slices.Contains(newServices, svc) })
	existingDependentsMap := make(map[string][]string)

	if len(servicesToRemove) > 0 {
		reverseDependencies := reverseDependencyMap(dependentServicesMap)
		existingServices := slices.DeleteFunc(slices.Clone(newServices), func(svc string) bool {
			return !slices.Contains(oldServices, svc)
		})
		for _, svc := range servicesToRemove {
			// looking for existing dependents in terraform configuration - remove the ones which are not in the previous terraform config (ignore the new ones)
			existingDependents := slices.DeleteFunc(getTransitiveDependents(reverseDependencies, svc), func(svc string) bool { return !slices.Contains(existingServices, svc) })
			if len(existingDependents) > 0 {
				existingDependentsMap[svc] = existingDependents
				// remove existing missing dependencies for removed dependant
				for k, v := range missingDependenciesMap {
					if slices.Contains(v, svc) {
						delete(missingDependenciesMap, k)
					}
				}
			}
		}
	}

	if len(missingDependenciesMap) > 0 || len(existingDependentsMap) > 0 {
		var sb strings.Builder
		if len(missingDependenciesMap) > 0 {
			sb.WriteString("There are additional services for which all necessary dependencies haven't been added. Please add these missing dependencies:\n\n")
			for svc, deps := range missingDependenciesMap {
				depsString := "\"" + strings.Join(deps, "\", \"") + "\""
				sb.WriteString(fmt.Sprintf("Added service: [\"%v\"]\nMissing dependencies: [%v]\n", svc, depsString))
			}
			sb.WriteString("\n")
		}
		if len(existingDependentsMap) > 0 {
			sb.WriteString("There are existing services in configuration which depend on the services to be removed. Please remove existing dependent services:\n\n")
			for svc, deps := range existingDependentsMap {
				depsString := "\"" + strings.Join(deps, "\", \"") + "\""
				sb.WriteString(fmt.Sprintf("Removed service: [\"%v\"]\nExisting dependents: [%v]\n", svc, depsString))
			}
		}
		return fmt.Errorf("%v\nIf you don't want to validate dependencies, set validate_dependencies to false to override.\n", sb.String())
	}
	return nil
}

func isError(err error, apiError ApiError) bool {
	if errInfo, ok := err.(*googleapi.Error); ok {
		for _, detail := range errInfo.Details {
			if errInfo, ok := detail.(map[string]interface{}); ok {
				if errInfo["reason"] == apiError {
					return true
				}
			}
		}
	}
	return false
}

func isCommonOpError(err error, apiError ApiError) (bool, string) {
	if opError, ok := errors.Unwrap(err).(*tpgresource.CommonOpError); ok {
		for _, detail := range opError.Details {
			errInfo := make(map[string]interface{})
			if err := json.Unmarshal(detail, &errInfo); err != nil {
				return false, ""
			}
			if errInfo["reason"] == apiError {
				return true, opError.Message
			}
		}
	}
	return false, ""
}

func isTimeoutError(err error) bool {
	_, ok := errors.Unwrap(err).(*retry.TimeoutError)
	return ok
}

func getServicesFromString(msg string) []string {
	return servicesRegexp.FindAllString(msg, -1)
}

func getErrorMessageForServicesInUse(msg string) string {
	services := getServicesFromString(msg)
	return fmt.Sprintf("The service{s} %v has been used in the last 30 days or was enabled in the past 3 days. If you still wish to remove the service{s}, please set the check_usage_on_remove flag to false to proceed.", strings.Join(services, ", "))
}

func ResourceServiceUsageV2ConsumerPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceUsageV2ConsumerPolicyCreate,
		Read:   resourceServiceUsageV2ConsumerPolicyRead,
		Update: resourceServiceUsageV2ConsumerPolicyUpdate,
		Delete: resourceServiceUsageV2ConsumerPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceUsageV2ConsumerPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enable_rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `(Required) The consumer policy rule that defines enabled services. The structure is documented below.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"services": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `(Optional): List of service names to be enabled in the format of services/<service_name>`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateRegexp(`^default$`),
				Description:  `(Required) The name of the policy. Currently only the “default” policy name is supported.`,
			},
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `(Required) The name of the parent. It can be a project, folder or organization in the format of  projects/<project_id or project_number>, folders/<folder_number> or organizations/<org_number> .`,
			},
			"check_usage_on_remove": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `(Optional) Default value is false. If true, the usage of the service to be removed will be checked. If the service has been used within the past 30 days or was enabled in the last 3 days, an error will be thrown.`,
				Default:     false,
			},
			"validate_dependencies": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `(Optional) Default value is true. If true, this flag enforces dependency management within the consumer policy. When adding a new service, it verifies that all its dependencies are already present/added in the policy. Conversely, when removing a service, it ensures that no other services within the policy depend on the service to be removed. If the validation fails, a comprehensive message will be presented, outlining the missing dependencies and providing instructions on how to address the issue.`,
				Default:     true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceServiceUsageV2ConsumerPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	d.SetId("")
	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceUsageV2BasePath}}{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ConsumerPolicy: %#v", obj)

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Couldn't find project id: %s", err)
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	parent := d.Get("parent").(string)
	name := d.Get("name").(string)
	consumerPolicyName := fmt.Sprintf("%v/consumerPolicies/%v", parent, name)

	oldRules, newRules := d.GetChange("enable_rules")

	// check state with API
	oldServices, err := getServices(oldRules, []string{})
	if err != nil {
		return err
	}
	newServices, err := getServices(newRules, []string{})
	if err != nil {
		return err
	}

	newServices, duplicates := removeDuplicateElementsWithDuplicates(newServices)
	if len(duplicates) > 0 {
		return fmt.Errorf("The services (%v) are listed multiple times in the consumer policy %v. "+
			"Please remove duplicate entries and try again.\n", strings.Join(duplicates, ", "), consumerPolicyName)
	}
	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return err
	}

	// checking already enabled apis on the api side and filtering services for a removal
	var apiServices []string
	if rules, ok := resp["enableRules"]; ok {
		if apiServices, err = getServices(rules, apiServices); err != nil {
			return err
		}
	}

	oldAndNewServices := removeDuplicateElements(slices.Concat(oldServices, newServices))
	apiExtraServices := slices.DeleteFunc(apiServices, func(svc string) bool { return slices.Contains(oldAndNewServices, svc) })

	if len(apiExtraServices) > 0 {
		return fmt.Errorf("Found services (%v) in consumer policy %v/consumerPolicies/%v but don't exist in terraform config. "+
			"It may be due to policy change via other toolings like gcloud and causes policy out of sync. You can either add these services "+
			"to terraform config (recommended) to make it up to date or you can remove these services via other toolings.\n", strings.Join(apiExtraServices, ", "), parent, name)
	}

	// check dependencies if the "respect_dependencies" flag is enabled
	validateDependenciesEnabled := d.Get("validate_dependencies").(bool)

	if validateDependenciesEnabled {
		dependentServicesMap, err := getDependentServicesMap(url, oldAndNewServices, func(dependenciesUrl string) (map[string]interface{}, error) {
			return transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    dependenciesUrl,
				UserAgent: userAgent,
				Timeout:   d.Timeout(schema.TimeoutUpdate),
				Headers:   make(http.Header),
			})
		})
		if err != nil {
			return err
		}
		if err := validateDependencies(oldServices, newServices, dependentServicesMap); err != nil {
			return err
		}
	}

	if !d.Get("check_usage_on_remove").(bool) {
		url, err = transport_tpg.AddQueryParams(url, map[string]string{"force": "true"})
	}

	newObj := make(map[string]interface{})
	newObj["enableRules"] = []map[string][]string{
		{"services": removeDuplicateElements(newServices)},
	}
	obj = newObj
	obj["etag"] = resp["etag"]
	log.Printf("[DEBUG] Update policy request: %v", obj)

	resp, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating ConsumerPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	timeout := d.Timeout(schema.TimeoutCreate)
	var opRes map[string]interface{}
	err = ServiceUsageV2OperationWaitTimeWithResponse(
		config, resp, &opRes, "Updating ConsumerPolicy", config.UserAgent,
		timeout)
	if err != nil {
		// The resource wasn't actually created
		d.SetId("")
		usageError := ApiErrorCommonSuServicesHaveHierarchicalUsage
		if strings.Contains(parent, "projects/") {
			usageError = ApiErrorCommonSuServicesHaveUsage
		}
		if ok, _ := isCommonOpError(err, ApiErrorSuConflictingConcurrentModification); ok || isTimeoutError(err) {
			return errors.New("The consumer policy might have been changed from other toolings. Please try again.")
		} else if ok, msg := isCommonOpError(err, usageError); ok {
			return errors.New(getErrorMessageForServicesInUse(msg))
		}
		return fmt.Errorf("Error while waiting to update ConsumerPolicy: %s", err)
	}
	log.Printf("[DEBUG] Finished updating ConsumerPolicy %q: %#v", d.Id(), opRes)
	return resourceServiceUsageV2ConsumerPolicyRead(d, meta)
}

func resourceServiceUsageV2ConsumerPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceUsageV2BasePath}}{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ServiceUsageV2ConsumerPolicy %q", d.Id()))
	}

	res, err = resourceServiceUsageV2ConsumerPolicyDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ServiceUsageV2ConsumerPolicy because it no longer exists.")
		d.SetId("")
		return nil
	}

	return nil
}

func resourceServiceUsageV2ConsumerPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	d.SetId("")
	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceUsageV2BasePath}}{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ConsumerPolicy: %#v", obj)

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Couldn't find project id: %s", err)
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	parent := d.Get("parent").(string)
	name := d.Get("name").(string)
	consumerPolicyName := fmt.Sprintf("%v/consumerPolicies/%v", parent, name)

	oldRules, newRules := d.GetChange("enable_rules")

	// check state with API
	oldServices, err := getServices(oldRules, []string{})
	if err != nil {
		return err
	}
	newServices, err := getServices(newRules, []string{})
	if err != nil {
		return err
	}

	newServices, duplicates := removeDuplicateElementsWithDuplicates(newServices)
	if len(duplicates) > 0 {
		return fmt.Errorf("The services (%v) are listed multiple times in the consumer policy %v. "+
			"Please remove duplicate entries and try again.\n", strings.Join(duplicates, ", "), consumerPolicyName)
	}
	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return err
	}

	// checking already enabled apis on the api side and filtering services for a removal
	var apiServices []string
	if rules, ok := resp["enableRules"]; ok {
		if apiServices, err = getServices(rules, apiServices); err != nil {
			return err
		}
	}

	oldAndNewServices := removeDuplicateElements(slices.Concat(oldServices, newServices))
	apiExtraServices := slices.DeleteFunc(apiServices, func(svc string) bool { return slices.Contains(oldAndNewServices, svc) })

	if len(apiExtraServices) > 0 {
		return fmt.Errorf("Found services (%v) in consumer policy %v/consumerPolicies/%v but don't exist in terraform config. "+
			"It may be due to policy change via other toolings like gcloud and causes policy out of sync. You can either add these services "+
			"to terraform config (recommended) to make it up to date or you can remove these services via other toolings.\n", strings.Join(apiExtraServices, ", "), parent, name)
	}

	// check dependencies if the "respect_dependencies" flag is enabled
	validateDependenciesEnabled := d.Get("validate_dependencies").(bool)

	if validateDependenciesEnabled {
		dependentServicesMap, err := getDependentServicesMap(url, oldAndNewServices, func(dependenciesUrl string) (map[string]interface{}, error) {
			return transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    dependenciesUrl,
				UserAgent: userAgent,
				Timeout:   d.Timeout(schema.TimeoutUpdate),
				Headers:   make(http.Header),
			})
		})
		if err != nil {
			return err
		}
		if err := validateDependencies(oldServices, newServices, dependentServicesMap); err != nil {
			return err
		}
	}

	if !d.Get("check_usage_on_remove").(bool) {
		url, err = transport_tpg.AddQueryParams(url, map[string]string{"force": "true"})
	}

	newObj := make(map[string]interface{})
	newObj["enableRules"] = []map[string][]string{
		{"services": removeDuplicateElements(newServices)},
	}
	obj = newObj
	obj["etag"] = resp["etag"]
	log.Printf("[DEBUG] Update policy request: %v", obj)

	resp, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating ConsumerPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	timeout := d.Timeout(schema.TimeoutCreate)
	var opRes map[string]interface{}
	err = ServiceUsageV2OperationWaitTimeWithResponse(
		config, resp, &opRes, "Updating ConsumerPolicy", config.UserAgent,
		timeout)
	if err != nil {
		// The resource wasn't actually created
		d.SetId("")
		usageError := ApiErrorCommonSuServicesHaveHierarchicalUsage
		if strings.Contains(parent, "projects/") {
			usageError = ApiErrorCommonSuServicesHaveUsage
		}
		if ok, _ := isCommonOpError(err, ApiErrorSuConflictingConcurrentModification); ok || isTimeoutError(err) {
			return errors.New("The consumer policy might have been changed from other toolings. Please try again.")
		} else if ok, msg := isCommonOpError(err, usageError); ok {
			return errors.New(getErrorMessageForServicesInUse(msg))
		}
		return fmt.Errorf("Error while waiting to update ConsumerPolicy: %s", err)
	}
	log.Printf("[DEBUG] Finished updating ConsumerPolicy %q: %#v", d.Id(), opRes)
	return resourceServiceUsageV2ConsumerPolicyRead(d, meta)
}

func resourceServiceUsageV2ConsumerPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceUsageV2BasePath}}{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ConsumerPolicy %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ConsumerPolicy")
	}

	log.Printf("[DEBUG] Finished deleting ConsumerPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceServiceUsageV2ConsumerPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/consumerPolicies/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/consumerPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandServiceUsageV2ConsumerPolicyEnableRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServices, err := expandServiceUsageV2ConsumerPolicyEnableRulesServices(original["services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServices); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["services"] = transformedServices
	}

	return transformed, nil
}

func expandServiceUsageV2ConsumerPolicyEnableRulesServices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandServiceUsageV2ConsumerPolicyValidateDependencies(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandServiceUsageV2ConsumerPolicyCheckUsageOnRemove(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceServiceUsageV2ConsumerPolicyDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["enableRules"]; ok {
		original := v.([]interface{})
		res["enable_rules"] = original[0]
	}

	return res, nil
}
