package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SNMP_RESOURCE = `
resource "thunder_slb_template_snmp" "snmp" {
	user_tag = "test_tag"
	priv_proto = "aes"
	context_name = "testcont"
	interval = 10
	security_level = "no-auth"
	community = "tttest"
	auth_proto = "sha"
	version = "v1"
	interface = 0
	port = 1770
	snmp_name = "testsnmp"

}
`

//Acceptance test
func TestAccSlbTemplateSNMP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SNMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "priv_proto", "aes"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "context_name", "testcont"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "interval", "10"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "security_level", "no-auth"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "community", "tttest"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "auth_proto", "sha"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "version", "v1"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "interface", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "port", "1770"),
					resource.TestCheckResourceAttr("thunder_slb_template_snmp.snmp", "snmp_name", "testsnmp"),
				),
			},
		},
	})
}
