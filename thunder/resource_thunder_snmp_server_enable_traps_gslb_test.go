package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_ENABLE_TRAPS_GSLB_RESOURCE = `
resource "thunder_snmp_server_enable_traps_gslb" "SnmpServerEnableTrapsTest" {
	all = 1
 
}
`

//Acceptance test
func TestAccSnmpServerEnableTrapsGslb_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_ENABLE_TRAPS_GSLB_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_enable_traps_gslb.SnmpServerEnableTrapsTest", "all", "1"),
				),
			},
		},
	})
}
