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

package netapp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetappactiveDirectory_netappActiveDirectoryFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappactiveDirectoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappactiveDirectory_netappActiveDirectoryFullExample(context),
			},
			{
				ResourceName:            "google_netapp_active_directory.test_active_directory_full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "password", "terraform_labels"},
			},
		},
	})
}

func testAccNetappactiveDirectory_netappActiveDirectoryFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_active_directory" "test_active_directory_full" {
    name = "tf-test-test-active-directory-full%{random_suffix}"
    location = "us-central1"
    domain = "ad.internal"
    dns = "172.30.64.3"
    net_bios_prefix = "smbserver"
    username = "user"
    password = "pass"
    aes_encryption         = false
    backup_operators       = ["test1", "test2"]
    administrators         = ["test1", "test2"]
    description            = "ActiveDirectory is the public representation of the active directory config."
    encrypt_dc_connections = false
    kdc_hostname           = "hostname"
    kdc_ip                 = "10.10.0.11"
    labels                 = { 
        "foo": "bar"
    }
    ldap_signing           = false
    nfs_users_with_ldap    = false
    organizational_unit    = "CN=Computers"
    security_operators     = ["test1", "test2"]
    site                   = "test-site"
  }
`, context)
}

func testAccCheckNetappactiveDirectoryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_netapp_active_directory" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetappactiveDirectory still exists at %s", url)
			}
		}

		return nil
	}
}
