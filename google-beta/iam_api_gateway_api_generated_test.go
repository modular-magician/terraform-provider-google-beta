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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayApiIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccApiGatewayApiIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccApiGatewayApiIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccApiGatewayApiIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccApiGatewayApiIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccApiGatewayApiIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccApiGatewayApiIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
}

resource "google_api_gateway_api_iam_member" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccApiGatewayApiIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_api_gateway_api_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayApiIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_api_gateway_api_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayApiIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
}

resource "google_api_gateway_api_iam_binding" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccApiGatewayApiIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
}

resource "google_api_gateway_api_iam_binding" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
