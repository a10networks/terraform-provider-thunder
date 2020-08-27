package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_APPLY_CHANGES_RESOURCE = `
{'resource': {'vthunder_fw_apply_changes': {'apply_changes': {'forced': 1}}}}
`

//Acceptance test
func TestAccFwApplyChanges_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_APPLY_CHANGES_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_apply_changes.FwTest", "apply_changes", "{'forced': 1}"),
				),
			},
		},
	})
}
