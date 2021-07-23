package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_UDP_RESOURCE = `
resource "thunder_slb_template_udp" "udp"{
qos = 12
name = "udp2"
stateless_conn_timeout = 120
idle_timeout = 120
re_select_if_server_down = 1
immediate = 1
user_tag = "tag1"
}
`

//Acceptance test
func TestAccThunderUdp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_UDP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "name", "udp2"),
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "qos", "12"),
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "stateless_conn_timeout", "120"),
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "idle_timeout", "120"),
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "re_select_if_server_down", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "immediate", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_udp.udp", "user_tag", "tag1"),
				),
			},
		},
	})
}
