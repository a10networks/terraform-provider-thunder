package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_PLAYER_ID_GLOBAL_RESOURCE = `
resource "vthunder_slb_player_id_global" "testname" {
	enforcement_timer = 10
	abs_max_expiration = 10
	sampling_enable  {
	    counters1 = "all"
	}
	force_passive = 1
	pkt_activity_expiration = 10
	min_expiration = 10
	enable_64bit_player_id = 0 
}
`

//Acceptance test
func TestAccSlbPlayerIdGlobal_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_PLAYER_ID_GLOBAL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "enforcement_timer", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "abs_max_expiration", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "sampling_enable.0.counters1", "all"),
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "force_passive", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "pkt_activity_expiration", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "min_expiration", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_player_id_global.testname", "enable_64bit_player_id", "0"),
				),
			},
		},
	})
}
