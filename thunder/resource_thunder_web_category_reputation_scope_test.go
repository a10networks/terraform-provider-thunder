package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_WEB_CATEGORY_REPUTATION_SCOPE_RESOURCE = `
resource "thunder_web_category_reputation_scope" "WebCategoryTest" {
	name = "a"
user_tag = "a"
 
}
`

//Acceptance test
func TestAccWebCategoryReputationScope_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_WEB_CATEGORY_REPUTATION_SCOPE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_web_category_reputation_scope.WebCategoryTest", "name", "a"),
					resource.TestCheckResourceAttr("thunder_web_category_reputation_scope.WebCategoryTest", "user_tag", "a"),
				),
			},
		},
	})
}
