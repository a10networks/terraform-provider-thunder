package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_COMMON_CONN_RATE_LIMIT_SRC_IP_RESOURCE = `
resource "vthunder_slb_common_conn_rate_limit_src_ip" "testname" {
	protocol = "tcp"
	limit_period = 1000
	limit = 50
	exceed_action = 0
	shared = 0
}
`

//Acceptance test
func TestAccSlbCommonConnRateLimitSrcIP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_COMMON_CONN_RATE_LIMIT_SRC_IP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_common_conn_rate_limit_src_ip.testname", "protocol", "tcp"),
					resource.TestCheckResourceAttr("vthunder_slb_common_conn_rate_limit_src_ip.testname", "limit_period", "1000"),
					resource.TestCheckResourceAttr("vthunder_slb_common_conn_rate_limit_src_ip.testname", "limit", "50"),
					resource.TestCheckResourceAttr("vthunder_slb_common_conn_rate_limit_src_ip.testname", "exceed_action", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_common_conn_rate_limit_src_ip.testname", "shared", "0"),
				),
			},
		},
	})
}
