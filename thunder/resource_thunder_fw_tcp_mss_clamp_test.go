package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_TCP_MSS_CLAMP_RESOURCE = `
resource "thunder_fw_tcp_mss_clamp" "FwTcpTest" {
	mss_clamp_type = "fixed"
	mss_value = "0" 
}

`

//Acceptance test
func TestAccFwTcpMssClamp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TCP_MSS_CLAMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr("thunder_fw_tcp_mss_clamp.FwTcpTest", "mss_clamp_type", "fixed"),
					resource.TestCheckResourceAttr("thunder_fw_tcp_mss_clamp.FwTcpTest", "mss_value", "0"),
				),
			},
		},
	})
}
