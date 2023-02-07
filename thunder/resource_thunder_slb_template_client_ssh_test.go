package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_TEMPLATE_CLIENT_SSH_RESOURCE = `
resource "thunder_slb_template_client_ssh" "client_ssh1" {
	name = "testssh"
	user_tag = "test_tag"
	forward_proxy_enable = 1
	forward_proxy_hostkey = "test"
	passphrase = "testing123"
}
`

//Acceptance test
func TestAccTemplateClientSSH_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_CLIENT_SSH_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssh.client_ssh1", "name", "testssh"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssh.client_ssh1", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssh.client_ssh1", "forward_proxy_enable", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssh.client_ssh1", "forward_proxy_hostkey", "test"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssh.client_ssh1", "passphrase", "testing123"),
				),
			},
		},
	})
}
