package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SMTP_RESOURCE = `
resource "thunder_slb_template_smtp" "smtp" {
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
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "name", "testsmtp"),
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "server_domain", "gmail.com"),
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "client_starttls_type", "optional"),
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "command_disable.0.disable_type", "expn"),
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "server_starttls_type", "optional"),
					resource.TestCheckResourceAttr("thunder_slb_template_smtp.smtp", "service_ready_msg", "ttt test"),
				),
			},
		},
	})
}
