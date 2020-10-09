package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_SSL_EXPIRE_CHECK_RESOURCE = `
resource "thunder_slb_ssl_expire_check" "SSLExpireCheck" {
	before = 5
	expire_address1 = "abc@gmail.com"
	interval_days = 2
	ssl_expire_email_address = "abc@gmail.com"

}
`

//Acceptance test
func TestAccSlbSSLExpireCheck_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SSL_EXPIRE_CHECK_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_ssl_expire_check.SSLExpireCheck", "before", "5"),
					resource.TestCheckResourceAttr("thunder_slb_ssl_expire_check.SSLExpireCheck", "expire_address1", "abc@gmail.com"),
					resource.TestCheckResourceAttr("thunder_slb_ssl_expire_check.SSLExpireCheck", "interval_days", "2"),
					resource.TestCheckResourceAttr("thunder_slb_ssl_expire_check.SSLExpireCheck", "ssl_expire_email_address", "abc@gmail.com"),
				),
			},
		},
	})
}
