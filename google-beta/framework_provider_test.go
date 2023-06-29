// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccFrameworkProviderMeta_setModuleName(t *testing.T) {
	// TODO: https://github.com/hashicorp/terraform-provider-google/issues/14158
	acctest.SkipIfVcr(t)
	t.Parallel()

	moduleName := "my-module"
	managedZoneName := fmt.Sprintf("tf-test-zone-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducerFramework(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFrameworkProviderMeta_setModuleName(moduleName, managedZoneName, RandString(t, 10)),
			},
		},
	})
}

func TestAccFrameworkProviderBasePath_setInvalidBasePath(t *testing.T) {
	t.Parallel()

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { acctest.AccTestPreCheck(t) },
		CheckDestroy: testAccCheckComputeAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"google": {
						VersionConstraint: "4.58.0",
						Source:            "hashicorp/google-beta",
					},
				},
				Config:      testAccProviderBasePath_setBasePath("https://www.example.com/compute/beta/", RandString(t, 10)),
				ExpectError: regexp.MustCompile("got HTTP response code 404 with body"),
			},
			{
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Config:                   testAccProviderBasePath_setBasePath("https://www.example.com/compute/beta/", RandString(t, 10)),
				ExpectError:              regexp.MustCompile("got HTTP response code 404 with body"),
			},
		},
	})
}

func TestAccFrameworkProviderBasePath_setBasePath(t *testing.T) {
	// TODO: https://github.com/hashicorp/terraform-provider-google/issues/14158
	acctest.SkipIfVcr(t)
	t.Parallel()

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { acctest.AccTestPreCheck(t) },
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducerFramework(t),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"google": {
						VersionConstraint: "4.58.0",
						Source:            "hashicorp/google-beta",
					},
				},
				Config: testAccFrameworkProviderBasePath_setBasePath("https://www.googleapis.com/dns/v1beta2/", RandString(t, 10)),
			},
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"google": {
						VersionConstraint: "4.58.0",
						Source:            "hashicorp/google-beta",
					},
				},
				ResourceName:      "google_dns_managed_zone.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Config:                   testAccFrameworkProviderBasePath_setBasePath("https://www.googleapis.com/dns/v1beta2/", RandString(t, 10)),
			},
			{
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				ResourceName:             "google_dns_managed_zone.foo",
				ImportState:              true,
				ImportStateVerify:        true,
			},
			{
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Config:                   testAccFrameworkProviderBasePath_setBasePathstep3("https://www.googleapis.com/dns/v1beta2/", RandString(t, 10)),
			},
		},
	})
}

func testAccFrameworkProviderMeta_setModuleName(key, managedZoneName, recordSetName string) string {
	return fmt.Sprintf(`
terraform {
  provider_meta "google" {
    module_name = "%s"
  }
}

provider "google" {}

resource "google_dns_managed_zone" "zone" {
  name     = "%s-hashicorptest-com"
  dns_name = "%s.hashicorptest.com."
}

resource "google_dns_record_set" "rs" {
  managed_zone = google_dns_managed_zone.zone.name
  name         = "%s.${google_dns_managed_zone.zone.dns_name}"
  type         = "A"
  ttl          = 300
  rrdatas      = [
  "192.168.1.0",
  ]
}

data "google_dns_record_set" "rs" {
  managed_zone = google_dns_record_set.rs.managed_zone
  name         = google_dns_record_set.rs.name
  type         = google_dns_record_set.rs.type
}`, key, managedZoneName, managedZoneName, recordSetName)
}

func testAccFrameworkProviderBasePath_setBasePath(endpoint, name string) string {
	return fmt.Sprintf(`
provider "google" {
  alias               = "dns_custom_endpoint"
  dns_custom_endpoint = "%s"
}

resource "google_dns_managed_zone" "foo" {
  provider    = google.dns_custom_endpoint
  name        = "tf-test-zone-%s"
  dns_name    = "tf-test-zone-%s.hashicorptest.com."
  description = "QA DNS zone"
}

data "google_dns_managed_zone" "qa" {
  provider    = google.dns_custom_endpoint
  name = google_dns_managed_zone.foo.name
}`, endpoint, name, name)
}

func testAccFrameworkProviderBasePath_setBasePathstep3(endpoint, name string) string {
	return fmt.Sprintf(`
provider "google" {
  alias               = "dns_custom_endpoint"
  dns_custom_endpoint = "%s"
}

resource "google_dns_managed_zone" "foo" {
  provider    = google.dns_custom_endpoint
  name        = "tf-test-zone-%s"
  dns_name    = "tf-test-zone-%s.hashicorptest.com."
  description = "QA DNS zone"
}
`, endpoint, name, name)
}
