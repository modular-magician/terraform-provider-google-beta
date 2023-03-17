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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataCatalogTaxonomyIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTaxonomyIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccDataCatalogTaxonomyIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccDataCatalogTaxonomyIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataCatalogTaxonomyIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccDataCatalogTaxonomyIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTaxonomyIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccDataCatalogTaxonomyIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccDataCatalogTaxonomyIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_taxonomy" "basic_taxonomy" {
  display_name =  "tf_test_my_taxonomy%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_data_catalog_taxonomy_iam_member" "foo" {
  taxonomy = google_data_catalog_taxonomy.basic_taxonomy.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataCatalogTaxonomyIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_taxonomy" "basic_taxonomy" {
  display_name =  "tf_test_my_taxonomy%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_data_catalog_taxonomy_iam_policy" "foo" {
  taxonomy = google_data_catalog_taxonomy.basic_taxonomy.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataCatalogTaxonomyIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_taxonomy" "basic_taxonomy" {
  display_name =  "tf_test_my_taxonomy%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

data "google_iam_policy" "foo" {
}

resource "google_data_catalog_taxonomy_iam_policy" "foo" {
  taxonomy = google_data_catalog_taxonomy.basic_taxonomy.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataCatalogTaxonomyIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_taxonomy" "basic_taxonomy" {
  display_name =  "tf_test_my_taxonomy%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_data_catalog_taxonomy_iam_binding" "foo" {
  taxonomy = google_data_catalog_taxonomy.basic_taxonomy.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataCatalogTaxonomyIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_taxonomy" "basic_taxonomy" {
  display_name =  "tf_test_my_taxonomy%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_data_catalog_taxonomy_iam_binding" "foo" {
  taxonomy = google_data_catalog_taxonomy.basic_taxonomy.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
