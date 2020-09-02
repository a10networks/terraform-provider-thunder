package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_URPF_RESOURCE = `
resource "vthunder_fw_urpf" "FwTest" {
	status = "strict" 
}
`

//Acceptance test
func TestAccFwUrpf_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_URPF_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_urpf.FwTest", "status", "strict"),
				),
			},
		},
	})
}
