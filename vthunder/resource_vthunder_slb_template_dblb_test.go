package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_DBLB_RESOURCE = `
resource "vthunder_slb_template_dblb" "testname" {
	name = "testdblb"
	user_tag = "test_tag"
	server_version = "MSSQL2008"
}
`

//Acceptance test
func TestTemplateDBLB_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_DBLB_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_dblb.testname", "name", "testdblb"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dblb.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dblb.testname", "server_version", "MSSQL2008"),
				),
			},
		},
	})
}
