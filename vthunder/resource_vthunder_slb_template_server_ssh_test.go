package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SERVER_SSH_RESOURCE = `
resource "vthunder_slb_template_server_ssh" "testname" {
	name = "testserverssh"
	user_tag = "test_tag"
	forward_proxy_enable = 1
	
}
`

//Acceptance test
func TestSlbTemplateServerSSH_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SERVER_SSH_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssh.testname", "name", "testserverssh"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssh.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssh.testname", "forward_proxy_enable", "1"),
				),
			},
		},
	})
}
