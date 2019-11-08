package vthunder

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TEST_PARTITION = "Common"

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var testProviders = map[string]terraform.ResourceProvider{
	"vthunder": Provider(),
}

func init() {

	testAccProvider = ProviderTest().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"vthunder": testAccProvider,
	}

}

func ProviderTest() terraform.ResourceProvider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"vthunder_virtual_server": resourceVirtualServer(),
			"vthunder_service_group":  resourceServiceGroup(),
			"vthunder_server":         resourceServer(),
			"vthunder_ethernet":       resourceEthernet(),
			"vthunder_rib_route":      resourceRibRoute(),
		},

		ConfigureFunc: providerConfigureTest,
	}
}

func providerConfigureTest(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:  os.Getenv("ADDRESS"),  //hostname
		Username: os.Getenv("USERNAME"), //username
		Password: os.Getenv("PASSWORD"), //password
	}

	return config.Client()
}

func TestAccProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAcctPreCheck(t *testing.T) {
	if v := os.Getenv("ADDRESS"); v == "" {
		t.Fatal("ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("USERNAME"); v == "" {
		t.Fatal("USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("PASSWORD"); v == "" {
		t.Fatal("PASSWORD must be set for acceptance tests")
	}
}
