package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_FTP_RESOURCE = `
resource "vthunder_slb_template_ftp" "testname" {
	name = "testftp"
	user_tag = "test_tag"
	active_mode_port = 1
	active_mode_port_val = 1
}
`

//Acceptance test
func TestAccSlbTemplateFTP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_FTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_ftp.testname", "name", "testftp"),
					resource.TestCheckResourceAttr("vthunder_slb_template_ftp.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_ftp.testname", "active_mode_port", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_ftp.testname", "active_mode_port_val", "1"),
				),
			},
		},
	})
}
