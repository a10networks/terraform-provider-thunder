package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SNMP_SERVER_COMMUNITY_READ_RESOURCE = `
resource "thunder_snmp_server_community_read" "SnmpServerCommunityTest" {
	user = "a"
user_tag = "a"
 
}
`

//Acceptance test
func TestAccSnmpServerCommunityRead_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SNMP_SERVER_COMMUNITY_READ_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_snmp_server_community_read.SnmpServerCommunityTest", "user", "a"),
					resource.TestCheckResourceAttr("thunder_snmp_server_community_read.SnmpServerCommunityTest", "user_tag", "a"),
				),
			},
		},
	})
}
