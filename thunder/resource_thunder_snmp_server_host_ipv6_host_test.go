package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_HOST_IPV6_HOST_RESOURCE = `
resource "thunder_snmp_server_host_ipv6_host" "SnmpServerHostTest" {
	version = "v1"
ipv6_addr = "2003::1"
v1_v2c_comm = "a"
 
}
`

// Acceptance test
func TestAccSnmpServerHostIpv6Host_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_HOST_IPV6_HOST_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_host_ipv6_host.SnmpServerHostTest", "version", "v1"),
					resource.TestCheckResourceAttr("thunder_snmp_server_host_ipv6_host.SnmpServerHostTest", "ipv6_addr", "2003::1"),
					resource.TestCheckResourceAttr("thunder_snmp_server_host_ipv6_host.SnmpServerHostTest", "v1_v2c_comm", "a"),
				),
			},
		},
	})
}
