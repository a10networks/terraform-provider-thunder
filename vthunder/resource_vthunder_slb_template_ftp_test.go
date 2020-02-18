package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_FTP_RESOURCE = `
resource "vthunder_TemplateFTP" "testname" {
	name = "testftp"
	user_tag = "test_tag"
	active_mode_port = 1
	active_mode_port_val = 4048
	to = 4050
}
`

//Acceptance test
func TestTemplateFTP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_FTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_TemplateFTP.testname", "name", "testftp"),
					resource.TestCheckResourceAttr("vthunder_TemplateFTP.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_TemplateFTP.testname", "active_mode_port", 1),
					resource.TestCheckResourceAttr("vthunder_TemplateFTP.testname", "active_mode_port_val", 4048),
					resource.TestCheckResourceAttr("vthunder_TemplateFTP.testname", "to", 4050),
				),
			},
		},
	})
}
