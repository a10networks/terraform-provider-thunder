package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_USER_RESOURCE = `
resource "thunder_snmp_server_user" "SnmpServerTest" {
	username = "a"
auth_val = "md5"
v3 = "auth"
group = "a"
passwd = "abcdefgh"
priv = "des"
encpasswd = "abcdefgh"
 
}
`

// Acceptance test
func TestAccSnmpServerUser_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_USER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "username", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "auth_val", "md5"),
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "v3", "auth"),
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "group", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "passwd", "abcdefgh"),
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "priv", "des"),
					resource.TestCheckResourceAttr("thunder_snmp_server_user.SnmpServerTest", "encpasswd", "abcdefgh"),
				),
			},
		},
	})
}
