package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_EXTERNAL_SERVICE_RESOURCE = `
resource "thunder_slb_template_external_service" "external_service" {
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
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "name", "testexternalservice"),
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "type", "skyfire-icap"),
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "action", "continue"),
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "failure_action", "continue"),
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "timeout", "10"),
					resource.TestCheckResourceAttr("thunder_slb_template_external_service.external_service", "request_header_forward_list.0.request_header_forward", "testf"),
				),
			},
		},
	})
}
