// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccProjectServiceIdentity_basic(t *testing.T) {
	t.Parallel()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testGoogleProjectServiceIdentity_basic(),
				Check: resource.ComposeTestCheckFunc(
					// Email field for healthcare service account should be non-empty and contain at least an "@".
					resource.TestCheckResourceAttrWith("google_project_service_identity.hc_sa", "email", func(value string) error {
						if strings.Contains(value, "@") {
							return nil
						}
						return fmt.Errorf("hc_sa service identity email value was %s, expected a valid email", value)
					}),
					// Member field for healthcare service account should be non-empty, start with "serviceAccount:" and contain at least an "@".
					resource.TestCheckResourceAttrWith("google_project_service_identity.hc_sa", "member", func(value string) error {
						if strings.HasPrefix(value, "serviceAccount:") && strings.Contains(value, "@") {
							return nil
						}
						return fmt.Errorf("hc_sa service identity member value was %s, expected a valid email with prefix serviceAccount:", value)
					}),
					// Email field for logging service identity will be empty for as long as
					// `gcloud beta services identity create --service=logging.googleapis.com` doesn't return an email address
					resource.TestCheckNoResourceAttr("google_project_service_identity.log_sa", "email"),
					// Member field for logging service identity will be empty for as long as
					// `gcloud beta services identity create --service=logging.googleapis.com` doesn't return an email address
					resource.TestCheckNoResourceAttr("google_project_service_identity.log_sa", "member"),
				),
			},
		},
	})
}

func testGoogleProjectServiceIdentity_basic() string {
	return `
data "google_project" "project" {}

# Service which has an email returned from service identity API
resource "google_project_service_identity" "hc_sa" {
  project = data.google_project.project.project_id
  service = "healthcare.googleapis.com"
}

# Service which does NOT have email returned from service identity API
# Email attribute will be null - correct as of 2022-12-13
resource "google_project_service_identity" "log_sa" {
  project = data.google_project.project.project_id
  service = "logging.googleapis.com"
}
`
}
