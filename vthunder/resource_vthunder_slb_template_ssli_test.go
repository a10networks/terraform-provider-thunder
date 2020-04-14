package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_SSLI_RESOURCE = `
resource "vthunder_slb_template_ssli" "ssli" {
	name = "testssli"
	type = "init"
	user_tag = "test_user"
}
`

//Acceptance test
func TestAccTemplateSsli_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_SSLI_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_ssli.ssli", "name", "testssli"),
					resource.TestCheckResourceAttr("vthunder_slb_template_ssli.ssli", "type", "init"),
					resource.TestCheckResourceAttr("vthunder_slb_template_ssli.ssli", "user_tag", "test_user"),
				),
			},
		},
	})
}
