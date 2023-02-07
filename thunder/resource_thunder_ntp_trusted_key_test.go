package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_NTP_TRUSTED_KEY_RESOURCE = `
resource "thunder_ntp_trusted_key" "NtpTest" {
	
key = 1
 
}
`

//Acceptance test
func TestAccNtpTrustedKey_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_NTP_TRUSTED_KEY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ntp_trusted_key.NtpTest", "key", "1"),
				),
			},
		},
	})
}
