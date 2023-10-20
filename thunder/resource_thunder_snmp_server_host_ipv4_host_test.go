package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_HOST_IPV4_HOST_RESOURCE = `
resource "thunder_snmp_server_host_ipv4_host" "SnmpServerHostTest" {
	ipv4_addr = "3.3.3.3"
version = "v1"
v1_v2c_comm = "a"
 
}
`

// Acceptance test
func TestAccSnmpServerHostIpv4Host_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_HOST_IPV4_HOST_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_host_ipv4_host.SnmpServerHostTest", "ipv4_addr", "3.3.3.3"),
					resource.TestCheckResourceAttr("thunder_snmp_server_host_ipv4_host.SnmpServerHostTest", "version", "v1"),
					resource.TestCheckResourceAttr("thunder_snmp_server_host_ipv4_host.SnmpServerHostTest", "v1_v2c_comm", "a"),
				),
			},
		},
	})
}
