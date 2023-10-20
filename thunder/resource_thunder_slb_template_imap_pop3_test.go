package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_TEMPLATE_IMAPPOP_RESOURCE = `
resource "thunder_slb_template_imap_pop3" "imap_pop3" {
	name = "Testimap"
	logindisabled = 0
	starttls = "disabled"
	user_tag = "test_user"
}
`

// Acceptance test
func TestAccSlbTemplateImapPOP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_IMAPPOP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_imap_pop3.imap_pop3", "name", "Testimap"),
					resource.TestCheckResourceAttr("thunder_slb_template_imap_pop3.imap_pop3", "logindisabled", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_imap_pop3.imap_pop3", "starttls", "disabled"),
					resource.TestCheckResourceAttr("thunder_slb_template_imap_pop3.imap_pop3", "user_tag", "test_user"),
				),
			},
		},
	})
}
