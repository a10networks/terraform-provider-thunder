package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_NTP_SERVER_HOSTNAME_RESOURCE = `
resource "thunder_ntp_server_hostname" "NtpServerTest" {

host_servername = "a"
key = 1
 
}
`

//Acceptance test
func TestAccNtpServerHostname_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_NTP_SERVER_HOSTNAME_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ntp_server_hostname.NtpServerTest", "host_servername", "a"),
					resource.TestCheckResourceAttr("thunder_ntp_server_hostname.NtpServerTest", "key", "1"),
				),
			},
		},
	})
}
