package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_SERVICE_GROUP_RESOURCE = `
{'resource': {'vthunder_fw_service_group': {'service_group': {'protocol': 'tcp', 'name': 'a', 'user_tag': 'a'}}}}
`

//Acceptance test
func TestAccFwServiceGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_SERVICE_GROUP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_service_group.FwTest", "service_group", "{'protocol': 'tcp', 'name': 'a', 'user-tag': 'a'}"),
				),
			},
		},
	})
}
