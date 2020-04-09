package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_DIAMETER_RESOURCE = `
resource "vthunder_slb_template_diameter" "testname" {
	name = "testdiameter"
	user_tag = "test_tag"
	terminate_on_cca_t = 0
	message_code_list {
		message_code = 100
	}
	avp_list {
		int32 = 0
		mandatory = 0
		avp = 0
	}
	idle_timeout = 10
	customize_cea = 0
	product_name = "yyeuey"
	dwr_up_retry = 4
	forward_unknown_session_id = 0
	vendor_id = 0
	session_age = 50
	load_balance_on_session_id = 0
	dwr_time = 0
	origin_realm = "tttest"
	origin_host {
		origin_host_name = "testtt"
	}
	multiple_origin_host = 0
	forward_to_latest_server = 0
}
`

//Acceptance test
func TestAccSlbTemplateDiameter_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_DIAMETER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "name", "testdiameter"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "terminate_on_cca_t", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "idle_timeout", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "customize_cea", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "product_name", "yyeuey"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "dwr_up_retry", "4"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "forward_unknown_session_id", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "vendor_id", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "session_age", "50"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "load_balance_on_session_id", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "dwr_time", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "origin_realm", "tttest"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "multiple_origin_host", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "forward_to_latest_server", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "message_code_list.0.message_code", "100"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "avp_list.0.int32", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "avp_list.0.mandatory", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "avp_list.0.avp", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_diameter.testname", "origin_host.0.origin_host_name", "testtt"),
				),
			},
		},
	})
}
