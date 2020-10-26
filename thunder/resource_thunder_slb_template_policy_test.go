package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_POLICY_RESOURCE = `
resource "thunder_slb_template_policy" "policy" {
	name = "testdynamicpolicy"
	user_tag = "test_tag"
	use_destination_ip = 1
	over_limit = 1
	timeout = 1
	over_limit_lockup = 2
	over_limit_reset = 1
	overlap = 1
}
`

//Acceptance test
func TestAccSlbTemplatePolicy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_POLICY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "name", "testdynamicpolicy"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "use_destination_ip", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "over_limit", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "timeout", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "over_limit_lockup", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "over_limit_reset", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_policy.policy", "overlap", "1"),
				),
			},
		},
	})
}
