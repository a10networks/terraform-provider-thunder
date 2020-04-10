package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TRANSPARENT_ACL_TEMPLATE_RESOURCE = `
resource "vthunder_slb_transparent_acl_template" "testname" {
	name = "testtransparentacltemplate"
}
`

//Acceptance test
func TestAccSlbTransparentAclTemplate_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TRANSPARENT_ACL_TEMPLATE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_transparent_acl_template.testname", "name", "testtransparentacltemplate"),
				),
			},
		},
	})
}
