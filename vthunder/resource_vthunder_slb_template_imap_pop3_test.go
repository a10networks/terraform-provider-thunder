package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_IMAPPOP_RESOURCE = `
resource "vthunder_slb_template_imap_pop3" "imap_pop3" {
	name = "Testimap"
	logindisabled = 0
	starttls = "disabled"
	user_tag = "test_user"
}
`

//Acceptance test
func TestAccSlbTemplateImapPOP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_IMAPPOP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_imap_pop3.imap_pop3", "name", "Testimap"),
					resource.TestCheckResourceAttr("vthunder_slb_template_imap_pop3.imap_pop3", "logindisabled", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_imap_pop3.imap_pop3", "starttls", "disabled"),
					resource.TestCheckResourceAttr("vthunder_slb_template_imap_pop3.imap_pop3", "user_tag", "test_user"),
				),
			},
		},
	})
}
