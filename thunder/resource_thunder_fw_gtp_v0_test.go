package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_GTP_V0_RESOURCE = `
resource "thunder_fw_gtp_v0" "FwTest" {
	gtpv0_value = "enable" 
}
`

// Acceptance test
func TestAccFwGtpV0_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_GTP_V0_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_gtp_v0.FwTest", "gtpv0_value", "enable"),
				),
			},
		},
	})
}
