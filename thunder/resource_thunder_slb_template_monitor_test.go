package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_MONITOR_RESOURCE = `
resource "thunder_slb_template_monitor" "monitor" {
	user_tag = "test_tag"
	link_disable_cfg {
		dis_sequence = 1 
		diseth = 1
		}
	monitor_relation = "monitor-and"
	id2 = 1
}
`

// Acceptance test
func TestAccSlbTemplateMonitor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_MONITOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_monitor.monitor", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_monitor.monitor", "link_disable_cfg.0.dis_sequence", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_monitor.monitor", "link_disable_cfg.0.diseth", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_monitor.monitor", "monitor_relation", "monitor-and"),
					resource.TestCheckResourceAttr("thunder_slb_template_monitor.monitor", "id", "1"),
				),
			},
		},
	})
}
