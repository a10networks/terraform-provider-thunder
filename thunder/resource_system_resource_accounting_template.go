package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceAccountingTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_accounting_template`: Create resource accounting template\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceAccountingTemplateCreate,
		UpdateContext: resourceSystemResourceAccountingTemplateUpdate,
		ReadContext:   resourceSystemResourceAccountingTemplateRead,
		DeleteContext: resourceSystemResourceAccountingTemplateDelete,

		Schema: map[string]*schema.Schema{
			"app_resources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_device_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_device_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-device allowed (gslb-device count (default is max-value))",
									},
									"gslb_device_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_geo_location_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_geo_location_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-geo-location allowed (gslb-geo-location count (default is max-value))",
									},
									"gslb_geo_location_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_ip_list_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_ip_list_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-ip-list allowed (gslb-ip-list count (default is max-value))",
									},
									"gslb_ip_list_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_policy_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_policy_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-policy allowed (gslb-policy count (default is max-value))",
									},
									"gslb_policy_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_service_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_service_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service allowed (gslb-service count (default is max-value)",
									},
									"gslb_service_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_service_ip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_service_ip_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-ip allowed (gslb-service-ip count (default is max-value))",
									},
									"gslb_service_ip_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_service_port_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_service_port_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-port allowed ( gslb-service-port count (default is max-value))",
									},
									"gslb_service_port_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_site_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_site_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-site allowed (gslb-site count (default is max-value))",
									},
									"gslb_site_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_svc_group_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_svc_group_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-svc-group allowed (gslb-svc-group count (default is max-value))",
									},
									"gslb_svc_group_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-template allowed (gslb-template count (default is max-value))",
									},
									"gslb_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"gslb_zone_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_zone_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-zone allowed (gslb-zone count (default is max-value))",
									},
									"gslb_zone_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"health_monitor_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"health_monitor_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of health monitors allowed (health-monitor count (default is max-value))",
									},
									"health_monitor_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"real_port_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"real_port_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-ports allowed (real-port count (default is max-value))",
									},
									"real_port_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"real_server_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"real_server_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-servers allowed (real-server count (default is max-value))",
									},
									"real_server_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"service_group_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_group_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of service groups allowed (service-group count (default is max-value))",
									},
									"service_group_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"virtual_server_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_server_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-servers allowed (virtual-server count (default is max-value))",
									},
									"virtual_server_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"virtual_port_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_port_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-port allowed (virtual-port count (default is max-value))",
									},
									"virtual_port_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"cache_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cache_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of cache-template allowed (cache-template count (default is max-value))",
									},
									"cache_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"client_ssl_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ssl_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of client-ssl-template allowed (client-ssl-template count (default is max-value))",
									},
									"client_ssl_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"conn_reuse_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"conn_reuse_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of conn-reuse-template allowed (conn-reuse-template count (default is max-value))",
									},
									"conn_reuse_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"fast_tcp_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fast_tcp_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-tcp-template allowed (fast-tcp-template count (default is max-value))",
									},
									"fast_tcp_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"fast_udp_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fast_udp_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-udp-template allowed (fast-udp-template count (default is max-value))",
									},
									"fast_udp_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"fix_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fix_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of fix-template allowed (fix-template count (default is max-value))",
									},
									"fix_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"http_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of http-template allowed (http-template count (default is max-value))",
									},
									"http_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"link_cost_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"link_cost_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of link-cost-template allowed (link-cost-template count (default is max-value))",
									},
									"link_cost_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"pbslb_entry_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pbslb_entry_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of pbslb-entry allowed (pbslb-entry count)",
									},
									"pbslb_entry_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"persist_cookie_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"persist_cookie_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-cookie-template allowed (persist-cookie-template count (default is max-value))",
									},
									"persist_cookie_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"persist_srcip_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"persist_srcip_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-srcip-template allowed (persist-source-ip-template count (default is max-value))",
									},
									"persist_srcip_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"server_ssl_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_ssl_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of server-ssl-template allowed (server-ssl-template count (default is max-value))",
									},
									"server_ssl_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"proxy_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proxy_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of proxy-template allowed (server-ssl-template count (default is max-value))",
									},
									"proxy_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"stream_template_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stream_template_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of stream-template allowed (server-ssl-template count (default is max-value))",
									},
									"stream_template_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Enter template name",
			},
			"network_resources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_ipv4_route_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_ipv4_route_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv4 routes allowed (Static ipv4 routes (default is max-value))",
									},
									"static_ipv4_route_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"static_ipv6_route_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_ipv6_route_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv6 routes allowed (Static ipv6 routes (default is max-value))",
									},
									"static_ipv6_route_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"ipv4_acl_line_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_acl_line_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV4 ACL lines (default is max-value))",
									},
									"ipv4_acl_line_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"ipv6_acl_line_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_acl_line_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV6 ACL lines (default is max-value))",
									},
									"ipv6_acl_line_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"static_arp_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_arp_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of static arp entries allowed (Static arp (default is max-value))",
									},
									"static_arp_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"static_neighbor_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_neighbor_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of static neighbor entries allowed (Static neighbors (default is max-value))",
									},
									"static_neighbor_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"static_mac_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_mac_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of static MAC entries allowed (Static MACs (default is max-value))",
									},
									"static_mac_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"object_group_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_group_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of object groups allowed (Object group (default is max-value))",
									},
									"object_group_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"object_group_clause_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_group_clause_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the number of object group clauses allowed (Object group clauses (default is max-value))",
									},
									"object_group_clause_min_guarantee": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
									},
								},
							},
						},
						"threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"system_resources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bw_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bw_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the bandwidth limit in mbps (Bandwidth limit in Mbit/s (no limits applied by default))",
									},
									"bw_limit_watermark_disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions)",
									},
								},
							},
						},
						"concurrent_session_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"concurrent_session_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the Concurrent Session limit (cps) (Concurrent-Session cps limit (no limits applied by default))",
									},
								},
							},
						},
						"l4_session_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l4_session_limit_max": {
										Type: schema.TypeString, Optional: true, Description: "Enter the l4 session limit in % (0.01% to 99.99%) (Enter a number from 0.01 to 99.99 (up to 2 digits precision))",
									},
									"l4_session_limit_min_guarantee": {
										Type: schema.TypeString, Optional: true, Default: "0", Description: "minimum guaranteed value in % (up to 2 digits precision) (Enter a number from 0 to 99.99)",
									},
								},
							},
						},
						"l4cps_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l4cps_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the L4 cps limit (L4 cps limit (no limits applied by default))",
									},
								},
							},
						},
						"l7cps_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l7cps_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "L7cps-limit (L7 cps limit (no limits applied by default))",
									},
								},
							},
						},
						"natcps_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"natcps_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the Nat cps limit (NAT cps limit (no limits applied by default))",
									},
								},
							},
						},
						"fwcps_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fwcps_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the Firewall cps limit (Firewall cps limit (no limits applied by default))",
									},
								},
							},
						},
						"ssl_throughput_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssl_throughput_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the ssl throughput limit in mbps (SSL Througput limit in Mbit/s (no limits applied by default))",
									},
									"ssl_throughput_limit_watermark_disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions)",
									},
								},
							},
						},
						"sslcps_limit_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sslcps_limit_max": {
										Type: schema.TypeInt, Optional: true, Description: "Enter the SSL cps limit (SSL cps limit (no limits applied by default))",
									},
								},
							},
						},
						"threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemResourceAccountingTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceAccountingTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceAccountingTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceAccountingTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemResourceAccountingTemplateAppResources1701(d []interface{}) edpt.SystemResourceAccountingTemplateAppResources1701 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResources1701
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1702(in["gslb_device_cfg"].([]interface{}))
		ret.GslbGeoLocationCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1703(in["gslb_geo_location_cfg"].([]interface{}))
		ret.GslbIpListCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbIpListCfg1704(in["gslb_ip_list_cfg"].([]interface{}))
		ret.GslbPolicyCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1705(in["gslb_policy_cfg"].([]interface{}))
		ret.GslbServiceCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceCfg1706(in["gslb_service_cfg"].([]interface{}))
		ret.GslbServiceIpCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1707(in["gslb_service_ip_cfg"].([]interface{}))
		ret.GslbServicePortCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1708(in["gslb_service_port_cfg"].([]interface{}))
		ret.GslbSiteCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbSiteCfg1709(in["gslb_site_cfg"].([]interface{}))
		ret.GslbSvcGroupCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1710(in["gslb_svc_group_cfg"].([]interface{}))
		ret.GslbTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1711(in["gslb_template_cfg"].([]interface{}))
		ret.GslbZoneCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbZoneCfg1712(in["gslb_zone_cfg"].([]interface{}))
		ret.HealthMonitorCfg = getObjectSystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1713(in["health_monitor_cfg"].([]interface{}))
		ret.RealPortCfg = getObjectSystemResourceAccountingTemplateAppResourcesRealPortCfg1714(in["real_port_cfg"].([]interface{}))
		ret.RealServerCfg = getObjectSystemResourceAccountingTemplateAppResourcesRealServerCfg1715(in["real_server_cfg"].([]interface{}))
		ret.ServiceGroupCfg = getObjectSystemResourceAccountingTemplateAppResourcesServiceGroupCfg1716(in["service_group_cfg"].([]interface{}))
		ret.VirtualServerCfg = getObjectSystemResourceAccountingTemplateAppResourcesVirtualServerCfg1717(in["virtual_server_cfg"].([]interface{}))
		ret.VirtualPortCfg = getObjectSystemResourceAccountingTemplateAppResourcesVirtualPortCfg1718(in["virtual_port_cfg"].([]interface{}))
		ret.CacheTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1719(in["cache_template_cfg"].([]interface{}))
		ret.ClientSslTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1720(in["client_ssl_template_cfg"].([]interface{}))
		ret.ConnReuseTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1721(in["conn_reuse_template_cfg"].([]interface{}))
		ret.FastTcpTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1722(in["fast_tcp_template_cfg"].([]interface{}))
		ret.FastUdpTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1723(in["fast_udp_template_cfg"].([]interface{}))
		ret.FixTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesFixTemplateCfg1724(in["fix_template_cfg"].([]interface{}))
		ret.HttpTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1725(in["http_template_cfg"].([]interface{}))
		ret.LinkCostTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1726(in["link_cost_template_cfg"].([]interface{}))
		ret.PbslbEntryCfg = getObjectSystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1727(in["pbslb_entry_cfg"].([]interface{}))
		ret.PersistCookieTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1728(in["persist_cookie_template_cfg"].([]interface{}))
		ret.PersistSrcipTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1729(in["persist_srcip_template_cfg"].([]interface{}))
		ret.ServerSslTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1730(in["server_ssl_template_cfg"].([]interface{}))
		ret.ProxyTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1731(in["proxy_template_cfg"].([]interface{}))
		ret.StreamTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1732(in["stream_template_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1702(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1702 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1702
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceMax = in["gslb_device_max"].(int)
		ret.GslbDeviceMinGuarantee = in["gslb_device_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1703(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1703 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1703
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbGeoLocationMax = in["gslb_geo_location_max"].(int)
		ret.GslbGeoLocationMinGuarantee = in["gslb_geo_location_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbIpListCfg1704(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbIpListCfg1704 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbIpListCfg1704
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbIpListMax = in["gslb_ip_list_max"].(int)
		ret.GslbIpListMinGuarantee = in["gslb_ip_list_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1705(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1705 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1705
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbPolicyMax = in["gslb_policy_max"].(int)
		ret.GslbPolicyMinGuarantee = in["gslb_policy_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceCfg1706(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceCfg1706 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceCfg1706
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceMax = in["gslb_service_max"].(int)
		ret.GslbServiceMinGuarantee = in["gslb_service_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1707(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1707 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1707
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceIpMax = in["gslb_service_ip_max"].(int)
		ret.GslbServiceIpMinGuarantee = in["gslb_service_ip_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1708(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1708 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1708
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServicePortMax = in["gslb_service_port_max"].(int)
		ret.GslbServicePortMinGuarantee = in["gslb_service_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbSiteCfg1709(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbSiteCfg1709 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbSiteCfg1709
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSiteMax = in["gslb_site_max"].(int)
		ret.GslbSiteMinGuarantee = in["gslb_site_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1710(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1710 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1710
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSvcGroupMax = in["gslb_svc_group_max"].(int)
		ret.GslbSvcGroupMinGuarantee = in["gslb_svc_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1711(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1711 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1711
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbTemplateMax = in["gslb_template_max"].(int)
		ret.GslbTemplateMinGuarantee = in["gslb_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbZoneCfg1712(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbZoneCfg1712 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbZoneCfg1712
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbZoneMax = in["gslb_zone_max"].(int)
		ret.GslbZoneMinGuarantee = in["gslb_zone_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1713(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1713 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1713
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HealthMonitorMax = in["health_monitor_max"].(int)
		ret.HealthMonitorMinGuarantee = in["health_monitor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesRealPortCfg1714(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesRealPortCfg1714 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesRealPortCfg1714
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealPortMax = in["real_port_max"].(int)
		ret.RealPortMinGuarantee = in["real_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesRealServerCfg1715(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesRealServerCfg1715 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesRealServerCfg1715
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealServerMax = in["real_server_max"].(int)
		ret.RealServerMinGuarantee = in["real_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesServiceGroupCfg1716(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesServiceGroupCfg1716 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesServiceGroupCfg1716
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceGroupMax = in["service_group_max"].(int)
		ret.ServiceGroupMinGuarantee = in["service_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesVirtualServerCfg1717(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesVirtualServerCfg1717 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesVirtualServerCfg1717
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualServerMax = in["virtual_server_max"].(int)
		ret.VirtualServerMinGuarantee = in["virtual_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesVirtualPortCfg1718(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesVirtualPortCfg1718 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesVirtualPortCfg1718
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualPortMax = in["virtual_port_max"].(int)
		ret.VirtualPortMinGuarantee = in["virtual_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1719(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1719 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1719
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheTemplateMax = in["cache_template_max"].(int)
		ret.CacheTemplateMinGuarantee = in["cache_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1720(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1720 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1720
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSslTemplateMax = in["client_ssl_template_max"].(int)
		ret.ClientSslTemplateMinGuarantee = in["client_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1721(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1721 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1721
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnReuseTemplateMax = in["conn_reuse_template_max"].(int)
		ret.ConnReuseTemplateMinGuarantee = in["conn_reuse_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1722(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1722 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1722
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastTcpTemplateMax = in["fast_tcp_template_max"].(int)
		ret.FastTcpTemplateMinGuarantee = in["fast_tcp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1723(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1723 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1723
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastUdpTemplateMax = in["fast_udp_template_max"].(int)
		ret.FastUdpTemplateMinGuarantee = in["fast_udp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesFixTemplateCfg1724(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesFixTemplateCfg1724 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesFixTemplateCfg1724
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixTemplateMax = in["fix_template_max"].(int)
		ret.FixTemplateMinGuarantee = in["fix_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1725(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1725 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1725
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpTemplateMax = in["http_template_max"].(int)
		ret.HttpTemplateMinGuarantee = in["http_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1726(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1726 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1726
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LinkCostTemplateMax = in["link_cost_template_max"].(int)
		ret.LinkCostTemplateMinGuarantee = in["link_cost_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1727(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1727 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1727
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PbslbEntryMax = in["pbslb_entry_max"].(int)
		ret.PbslbEntryMinGuarantee = in["pbslb_entry_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1728(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1728 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1728
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistCookieTemplateMax = in["persist_cookie_template_max"].(int)
		ret.PersistCookieTemplateMinGuarantee = in["persist_cookie_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1729(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1729 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1729
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistSrcipTemplateMax = in["persist_srcip_template_max"].(int)
		ret.PersistSrcipTemplateMinGuarantee = in["persist_srcip_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1730(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1730 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1730
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerSslTemplateMax = in["server_ssl_template_max"].(int)
		ret.ServerSslTemplateMinGuarantee = in["server_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1731(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1731 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1731
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyTemplateMax = in["proxy_template_max"].(int)
		ret.ProxyTemplateMinGuarantee = in["proxy_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1732(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1732 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1732
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamTemplateMax = in["stream_template_max"].(int)
		ret.StreamTemplateMinGuarantee = in["stream_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResources1733(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResources1733 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResources1733
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1734(in["static_ipv4_route_cfg"].([]interface{}))
		ret.StaticIpv6RouteCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1735(in["static_ipv6_route_cfg"].([]interface{}))
		ret.Ipv4AclLineCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1736(in["ipv4_acl_line_cfg"].([]interface{}))
		ret.Ipv6AclLineCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1737(in["ipv6_acl_line_cfg"].([]interface{}))
		ret.StaticArpCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1738(in["static_arp_cfg"].([]interface{}))
		ret.StaticNeighborCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1739(in["static_neighbor_cfg"].([]interface{}))
		ret.StaticMacCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1740(in["static_mac_cfg"].([]interface{}))
		ret.ObjectGroupCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1741(in["object_group_cfg"].([]interface{}))
		ret.ObjectGroupClauseCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1742(in["object_group_clause_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1734(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1734 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1734
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteMax = in["static_ipv4_route_max"].(int)
		ret.StaticIpv4RouteMinGuarantee = in["static_ipv4_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1735(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1735 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1735
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv6RouteMax = in["static_ipv6_route_max"].(int)
		ret.StaticIpv6RouteMinGuarantee = in["static_ipv6_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1736(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1736 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1736
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4AclLineMax = in["ipv4_acl_line_max"].(int)
		ret.Ipv4AclLineMinGuarantee = in["ipv4_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1737(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1737 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1737
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6AclLineMax = in["ipv6_acl_line_max"].(int)
		ret.Ipv6AclLineMinGuarantee = in["ipv6_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1738(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1738 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1738
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticArpMax = in["static_arp_max"].(int)
		ret.StaticArpMinGuarantee = in["static_arp_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1739(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1739 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1739
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNeighborMax = in["static_neighbor_max"].(int)
		ret.StaticNeighborMinGuarantee = in["static_neighbor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1740(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1740 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1740
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticMacMax = in["static_mac_max"].(int)
		ret.StaticMacMinGuarantee = in["static_mac_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1741(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1741 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1741
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupMax = in["object_group_max"].(int)
		ret.ObjectGroupMinGuarantee = in["object_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1742(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1742 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1742
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupClauseMax = in["object_group_clause_max"].(int)
		ret.ObjectGroupClauseMinGuarantee = in["object_group_clause_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResources1743(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResources1743 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResources1743
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesBwLimitCfg1744(in["bw_limit_cfg"].([]interface{}))
		ret.ConcurrentSessionLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1745(in["concurrent_session_limit_cfg"].([]interface{}))
		ret.L4SessionLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1746(in["l4_session_limit_cfg"].([]interface{}))
		ret.L4cpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1747(in["l4cps_limit_cfg"].([]interface{}))
		ret.L7cpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1748(in["l7cps_limit_cfg"].([]interface{}))
		ret.NatcpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1749(in["natcps_limit_cfg"].([]interface{}))
		ret.FwcpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1750(in["fwcps_limit_cfg"].([]interface{}))
		ret.SslThroughputLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1751(in["ssl_throughput_limit_cfg"].([]interface{}))
		ret.SslcpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1752(in["sslcps_limit_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesBwLimitCfg1744(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesBwLimitCfg1744 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesBwLimitCfg1744
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitMax = in["bw_limit_max"].(int)
		ret.BwLimitWatermarkDisable = in["bw_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1745(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1745 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1745
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentSessionLimitMax = in["concurrent_session_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1746(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1746 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1746
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4SessionLimitMax = in["l4_session_limit_max"].(string)
		ret.L4SessionLimitMinGuarantee = in["l4_session_limit_min_guarantee"].(string)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1747(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1747 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1747
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4cpsLimitMax = in["l4cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1748(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1748 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1748
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L7cpsLimitMax = in["l7cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1749(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1749 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1749
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatcpsLimitMax = in["natcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1750(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1750 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1750
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FwcpsLimitMax = in["fwcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1751(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1751 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1751
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslThroughputLimitMax = in["ssl_throughput_limit_max"].(int)
		ret.SslThroughputLimitWatermarkDisable = in["ssl_throughput_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1752(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1752 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1752
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslcpsLimitMax = in["sslcps_limit_max"].(int)
	}
	return ret
}

func dataToEndpointSystemResourceAccountingTemplate(d *schema.ResourceData) edpt.SystemResourceAccountingTemplate {
	var ret edpt.SystemResourceAccountingTemplate
	ret.Inst.AppResources = getObjectSystemResourceAccountingTemplateAppResources1701(d.Get("app_resources").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NetworkResources = getObjectSystemResourceAccountingTemplateNetworkResources1733(d.Get("network_resources").([]interface{}))
	ret.Inst.SystemResources = getObjectSystemResourceAccountingTemplateSystemResources1743(d.Get("system_resources").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
