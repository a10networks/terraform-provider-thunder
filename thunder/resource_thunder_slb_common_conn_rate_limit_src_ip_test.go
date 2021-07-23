package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_COMMON_CONN_RATE_LIMIT_SRC_IP_RESOURCE = `
resource "thunder_slb_common_conn_rate_limit_src_ip" "src_ip" {
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
					resource.TestCheckResourceAttr("thunder_slb_common_conn_rate_limit_src_ip.src_ip", "protocol", "tcp"),
					resource.TestCheckResourceAttr("thunder_slb_common_conn_rate_limit_src_ip.src_ip", "limit_period", "1000"),
					resource.TestCheckResourceAttr("thunder_slb_common_conn_rate_limit_src_ip.src_ip", "limit", "50"),
					resource.TestCheckResourceAttr("thunder_slb_common_conn_rate_limit_src_ip.src_ip", "exceed_action", "0"),
					resource.TestCheckResourceAttr("thunder_slb_common_conn_rate_limit_src_ip.src_ip", "shared", "0"),
				),
			},
		},
	})
}
