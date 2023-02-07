package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TRANSPARENT_TCP_TEMPLATE_RESOURCE = `
resource "thunder_slb_transparent_tcp_template" "tcp_template" {
	name = "testtransperenttcptemplate"
}
`

//Acceptance test
func TestAccSlbTransperentTcpTemplate_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TRANSPARENT_TCP_TEMPLATE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_transparent_tcp_template.tcp_template", "name", "testtransperenttcptemplate"),
				),
			},
		},
	})
}
