package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)


var TEST_PARTITION_RESOURCE = `
resource "vthunder_partition" "partition"{
user_tag="tag1"
partition_name="part8"
application_type="adc"
id2=8
}
`

//Acceptance test
func TestAccVthunderPartition_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_PARTITION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_partition.partition", "user_tag", "tag1"),
					resource.TestCheckResourceAttr("vthunder_partition.partition", "partition_name", "part8"),
					resource.TestCheckResourceAttr("vthunder_partition.partition", "application_type", "adc"),
					resource.TestCheckResourceAttr("vthunder_partition.partition", "id2", "8"),
				),
			},
		},
	})
}
