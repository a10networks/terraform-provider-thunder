package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_TIMEZONE_RESOURCE = `
resource "thunder_timezone" "timezoneTest" {
	timezone_index_cfg {  
 	timezone_index =  "America/Phoenix" 
	nodst =  0 
	}
uuid = "b2dc8f28-bd4f-11eb-a0ec-91d97073b9cc"
a10_url = "/axapi/v3/timezone"
 
}
`

// Acceptance test
func TestAccTimezone_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TIMEZONE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_timezone.timezoneTest", "timezone_index_cfg.0.timezone_index", "America/Phoenix"),
					resource.TestCheckResourceAttr("thunder_timezone.timezoneTest", "timezone_index_cfg.0.nodst", "0"),
					resource.TestCheckResourceAttr("thunder_timezone.timezoneTest", "uuid", "b2dc8f28-bd4f-11eb-a0ec-91d97073b9cc"),
					resource.TestCheckResourceAttr("thunder_timezone.timezoneTest", "a10_url", "/axapi/v3/timezone"),
				),
			},
		},
	})
}
