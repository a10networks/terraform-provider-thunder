package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_SERVER_SSL_RESOURCE = `
resource "thunder_slb_template_server_ssl" "server_ssl" {
	name = "testserverssl"
	user_tag = "test_tag"
	sslilogging = "disable"
	ocsp_stapling = 1
	ssli_logging = 1
	session_cache_size = 1
	handshake_logging_enable = 1
	close_notify = 1
}
`

//Acceptance test
func TestAccSlbTemplateServerSSL_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SERVER_SSL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "name", "testserverssl"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "sslilogging", "disable"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "ocsp_stapling", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "ssli_logging", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "session_cache_size", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "handshake_logging_enable", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server_ssl.server_ssl", "close_notify", "1"),
				),
			},
		},
	})
}
