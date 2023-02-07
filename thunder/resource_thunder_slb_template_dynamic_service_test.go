package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_DYNAMIC_SERVICE_RESOURCE = `
resource "thunder_slb_template_dynamic_service" "dynamic_service" {
	name = "testdynamicservice"
	user_tag = "test_tag"
	dns_server {
		ipv4_dns_server = "10.0.0.10"
	}
}
`

//Acceptance test
func TestAccSlbTemplateDynamicService_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_DYNAMIC_SERVICE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_dynamic_service.dynamic_service", "name", "testdynamicservice"),
					resource.TestCheckResourceAttr("thunder_slb_template_dynamic_service.dynamic_service", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_dynamic_service.dynamic_service", "dns_server.0.ipv4_dns_server", "10.0.0.10"),
				),
			},
		},
	})
}
