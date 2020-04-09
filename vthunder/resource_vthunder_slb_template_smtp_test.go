package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SMTP_RESOURCE = `
resource "vthunder_slb_template_smtp" "testname" {
	name = "testsmtp"
	user_tag = "test_tag"
	server_domain= "gmail.com"
	client_starttls_type = "optional"
	command_disable {
		disable_type = "expn"
	}
	server_starttls_type = "optional"
	service_ready_msg = "ttt test"
}
`

//Acceptance test
func TestAccSlbTemplateSMTP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SMTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "name", "testsmtp"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "server_domain", "gmail.com"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "client_starttls_type", "optional"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "command_disable.0.disable_type", "expn"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "server_starttls_type", "optional"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smtp.testname", "service_ready_msg", "ttt test"),
				),
			},
		},
	})
}
