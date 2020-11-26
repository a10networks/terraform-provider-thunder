package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_VIEW_RESOURCE = `
resource "thunder_snmp_server_view" "SnmpServerTest" {
	oid = "a"
viewname = "a"
type = "included"
 
}
`

//Acceptance test
func TestAccSnmpServerView_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_VIEW_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_view.SnmpServerTest", "oid", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_view.SnmpServerTest", "viewname", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_view.SnmpServerTest", "type", "included"),
				),
			},
		},
	})
}
