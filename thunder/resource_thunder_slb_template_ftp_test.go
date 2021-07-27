package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_FTP_RESOURCE = `
resource "thunder_slb_template_ftp" "template_ftp" {
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
					resource.TestCheckResourceAttr("thunder_slb_template_ftp.template_ftp", "name", "testftp"),
					resource.TestCheckResourceAttr("thunder_slb_template_ftp.template_ftp", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_ftp.template_ftp", "active_mode_port", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_ftp.template_ftp", "active_mode_port_val", "1"),
				),
			},
		},
	})
}
