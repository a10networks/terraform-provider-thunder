package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_FIX_RESOURCE = `
resource "vthunder_slb_template_fix" "template_fix" {
	name = "testfix"
	logging = "init"
	tag_switching {
		equals = "test"
		service_group = "testsg"
		switching_type = "sender-comp-id"
	} 
	insert_client_ip = 1111
	user_tag = "test_user"
}
`

//Acceptance test
func TestAccSlbTemplateFix_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_FIX_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "name", "testfix"),
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "tag_switching.0.equals", "test"),
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "tag_switching.0.service_group", "testsg"),
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "tag_switching.0.switching_type", "sender-comp-id"),
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "logging", "init"),
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "insert_client_ip", "1111"),
					resource.TestCheckResourceAttr("vthunder_slb_template_fix.template_fix", "user_tag", "test_user"),
				),
			},
		},
	})
}
