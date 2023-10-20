package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_SNMPV3_USER_RESOURCE = `
resource "thunder_snmp_server_snmpv3_user" "SnmpServerSNMPv3Test" {
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
func TestAccSnmpServerSNMPv3User_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_SNMPV3_USER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.username", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.auth_val", "md5"),
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.v3", "auth"),
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.group", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.passwd", "abcdefgh"),
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.priv", "des"),
					resource.TestCheckResourceAttr("thunder_snmp_server_SNMPv3_user.SnmpServerSNMPv3Test", "user.0.encpasswd", "abcdefgh"),
				),
			},
		},
	})
}
