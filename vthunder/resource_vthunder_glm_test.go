package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_GLM_RESOURCE = `
	resource "vthunder_glm" "glm1" {
		token = "vTh7aa90d73d"
		use_mgmt_port = 1
	    enable_requests = 1
	    allocate_bandwidth = 1000
}
`

//Acceptance test
func TestAccVthunderGlm_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GLM_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_glm.glm1", "token", "vTh7aa90d73d"),
					resource.TestCheckResourceAttr("vthunder_glm.glm1", "use_mgmt_port", "1"),
					resource.TestCheckResourceAttr("vthunder_glm.glm1", "enable_requests", "1"),
					resource.TestCheckResourceAttr("vthunder_glm.glm1", "allocate_bandwidth", "1000"),
				),
			},
		},
	})
}
