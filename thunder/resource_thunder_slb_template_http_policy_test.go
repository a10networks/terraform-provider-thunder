package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_HTTP_POLICY_RESOURCE = `
resource "thunder_slb_template_http_policy" "http_policy"{
name= "httppol1"
cookie_name="cookie-two"
user_tag="u1"
}
`

//Acceptance test
func TestAccVthuderHttpPolicy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_HTTP_POLICY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_http_policy.http_policy", "name", "httppol1"),
					resource.TestCheckResourceAttr("thunder_slb_template_http_policy.http_policy", "cookie_name", "cookie-two"),
					resource.TestCheckResourceAttr("thunder_slb_template_http_policy.http_policy", "user_tag", "u1"),
				),
			},
		},
	})
}
