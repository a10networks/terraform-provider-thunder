package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_SSLI_RESOURCE = `
resource "vthunder_TemplateSsli" "testname" {
	name = "testssli"
	type = "init"
	user_tag = 1
}
`

//Acceptance test
func TestTemplateSsli_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_SSLI_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_TemplateSsli.testname", "name", "testssli"),
					resource.TestCheckResourceAttr("vthunder_TemplateSsli.testname", "type", "http"),
					resource.TestCheckResourceAttr("vthunder_TemplateSsli.testname", "user_tag", "test_user"),
				),
			},
		},
	})
}
