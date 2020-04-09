package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SMPP_RESOURCE = `
resource "vthunder_slb_template_smpp" "smpp"{
name = "smpp2"
server_enquire_link = 1
server_selection_per_request = 1
client_enquire_link = 1
user_tag = "utag1"
server_enquire_link_val = 6
user = "u1"
password = "pwd1"
}
`

//Acceptance test
func TestAccAccVthunderSmpp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SMPP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "name", "smpp2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "server_enquire_link", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "server_selection_per_request", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "client_enquire_link", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "user_tag", "utag1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "server_enquire_link_val", "6"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "user", "u1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_smpp.smpp", "password", "pwd1"),
				),
			},
		},
	})
}
