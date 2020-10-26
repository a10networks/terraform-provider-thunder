package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_TCP_RESOURCE = `
resource "thunder_ip_tcp" "Iptcp" {
	syn_cookie  {
	    threshold = 4
	}
}
`

//Acceptance test
func TestAccIpTcp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_TCP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ip_tcp.Iptcp", "syn_cookie.0.threshold", "4"),
				),
			},
		},
	})
}
