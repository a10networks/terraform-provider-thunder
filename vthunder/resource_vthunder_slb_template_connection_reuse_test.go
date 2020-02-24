package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_CONNECTION_REUSE_RESOURCE = `
resource "vthunder_template_connection_reuse" "testname" {
	name = "testConn"
	keep_alive_conn = 0
	limit_per_server = 10
	timeout = 120
	user_tag = "testtag"
}
`

//Acceptance test
func TestTemplateConnReuse_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_CONNECTION_REUSE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_template_connection_reuse.testname", "name", "testConn"),
					resource.TestCheckResourceAttr("vthunder_template_connection_reuse.testname", "keep_alive_conn", 0),
					resource.TestCheckResourceAttr("vthunder_template_connection_reuse.testname", "limit_per_server", 10),
					resource.TestCheckResourceAttr("vthunder_template_connection_reuse.testname", "user_tag", "tag_tester"),
					resource.TestCheckResourceAttr("vthunder_template_connection_reuse.testname", "timeout", 120),
				),
			},
		},
	})
}
