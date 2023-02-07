package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_RC_CACHE_GLOBAL_RESOURCE = `
resource "thunder_slb_rc_cache_global" "rc_cache1" {

	sampling_enable  {
	    counters1 = "all"
	 }
}
`

//Acceptance test
func TestAccSlbRcCacheGlobal_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_RC_CACHE_GLOBAL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_rc_cache_global.rc_cache1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
