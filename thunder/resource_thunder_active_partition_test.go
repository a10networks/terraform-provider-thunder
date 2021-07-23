package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ACTIVE_PARTITION_RESOURCE = `
resource "thunder_active_partition" "Test" {
	curr_part_name = "new_part"
 
}
`

//Acceptance test
func TestAccActivePartition_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ACTIVE_PARTITION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_active_partition.Test", "curr_part_name", "new_part"),
				),
			},
		},
	})
}
