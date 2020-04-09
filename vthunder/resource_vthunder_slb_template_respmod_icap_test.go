package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_RESPMOD_ICAP_RESOURCE = `
resource "vthunder_slb_template_respmod_icap" "testname" {
	name = "testrespmodicap"
	user_tag = "test_tag"
	min_payload_size = 30
	preview = 200
	disable_http_server_reset = 0
	fail_close = 0
	include_protocol_in_uri = 0
	log_only_allowed_method = 0
	cylance = 0

}
`

//Acceptance test
func TestAccSlbTemplateRespmodIcap_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_RESPMOD_ICAP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "name", "testrespmodicap"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "min_payload_size", "30"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "preview", "200"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "disable_http_server_reset", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "fail_close", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "include_protocol_in_uri", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "log_only_allowed_method", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_respmod_icap.testname", "cylance", "0"),
				),
			},
		},
	})
}
