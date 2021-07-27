package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_GTP_IN_GTP_FILTERING_RESOURCE = `
resource "thunder_fw_gtp_in_gtp_filtering" "FwTest" {
	gtp_in_gtp_value = "disable" 
}
`

//Acceptance test
func TestAccFwGtpInGtpFiltering_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_GTP_IN_GTP_FILTERING_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_gtp_in_gtp_filtering.FwTest", "gtp_in_gtp_value", "disable"),
				),
			},
		},
	})
}
