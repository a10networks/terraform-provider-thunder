package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_ACTIVE_RULE_SET_RESOURCE = `
{'resource': {'vthunder_fw_active_rule_set': {'active_rule_set': {'name': 'a', 'session_aging': 'a'}}}}
`

//Acceptance test
func TestAccFwActiveRuleSet_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_ACTIVE_RULE_SET_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_active_rule_set.FwTest", "active_rule_set", "{'name': 'a', 'session-aging': 'a'}"),
				),
			},
		},
	})
}
