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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataprocAutoscalingPolicyIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocAutoscalingPolicyIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_autoscaling_policy_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-dataproc-policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataprocAutoscalingPolicyIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_autoscaling_policy_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-dataproc-policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocAutoscalingPolicyIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataprocAutoscalingPolicyIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_autoscaling_policy_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-dataproc-policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocAutoscalingPolicyIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocAutoscalingPolicyIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataproc_autoscaling_policy_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataproc_autoscaling_policy_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-dataproc-policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataprocAutoscalingPolicyIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataproc_autoscaling_policy_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-dataproc-policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataprocAutoscalingPolicyIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_autoscaling_policy" "basic" {
  policy_id = "tf-test-dataproc-policy%{random_suffix}"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}

resource "google_dataproc_autoscaling_policy_iam_member" "foo" {
  project = google_dataproc_autoscaling_policy.basic.project
  location = google_dataproc_autoscaling_policy.basic.location
  policy_id = google_dataproc_autoscaling_policy.basic.policy_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataprocAutoscalingPolicyIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_autoscaling_policy" "basic" {
  policy_id = "tf-test-dataproc-policy%{random_suffix}"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataproc_autoscaling_policy_iam_policy" "foo" {
  project = google_dataproc_autoscaling_policy.basic.project
  location = google_dataproc_autoscaling_policy.basic.location
  policy_id = google_dataproc_autoscaling_policy.basic.policy_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataproc_autoscaling_policy_iam_policy" "foo" {
  project = google_dataproc_autoscaling_policy.basic.project
  location = google_dataproc_autoscaling_policy.basic.location
  policy_id = google_dataproc_autoscaling_policy.basic.policy_id
  depends_on = [
    google_dataproc_autoscaling_policy_iam_policy.foo
  ]
}
`, context)
}

func testAccDataprocAutoscalingPolicyIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_autoscaling_policy" "basic" {
  policy_id = "tf-test-dataproc-policy%{random_suffix}"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_dataproc_autoscaling_policy_iam_policy" "foo" {
  project = google_dataproc_autoscaling_policy.basic.project
  location = google_dataproc_autoscaling_policy.basic.location
  policy_id = google_dataproc_autoscaling_policy.basic.policy_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataprocAutoscalingPolicyIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_autoscaling_policy" "basic" {
  policy_id = "tf-test-dataproc-policy%{random_suffix}"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}

resource "google_dataproc_autoscaling_policy_iam_binding" "foo" {
  project = google_dataproc_autoscaling_policy.basic.project
  location = google_dataproc_autoscaling_policy.basic.location
  policy_id = google_dataproc_autoscaling_policy.basic.policy_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataprocAutoscalingPolicyIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_autoscaling_policy" "basic" {
  policy_id = "tf-test-dataproc-policy%{random_suffix}"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}

resource "google_dataproc_autoscaling_policy_iam_binding" "foo" {
  project = google_dataproc_autoscaling_policy.basic.project
  location = google_dataproc_autoscaling_policy.basic.location
  policy_id = google_dataproc_autoscaling_policy.basic.policy_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
