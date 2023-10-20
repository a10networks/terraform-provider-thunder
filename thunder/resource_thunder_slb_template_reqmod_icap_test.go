package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_REQMOD_ICAP_RESOURCE = `
resource "thunder_slb_template_reqmod_icap" "reqmod_icap" {
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

// Acceptance test
func TestAccSlbTemplateReqmodIcap_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_REQMOD_ICAP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "name", "testreqmodicap"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "min_payload_size", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "allowed_http_methods", "POST"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "service_url", "10.0.0.11"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "preview", "200"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "disable_http_server_reset", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "fail_close", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "include_protocol_in_uri", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "log_only_allowed_method", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_reqmod_icap.reqmod_icap", "cylance", "0"),
				),
			},
		},
	})
}
