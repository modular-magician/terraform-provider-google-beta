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
)

func TestAccFirebaseHostingChannel_firebasehostingChannelBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseHostingChannelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
		},
	})
}

func testAccFirebaseHostingChannel_firebasehostingChannelBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "default" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-basic%{random_suffix}"
}
`, context)
}

func TestAccFirebaseHostingChannel_firebasehostingChannelFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseHostingChannelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelFullExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
		},
	})
}

func testAccFirebaseHostingChannel_firebasehostingChannelFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "full" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-full%{random_suffix}"
  ttl = "86400s"
  retained_release_count = 20
  labels = {
    "some-key": "some-value"
  }
}
`, context)
}

func testAccCheckFirebaseHostingChannelDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_hosting_channel" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("FirebaseHostingChannel still exists at %s", url)
			}
		}

		return nil
	}
}
