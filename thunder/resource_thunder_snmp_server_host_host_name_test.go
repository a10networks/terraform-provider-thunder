package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_HOST_HOST_NAME_RESOURCE = `
resource "thunder_snmp_server_host_host_name" "SnmpServerHostTest" {
	hostname = "a"
version = "v1"
v1_v2c_comm = "a"
 
}
`

//Acceptance test
func TestAccSnmpServerHostHostName_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_HOST_HOST_NAME_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_host_host_name.SnmpServerHostTest", "hostname", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_host_host_name.SnmpServerHostTest", "version", "v1"),
					resource.TestCheckResourceAttr("thunder_snmp_server_host_host_name.SnmpServerHostTest", "v1_v2c_comm", "a"),
				),
			},
		},
	})
}
