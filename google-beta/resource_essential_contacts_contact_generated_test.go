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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccEssentialContactsContact_essentialContactExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckEssentialContactsContactDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEssentialContactsContact_essentialContactExample(context),
			},
			{
				ResourceName:            "google_essential_contacts_contact.contact",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccEssentialContactsContact_essentialContactExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_essential_contacts_contact" "contact" {
  parent = data.google_project.project.id
  email = "foo@bar.com"
  language_tag = "en-GB"
  notification_category_subscriptions = ["ALL"]
}
`, context)
}

func testAccCheckEssentialContactsContactDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_essential_contacts_contact" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{EssentialContactsBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("EssentialContactsContact still exists at %s", url)
			}
		}

		return nil
	}
}
