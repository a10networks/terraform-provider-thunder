package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_URPF_RESOURCE = `
resource "thunder_fw_urpf" "FwTest" {
	status = "strict" 
}
`

//Acceptance test
func TestAccFwUrpf_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_URPF_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_urpf.FwTest", "status", "strict"),
				),
			},
		},
	})
}
