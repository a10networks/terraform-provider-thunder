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
			"vthunder_slb_template_udp": resourceSlbTemplateUdp(),
			"vthunder_slb_template_smpp":resourceSlbTemplateSmpp(),
			"vthunder_slb_template_fix":resourceTemplateFix(),
			"vthunder_slb_template_ftp":resourceTemplateFTP(),
			"vthunder_slb_template_mqtt":resourceSlbTemplateMqtt(),
			"vthunder_slb_template_http_policy":resourceSlbTemplateHttpPolicy(),
			"vthunder_slb_template_ssli":resourceTemplateSSLI(),
			"vthunder_slb_template_cipher":resourceTemplateCipher(),
			"vthunder_slb_template_imap_pop3":resourceTemplateImap_POP3(),
			"vthunder_slb_template_dns":resourceTemplateDNS(),
			"vthunder_slb_template_connection_reuse":resourceTemplateConnectionReuse(),
			"vthunder_slb_template_client_ssh":resourceTemplateClientSsh(),
			"vthunder_slb_template_dblb":resourceTemplateDBLB(),
			"vthunder_slb_template_port":resourceTemplatePort(),
			"vthunder_slb_template_smtp":resourceSlbTemplateSMTP(),
			"vthunder_slb_template_tcp_proxy":resourceSlbTemplateTcpProxy(),
			"vthunder_slb_template_server":resourceSlbTemplateServer(),
			"vthunder_slb_template_virtual_port":resourceSlbTemplateVirtualPort(),
			"vthunder_slb_template_diameter":         resourceSlbTemplateDiameter(),
			"vthunder_slb_template_cache":            resourceSlbTemplateCache(),
			"vthunder_slb_template_csv":              resourceSlbTemplateCSV(),
			"vthunder_slb_template_reqmod_icap":      resourceSlbTemplateReqmodIcap(),
			"vthunder_slb_template_respmod_icap":     resourceSlbTemplateRespmodIcap(),
			"vthunder_slb_template_dynamic_service":  resourceSlbTemplateDynamicService(),
			"vthunder_slb_template_server_ssh":       resourceSlbTemplateServerSSH(),
			"vthunder_slb_template_policy":           resourceSlbTemplatePolicy(),
			"vthunder_slb_template_server_ssl":       resourceSlbTemplateServerSSL(),
			"vthunder_slb_template_client_ssl":       resourceSlbTemplateClientSSL(),
			"vthunder_slb_template_sip":              resourceSlbTemplateSIP(),
			"vthunder_slb_template_http":             resourceSlbTemplateHTTP(),
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
