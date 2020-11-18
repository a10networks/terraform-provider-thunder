package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_SNMPV1_V2C_USER_OID_RESOURCE = `
resource "thunder_snmp_server_snmpv1_v2c_user_oid" "SnmpServerSNMPv1V2cUserTest" {
	oid_val = "a"
user_tag = "a"
 
}
`

//Acceptance test
func TestAccSnmpServerSNMPv1V2cUserOid_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_SNMPV1_V2C_USER_OID_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_snmpv1_v2c_user_oid.SnmpServerSNMPv1V2cUserTest", "oid_val", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_snmpv1_v2c_user_oid.SnmpServerSNMPv1V2cUserTest", "user_tag", "a"),
				),
			},
		},
	})
}
