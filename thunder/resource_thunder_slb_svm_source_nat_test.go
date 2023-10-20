package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_SVM_SOURCE_NAT_RESOURCE = `
resource "thunder_slb_svm_source_nat" "svm_source" {
	pool = "test" 
}
`

// Acceptance test
func TestAccSlbSvmSourceNat_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SVM_SOURCE_NAT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_svm_source_nat.svm_source", "pool", "test"),
				),
			},
		},
	})
}
