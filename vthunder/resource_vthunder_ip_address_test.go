package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_ADDRESS_RESOURCE = `
resource "vthunder_ip_address" "testname" {
	name = "testaddress"
	user_tag = "test_tag"
	ip_addr = string
	ip_mask = string
	uuid = string 
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
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "name", "testaddress"),
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "ip_addr", "string"),
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "ip_mask", "string"),
					resource.TestCheckResourceAttr("vthunder_ip_address.testname", "uuid", "string"),
				),
			},
		},
	})
}
