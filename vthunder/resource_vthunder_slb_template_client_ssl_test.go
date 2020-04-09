package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_CLIENT_SSL_RESOURCE = `
resource "vthunder_slb_template_client_ssl" "testname" {
	name = "testclientssl"
	user_tag = "test_tag"
	forward_proxy_ssl_version = 33
	forward_proxy_crl_disable = 1
	fp_cert_ext_aia_ocsp = "ocp" 
}
`

//Acceptance test
func TestAccSlbTemplateClientSSL_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_CLIENT_SSL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_client_ssl.testname", "name", "testclientssl"),
					resource.TestCheckResourceAttr("vthunder_slb_template_client_ssl.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_client_ssl.testname", "forward_proxy_ssl_version", "33"),
					resource.TestCheckResourceAttr("vthunder_slb_template_client_ssl.testname", "forward_proxy_crl_disable", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_client_ssl.testname", "fp_cert_ext_aia_ocsp", "ocp"),
				),
			},
		},
	})
}
