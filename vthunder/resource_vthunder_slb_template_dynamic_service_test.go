package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_DYNAMIC_SERVICE_RESOURCE = `
resource "vthunder_slb_template_dynamic_service" "testname" {
	name = "testdynamicservice"
	user_tag = "test_tag"
	dns_server {
		ipv4_dns_server = "10.0.0.10"
	}
}
`

//Acceptance test
func TestSlbTemplateDynamicService_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_DYNAMIC_SERVICE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_dynamic_service.testname", "name", "testdynamicservice"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dynamic_service.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dynamic_service.testname", "dns_server.0.ipv4_dns_server", "10.0.0.10"),
				),
			},
		},
	})
}
