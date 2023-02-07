package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_CONTACT_RESOURCE = `
resource "thunder_snmp_server_contact" "SnmpServerTest" {
	contact_name = "a"
 
}
`

//Acceptance test
func TestAccSnmpServerContact_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_CONTACT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_contact.SnmpServerTest", "contact_name", "a"),
				),
			},
		},
	})
}
