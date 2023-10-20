package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ACCESS_LIST_EXTENDED_RESOURCE = `
resource "thunder_access_list_extended" "AccessListTest" {
	extd = 100
rules {   
	extd_seq_num =  1 
	extd_action =  "deny" 
	tcp =  1 
	src_any =  1 
	src_eq =  1 
	dst_any =  1 
	dst_eq =  1 
	ethernet =  1 
	acl_log =  1 
	transparent_session_only =  1 
	}
 
}
`

// Acceptance test
func TestAccAccessListExtended_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ACCESS_LIST_EXTENDED_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "extd", "100"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.extd_seq_num", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.extd_action", "deny"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.tcp", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.src_any", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.src_eq", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.dst_any", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.dst_eq", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.ethernet", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.acl_log", "1"),
					resource.TestCheckResourceAttr("thunder_access_list_extended.AccessListTest", "rules.0.transparent_session_only", "1"),
				),
			},
		},
	})
}
