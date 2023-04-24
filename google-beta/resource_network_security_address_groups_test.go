package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkSecurityAddressGroups_update(t *testing.T) {
	t.Parallel()

	addressGroupsName := fmt.Sprintf("tf-test-address-group-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroups_basic(addressGroupsName),
			},
			{
				ResourceName:      "google_network_security_address_group.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityAddressGroups_update(addressGroupsName),
			},
			{
				ResourceName:      "google_network_security_address_group.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityAddressGroups_basic(addressGroupsName string) string {
	return fmt.Sprintf(`
resource "google_network_security_address_group" "foobar" {
    name        = "%s"
    location    = "us-central1"
    description = "my address groups"
    type        = "IPV4"
    capacity    = "100"
    labels      = {
		foo = "bar"
    }
    items = ["208.80.154.224/32"]
}
`, addressGroupsName)
}

func testAccNetworkSecurityAddressGroups_update(addressGroupsName string) string {
	return fmt.Sprintf(`
resource "google_network_security_address_group" "foobar" {
    name        = "%s"
    location    = "us-central1"
    description = "my address groups. Update"
    type        = "IPV4"
    capacity    = "100"
    labels      = {
		foo = "foo"
    }
    items = ["208.80.155.224/32", "208.80.154.224/32"]
}
`, addressGroupsName)
}
