package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_IP_NAT_ALG_PPTP_RESOURCE = `
resource "thunder_ip_nat_alg_pptp" "pptp" {
	pptp = "disable"
    sampling_enable {
        counters1 = "all"
      }
}
`

//Acceptance test
func TestAccIpNatAlgPptp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_NAT_ALG_PPTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ip_nat_alg_pptp.pptp", "pptp", "disable"),
					resource.TestCheckResourceAttr("thunder_ip_nat_alg_pptp.pptp", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
