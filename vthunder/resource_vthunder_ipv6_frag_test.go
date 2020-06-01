package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IPV6_FRAG_RESOURCE = `
resource "vthunder_ipv6_frag" "testname" {

	sampling_enable {
	  counters1 = "all"
	  }
	frag_timeout = 10 
}
`

//Acceptance test
func TestAccIpv6Frag_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IPV6_FRAG_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ipv6_frag.testname", "sampling_enable.0.counters1", "all"),
					resource.TestCheckResourceAttr("vthunder_ipv6_frag.testname", "frag_timeout", "10"),
				),
			},
		},
	})
}
