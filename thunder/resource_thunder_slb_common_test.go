package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_COMMON_RESOURCE = `
resource "thunder_slb_common" "common" {
	low_latency = 0
	use_mss_tab = 0
	stats_data_disable = 0
	compress_block_size = 7000
	player_id_check_enable = 0
	msl_time = 20
	graceful_shutdown_enable = 0
	hw_syn_rr = 2000 
	conn_rate_limit {
		src_ip_list {
			protocol = "tcp"
			limit_period = "100"
			limit = 50
			exceed_action = 0
			shared = 0
		}
	}
}
`

//Acceptance test
func TestAccSlbCommon_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_COMMON_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_common.common", "low_latency", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "use_mss_tab", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "stats_data_disable", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "compress_block_size", "7000"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "player_id_check_enable", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "msl_time", "20"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "graceful_shutdown_enable", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "hw_syn_rr", "2000"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "conn_rate_limit.0.src_ip_list.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "conn_rate_limit.0.src_ip_list.0.limit_period", "100"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "conn_rate_limit.0.src_ip_list.0.limit", "50"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "conn_rate_limit.0.src_ip_list.0.exceed_action", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common.common", "conn_rate_limit.0.src_ip_list.0.shared", "0"),
				),
			},
		},
	})
}
