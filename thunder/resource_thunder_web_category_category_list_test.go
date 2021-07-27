package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_WEB_CATEGORY_CATEGORY_LIST_RESOURCE = `
resource "thunder_web_category_category_list" "WebCategoryTest" {
	name = "a"
user_tag = "a"
 
}
`

//Acceptance test
func TestAccWebCategoryCategoryList_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_WEB_CATEGORY_CATEGORY_LIST_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_web_category_category_list.WebCategoryTest", "name", "a"),
					resource.TestCheckResourceAttr("thunder_web_category_category_list.WebCategoryTest", "user_tag", "a"),
				),
			},
		},
	})
}
