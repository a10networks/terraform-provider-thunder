package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_ACCESS_LIST_STANDARD_RESOURCE = `
resource "thunder_access_list_standard" "AccessListTest" {
	std = 1
stdrules {   
	seq_num =  1 
	action =  "deny" 
	any =  1 
	log =  1 
	transparent_session_only =  1 
	}
 
}
`

//Acceptance test
func TestAccAccessListStandard_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ACCESS_LIST_STANDARD_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_access_list_standard.AccessListTest", "std", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_standard.AccessListTest", "stdrules.0.seq_num", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_standard.AccessListTest", "stdrules.0.action", "deny"),
					resource.TestCheckResourceAttr("thunder_access_list_standard.AccessListTest", "stdrules.0.any", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_standard.AccessListTest", "stdrules.0.log", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_standard.AccessListTest", "stdrules.0.transparent_session_only", "1"),
				),
			},
		},
	})
}
