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
)

func resourceBillingBudget() *schema.Resource {
	return &schema.Resource{
		Create: resourceBillingBudgetCreate,
		Read:   resourceBillingBudgetRead,
		Update: resourceBillingBudgetUpdate,
		Delete: resourceBillingBudgetDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
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
			"billing_account": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the billing account to set a budget on.`,
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
						"disable_default_iam_recipients": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: `Boolean. When set to true, disables default notifications sent
when a threshold is exceeded. Default recipients are
those with Billing Account Administrators and Billing
Account Users IAM roles for the target account.`,
							Default: false,
						},
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
				Computed: true,
				Optional: true,
				Description: `Filters that define which resources are used to compute the actual
spend against the budget.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credit_types": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `A set of subaccounts of the form billingAccounts/{account_id},
specifying that usage from only this set of subaccounts should
be included in the budget. If a subaccount is set to the name of
the parent account, usage from the parent account will be included.
If the field is omitted, the report will include usage from the parent
account and all subaccounts, if they exist.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"budget_filter.0.projects", "budget_filter.0.credit_types_treatment", "budget_filter.0.services", "budget_filter.0.subaccounts", "budget_filter.0.labels"},
						},
						"credit_types_treatment": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", "INCLUDE_SPECIFIED_CREDITS", ""}, false),
							Description: `Specifies how credits should be treated when determining spend
for threshold calculations. Default value: "INCLUDE_ALL_CREDITS" Possible values: ["INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", "INCLUDE_SPECIFIED_CREDITS"]`,
							Default:      "INCLUDE_ALL_CREDITS",
							AtLeastOneOf: []string{"budget_filter.0.projects", "budget_filter.0.credit_types_treatment", "budget_filter.0.services", "budget_filter.0.subaccounts", "budget_filter.0.labels"},
						},
						"labels": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Description: `A single label and value pair specifying that usage from only
this set of labeled resources should be included in the budget.`,
							Elem:         &schema.Schema{Type: schema.TypeString},
							AtLeastOneOf: []string{"budget_filter.0.projects", "budget_filter.0.credit_types_treatment", "budget_filter.0.services", "budget_filter.0.subaccounts", "budget_filter.0.labels"},
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
							AtLeastOneOf: []string{"budget_filter.0.projects", "budget_filter.0.credit_types_treatment", "budget_filter.0.services", "budget_filter.0.subaccounts", "budget_filter.0.labels"},
						},
						"services": {
							Type:     schema.TypeList,
							Computed: true,
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
							AtLeastOneOf: []string{"budget_filter.0.projects", "budget_filter.0.credit_types_treatment", "budget_filter.0.services", "budget_filter.0.subaccounts", "budget_filter.0.labels"},
						},
						"subaccounts": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `A set of subaccounts of the form billingAccounts/{account_id},
specifying that usage from only this set of subaccounts should
be included in the budget. If a subaccount is set to the name of
the parent account, usage from the parent account will be included.
If the field is omitted, the report will include usage from the parent
account and all subaccounts, if they exist.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"budget_filter.0.projects", "budget_filter.0.credit_types_treatment", "budget_filter.0.services", "budget_filter.0.subaccounts", "budget_filter.0.labels"},
						},
					},
				},
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
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandBillingBudgetDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	budgetFilterProp, err := expandBillingBudgetBudgetFilter(d.Get("budget_filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("budget_filter"); !isEmptyValue(reflect.ValueOf(budgetFilterProp)) && (ok || !reflect.DeepEqual(v, budgetFilterProp)) {
		obj["budgetFilter"] = budgetFilterProp
	}
	amountProp, err := expandBillingBudgetAmount(d.Get("amount"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("amount"); !isEmptyValue(reflect.ValueOf(amountProp)) && (ok || !reflect.DeepEqual(v, amountProp)) {
		obj["amount"] = amountProp
	}
	thresholdRulesProp, err := expandBillingBudgetThresholdRules(d.Get("threshold_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threshold_rules"); !isEmptyValue(reflect.ValueOf(thresholdRulesProp)) && (ok || !reflect.DeepEqual(v, thresholdRulesProp)) {
		obj["thresholdRules"] = thresholdRulesProp
	}
	notificationsRuleProp, err := expandBillingBudgetAllUpdatesRule(d.Get("all_updates_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("all_updates_rule"); !isEmptyValue(reflect.ValueOf(notificationsRuleProp)) && (ok || !reflect.DeepEqual(v, notificationsRuleProp)) {
		obj["notificationsRule"] = notificationsRuleProp
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

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
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
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	return resourceBillingBudgetRead(d, meta)
}

func resourceBillingBudgetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BillingBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BillingBudget %q", d.Id()))
	}

	if err := d.Set("name", flattenBillingBudgetName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}
	if err := d.Set("display_name", flattenBillingBudgetDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}
	if err := d.Set("budget_filter", flattenBillingBudgetBudgetFilter(res["budgetFilter"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}
	if err := d.Set("amount", flattenBillingBudgetAmount(res["amount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}
	if err := d.Set("threshold_rules", flattenBillingBudgetThresholdRules(res["thresholdRules"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}
	if err := d.Set("all_updates_rule", flattenBillingBudgetAllUpdatesRule(res["notificationsRule"], d, config)); err != nil {
		return fmt.Errorf("Error reading Budget: %s", err)
	}

	return nil
}

func resourceBillingBudgetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandBillingBudgetDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	budgetFilterProp, err := expandBillingBudgetBudgetFilter(d.Get("budget_filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("budget_filter"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, budgetFilterProp)) {
		obj["budgetFilter"] = budgetFilterProp
	}
	amountProp, err := expandBillingBudgetAmount(d.Get("amount"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("amount"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, amountProp)) {
		obj["amount"] = amountProp
	}
	thresholdRulesProp, err := expandBillingBudgetThresholdRules(d.Get("threshold_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threshold_rules"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, thresholdRulesProp)) {
		obj["thresholdRules"] = thresholdRulesProp
	}
	notificationsRuleProp, err := expandBillingBudgetAllUpdatesRule(d.Get("all_updates_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("all_updates_rule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationsRuleProp)) {
		obj["notificationsRule"] = notificationsRuleProp
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

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Budget %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Budget %q: %#v", d.Id(), res)
	}

	return resourceBillingBudgetRead(d, meta)
}

func resourceBillingBudgetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

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

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Budget")
	}

	log.Printf("[DEBUG] Finished deleting Budget %q: %#v", d.Id(), res)
	return nil
}

func flattenBillingBudgetName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetFilter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["projects"] =
		flattenBillingBudgetBudgetFilterProjects(original["projects"], d, config)
	transformed["credit_types_treatment"] =
		flattenBillingBudgetBudgetFilterCreditTypesTreatment(original["creditTypesTreatment"], d, config)
	transformed["services"] =
		flattenBillingBudgetBudgetFilterServices(original["services"], d, config)
	transformed["credit_types"] =
		flattenBillingBudgetBudgetFilterCreditTypes(original["creditTypes"], d, config)
	transformed["subaccounts"] =
		flattenBillingBudgetBudgetFilterSubaccounts(original["subaccounts"], d, config)
	transformed["labels"] =
		flattenBillingBudgetBudgetFilterLabels(original["labels"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetBudgetFilterProjects(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetFilterCreditTypesTreatment(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetFilterServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetFilterCreditTypes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetFilterSubaccounts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetBudgetFilterLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetAmount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["specified_amount"] =
		flattenBillingBudgetAmountSpecifiedAmount(original["specifiedAmount"], d, config)
	transformed["last_period_amount"] =
		flattenBillingBudgetAmountLastPeriodAmount(original["lastPeriodAmount"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetAmountSpecifiedAmount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["currency_code"] =
		flattenBillingBudgetAmountSpecifiedAmountCurrencyCode(original["currencyCode"], d, config)
	transformed["units"] =
		flattenBillingBudgetAmountSpecifiedAmountUnits(original["units"], d, config)
	transformed["nanos"] =
		flattenBillingBudgetAmountSpecifiedAmountNanos(original["nanos"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetAmountSpecifiedAmountCurrencyCode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetAmountSpecifiedAmountUnits(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetAmountSpecifiedAmountNanos(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenBillingBudgetAmountLastPeriodAmount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v != nil
}

func flattenBillingBudgetThresholdRules(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"threshold_percent": flattenBillingBudgetThresholdRulesThresholdPercent(original["thresholdPercent"], d, config),
			"spend_basis":       flattenBillingBudgetThresholdRulesSpendBasis(original["spendBasis"], d, config),
		})
	}
	return transformed
}
func flattenBillingBudgetThresholdRulesThresholdPercent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetThresholdRulesSpendBasis(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetAllUpdatesRule(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_topic"] =
		flattenBillingBudgetAllUpdatesRulePubsubTopic(original["pubsubTopic"], d, config)
	transformed["schema_version"] =
		flattenBillingBudgetAllUpdatesRuleSchemaVersion(original["schemaVersion"], d, config)
	transformed["monitoring_notification_channels"] =
		flattenBillingBudgetAllUpdatesRuleMonitoringNotificationChannels(original["monitoringNotificationChannels"], d, config)
	transformed["disable_default_iam_recipients"] =
		flattenBillingBudgetAllUpdatesRuleDisableDefaultIamRecipients(original["disableDefaultIamRecipients"], d, config)
	return []interface{}{transformed}
}
func flattenBillingBudgetAllUpdatesRulePubsubTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetAllUpdatesRuleSchemaVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil || isEmptyValue(reflect.ValueOf(v)) {
		return "1.0"
	}

	return v
}

func flattenBillingBudgetAllUpdatesRuleMonitoringNotificationChannels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBillingBudgetAllUpdatesRuleDisableDefaultIamRecipients(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBillingBudgetDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjects, err := expandBillingBudgetBudgetFilterProjects(original["projects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjects); val.IsValid() && !isEmptyValue(val) {
		transformed["projects"] = transformedProjects
	}

	transformedCreditTypesTreatment, err := expandBillingBudgetBudgetFilterCreditTypesTreatment(original["credit_types_treatment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreditTypesTreatment); val.IsValid() && !isEmptyValue(val) {
		transformed["creditTypesTreatment"] = transformedCreditTypesTreatment
	}

	transformedServices, err := expandBillingBudgetBudgetFilterServices(original["services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["services"] = transformedServices
	}

	transformedCreditTypes, err := expandBillingBudgetBudgetFilterCreditTypes(original["credit_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreditTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["creditTypes"] = transformedCreditTypes
	}

	transformedSubaccounts, err := expandBillingBudgetBudgetFilterSubaccounts(original["subaccounts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubaccounts); val.IsValid() && !isEmptyValue(val) {
		transformed["subaccounts"] = transformedSubaccounts
	}

	transformedLabels, err := expandBillingBudgetBudgetFilterLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandBillingBudgetBudgetFilterProjects(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetFilterCreditTypesTreatment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetFilterServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetFilterCreditTypes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetFilterSubaccounts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetBudgetFilterLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBillingBudgetAmount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSpecifiedAmount, err := expandBillingBudgetAmountSpecifiedAmount(original["specified_amount"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpecifiedAmount); val.IsValid() && !isEmptyValue(val) {
		transformed["specifiedAmount"] = transformedSpecifiedAmount
	}

	transformedLastPeriodAmount, err := expandBillingBudgetAmountLastPeriodAmount(original["last_period_amount"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLastPeriodAmount); val.IsValid() && !isEmptyValue(val) {
		transformed["lastPeriodAmount"] = transformedLastPeriodAmount
	}

	return transformed, nil
}

func expandBillingBudgetAmountSpecifiedAmount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCurrencyCode, err := expandBillingBudgetAmountSpecifiedAmountCurrencyCode(original["currency_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrencyCode); val.IsValid() && !isEmptyValue(val) {
		transformed["currencyCode"] = transformedCurrencyCode
	}

	transformedUnits, err := expandBillingBudgetAmountSpecifiedAmountUnits(original["units"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnits); val.IsValid() && !isEmptyValue(val) {
		transformed["units"] = transformedUnits
	}

	transformedNanos, err := expandBillingBudgetAmountSpecifiedAmountNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !isEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandBillingBudgetAmountSpecifiedAmountCurrencyCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAmountSpecifiedAmountUnits(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAmountSpecifiedAmountNanos(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAmountLastPeriodAmount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return struct{}{}, nil
}

func expandBillingBudgetThresholdRules(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedThresholdPercent, err := expandBillingBudgetThresholdRulesThresholdPercent(original["threshold_percent"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["thresholdPercent"] = transformedThresholdPercent
		}

		transformedSpendBasis, err := expandBillingBudgetThresholdRulesSpendBasis(original["spend_basis"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpendBasis); val.IsValid() && !isEmptyValue(val) {
			transformed["spendBasis"] = transformedSpendBasis
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBillingBudgetThresholdRulesThresholdPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetThresholdRulesSpendBasis(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAllUpdatesRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandBillingBudgetAllUpdatesRulePubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	transformedSchemaVersion, err := expandBillingBudgetAllUpdatesRuleSchemaVersion(original["schema_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchemaVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["schemaVersion"] = transformedSchemaVersion
	}

	transformedMonitoringNotificationChannels, err := expandBillingBudgetAllUpdatesRuleMonitoringNotificationChannels(original["monitoring_notification_channels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonitoringNotificationChannels); val.IsValid() && !isEmptyValue(val) {
		transformed["monitoringNotificationChannels"] = transformedMonitoringNotificationChannels
	}

	transformedDisableDefaultIamRecipients, err := expandBillingBudgetAllUpdatesRuleDisableDefaultIamRecipients(original["disable_default_iam_recipients"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisableDefaultIamRecipients); val.IsValid() && !isEmptyValue(val) {
		transformed["disableDefaultIamRecipients"] = transformedDisableDefaultIamRecipients
	}

	return transformed, nil
}

func expandBillingBudgetAllUpdatesRulePubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAllUpdatesRuleSchemaVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAllUpdatesRuleMonitoringNotificationChannels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBillingBudgetAllUpdatesRuleDisableDefaultIamRecipients(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
