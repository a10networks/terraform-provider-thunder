package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_MANAGEMENT_INDEX_RESOURCE = `
resource "thunder_snmp_server_management_index" "SnmpServerTest" {
	mgmt_index = 1
 
}
`

//Acceptance test
func TestAccSnmpServerManagementIndex_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_MANAGEMENT_INDEX_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_management_index.SnmpServerTest", "mgmt_index", "1"),
				),
			},
		},
	})
}
