package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_COMMON_RESOURCE = `
resource "vthunder_slb_common" "testname" {
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
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "low_latency", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "use_mss_tab", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "stats_data_disable", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "compress_block_size", "7000"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "player_id_check_enable", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "msl_time", "20"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "graceful_shutdown_enable", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "hw_syn_rr", "2000"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "conn_rate_limit.0.src_ip_list.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "conn_rate_limit.0.src_ip_list.0.limit_period", "100"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "conn_rate_limit.0.src_ip_list.0.limit", "50"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "conn_rate_limit.0.src_ip_list.0.exceed_action", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common.testname", "conn_rate_limit.0.src_ip_list.0.shared", "0"),
				),
			},
		},
	})
}
