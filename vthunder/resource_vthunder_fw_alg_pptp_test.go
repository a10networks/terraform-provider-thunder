package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_ALG_PPTP_RESOURCE = `
{'resource': {'vthunder_fw_alg_pptp': {'pptp': {'default_port_disable': 'default-port-disable'}}}}
`

//Acceptance test
func TestAccFwAlgPptp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_ALG_PPTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_alg_pptp.FwAlgTest", "pptp", "{'default-port-disable': 'default-port-disable'}"),
				),
			},
		},
	})
}
