package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_GLM_RESOURCE = `
	resource "thunder_glm" "glm1" {
		token = "vTh7aa90d73d"
		use_mgmt_port = 1
	    enable_requests = 1
	    allocate_bandwidth = 1000
}
`

// Acceptance test
func TestAccThunderGlm_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GLM_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_glm.glm1", "token", "vTh7aa90d73d"),
					resource.TestCheckResourceAttr("thunder_glm.glm1", "use_mgmt_port", "1"),
					resource.TestCheckResourceAttr("thunder_glm.glm1", "enable_requests", "1"),
					resource.TestCheckResourceAttr("thunder_glm.glm1", "allocate_bandwidth", "1000"),
				),
			},
		},
	})
}
