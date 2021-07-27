package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SYSTEM_RESOURCE = `
resource "thunder_system" "Test" {
	anomaly_log = 1
	promiscuous_mode = 0
 
}
`

//Acceptance test
func TestAccSystem_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SYSTEM_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_system.Test", "anomaly_log", "1"),
					resource.TestCheckResourceAttr("thunder_system.Test", "promiscuous_mode", "0"),
				),
			},
		},
	})
}
