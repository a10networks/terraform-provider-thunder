package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SERVER_SSL_RESOURCE = `
resource "vthunder_slb_template_server_ssl" "testname" {
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
func TestSlbTemplateServerSSL_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SERVER_SSL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "name", "testserverssl"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "sslilogging", "disable"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "ocsp_stapling", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "ssli_logging", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "session_cache_size", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "handshake_logging_enable", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server_ssl.testname", "close_notify", "1"),
				),
			},
		},
	})
}
