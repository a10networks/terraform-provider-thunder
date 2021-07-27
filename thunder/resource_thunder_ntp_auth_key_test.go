package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_NTP_AUTH_KEY_RESOURCE = `
resource "thunder_ntp_auth_key" "NtpTest" {

key_type = "ascii"
alg_type = "M"
key = 1
asc_key = "a"
 
}
`

//Acceptance test
func TestAccNtpAuthKey_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_NTP_AUTH_KEY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ntp_auth_key.NtpTest", "key_type", "ascii"),
					resource.TestCheckResourceAttr("thunder_ntp_auth_key.NtpTest", "alg_type", "M"),
					resource.TestCheckResourceAttr("thunder_ntp_auth_key.NtpTest", "key", "1"),
					resource.TestCheckResourceAttr("thunder_ntp_auth_key.NtpTest", "asc_key", "a"),
				),
			},
		},
	})
}
