package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_SERVICE_GROUP_RESOURCE = `
resource "thunder_fw_service_group" "FwTest" {
	protocol = "tcp"
	name = "a"
	user_tag = "a" 
}
`

// Acceptance test
func TestAccFwServiceGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_SERVICE_GROUP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr("thunder_fw_service_group.FwTest", "protocol", "tcp"),
					resource.TestCheckResourceAttr("thunder_fw_service_group.FwTest", "name", "a"),
					resource.TestCheckResourceAttr("thunder_fw_service_group.FwTest", "user_tag", "a"),
				),
			},
		},
	})
}
