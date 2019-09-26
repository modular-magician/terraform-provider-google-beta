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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIapWebIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
	project_id = "tf-test%{random_suffix}"
	name = "tf-test%{random_suffix}"
	org_id = "%{org_id}"
}

resource "google_project_service" "project_service" {
	project = "${google_project.project.project_id}"
	service = "iap.googleapis.com"
}

resource "google_iap_web_iam_member" "foo" {
	project = "${google_project_service.project_service.project}"
	role = "%{role}"
	member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
	project_id = "tf-test%{random_suffix}"
	name = "tf-test%{random_suffix}"
	org_id = "%{org_id}"
}

resource "google_project_service" "project_service" {
	project = "${google_project.project.project_id}"
	service = "iap.googleapis.com"
}

data "google_iam_policy" "foo" {
	binding {
		role = "%{role}"
		members = ["user:admin@hashicorptest.com"]
	}
}

resource "google_iap_web_iam_policy" "foo" {
	project = "${google_project_service.project_service.project}"
	policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccIapWebIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
	project_id = "tf-test%{random_suffix}"
	name = "tf-test%{random_suffix}"
	org_id = "%{org_id}"
}

resource "google_project_service" "project_service" {
	project = "${google_project.project.project_id}"
	service = "iap.googleapis.com"
}

resource "google_iap_web_iam_binding" "foo" {
	project = "${google_project_service.project_service.project}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
	project_id = "tf-test%{random_suffix}"
	name = "tf-test%{random_suffix}"
	org_id = "%{org_id}"
}

resource "google_project_service" "project_service" {
	project = "${google_project.project.project_id}"
	service = "iap.googleapis.com"
}

resource "google_iap_web_iam_binding" "foo" {
	project = "${google_project_service.project_service.project}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
