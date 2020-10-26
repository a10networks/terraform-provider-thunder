package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_MQTT_RESOURCE = `
resource "thunder_slb_template_mqtt" "mqtt"{
user_tag="tag2"
clientid_hash_persist=0
name="mqtt2"
}
`

//Acceptance test
func TestAccThunderMqtt_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_MQTT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_mqtt.mqtt", "name", "mqtt2"),
					resource.TestCheckResourceAttr("thunder_slb_template_mqtt.mqtt", "user_tag", "tag2"),
					resource.TestCheckResourceAttr("thunder_slb_template_mqtt.mqtt", "clientid_hash_persist", "0"),
				),
			},
		},
	})
}
