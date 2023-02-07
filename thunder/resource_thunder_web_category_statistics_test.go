package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_WEB_CATEGORY_STATISTICS_RESOURCE = `
resource "thunder_web_category_statistics" "WebCategoryTest" {
	sampling-enable {   
	counters1 =  "all" 
	}
 
}
`

//Acceptance test
func TestAccWebCategoryStatistics_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_WEB_CATEGORY_STATISTICS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_web_category_statistics.WebCategoryTest", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
