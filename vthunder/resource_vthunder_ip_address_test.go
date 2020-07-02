package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_ADDRESS_RESOURCE = `
resource "vthunder_ip_address" "testname" {
	ip_addr = "3.3.3.3"
	ip_mask = "255.255.0.0"
}
`

//Acceptance test
func TestIPAddress_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_ADDRESS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "ip_addr", "3.3.3.3"),
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "ip_mask", "255.255.0.0"),
				),
			},
		},
	})
}
