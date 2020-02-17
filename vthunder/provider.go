package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"reflect"
	"strings"
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
			"vthunder_virtual_server":    resourceVirtualServer(),
			"vthunder_service_group":     resourceServiceGroup(),
			"vthunder_server":            resourceServer(),
			"vthunder_ethernet":          resourceEthernet(),
			"vthunder_rib_route":         resourceRibRoute(),
			"vthunder_vrrp_common":       resourceVrrpCommon(),
			"vthunder_vrrp_vrid":         resourceVrrpVrid(),
			"vthunder_vrrp_peer_group":   resourceVrrpPeerGroup(),
			"vthunder_dns_primary":       resourceDnsPrimary(),
			"vthunder_import":            resourceImport(),
			"vthunder_reboot":            resourceReboot(),
			"vthunder_vrrp_session_sync": resourceVrrpSessionSync(),
			"vthunder_glm": resourceGlm(),
			"vthunder_configure_sync": resourceConfigureSync(),
			"vthunder_harmony_controller_profile": resourceHarmonyControllerProfile(),
			"vthunder_overlay_tunnel_options": resourceOverlayTunnelOptions(),
			"vthunder_overlay_tunnel_vtep": resourceOverlayTunnelVtep(),
			"vthunder_partition":resourceOverlayTunnelPartition(),
			"vthunder_slb_template_tcp":resourceSlbTemplateTcp(),
			"vthunder_slb_template_udp":resourceSlbTemplateUdp(),
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
