package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_CSV_RESOURCE = `
resource "thunder_slb_template_csv" "template_csv1" {
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

// Acceptance test
func TestAccSlbTemplateCSV_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_CSV_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_csv.template_csv1", "csv_name", "testcsv"),
					resource.TestCheckResourceAttr("thunder_slb_template_csv.template_csv1", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_csv.template_csv1", "ipv6_enable", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_csv.template_csv1", "delim_num", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_csv.template_csv1", "multiple_fields.0.field", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_csv.template_csv1", "multiple_fields.0.csv_type", "ip-from"),
				),
			},
		},
	})
}
