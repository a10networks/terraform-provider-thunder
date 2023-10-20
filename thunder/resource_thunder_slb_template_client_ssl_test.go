package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_CLIENT_SSL_RESOURCE = `
resource "thunder_slb_template_client_ssl" "client_ssl1" {
	name = "testclientssl"
	user_tag = "test_tag"
	forward_proxy_ssl_version = 33
	forward_proxy_crl_disable = 1
	fp_cert_ext_aia_ocsp = "ocp" 
}
`

// Acceptance test
func TestAccSlbTemplateClientSSL_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_CLIENT_SSL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssl.client_ssl1", "name", "testclientssl"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssl.client_ssl1", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssl.client_ssl1", "forward_proxy_ssl_version", "33"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssl.client_ssl1", "forward_proxy_crl_disable", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_client_ssl.client_ssl1", "fp_cert_ext_aia_ocsp", "ocp"),
				),
			},
		},
	})
}
