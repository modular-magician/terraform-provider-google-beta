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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccIapWebTypeAppEngineIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebTypeAppEngineIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebTypeAppEngineIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapWebTypeAppEngineIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":        acctest.GetTestOrgFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_iap_web_type_app_engine_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebTypeAppEngineIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_member" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebTypeAppEngineIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_web_type_app_engine_iam_policy" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebTypeAppEngineIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

data "google_iam_policy" "foo" {
}

resource "google_iap_web_type_app_engine_iam_policy" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_web_type_app_engine_iam_binding" "foo2" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_web_type_app_engine_iam_binding" "foo3" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_member" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_member" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_iap_web_type_app_engine_iam_member" "foo2" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_web_type_app_engine_iam_member" "foo3" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_iap_web_type_app_engine_iam_policy" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
