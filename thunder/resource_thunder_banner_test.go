package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_BANNER_RESOURCE = `
resource "thunder_banner" "bannerTest" {
	exec_banner_cfg {  
 	exec =  1 
	exec_banner =  "Sample Exec Banner" 
	}
login_banner_cfg {  
 	login =  1 
	login_banner =  "Sample Login Banner" 
	}
 
}
`

//Acceptance test
func TestAccBanner_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_BANNER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_banner.bannerTest", "exec_banner_cfg.0.exec", "1"),
					resource.TestCheckResourceAttr("thunder_banner.bannerTest", "exec_banner_cfg.0.exec_banner", "Sample Exec Banner"),
					resource.TestCheckResourceAttr("thunder_banner.bannerTest", "login_banner_cfg.0.login", "1"),
					resource.TestCheckResourceAttr("thunder_banner.bannerTest", "login_banner_cfg.0.login_banner", "Sample Login Banner"),
				),
			},
		},
	})
}
