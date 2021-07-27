package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FTP_PROXY_RESOURCE = `
resource "thunder_slb_ftp_proxy" "ftp_proxy1" {
	sampling_enable {
		counters1 = "all"
	}
}

`

//Acceptance test
func TestAccThunderFTPProxy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FTP_PROXY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_ftp_proxy.ftp_proxy1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
