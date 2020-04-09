package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SIP_RESOURCE = `
resource "vthunder_slb_template_sip" "testname" {
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
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "name", "testsip"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "smp_call_id_rtp_session", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "client_keep_alive", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "alg_source_nat", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "alg_dest_nat", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "server_keep_alive", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "interval", "11"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "dialog_aware", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "failed_server_selection", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "timeout", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_sip.testname", "drop_when_server_fail", "1"),
				),
			},
		},
	})
}
