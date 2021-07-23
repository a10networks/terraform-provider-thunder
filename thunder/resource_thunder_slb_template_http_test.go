package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_HTTP_RESOURCE = `
resource "thunder_slb_template_http" "http" {
	name = "testhttp"
	user_tag = "test_tag"
	keep_client_alive = 1
	req_hdr_wait_time = 1
}
`

//Acceptance test
func TestAccSlbTemplateHTTP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_HTTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_http.http", "name", "testhttp"),
					resource.TestCheckResourceAttr("thunder_slb_template_http.http", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_http.http", "keep_client_alive", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_http.http", "req_hdr_wait_time", "1"),
				),
			},
		},
	})
}
