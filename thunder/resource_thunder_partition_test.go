package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_PARTITION_RESOURCE = `
resource "thunder_partition" "Test" {
	partition_name = "new_part"
	id1 = 20
	application_type = "adc"
 }
`

//Acceptance test
func TestAccPartition_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_PARTITION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_partition.Test", "partition_name", "new_part"),
					resource.TestCheckResourceAttr("thunder_partition.Test", "id1", "20"),
					resource.TestCheckResourceAttr("thunder_partition.Test", "application_type", "adc"),
				),
			},
		},
	})
}
