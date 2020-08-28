package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_LOCAL_LOG_RESOURCE = `
{'resource': {'vthunder_fw_local_log': {'local_log': {'local_logging': 1}}}}
`

//Acceptance test
func TestAccFwLocalLog_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_LOCAL_LOG_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_local_log.FwTest", "local_log", "{'local-logging': 1}"),
				),
			},
		},
	})
}
