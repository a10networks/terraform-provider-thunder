package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_RESPMOD_ICAP_RESOURCE = `
resource "thunder_slb_template_respmod_icap" "respmod_icap" {
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

// Acceptance test
func TestAccSlbTemplateRespmodIcap_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_RESPMOD_ICAP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "name", "testrespmodicap"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "min_payload_size", "30"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "preview", "200"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "disable_http_server_reset", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "fail_close", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "include_protocol_in_uri", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "log_only_allowed_method", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_respmod_icap.respmod_icap", "cylance", "0"),
				),
			},
		},
	})
}
