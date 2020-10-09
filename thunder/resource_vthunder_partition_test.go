package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_PARTITION_RESOURCE = `
resource "thunder_partition" "partition"{
user_tag="tag1"
partition_name="part8"
application_type="adc"
id2=8
}
`

//Acceptance test
func TestAccThunderPartition_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_PARTITION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_partition.partition", "user_tag", "tag1"),
					resource.TestCheckResourceAttr("thunder_partition.partition", "partition_name", "part8"),
					resource.TestCheckResourceAttr("thunder_partition.partition", "application_type", "adc"),
					resource.TestCheckResourceAttr("thunder_partition.partition", "id2", "8"),
				),
			},
		},
	})
}
