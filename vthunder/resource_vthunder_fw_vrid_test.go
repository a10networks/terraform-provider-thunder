package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_VRID_RESOURCE = `
{'resource': {'vthunder_fw_vrid': {'vrid': {'vrid': 1}}}}
`

//Acceptance test
func TestAccFwVrid_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_VRID_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_vrid.FwTest", "vrid", "{'vrid': 1}"),
				),
			},
		},
	})
}
