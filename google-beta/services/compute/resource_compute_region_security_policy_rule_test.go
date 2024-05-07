// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleBasicUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePreUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePostUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePreUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "default" {
  region      = "us-west2"
  name        = "tf-test%{random_suffix}"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.default.name
  region          = "us-west2"
  description     = "basic rule pre update"
  action          = "allow"
  priority        = 100
  preview         = false
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePostUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "default" {
  region      = "us-west2"
  name        = "tf-test%{random_suffix}"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.default.name
  region          = "us-west2"
  description     = "basic rule post update"
  action          = "deny(403)"
  priority        = 100
  preview         = true
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["172.16.0.0/12"]
    }
  }
}
`, context)
}

func TestAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate(t *testing.T) {
	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchBasic(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate2(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchBasic(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-west2"
  name        = "tf-test-policyddos%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}
  
resource "google_compute_network_edge_security_service" "edge_sec_service" {
  region          = "us-west2"
  name            = "tf-test-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}
  
resource "google_compute_region_security_policy" "policynetworkmatch" {
  region      = "us-west2"
  name        = "tf-test-polnetmatch%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}
  
resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  priority = 100
  network_match {
    src_ip_ranges    = ["10.10.0.0/16"]
  }
  action = "allow"
  preview = true
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-west2"
  name        = "tf-test-policyddos%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}
  
resource "google_compute_network_edge_security_service" "edge_sec_service" {
  region          = "us-west2"
  name            = "tf-test-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}
  
resource "google_compute_region_security_policy" "policynetworkmatch" {
  region      = "us-west2"
  name        = "tf-test-polnetmatch%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}
  
resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  priority = 100
  network_match {
    src_ip_ranges    = ["10.10.0.0/16"]
    src_asns         = [6939]
    src_ports        = [443]
    src_region_codes = ["US"]
    ip_protocols     = ["UDP"]
    dest_ip_ranges   = ["10.0.0.0/8"]
    dest_ports       = [80]
    user_defined_fields {
      name = "SIG1_AT_0"
      values = ["0x8700"]
    }
  }
  action = "allow"
  preview = true
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate2(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-west2"
  name        = "tf-test-policyddos%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}
  
resource "google_compute_network_edge_security_service" "edge_sec_service" {
  region          = "us-west2"
  name            = "tf-test-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}
  
resource "google_compute_region_security_policy" "policynetworkmatch" {
  region      = "us-west2"
  name        = "tf-test-polnetmatch%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}
  
resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  priority = 100
  network_match {
    src_ip_ranges    = ["10.0.0.0/8"]
    src_asns         = [15169]
    src_ports        = [80]
    src_region_codes = ["AU"]
    ip_protocols     = ["TCP"]
    dest_ip_ranges   = ["10.10.0.0/16"]
    dest_ports       = [443]
    user_defined_fields {
      name = "SIG2_AT_8"
      values = ["0x8700","0x8F00"]
    }
  }
  action = "allow"
  preview = true
}
`, context)
}

func TestAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig_create(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig_update(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig_clear(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("google_compute_region_security_policy_rule.policy_rule", "preconfigured_waf_config.0"),
				),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig_create(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policy" {
  name        = "tf-test%{random_suffix}"
  region      = "us-west2"
  type        = "CLOUD_ARMOR"
  description = "Regional security policy - create"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "Rule with preconfiguredWafConfig - create"
  action   = "deny"
  priority = "1000"
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
  preconfigured_waf_config {
    exclusion {
      request_cookie {
        operator = "EQUALS_ANY"
      }
      request_header {
        operator = "EQUALS"
        value    = "Referer"
      }
      request_uri {
        operator = "STARTS_WITH"
        value    = "/admin"
      }
      request_query_param {
        operator = "EQUALS"
        value    = "password"
      }
      request_query_param {
        operator = "STARTS_WITH"
        value    = "freeform"
      }
      target_rule_set = "sqli-stable"
    }
    exclusion {
      request_query_param {
        operator = "CONTAINS"
        value    = "password"
      }
      request_query_param {
        operator = "STARTS_WITH"
        value    = "freeform"
      }
      target_rule_set = "xss-stable"
    }
  }
  preview = false
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policy" {
  name        = "tf-test%{random_suffix}"
  region      = "us-west2"
  type        = "CLOUD_ARMOR"
  description = "Regional security policy - update"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "Rule with preconfiguredWafConfig - update"
  action   = "deny"
  priority = "1000"
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
  preconfigured_waf_config {
    exclusion {
      request_uri {
        operator = "STARTS_WITH"
        value    = "/admin"
      }
      target_rule_set = "rce-stable"
    }
    exclusion {
      request_query_param {
        operator = "CONTAINS"
        value    = "password"
      }
      request_query_param {
        operator = "STARTS_WITH"
        value    = "freeform"
      }
      request_query_param {
        operator = "EQUALS"
        value    = "description"
      }
      request_cookie {
        operator = "CONTAINS"
        value    = "TokenExpired"
      }
      target_rule_set = "xss-stable"
      target_rule_ids = [
        "owasp-crs-v030001-id941330-xss",
        "owasp-crs-v030001-id941340-xss",
      ]
    }
  }
  preview = false
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_withPreconfiguredWafConfig_clear(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policy" {
  name        = "tf-test%{random_suffix}"
  region      = "us-west2"
  type        = "CLOUD_ARMOR"
  description = "Regional security policy - clear"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "Rule with preconfiguredWafConfig - clear"
  action   = "deny"
  priority = "1000"
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
  preview = false
}
`, context)
}

func TestAccComputeRegionSecurityPolicyRule_withRateLimitOptions(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptionsCreate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptionsUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOptionsCreate(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_compute_region_security_policy" "default" {
    region      = "us-west2"
    name        = "tf-test%{random_suffix}"
    description = "basic region security policy"
    type        = "CLOUD_ARMOR"
  }
  
  resource "google_compute_region_security_policy_rule" "policy_rule" {
    security_policy = google_compute_region_security_policy.default.name
    region          = "us-west2"
    description     = "rule create with rate limit"
    priority        = 101
    action          = "rate_based_ban"
    rate_limit_options {
      rate_limit_threshold {
        count = 500
        interval_sec = 10
      }
      conform_action = "allow"
      exceed_action = "deny(404)"
      enforce_on_key = "ALL"
      ban_threshold {
        count = 750
        interval_sec = 180
      }
      ban_duration_sec = 180
    }
    match {
      config {
        src_ip_ranges = [
          "*"
        ]
      }
      versioned_expr = "SRC_IPS_V1"
    }
  }
`, context)
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOptionsUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_compute_region_security_policy" "default" {
    region      = "us-west2"
    name        = "tf-test%{random_suffix}"
    description = "basic region security policy"
    type        = "CLOUD_ARMOR"
  }
  
  resource "google_compute_region_security_policy_rule" "policy_rule" {
    security_policy = google_compute_region_security_policy.default.name
    region          = "us-west2"
    description     = "rule update with rate limit"
    priority        = 101
    action          = "rate_based_ban"
    rate_limit_options {
      rate_limit_threshold {
        count = 1000
        interval_sec = 30
      }
      conform_action = "allow"
      exceed_action = "deny(404)"
      enforce_on_key = "ALL"
      ban_threshold {
        count = 2000
        interval_sec = 180
      }
      ban_duration_sec = 300
    }
    match {
      config {
        src_ip_ranges = [
          "*"
        ]
      }
      versioned_expr = "SRC_IPS_V1"
    }
  }
`, context)
}

func TestAccComputeRegionSecurityPolicyRule_withRateLimit_withEnforceOnKeyConfigs(t *testing.T) {
	t.Parallel()

	spName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKeyConfigs(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionSecurityPolicyRule_withRateLimitOption_withMultipleEnforceOnKeyConfigs(t *testing.T) {
	t.Parallel()

	spName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOption_withMultipleEnforceOnKeyConfigs(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOption_withMultipleEnforceOnKeyConfigs2(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionSecurityPolicyRule_EnforceOnKeyUpdates(t *testing.T) {
	t.Parallel()

	spName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withoutRateLimitOptions(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKeyName(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKey(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKeyConfigs(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKey(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKeyName(spName),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKey(spName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policy" {
  region      = "us-west2"
  name        = "%s"
  description = "basic regional policy base"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "throttle rule withEnforceOnKey"
  action          = "throttle"
  priority        = "100"
  
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }

  rate_limit_options {
    conform_action = "allow"
    exceed_action = "deny(403)"

    enforce_on_key = "IP"

    rate_limit_threshold {
      count = 10
      interval_sec = 60
    }
  }
}
`, spName)
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKeyConfigs(spName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policy" {
  region      = "us-west2"
  name        = "%s"
  description = "basic regional policy base"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "throttle rule withEnforceOnKeyConfigs"
  action          = "throttle"
  priority        = "100"

  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }

  rate_limit_options {
    conform_action = "allow"
    exceed_action = "deny(403)"

    enforce_on_key_configs {
      enforce_on_key_type = "IP"
    }

    rate_limit_threshold {
      count = 10
      interval_sec = 60
    }
  }
}
`, spName)
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOption_withMultipleEnforceOnKeyConfigs(spName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policy" {
  region      = "us-west2"
  name        = "%s"
  description = "basic regional policy base"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "throttle rule with withMultipleEnforceOnKeyConfigs"
  action          = "throttle"
  priority        = "100"

  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }

  rate_limit_options {
    conform_action = "allow"
    exceed_action = "deny(429)"

    rate_limit_threshold {
      count = 10
      interval_sec = 60
    }

    enforce_on_key_configs {
      enforce_on_key_type = "HTTP_PATH"
    }

    enforce_on_key_configs {
      enforce_on_key_type = "HTTP_HEADER"
      enforce_on_key_name = "user-agent"
    }

    enforce_on_key_configs {
      enforce_on_key_type = "REGION_CODE"
    }
  }
}
`, spName)
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOption_withMultipleEnforceOnKeyConfigs2(spName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policy" {
  region      = "us-west2"
  name        = "%s"
  description = "basic regional policy base"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "throttle rule withMultipleEnforceOnKeyConfigs2"
  action          = "throttle"
  priority        = "100"

  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }

  rate_limit_options {
    conform_action = "allow"
    exceed_action = "deny(429)"

    rate_limit_threshold {
      count = 10
      interval_sec = 60
    }

    enforce_on_key_configs {
      enforce_on_key_type = "REGION_CODE"
    }

    enforce_on_key_configs {
      enforce_on_key_type = "TLS_JA3_FINGERPRINT"
    }

    enforce_on_key_configs {
      enforce_on_key_type = "USER_IP"
    }
  }
}

`, spName)
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withoutRateLimitOptions(spName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policy" {
  region      = "us-west2"
  name        = "%s"
  description = "basic regional policy base"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "basic policy rule withoutRateLimitOptions"
  action          = "deny(403)"
  priority        = "100"
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }
}

`, spName)
}

func testAccComputeRegionSecurityPolicyRule_withRateLimitOptions_withEnforceOnKeyName(spName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policy" {
  region      = "us-west2"
  name        = "%s"
  description = "basic regional policy base"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.policy.name
  region          = "us-west2"
  description     = "throttle rule withEnforceOnKeyName"
  action          = "throttle"
  priority        = "100"

  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }

  rate_limit_options {
    conform_action = "allow"
    exceed_action = "deny(403)"

    enforce_on_key = "HTTP_HEADER"
    enforce_on_key_name = "user-agent"

    rate_limit_threshold {
      count = 10
      interval_sec = 60
    }
  }
}
`, spName)
}
