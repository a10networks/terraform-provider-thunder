package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FILE_SSL_CERT_RESOURCE = `
resource "thunder_file_ssl_cert" "FileTest" {
	file = "pkiCertName"      
	action = "import"
	certificate_type = "pem"
	file_handle = "pkiCertName"
	file_local_path = "/home/uername/pkiCertName" 
}
`

//Acceptance test
func TestAccFileSslCert_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FILE_SSL_CERT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_file_ssl_cert.FileTest", "file", "pkiCertName"),
					resource.TestCheckResourceAttr("thunder_file_ssl_cert.FileTest", "action", "import"),
					resource.TestCheckResourceAttr("thunder_file_ssl_cert.FileTest", "certificate_type", "pem"),
					resource.TestCheckResourceAttr("thunder_file_ssl_cert.FileTest", "file_handle", "pkiCertName"),
					resource.TestCheckResourceAttr("thunder_file_ssl_cert.FileTest", "file_local_path", "/home/uername/pkiCertName"),
				),
			},
		},
	})
}
