package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_IMPORT_RESOURCE = `
	resource "vthunder_import" "import1" {
		password = "arun123"
	    cloud_creds = "oci_api_key.pem"
	    use_mgmt_port = 1
	    overwrite = 1
	    remote_file = "sftp://ubuntu@13.52.180.11/home/ubuntu/oci_api_key.pem"  
}
`

//Acceptance test
func TestAccVthunderImport_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IMPORT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_import.import1", "password", "arun123"),
					resource.TestCheckResourceAttr("vthunder_import.import1", "cloud_creds", "oci_api_key.pem"),
					resource.TestCheckResourceAttr("vthunder_import.import1", "use_mgmt_port", "1"),
					resource.TestCheckResourceAttr("vthunder_import.import1", "overwrite", "1"),
					resource.TestCheckResourceAttr("vthunder_import.import1", "remote_file", "sftp://ubuntu@13.52.180.11/home/ubuntu/oci_api_key.pem"),
				),
			},
		},
	})
}
