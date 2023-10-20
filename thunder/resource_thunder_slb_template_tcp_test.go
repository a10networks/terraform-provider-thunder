package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_TCP_RESOURCE = `
resource "thunder_slb_template_tcp" "tcp"{
name= "tcp2"
idle_timeout= 120
insert_client_ip= 0
lan_fast_ack= 0
reset_fwd= 0
reset_rev= 0
reset_follow_fin= 0
del_session_on_server_down= 0
}
`

// Acceptance test
func TestAccThunderTcp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TCP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "name", "tcp2"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "idle_timeout", "120"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "insert_client_ip", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "lan_fast_ack", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "reset_fwd", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "reset_rev", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "reset_follow_fin", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp.tcp", "del_session_on_server_down", "0"),
				),
			},
		},
	})
}
