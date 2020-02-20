package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_IMAPPOP_RESOURCE = `
resource "resourceTemplateImap_POP3" "testname" {
	name = "Testimap"
	logindisabled = 0
	starttls = "disabled"
	user_tag = "test_user"
}
`

//Acceptance test
func TestTemplateImapPOP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_IMAPPOP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("resourceTemplateImap_POP3.testname", "name", "Testimap"),
					resource.TestCheckResourceAttr("resourceTemplateImap_POP3.testname", "logindisabled", 0),
					resource.TestCheckResourceAttr("resourceTemplateImap_POP3.testname", "starttls", "disabled"),
					resource.TestCheckResourceAttr("resourceTemplateImap_POP3.testname", "user_tag", "test_user"),
				),
			},
		},
	})
}
