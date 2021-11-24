package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_HEALTH_MONITOR_RESOURCE = `
resource "thunder_health_monitor" "HealthTest" {
	name = "hm-test"
 
}
`

//Acceptance test
func TestAccHealthMonitor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_HEALTH_MONITOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_health_monitor.HealthTest", "name", "hm-test"),
				),
			},
		},
	})
}
