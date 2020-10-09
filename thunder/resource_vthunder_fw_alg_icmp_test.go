package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_ALG_ICMP_RESOURCE = `
resource "thunder_fw_alg_icmp" "FwAlgTest" {
	disable = "disable" 
}

`

//Acceptance test
func TestAccFwAlgIcmp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_ALG_ICMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_alg_icmp.FwAlgTest", "disable", "disable"),
				),
			},
		},
	})
}
