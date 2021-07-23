package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_VIRTUAL_SERVER_RESOURCE = `
resource "thunder_slb_template_virtual_server" "virtual_server" {
	name = "testvirtualserver"
	user_tag = "test_tag"
	conn_limit = 1
	conn_limit_reset = 1
	icmpv6_rate_limit = 1
	subnet_gratuitous_arp = 1
	tcp_stack_tfo_backoff_time = 1 
}
`

//Acceptance test
func TestAccSlbTemplateVirtualServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_VIRTUAL_SERVER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "name", "testvirtualserver"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "conn_limit", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "conn_limit_reset", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "icmpv6_rate_limit", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "subnet_gratuitous_arp", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_server.virtual_server", "tcp_stack_tfo_backoff_time", "1"),
				),
			},
		},
	})
}
