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

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceActiveDirectoryDomainTrust() *schema.Resource {
	return &schema.Resource{
		Create: resourceActiveDirectoryDomainTrustCreate,
		Read:   resourceActiveDirectoryDomainTrustRead,
		Update: resourceActiveDirectoryDomainTrustUpdate,
		Delete: resourceActiveDirectoryDomainTrustDelete,

		Importer: &schema.ResourceImporter{
			State: resourceActiveDirectoryDomainTrustImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.`,
			},
			"target_dns_ip_addresses": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: `The target DNS server IP addresses which can resolve the remote domain involved in the trust.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"target_domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The fully qualified target domain name which will be in trust with the current domain.`,
			},
			"trust_direction": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"INBOUND", "OUTBOUND", "BIDIRECTIONAL"}),
				Description:  `The trust direction, which decides if the current domain is trusted, trusting, or both. Possible values: ["INBOUND", "OUTBOUND", "BIDIRECTIONAL"]`,
			},
			"trust_handshake_secret": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The trust secret used for the handshake with the target domain. This will not be stored.`,
				Sensitive:   true,
			},
			"trust_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"FOREST", "EXTERNAL"}),
				Description:  `The type of trust represented by the trust resource. Possible values: ["FOREST", "EXTERNAL"]`,
			},
			"selective_authentication": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceActiveDirectoryDomainTrustCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	targetDomainNameProp, err := expandNestedActiveDirectoryDomainTrustTargetDomainName(d.Get("target_domain_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_domain_name"); !isEmptyValue(reflect.ValueOf(targetDomainNameProp)) && (ok || !reflect.DeepEqual(v, targetDomainNameProp)) {
		obj["targetDomainName"] = targetDomainNameProp
	}
	trustTypeProp, err := expandNestedActiveDirectoryDomainTrustTrustType(d.Get("trust_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_type"); !isEmptyValue(reflect.ValueOf(trustTypeProp)) && (ok || !reflect.DeepEqual(v, trustTypeProp)) {
		obj["trustType"] = trustTypeProp
	}
	trustDirectionProp, err := expandNestedActiveDirectoryDomainTrustTrustDirection(d.Get("trust_direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_direction"); !isEmptyValue(reflect.ValueOf(trustDirectionProp)) && (ok || !reflect.DeepEqual(v, trustDirectionProp)) {
		obj["trustDirection"] = trustDirectionProp
	}
	selectiveAuthenticationProp, err := expandNestedActiveDirectoryDomainTrustSelectiveAuthentication(d.Get("selective_authentication"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("selective_authentication"); !isEmptyValue(reflect.ValueOf(selectiveAuthenticationProp)) && (ok || !reflect.DeepEqual(v, selectiveAuthenticationProp)) {
		obj["selectiveAuthentication"] = selectiveAuthenticationProp
	}
	targetDnsIpAddressesProp, err := expandNestedActiveDirectoryDomainTrustTargetDnsIpAddresses(d.Get("target_dns_ip_addresses"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_dns_ip_addresses"); !isEmptyValue(reflect.ValueOf(targetDnsIpAddressesProp)) && (ok || !reflect.DeepEqual(v, targetDnsIpAddressesProp)) {
		obj["targetDnsIpAddresses"] = targetDnsIpAddressesProp
	}
	trustHandshakeSecretProp, err := expandNestedActiveDirectoryDomainTrustTrustHandshakeSecret(d.Get("trust_handshake_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_handshake_secret"); !isEmptyValue(reflect.ValueOf(trustHandshakeSecretProp)) && (ok || !reflect.DeepEqual(v, trustHandshakeSecretProp)) {
		obj["trustHandshakeSecret"] = trustHandshakeSecretProp
	}

	obj, err = resourceActiveDirectoryDomainTrustEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/domains/{{domain}}:attachTrust")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DomainTrust: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainTrust: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating DomainTrust: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ActiveDirectoryOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating DomainTrust", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create DomainTrust: %s", err)
	}

	opRes, err = resourceActiveDirectoryDomainTrustDecoder(d, meta, opRes)
	if err != nil {
		return fmt.Errorf("Error decoding response from operation: %s", err)
	}
	if opRes == nil {
		return fmt.Errorf("Error decoding response from operation, could not find object")
	}

	if _, ok := opRes["trusts"]; ok {
		opRes, err = flattenNestedActiveDirectoryDomainTrust(d, meta, opRes)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if opRes == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("target_domain_name", flattenNestedActiveDirectoryDomainTrustTargetDomainName(opRes["targetDomainName"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DomainTrust %q: %#v", d.Id(), res)

	return resourceActiveDirectoryDomainTrustRead(d, meta)
}

func resourceActiveDirectoryDomainTrustRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/domains/{{domain}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainTrust: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ActiveDirectoryDomainTrust %q", d.Id()))
	}

	res, err = flattenNestedActiveDirectoryDomainTrust(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ActiveDirectoryDomainTrust because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceActiveDirectoryDomainTrustDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ActiveDirectoryDomainTrust because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DomainTrust: %s", err)
	}

	if err := d.Set("target_domain_name", flattenNestedActiveDirectoryDomainTrustTargetDomainName(res["targetDomainName"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainTrust: %s", err)
	}
	if err := d.Set("trust_type", flattenNestedActiveDirectoryDomainTrustTrustType(res["trustType"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainTrust: %s", err)
	}
	if err := d.Set("trust_direction", flattenNestedActiveDirectoryDomainTrustTrustDirection(res["trustDirection"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainTrust: %s", err)
	}
	if err := d.Set("selective_authentication", flattenNestedActiveDirectoryDomainTrustSelectiveAuthentication(res["selectiveAuthentication"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainTrust: %s", err)
	}
	if err := d.Set("target_dns_ip_addresses", flattenNestedActiveDirectoryDomainTrustTargetDnsIpAddresses(res["targetDnsIpAddresses"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainTrust: %s", err)
	}

	return nil
}

func resourceActiveDirectoryDomainTrustUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainTrust: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	targetDomainNameProp, err := expandNestedActiveDirectoryDomainTrustTargetDomainName(d.Get("target_domain_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_domain_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetDomainNameProp)) {
		obj["targetDomainName"] = targetDomainNameProp
	}
	trustTypeProp, err := expandNestedActiveDirectoryDomainTrustTrustType(d.Get("trust_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustTypeProp)) {
		obj["trustType"] = trustTypeProp
	}
	trustDirectionProp, err := expandNestedActiveDirectoryDomainTrustTrustDirection(d.Get("trust_direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_direction"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustDirectionProp)) {
		obj["trustDirection"] = trustDirectionProp
	}
	selectiveAuthenticationProp, err := expandNestedActiveDirectoryDomainTrustSelectiveAuthentication(d.Get("selective_authentication"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("selective_authentication"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, selectiveAuthenticationProp)) {
		obj["selectiveAuthentication"] = selectiveAuthenticationProp
	}
	targetDnsIpAddressesProp, err := expandNestedActiveDirectoryDomainTrustTargetDnsIpAddresses(d.Get("target_dns_ip_addresses"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_dns_ip_addresses"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetDnsIpAddressesProp)) {
		obj["targetDnsIpAddresses"] = targetDnsIpAddressesProp
	}
	trustHandshakeSecretProp, err := expandNestedActiveDirectoryDomainTrustTrustHandshakeSecret(d.Get("trust_handshake_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_handshake_secret"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustHandshakeSecretProp)) {
		obj["trustHandshakeSecret"] = trustHandshakeSecretProp
	}

	obj, err = resourceActiveDirectoryDomainTrustUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/domains/{{domain}}:reconfigureTrust")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DomainTrust %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating DomainTrust %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating DomainTrust %q: %#v", d.Id(), res)
	}

	err = ActiveDirectoryOperationWaitTime(
		config, res, project, "Updating DomainTrust", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceActiveDirectoryDomainTrustRead(d, meta)
}

func resourceActiveDirectoryDomainTrustDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/domains/{{domain}}:detachTrust")
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	targetDomainNameProp, err := expandNestedActiveDirectoryDomainTrustTargetDomainName(d.Get("target_domain_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_domain_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetDomainNameProp)) {
		obj["targetDomainName"] = targetDomainNameProp
	}
	trustTypeProp, err := expandNestedActiveDirectoryDomainTrustTrustType(d.Get("trust_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustTypeProp)) {
		obj["trustType"] = trustTypeProp
	}
	trustDirectionProp, err := expandNestedActiveDirectoryDomainTrustTrustDirection(d.Get("trust_direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_direction"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustDirectionProp)) {
		obj["trustDirection"] = trustDirectionProp
	}
	selectiveAuthenticationProp, err := expandNestedActiveDirectoryDomainTrustSelectiveAuthentication(d.Get("selective_authentication"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("selective_authentication"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, selectiveAuthenticationProp)) {
		obj["selectiveAuthentication"] = selectiveAuthenticationProp
	}
	targetDnsIpAddressesProp, err := expandNestedActiveDirectoryDomainTrustTargetDnsIpAddresses(d.Get("target_dns_ip_addresses"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_dns_ip_addresses"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetDnsIpAddressesProp)) {
		obj["targetDnsIpAddresses"] = targetDnsIpAddressesProp
	}
	trustHandshakeSecretProp, err := expandNestedActiveDirectoryDomainTrustTrustHandshakeSecret(d.Get("trust_handshake_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_handshake_secret"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustHandshakeSecretProp)) {
		obj["trustHandshakeSecret"] = trustHandshakeSecretProp
	}

	obj, err = resourceActiveDirectoryDomainTrustEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting DomainTrust %q", d.Id())

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", project, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "DomainTrust")
	}

	err = ActiveDirectoryOperationWaitTime(
		config, res, project, "Deleting DomainTrust", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting DomainTrust %q: %#v", d.Id(), res)
	return nil
}

func resourceActiveDirectoryDomainTrustImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/domains/(?P<domain>[^/]+)/(?P<target_domain_name>[^/]+)",
		"(?P<project>[^/]+)/(?P<domain>[^/]+)/(?P<target_domain_name>[^/]+)",
		"(?P<domain>[^/]+)/(?P<target_domain_name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedActiveDirectoryDomainTrustTargetDomainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedActiveDirectoryDomainTrustTrustType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedActiveDirectoryDomainTrustTrustDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedActiveDirectoryDomainTrustSelectiveAuthentication(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedActiveDirectoryDomainTrustTargetDnsIpAddresses(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func expandNestedActiveDirectoryDomainTrustTargetDomainName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedActiveDirectoryDomainTrustTrustType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedActiveDirectoryDomainTrustTrustDirection(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedActiveDirectoryDomainTrustSelectiveAuthentication(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedActiveDirectoryDomainTrustTargetDnsIpAddresses(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandNestedActiveDirectoryDomainTrustTrustHandshakeSecret(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceActiveDirectoryDomainTrustEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	wrappedReq := map[string]interface{}{
		"trust": obj,
	}
	return wrappedReq, nil
}

func resourceActiveDirectoryDomainTrustUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	wrappedReq := map[string]interface{}{
		"targetDomainName":     obj["targetDomainName"],
		"targetDnsIpAddresses": obj["targetDnsIpAddresses"],
	}
	return wrappedReq, nil
}

func flattenNestedActiveDirectoryDomainTrust(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["trusts"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value trusts. Actual value: %v", v)
	}

	_, item, err := resourceActiveDirectoryDomainTrustFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceActiveDirectoryDomainTrustFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedTargetDomainName, err := expandNestedActiveDirectoryDomainTrustTargetDomainName(d.Get("target_domain_name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedTargetDomainName := flattenNestedActiveDirectoryDomainTrustTargetDomainName(expectedTargetDomainName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceActiveDirectoryDomainTrustDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemTargetDomainName := flattenNestedActiveDirectoryDomainTrustTargetDomainName(item["targetDomainName"], d, meta.(*transport_tpg.Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemTargetDomainName)) && isEmptyValue(reflect.ValueOf(expectedFlattenedTargetDomainName))) && !reflect.DeepEqual(itemTargetDomainName, expectedFlattenedTargetDomainName) {
			log.Printf("[DEBUG] Skipping item with targetDomainName= %#v, looking for %#v)", itemTargetDomainName, expectedFlattenedTargetDomainName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceActiveDirectoryDomainTrustDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	v, ok := res["domainTrust"]
	if !ok || v == nil {
		return res, nil
	}

	return v.(map[string]interface{}), nil
}
