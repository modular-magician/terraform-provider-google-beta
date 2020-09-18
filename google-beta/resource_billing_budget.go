// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"google.golang.org/api/googleapi"
)

func resourceBillingBudget() *schema.Resource {
	return &schema.Resource{
		Create: resourceBillingBudgetCreate,
		Read:   resourceBillingBudgetRead,
		Update: resourceBillingBudgetUpdate,
		Delete: resourceBillingBudgetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBillingBudgetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"billing_account": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the billing account to set a budget on.`,
			},
			"amount": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The budgeted amount for each usage period.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_period_amount": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: `Configures a budget amount that is automatically set to 100% of
last period's spend.
Boolean. Set value to true to use. Do not set to false, instead
use the 'specified_amount' block.`,
							ExactlyOneOf: []string{"amount.0.specified_amount", "amount.0.last_period_amount"},
						},
						"specified_amount": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A specified amount to use as the budget. currencyCode is
optional. If specified, it must match the currency of the
billing account. The currencyCode is provided on output.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currency_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The 3-letter currency code defined in ISO 4217.`,
									},
									"nanos": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `Number of nano (10^-9) units of the amount.
The value must be between -999,999,999 and +999,999,999
inclusive. If units is positive, nanos must be positive or
zero. If units is zero, nanos can be positive, zero, or
negative. If units is negative, nanos must be negative or
zero. For example $-1.75 is represented as units=-1 and
nanos=-750,000,000.`,
									},
									"units": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The whole units of the amount. For example if currencyCode
is "USD", then 1 unit is one US dollar.`,
									},
								},
							},
							ExactlyOneOf: []string{"amount.0.specified_amount", "amount.0.last_period_amount"},
						},
					},
				},
			},
			"threshold_rules": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Rules that trigger alerts (notifications of thresholds being
crossed) when spend exceeds the specified percentages of the
budget.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_percent": {
							Type:     schema.TypeFloat,
							Required: true,
							Description: `Send an alert when this threshold is exceeded. This is a
1.0-based percentage, so 0.5 = 50%. Must be >= 0.`,
						},
						"spend_basis": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"CURRENT_SPEND", "FORECASTED_SPEND", ""}, false),
							Description: `The type of basis used to determine if spend has passed
the threshold. Default value: "CURRENT_SPEND" Possible values: ["CURRENT_SPEND", "FORECASTED_SPEND"]`,
							Default: "CURRENT_SPEND",
						},
					},
				},
			},
			"all_updates_rule": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Defines notifications that are sent on every update to the
billing account's spend, regardless of the thresholds defined
using threshold rules.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitoring_notification_channels": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `The full resource name of a monitoring notification
channel in the form
projects/{project_id}/notificationChannels/{channel_id}.
A maximum of 5 channels are allowed.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"all_updates_rule.0.pubsub_topic", "all_updates_rule.0.monitoring_notification_channels"},
						},
						"pubsub_topic": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The name of the Cloud Pub/Sub topic where budget related
messages will be published, in the form
projects/{project_id}/topics/{topic_id}. Updates are sent
at regular intervals to the topic.`,
							AtLeastOneOf: []string{"all_updates_rule.0.pubsub_topic", "all_updates_rule.0.monitoring_notification_channels"},
						},
						"schema_version": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The schema version of the notification. Only "1.0" is
accepted. It represents the JSON schema as defined in
https://cloud.google.com/billing/docs/how-to/budgets#notification_format.`,
							Default: "1.0",
						},
					},
				},
			},
			"budget_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Filters that define which resources are used to compute the actual
spend against the budget.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credit_types_treatment": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", ""}, false),
							Description: `Specifies how credits should be treated when determining spend
for threshold calculations. Default value: "INCLUDE_ALL_CREDITS" Possible values: ["INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS"]`,
							Default: "INCLUDE_ALL_CREDITS",
						},
						"projects": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A set of projects of the form projects/{project_id},
specifying that usage from only this set of projects should be
included in the budget. If omitted, the report will include
all usage for the billing account, regardless of which project
the usage occurred on. Only zero or one project can be
specified currently.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A set of services of the form services/{service_id},
specifying that usage from only this set of services should be
included in the budget. If omitted, the report will include
usage for all the services. The service names are available
through the Catalog API:
https://cloud.google.com/billing/v1/how-tos/catalog-api.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				AtLeastOneOf: []string{},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User data for display name in UI. Must be <= 60 chars.`,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Resource name of the budget. The resource name
implies the scope of a budget. Values are of the form
billingAccounts/{billingAccountId}/budgets/{budgetId}.`,
			},
		},
	}
}

func resourceBillingBudgetCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	obj := make(map[string]interface{})
	budgetProp, err := expandBillingBudgetBudget(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(budgetProp)) {
		obj["budget"] = budgetProp
	}

	url, err := replaceVars(d, config, "{{BillingBasePath}}billingAccounts/{{billing_account}}/budgets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Budget: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Budget: %s", err)
	}
	if err := d.Set("name", flattenBillingBudgetName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Budget %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceBillingBudgetRead(d, meta)
}

func resourceBillingBudgetRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	url, err := replaceVars(d, config, "{{BillingBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BillingBudget %q", d.Id()))
	}

	if err := d.Set("name", flattenBillingBudgetName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenBillingBudgetBudget(res["budget"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading Budget: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				d.Set(k, v)
			}
		}
	}

	return nil
}

func resourceBillingBudgetUpdate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	obj := make(map[string]interface{})
	budgetProp, err := expandBillingBudgetBudget(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(budgetProp)) {
		obj["budget"] = budgetProp
	}

	url, err := replaceVars(d, config, "{{BillingBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Budget %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Budget %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Budget %q: %#v", d.Id(), res)
	}

	return resourceBillingBudgetRead(d, meta)
}

func resourceBillingBudgetDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	url, err := replaceVars(d, config, "{{BillingBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Budget %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Budget")
	}

	log.Printf("[DEBUG] Finished deleting Budget %q: %#v", d.Id(), res)
	return nil
}

func resourceBillingBudgetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenBillingBudgetName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["display_name"] =
		flattenBillingBudgetBudgetDisplayName(original["displayName"], d, config)
	transformed["budget_filter"] =
		flattenBillingBudgetBudgetBudgetFilter(original["budgetFilter"], d, config)
	transformed["amount"] =
		flattenBillingBudgetBudgetAmount(original["amount"], d, config)
	transformed["threshold_rules"] =
		flattenBillingBudgetBudgetThresholdRules(original["thresholdRules"], d, config)
	transformed["all_updates_rule"] =
		flattenBillingBudgetBudgetAllUpdatesRule(original["allUpdatesRule"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetBudgetDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetBudgetFilter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["projects"] =
		flattenBillingBudgetBudgetBudgetFilterProjects(original["projects"], d, config)
	transformed["credit_types_treatment"] =
		flattenBillingBudgetBudgetBudgetFilterCreditTypesTreatment(original["creditTypesTreatment"], d, config)
	transformed["services"] =
		flattenBillingBudgetBudgetBudgetFilterServices(original["services"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetBudgetBudgetFilterProjects(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetBudgetFilterCreditTypesTreatment(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetBudgetFilterServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetAmount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["specified_amount"] =
		flattenBillingBudgetBudgetAmountSpecifiedAmount(original["specifiedAmount"], d, config)
	transformed["last_period_amount"] =
		flattenBillingBudgetBudgetAmountLastPeriodAmount(original["lastPeriodAmount"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetBudgetAmountSpecifiedAmount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["currency_code"] =
		flattenBillingBudgetBudgetAmountSpecifiedAmountCurrencyCode(original["currencyCode"], d, config)
	transformed["units"] =
		flattenBillingBudgetBudgetAmountSpecifiedAmountUnits(original["units"], d, config)
	transformed["nanos"] =
		flattenBillingBudgetBudgetAmountSpecifiedAmountNanos(original["nanos"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetBudgetAmountSpecifiedAmountCurrencyCode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetAmountSpecifiedAmountUnits(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetAmountSpecifiedAmountNanos(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBillingBudgetBudgetAmountLastPeriodAmount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v != nil
}

func flattenBillingBudgetBudgetThresholdRules(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"threshold_percent": flattenBillingBudgetBudgetThresholdRulesThresholdPercent(original["thresholdPercent"], d, config),
			"spend_basis":       flattenBillingBudgetBudgetThresholdRulesSpendBasis(original["spendBasis"], d, config),
		})
	}
	return transformed
}
func flattenBillingBudgetBudgetThresholdRulesThresholdPercent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetThresholdRulesSpendBasis(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetAllUpdatesRule(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_topic"] =
		flattenBillingBudgetBudgetAllUpdatesRulePubsubTopic(original["pubsubTopic"], d, config)
	transformed["schema_version"] =
		flattenBillingBudgetBudgetAllUpdatesRuleSchemaVersion(original["schemaVersion"], d, config)
	transformed["monitoring_notification_channels"] =
		flattenBillingBudgetBudgetAllUpdatesRuleMonitoringNotificationChannels(original["monitoringNotificationChannels"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetBudgetAllUpdatesRulePubsubTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetAllUpdatesRuleSchemaVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetAllUpdatesRuleMonitoringNotificationChannels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBillingBudgetBudget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedDisplayName, err := expandBillingBudgetBudgetDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
		transformed["displayName"] = transformedDisplayName
	}

	transformedBudgetFilter, err := expandBillingBudgetBudgetBudgetFilter(d.Get("budget_filter"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBudgetFilter); val.IsValid() && !isEmptyValue(val) {
		transformed["budgetFilter"] = transformedBudgetFilter
	}

	transformedAmount, err := expandBillingBudgetBudgetAmount(d.Get("amount"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAmount); val.IsValid() && !isEmptyValue(val) {
		transformed["amount"] = transformedAmount
	}

	transformedThresholdRules, err := expandBillingBudgetBudgetThresholdRules(d.Get("threshold_rules"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedThresholdRules); val.IsValid() && !isEmptyValue(val) {
		transformed["thresholdRules"] = transformedThresholdRules
	}

	transformedAllUpdatesRule, err := expandBillingBudgetBudgetAllUpdatesRule(d.Get("all_updates_rule"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllUpdatesRule); val.IsValid() && !isEmptyValue(val) {
		transformed["allUpdatesRule"] = transformedAllUpdatesRule
	}

	return transformed, nil
}

func expandBillingBudgetBudgetDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetBudgetFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjects, err := expandBillingBudgetBudgetBudgetFilterProjects(original["projects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjects); val.IsValid() && !isEmptyValue(val) {
		transformed["projects"] = transformedProjects
	}

	transformedCreditTypesTreatment, err := expandBillingBudgetBudgetBudgetFilterCreditTypesTreatment(original["credit_types_treatment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreditTypesTreatment); val.IsValid() && !isEmptyValue(val) {
		transformed["creditTypesTreatment"] = transformedCreditTypesTreatment
	}

	transformedServices, err := expandBillingBudgetBudgetBudgetFilterServices(original["services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["services"] = transformedServices
	}

	return transformed, nil
}

func expandBillingBudgetBudgetBudgetFilterProjects(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetBudgetFilterCreditTypesTreatment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetBudgetFilterServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAmount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSpecifiedAmount, err := expandBillingBudgetBudgetAmountSpecifiedAmount(original["specified_amount"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpecifiedAmount); val.IsValid() && !isEmptyValue(val) {
		transformed["specifiedAmount"] = transformedSpecifiedAmount
	}

	transformedLastPeriodAmount, err := expandBillingBudgetBudgetAmountLastPeriodAmount(original["last_period_amount"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLastPeriodAmount); val.IsValid() && !isEmptyValue(val) {
		transformed["lastPeriodAmount"] = transformedLastPeriodAmount
	}

	return transformed, nil
}

func expandBillingBudgetBudgetAmountSpecifiedAmount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCurrencyCode, err := expandBillingBudgetBudgetAmountSpecifiedAmountCurrencyCode(original["currency_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrencyCode); val.IsValid() && !isEmptyValue(val) {
		transformed["currencyCode"] = transformedCurrencyCode
	}

	transformedUnits, err := expandBillingBudgetBudgetAmountSpecifiedAmountUnits(original["units"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnits); val.IsValid() && !isEmptyValue(val) {
		transformed["units"] = transformedUnits
	}

	transformedNanos, err := expandBillingBudgetBudgetAmountSpecifiedAmountNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !isEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandBillingBudgetBudgetAmountSpecifiedAmountCurrencyCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAmountSpecifiedAmountUnits(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAmountSpecifiedAmountNanos(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAmountLastPeriodAmount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return struct{}{}, nil
}

func expandBillingBudgetBudgetThresholdRules(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedThresholdPercent, err := expandBillingBudgetBudgetThresholdRulesThresholdPercent(original["threshold_percent"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["thresholdPercent"] = transformedThresholdPercent
		}

		transformedSpendBasis, err := expandBillingBudgetBudgetThresholdRulesSpendBasis(original["spend_basis"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpendBasis); val.IsValid() && !isEmptyValue(val) {
			transformed["spendBasis"] = transformedSpendBasis
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBillingBudgetBudgetThresholdRulesThresholdPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetThresholdRulesSpendBasis(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAllUpdatesRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandBillingBudgetBudgetAllUpdatesRulePubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	transformedSchemaVersion, err := expandBillingBudgetBudgetAllUpdatesRuleSchemaVersion(original["schema_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchemaVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["schemaVersion"] = transformedSchemaVersion
	}

	transformedMonitoringNotificationChannels, err := expandBillingBudgetBudgetAllUpdatesRuleMonitoringNotificationChannels(original["monitoring_notification_channels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonitoringNotificationChannels); val.IsValid() && !isEmptyValue(val) {
		transformed["monitoringNotificationChannels"] = transformedMonitoringNotificationChannels
	}

	return transformed, nil
}

func expandBillingBudgetBudgetAllUpdatesRulePubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAllUpdatesRuleSchemaVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetAllUpdatesRuleMonitoringNotificationChannels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
