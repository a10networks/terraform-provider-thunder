package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_CSV_RESOURCE = `
resource "vthunder_slb_template_csv" "testname" {
	csv_name = "testcsv"
	user_tag = "test_tag"
	ipv6_enable = 0
	delim_num = 0
	multiple_fields {
		field = 5 
		csv_type = "ip-from"
	}
}
`

//Acceptance test
func TestSlbTemplateCSV_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_CSV_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_csv.testname", "csv_name", "testcsv"),
					resource.TestCheckResourceAttr("vthunder_slb_template_csv.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_csv.testname", "ipv6_enable", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_csv.testname", "delim_num", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_csv.testname", "multiple_fields.0.field", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_csv.testname", "multiple_fields.0.csv_type", "ip-from"),
				),
			},
		},
	})
}
