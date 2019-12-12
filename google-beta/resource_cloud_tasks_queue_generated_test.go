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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudTasksQueue_queueBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudTasksQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_queueBasicExample(context),
			},
			{
				ResourceName:      "google_cloud_tasks_queue.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudTasksQueue_queueBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_tasks_queue" "default" {
  name = "cloud-tasks-queue-test%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func testAccCheckCloudTasksQueueDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_cloud_tasks_queue" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("CloudTasksQueue still exists at %s", url)
		}
	}

	return nil
}
