package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_LOGGING_RESOURCE = `
resource "thunder_slb_template_logging" "logging"{
name = "log1"
format= "abc"
local_logging= 1
}
`

//Acceptance test
func TestAccThunderLogging_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_LOGGING_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_logging.logging", "name", "log1"),
					resource.TestCheckResourceAttr("thunder_slb_template_logging.logging", "format", "abc"),
					resource.TestCheckResourceAttr("thunder_slb_template_logging.logging", "local_logging", "1"),
				),
			},
		},
	})
}
