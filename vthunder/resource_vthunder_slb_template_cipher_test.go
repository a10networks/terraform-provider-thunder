package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_CIPHER_RESOURCE = `
resource "vthunder_slb_template_cipher" "testname" {
	cipher_cfg {
		priority = 10
		cipher_suite = "SSL3_RSA_DES_192_CBC3_SHA"
	} 
	name = "testcipher"
	user_tag = "test_user"
}
`

//Acceptance test
func TestTemplateCipher_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_CIPHER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_cipher.testname", "name", "testcipher"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cipher.testname", "cipher_cfg.0.priority", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cipher.testname", "cipher_cfg.0.cipher_suite", "SSL3_RSA_DES_192_CBC3_SHA"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cipher.testname", "user_tag", "test_user"),
				),
			},
		},
	})
}
