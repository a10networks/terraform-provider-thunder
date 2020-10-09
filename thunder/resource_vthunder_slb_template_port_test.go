package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_PORT_RESOURCE = `
resource "thunder_slb_template_port" "port" {
	name = "testport"
	user_tag = "test_tag"
	slow_start = 1
	conn_limit = 2
	weight = 5
	extended_stats = 1
	del_session_on_server_down = 1
}
`

//Acceptance test
func TestAccSlbTemplatePort_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_PORT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "name", "testport"),
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "slow_start", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "conn_limit", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "weight", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "extended_stats", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_port.port", "del_session_on_server_down", "1"),
				),
			},
		},
	})
}
