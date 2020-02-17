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
			"vthunder_vrrp_vrid": resourceVrrpVrid(),
			"vthunder_glm": resourceGlm(),
			"vthunder_import": resourceImport(),
			"vthunder_vrrp_common": resourceVrrpCommon(),
			"vthunder_vrrp_peer_group": resourceVrrpPeerGroup(),
			"vthunder_vrrp_session_sync": resourceVrrpSessionSync(),
			"vthunder_dns_primary": resourceDnsPrimary(),
			"vthunder_harmony_controller_profile":resourceHarmonyControllerProfile(),
			"vthunder_overlay_tunnel_options": resourceOverlayTunnelOptions(),
			"vthunder_overlay_tunnel_vtep": resourceOverlayTunnelVtep(),
			"vthunder_partition":resourceOverlayTunnelPartition(),
			"vthunder_slb_template_tcp":resourceSlbTemplateTcp(),			
		},

		ConfigureFunc: providerConfigureTest,
	}
}

func providerConfigureTest(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:  os.Getenv("VTHUNDER_HOST"),     //hostname
		Username: os.Getenv("VTHUNDER_USER"),     //username
		Password: os.Getenv("VTHUNDER_PASSWORD"), //password
	}

	return config.Client()
}

func TestAccProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAcctPreCheck(t *testing.T) {
	if v := os.Getenv("VTHUNDER_HOST"); v == "" {
		t.Fatal("ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("VTHUNDER_USER"); v == "" {
		t.Fatal("USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("VTHUNDER_PASSWORD"); v == "" {
		t.Fatal("PASSWORD must be set for acceptance tests")
	}
}
