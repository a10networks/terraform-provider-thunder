package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_DIAMETER_RESOURCE = `
resource "thunder_slb_template_diameter" "template_diameter" {
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
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "name", "testdiameter"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "terminate_on_cca_t", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "idle_timeout", "10"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "customize_cea", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "product_name", "yyeuey"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "dwr_up_retry", "4"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "forward_unknown_session_id", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "vendor_id", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "session_age", "50"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "load_balance_on_session_id", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "dwr_time", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "origin_realm", "tttest"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "multiple_origin_host", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "forward_to_latest_server", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "message_code_list.0.message_code", "100"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "avp_list.0.int32", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "avp_list.0.mandatory", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "avp_list.0.avp", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_diameter.template_diameter", "origin_host.0.origin_host_name", "testtt"),
				),
			},
		},
	})
}
