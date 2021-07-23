package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_TEMPLATE_FIX_RESOURCE = `
resource "thunder_slb_template_fix" "template_fix" {
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
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "name", "testfix"),
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "tag_switching.0.equals", "test"),
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "tag_switching.0.service_group", "testsg"),
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "tag_switching.0.switching_type", "sender-comp-id"),
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "logging", "init"),
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "insert_client_ip", "1111"),
					resource.TestCheckResourceAttr("thunder_slb_template_fix.template_fix", "user_tag", "test_user"),
				),
			},
		},
	})
}
