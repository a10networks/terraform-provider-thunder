package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_PROFILE_RESOURCE = `
	resource "thunder_harmony_controller_profile" "profile" {
		host = "129.213.20.171"
		port=8443
		user_name = "scampbell@a10networks.com"
		secret_value = "admin123"
		provider2 = "root"
		action = "register"
		use_mgmt_port = 1
		region = "India"
		availability_zone = "Pune"
		thunder_mgmt_ip {
		ip_address = "129.213.20.171"
		}
		
}
`

//Acceptance test
func TestAccThunderProfile_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_PROFILE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "host", "129.213.20.171"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "port", "8443"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "user_name", "scampbell@a10networks.com"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "secret_value", "admin123"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "provider2", "root"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "action", "register"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "use_mgmt_port", "1"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "region", "India"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "availability_zone", "Pune"),
					resource.TestCheckResourceAttr("thunder_harmony_controller_profile.profile", "thunder_mgmt_ip.0.ip_address", "129.213.20.171"),
				),
			},
		},
	})
}
