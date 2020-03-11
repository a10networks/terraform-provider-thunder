package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_REQMOD_ICAP_RESOURCE = `
resource "vthunder_slb_template_reqmod_icap" "testname" {
	name = "testreqmodicap"
	user_tag = "test_tag"
	min_payload_size = 0
	allowed_http_methods = "POST"
	service_url = "10.0.0.11"
	preview = 200
	disable_http_server_reset = 0
	fail_close = 0
	include_protocol_in_uri = 0
	log_only_allowed_method = 0
	cylance = 0
	 
}
`

//Acceptance test
func TestSlbTemplateReqmodIcap_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_REQMOD_ICAP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "name", "testreqmodicap"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "min_payload_size", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "allowed_http_methods", "POST"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "service_url", "10.0.0.11"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "preview", "200"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "disable_http_server_reset", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "fail_close", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "include_protocol_in_uri", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "log_only_allowed_method", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_reqmod_icap.testname", "cylance", "0"),
				),
			},
		},
	})
}
