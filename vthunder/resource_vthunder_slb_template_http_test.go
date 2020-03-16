package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_HTTP_RESOURCE = `
resource "vthunder_slb_template_http" "testname" {
	name = "testhttp"
	user_tag = "test_tag"
	keep_client_alive = 1
	req_hdr_wait_time = 1
}
`

//Acceptance test
func TestSlbTemplateHTTP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_HTTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_http.testname", "name", "testhttp"),
					resource.TestCheckResourceAttr("vthunder_slb_template_http.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_http.testname", "keep_client_alive", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_http.testname", "req_hdr_wait_time", "1"),
				),
			},
		},
	})
}
