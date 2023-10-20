package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_GROUP_RESOURCE = `
resource "thunder_snmp_server_group" "SnmpServerTest" {
	read = "a"
groupname = "a"
v3 = "auth"
 
}
`

// Acceptance test
func TestAccSnmpServerGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_GROUP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_group.SnmpServerTest", "read", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_group.SnmpServerTest", "groupname", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_group.SnmpServerTest", "v3", "auth"),
				),
			},
		},
	})
}
