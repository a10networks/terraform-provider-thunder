package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_ENABLE_TRAPS_VCS_RESOURCE = `
resource "thunder_snmp_server_enable_traps_vcs" "SnmpServerEnableTrapsTest" {
	state_change = 1
 
}
`

//Acceptance test
func TestAccSnmpServerEnableTrapsVcs_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_ENABLE_TRAPS_VCS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_enable_traps_vcs.SnmpServerEnableTrapsTest", "state_change", "1"),
				),
			},
		},
	})
}
