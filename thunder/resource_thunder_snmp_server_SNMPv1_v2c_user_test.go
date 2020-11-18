package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_SNMPV1_V2C_USER_RESOURCE = `
resource "thunder_snmp_server_snmpv1_v2c_user" "thunder_snmp_server_snmpv1_v2c_user_test" {
	user = "a"
user_tag = "a"
 
}
`

//Acceptance test
func TestAccSnmpServerSNMPv1V2cUser_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_SNMPV1_V2C_USER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_snmpv1_v2c_user.thunder_snmp_server_snmpv1_v2c_user_test", "user", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_snmpv1_v2c_user.thunder_snmp_server_snmpv1_v2c_user_test", "user_tag", "a"),
				),
			},
		},
	})
}
