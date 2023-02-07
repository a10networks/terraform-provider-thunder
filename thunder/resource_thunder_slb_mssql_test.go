package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_MSSQL_RESOURCE = `
resource "thunder_slb_mssql" "mssql" {
	sampling_enable {
	    counters1 = "all"
	 }
}
`

//Acceptance test
func TestAccSlbMssql_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_MSSQL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_mssql.mssql", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
