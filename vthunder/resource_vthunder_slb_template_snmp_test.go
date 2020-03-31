package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SNMP_RESOURCE = `
resource "vthunder_slb_template_snmp" "testname" {
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
func TestSlbTemplateSNMP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SNMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "priv_proto", "aes"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "context_name", "testcont"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "interval", 10),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "security_level", "no-auth"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "community", "tttest"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "auth_proto", "sha"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "version", "v1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "interface", 0),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "port", 1770),
					resource.TestCheckResourceAttr("vthunder_slb_template_snmp.testname", "snmp_name", "testsnmp"),
				),
			},
		},
	})
}
