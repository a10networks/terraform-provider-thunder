package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_TEMPLATE_LOGGING_RESOURCE = `
resource "thunder_fw_template_logging" "FwTemplateTest" {
	name = "a"
	user_tag = "a" 
}
`

// Acceptance test
func TestAccFwTemplateLogging_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TEMPLATE_LOGGING_RESOURCE,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr("thunder_fw_template_logging.FwTemplateTest", "name", "a"),
					resource.TestCheckResourceAttr("thunder_fw_template_logging.FwTemplateTest", "user_tag", "a"),
				),
			},
		},
	})
}
