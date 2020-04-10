package vthunder

import (
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const DEFAULT_PARTITION = "Common"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name/IP of the VTHUNDER",
				DefaultFunc: schema.EnvDefaultFunc("VTHUNDER_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username with API access to the VTHUNDER",
				DefaultFunc: schema.EnvDefaultFunc("VTHUNDER_USER", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user's password",
				DefaultFunc: schema.EnvDefaultFunc("VTHUNDER_PASSWORD", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"vthunder_virtual_server":                    resourceVirtualServer(),
			"vthunder_service_group":                     resourceServiceGroup(),
			"vthunder_server":                            resourceServer(),
			"vthunder_ethernet":                          resourceEthernet(),
			"vthunder_rib_route":                         resourceRibRoute(),
			"vthunder_vrrp_common":                       resourceVrrpCommon(),
			"vthunder_vrrp_vrid":                         resourceVrrpVrid(),
			"vthunder_vrrp_peer_group":                   resourceVrrpPeerGroup(),
			"vthunder_dns_primary":                       resourceDnsPrimary(),
			"vthunder_import":                            resourceImport(),
			"vthunder_reboot":                            resourceReboot(),
			"vthunder_vrrp_session_sync":                 resourceVrrpSessionSync(),
			"vthunder_glm":                               resourceGlm(),
			"vthunder_configure_sync":                    resourceConfigureSync(),
			"vthunder_harmony_controller_profile":        resourceHarmonyControllerProfile(),
			"vthunder_overlay_tunnel_options":            resourceOverlayTunnelOptions(),
			"vthunder_overlay_tunnel_vtep":               resourceOverlayTunnelVtep(),
			"vthunder_partition":                         resourceOverlayTunnelPartition(),
			"vthunder_slb_template_tcp":                  resourceSlbTemplateTcp(),
			"vthunder_slb_template_udp":                  resourceSlbTemplateUdp(),
			"vthunder_slb_template_smpp":                 resourceSlbTemplateSmpp(),
			"vthunder_slb_template_fix":                  resourceTemplateFix(),
			"vthunder_slb_template_ftp":                  resourceSlbTemplateFTP(),
			"vthunder_slb_template_mqtt":                 resourceSlbTemplateMqtt(),
			"vthunder_slb_template_http_policy":          resourceSlbTemplateHttpPolicy(),
			"vthunder_slb_template_ssli":                 resourceTemplateSSLI(),
			"vthunder_slb_template_cipher":               resourceTemplateCipher(),
			"vthunder_slb_template_imap_pop3":            resourceTemplateImap_POP3(),
			"vthunder_slb_template_dns":                  resourceTemplateDNS(),
			"vthunder_slb_template_connection_reuse":     resourceTemplateConnectionReuse(),
			"vthunder_slb_template_client_ssh":           resourceTemplateClientSsh(),
			"vthunder_slb_template_dblb":                 resourceTemplateDBLB(),
			"vthunder_slb_template_port":                 resourceTemplatePort(),
			"vthunder_slb_template_smtp":                 resourceSlbTemplateSMTP(),
			"vthunder_slb_template_tcp_proxy":            resourceSlbTemplateTcpProxy(),
			"vthunder_slb_template_server":               resourceSlbTemplateServer(),
			"vthunder_slb_template_virtual_port":         resourceSlbTemplateVirtualPort(),
			"vthunder_slb_template_diameter":             resourceSlbTemplateDiameter(),
			"vthunder_slb_template_cache":                resourceSlbTemplateCache(),
			"vthunder_slb_template_csv":                  resourceSlbTemplateCSV(),
			"vthunder_slb_template_reqmod_icap":          resourceSlbTemplateReqmodIcap(),
			"vthunder_slb_template_respmod_icap":         resourceSlbTemplateRespmodIcap(),
			"vthunder_slb_template_dynamic_service":      resourceSlbTemplateDynamicService(),
			"vthunder_slb_template_server_ssh":           resourceSlbTemplateServerSSH(),
			"vthunder_slb_template_policy":               resourceSlbTemplatePolicy(),
			"vthunder_slb_template_server_ssl":           resourceSlbTemplateServerSSL(),
			"vthunder_slb_template_client_ssl":           resourceSlbTemplateClientSSL(),
			"vthunder_slb_template_sip":                  resourceSlbTemplateSIP(),
			"vthunder_slb_template_http":                 resourceSlbTemplateHTTP(),
			"vthunder_slb_template_monitor":              resourceSlbTemplateMonitor(),
			"vthunder_slb_template_logging":              resourceSlbTemplateLogging(),
			"vthunder_slb_template_snmp":                 resourceSlbTemplateSNMP(),
			"vthunder_slb_template_external_service":     resourceSlbTemplateExternalService(),
			"vthunder_slb_template_virtual_server":       resourceSlbTemplateVirtualServer(),
			"vthunder_slb_aflow":                         resourceSlbAflow(),
			"vthunder_slb_common":                        resourceSlbCommon(),
			"vthunder_slb_common_conn_rate_limit_src_ip": resourceSlbCommonConnRateLimitSrcIP(),
			"vthunder_slb_connection_reuse":              resourceSlbConnectionReuse(),
			"vthunder_slb_crl_srcip":                     resourceSlbCrlSrcip(),
			"vthunder_slb_dns":                           resourceSlbDNS(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func setToStringSlice(s *schema.Set) []string {
	list := make([]string, s.Len())
	for i, v := range s.List() {
		list[i] = v.(string)
	}
	return list
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:  d.Get("address").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	return config.Client()
}

func mapEntity(d map[string]interface{}, obj interface{}) {
	val := reflect.ValueOf(obj).Elem()
	for field := range d {
		f := val.FieldByName(strings.Title(field))
		if f.IsValid() {
			if f.Kind() == reflect.Slice {
				incoming := d[field].([]interface{})
				s := reflect.MakeSlice(f.Type(), len(incoming), len(incoming))
				for i := 0; i < len(incoming); i++ {
					s.Index(i).Set(reflect.ValueOf(incoming[i]))
				}
				f.Set(s)
			} else {
				f.Set(reflect.ValueOf(d[field]))
			}
		} else {
			log.Printf("[WARN] You probably weren't expecting %s to be an invalid field", field)
		}
	}
}
