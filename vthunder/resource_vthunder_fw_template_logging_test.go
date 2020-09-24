package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_TEMPLATE_LOGGING_RESOURCE = `
resource "vthunder_fw_template_logging" "FwTemplateTest" {
	name = "a"
	user_tag = "a" 
}
`

//Acceptance test
func TestAccFwTemplateLogging_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TEMPLATE_LOGGING_RESOURCE,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr("vthunder_fw_template_logging.FwTemplateTest", "name", "a"),
					resource.TestCheckResourceAttr("vthunder_fw_template_logging.FwTemplateTest", "user_tag", "a"),
				),
			},
		},
	})
}
