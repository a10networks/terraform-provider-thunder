package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_EXTERNAL_SERVICE_RESOURCE = `
resource "vthunder_slb_template_external_service" "testname" {
	name = "testexternalservice"
	user_tag = "test_tag"
	type = "skyfire-icap"
	action = "continue"
	failure_action = "continue"
	timeout = 10
	request_header_forward_list {
        request_header_forward = "testf"
      }
}
`

//Acceptance test
func TestAccSlbTemplateExternalService_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_EXTERNAL_SERVICE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "name", "testexternalservice"),
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "type", "skyfire-icap"),
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "action", "continue"),
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "failure_action", "continue"),
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "timeout", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_template_external_service.testname", "request_header_forward_list.0.request_header_forward", "testf"),
				),
			},
		},
	})
}
