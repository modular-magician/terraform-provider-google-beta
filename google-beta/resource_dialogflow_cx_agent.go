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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceDialogflowCXAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowCXAgentCreate,
		Read:   resourceDialogflowCXAgentRead,
		Update: resourceDialogflowCXAgentUpdate,
		Delete: resourceDialogflowCXAgentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowCXAgentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"default_language_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
for a list of the currently supported language codes. This field cannot be updated after creation.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The human-readable name of the agent, unique within the location.`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name of the location this agent is located in.

~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
 This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
 Another options is to use global location so you don't need to manually configure location settings.`,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
Europe/Paris.`,
			},
			"avatar_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.`,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 500),
				Description:  `The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.`,
			},
			"enable_spell_correction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Indicates if automatic spell correction is enabled in detect intent requests.`,
			},
			"enable_stackdriver_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines whether this agent should log conversation queries.`,
			},
			"security_settings": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.`,
			},
			"speech_to_text_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Settings related to speech recognition.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_speech_adaptation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether to use speech adaptation for speech recognition.`,
						},
					},
				},
			},
			"supported_language_codes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The list of all languages supported by this agent (except for the default_language_code).`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier of the agent.`,
			},
			"start_flow": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.`,
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

func resourceDialogflowCXAgentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXAgentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultLanguageCodeProp, err := expandDialogflowCXAgentDefaultLanguageCode(d.Get("default_language_code"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_language_code"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultLanguageCodeProp)) && (ok || !reflect.DeepEqual(v, defaultLanguageCodeProp)) {
		obj["defaultLanguageCode"] = defaultLanguageCodeProp
	}
	supportedLanguageCodesProp, err := expandDialogflowCXAgentSupportedLanguageCodes(d.Get("supported_language_codes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("supported_language_codes"); !tpgresource.IsEmptyValue(reflect.ValueOf(supportedLanguageCodesProp)) && (ok || !reflect.DeepEqual(v, supportedLanguageCodesProp)) {
		obj["supportedLanguageCodes"] = supportedLanguageCodesProp
	}
	timeZoneProp, err := expandDialogflowCXAgentTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandDialogflowCXAgentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	avatarUriProp, err := expandDialogflowCXAgentAvatarUri(d.Get("avatar_uri"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("avatar_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(avatarUriProp)) && (ok || !reflect.DeepEqual(v, avatarUriProp)) {
		obj["avatarUri"] = avatarUriProp
	}
	speechToTextSettingsProp, err := expandDialogflowCXAgentSpeechToTextSettings(d.Get("speech_to_text_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("speech_to_text_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(speechToTextSettingsProp)) && (ok || !reflect.DeepEqual(v, speechToTextSettingsProp)) {
		obj["speechToTextSettings"] = speechToTextSettingsProp
	}
	securitySettingsProp, err := expandDialogflowCXAgentSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(securitySettingsProp)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXAgentEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXAgentEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableSpellCorrectionProp)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/agents")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Agent: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Agent: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Agent: %s", err)
	}
	if err := d.Set("name", flattenDialogflowCXAgentName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Agent %q: %#v", d.Id(), res)

	return resourceDialogflowCXAgentRead(d, meta)
}

func resourceDialogflowCXAgentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Agent: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DialogflowCXAgent %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}

	if err := d.Set("name", flattenDialogflowCXAgentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowCXAgentDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("default_language_code", flattenDialogflowCXAgentDefaultLanguageCode(res["defaultLanguageCode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("supported_language_codes", flattenDialogflowCXAgentSupportedLanguageCodes(res["supportedLanguageCodes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("time_zone", flattenDialogflowCXAgentTimeZone(res["timeZone"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("description", flattenDialogflowCXAgentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("avatar_uri", flattenDialogflowCXAgentAvatarUri(res["avatarUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("speech_to_text_settings", flattenDialogflowCXAgentSpeechToTextSettings(res["speechToTextSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("start_flow", flattenDialogflowCXAgentStartFlow(res["startFlow"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("security_settings", flattenDialogflowCXAgentSecuritySettings(res["securitySettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("enable_stackdriver_logging", flattenDialogflowCXAgentEnableStackdriverLogging(res["enableStackdriverLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}
	if err := d.Set("enable_spell_correction", flattenDialogflowCXAgentEnableSpellCorrection(res["enableSpellCorrection"], d, config)); err != nil {
		return fmt.Errorf("Error reading Agent: %s", err)
	}

	return nil
}

func resourceDialogflowCXAgentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Agent: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXAgentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	supportedLanguageCodesProp, err := expandDialogflowCXAgentSupportedLanguageCodes(d.Get("supported_language_codes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("supported_language_codes"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, supportedLanguageCodesProp)) {
		obj["supportedLanguageCodes"] = supportedLanguageCodesProp
	}
	timeZoneProp, err := expandDialogflowCXAgentTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandDialogflowCXAgentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	avatarUriProp, err := expandDialogflowCXAgentAvatarUri(d.Get("avatar_uri"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("avatar_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, avatarUriProp)) {
		obj["avatarUri"] = avatarUriProp
	}
	speechToTextSettingsProp, err := expandDialogflowCXAgentSpeechToTextSettings(d.Get("speech_to_text_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("speech_to_text_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, speechToTextSettingsProp)) {
		obj["speechToTextSettings"] = speechToTextSettingsProp
	}
	securitySettingsProp, err := expandDialogflowCXAgentSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXAgentEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXAgentEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Agent %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("supported_language_codes") {
		updateMask = append(updateMask, "supportedLanguageCodes")
	}

	if d.HasChange("time_zone") {
		updateMask = append(updateMask, "timeZone")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("avatar_uri") {
		updateMask = append(updateMask, "avatarUri")
	}

	if d.HasChange("speech_to_text_settings") {
		updateMask = append(updateMask, "speechToTextSettings")
	}

	if d.HasChange("security_settings") {
		updateMask = append(updateMask, "securitySettings")
	}

	if d.HasChange("enable_stackdriver_logging") {
		updateMask = append(updateMask, "enableStackdriverLogging")
	}

	if d.HasChange("enable_spell_correction") {
		updateMask = append(updateMask, "enableSpellCorrection")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Agent %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Agent %q: %#v", d.Id(), res)
	}

	return resourceDialogflowCXAgentRead(d, meta)
}

func resourceDialogflowCXAgentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Agent: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Agent %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Agent")
	}

	log.Printf("[DEBUG] Finished deleting Agent %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowCXAgentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/agents/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowCXAgentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDialogflowCXAgentDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentDefaultLanguageCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentSupportedLanguageCodes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentTimeZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentAvatarUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentSpeechToTextSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_speech_adaptation"] =
		flattenDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(original["enableSpeechAdaptation"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentStartFlow(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentSecuritySettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentEnableStackdriverLogging(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXAgentEnableSpellCorrection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDialogflowCXAgentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDefaultLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSupportedLanguageCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAvatarUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSpeechToTextSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableSpeechAdaptation, err := expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(original["enable_speech_adaptation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSpeechAdaptation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableSpeechAdaptation"] = transformedEnableSpeechAdaptation
	}

	return transformed, nil
}

func expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSecuritySettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableSpellCorrection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
