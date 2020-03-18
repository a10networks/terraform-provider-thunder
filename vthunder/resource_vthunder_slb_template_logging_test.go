package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_MQTT_RESOURCE = `
resource "vthunder_slb_template_logging" "l"{
name = "log1"
format= "abc"
local_logging= 1
}
`

//Acceptance test
func TestAccVthunderMqtt_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_MQTT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_logging.logging", "name", "log1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_logging.logging", "format", "abc"),
					resource.TestCheckResourceAttr("vthunder_slb_template_logging.logging", "local_logging", "1"),
				),
			},
		},
	})
}
