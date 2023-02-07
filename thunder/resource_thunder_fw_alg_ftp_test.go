package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_ALG_FTP_RESOURCE = `
resource "thunder_fw_alg_ftp" "FwAlgTest" {
	default_port_disable = "default-port-disable" 
}
`

//Acceptance test
func TestAccFwAlgFtp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_ALG_FTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_alg_ftp.FwAlgTest", "default_port_disable", "default-port-disable"),
				),
			},
		},
	})
}
