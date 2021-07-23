package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_SIP_RESOURCE = `
resource "thunder_slb_template_sip" "sip2" {
	name = "testsip"
	user_tag = "test_tag"
	smp_call_id_rtp_session = 1
	client_keep_alive = 1
	alg_source_nat = 1
	alg_dest_nat = 1
	server_keep_alive = 1
	interval = 11
	dialog_aware = 1
	failed_server_selection = 1
	timeout = 1
	drop_when_server_fail = 1
}
`

//Acceptance test
func TestAccSlbTemplateSIP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SIP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "name", "testsip"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "smp_call_id_rtp_session", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "client_keep_alive", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "alg_source_nat", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "alg_dest_nat", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "server_keep_alive", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "interval", "11"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "dialog_aware", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "failed_server_selection", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "timeout", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_sip.sip2", "drop_when_server_fail", "1"),
				),
			},
		},
	})
}
