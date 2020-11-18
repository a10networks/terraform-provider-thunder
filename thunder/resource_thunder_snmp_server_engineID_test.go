package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SNMP_SERVER_ENGINEID_RESOURCE = `
resource "thunder_snmp_server_engine_id" "SnmpServerTest" {
	eng_id = "ABABABABAB"
 
}
`

//Acceptance test
func TestAccSnmpServerEngineID_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_ENGINEID_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_engine_id.SnmpServerTest", "eng_id", "ABABABABAB"),
				),
			},
		},
	})
}
