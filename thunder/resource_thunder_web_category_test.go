package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_WEB_CATEGORY_RESOURCE = `
resource "thunder_web_category" "Test" {
	server = "a"
 
}
`

//Acceptance test
func TestAccWebCategory_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_WEB_CATEGORY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_web_category.Test", "server", "a"),
				),
			},
		},
	})
}
