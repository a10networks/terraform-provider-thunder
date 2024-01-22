package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceAccounting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_accounting`: Create resource accounting template\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceAccountingCreate,
		UpdateContext: resourceSystemResourceAccountingUpdate,
		ReadContext:   resourceSystemResourceAccountingRead,
		DeleteContext: resourceSystemResourceAccountingDelete,

		Schema: map[string]*schema.Schema{
			"template_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Enter template name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
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
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemResourceAccountingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccounting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceAccountingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccounting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceAccountingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccounting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceAccountingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccounting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemResourceAccountingTemplateList(d []interface{}) []edpt.SystemResourceAccountingTemplateList {

	count1 := len(d)
	ret := make([]edpt.SystemResourceAccountingTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemResourceAccountingTemplateList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.AppResources = getObjectSystemResourceAccountingTemplateListAppResources(in["app_resources"].([]interface{}))
		oi.NetworkResources = getObjectSystemResourceAccountingTemplateListNetworkResources(in["network_resources"].([]interface{}))
		oi.SystemResources = getObjectSystemResourceAccountingTemplateListSystemResources(in["system_resources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResources(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResources {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg(in["gslb_device_cfg"].([]interface{}))
		ret.GslbGeoLocationCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg(in["gslb_geo_location_cfg"].([]interface{}))
		ret.GslbIpListCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbIpListCfg(in["gslb_ip_list_cfg"].([]interface{}))
		ret.GslbPolicyCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg(in["gslb_policy_cfg"].([]interface{}))
		ret.GslbServiceCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceCfg(in["gslb_service_cfg"].([]interface{}))
		ret.GslbServiceIpCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg(in["gslb_service_ip_cfg"].([]interface{}))
		ret.GslbServicePortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg(in["gslb_service_port_cfg"].([]interface{}))
		ret.GslbSiteCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbSiteCfg(in["gslb_site_cfg"].([]interface{}))
		ret.GslbSvcGroupCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg(in["gslb_svc_group_cfg"].([]interface{}))
		ret.GslbTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg(in["gslb_template_cfg"].([]interface{}))
		ret.GslbZoneCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbZoneCfg(in["gslb_zone_cfg"].([]interface{}))
		ret.HealthMonitorCfg = getObjectSystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg(in["health_monitor_cfg"].([]interface{}))
		ret.RealPortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesRealPortCfg(in["real_port_cfg"].([]interface{}))
		ret.RealServerCfg = getObjectSystemResourceAccountingTemplateListAppResourcesRealServerCfg(in["real_server_cfg"].([]interface{}))
		ret.ServiceGroupCfg = getObjectSystemResourceAccountingTemplateListAppResourcesServiceGroupCfg(in["service_group_cfg"].([]interface{}))
		ret.VirtualServerCfg = getObjectSystemResourceAccountingTemplateListAppResourcesVirtualServerCfg(in["virtual_server_cfg"].([]interface{}))
		ret.VirtualPortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesVirtualPortCfg(in["virtual_port_cfg"].([]interface{}))
		ret.CacheTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg(in["cache_template_cfg"].([]interface{}))
		ret.ClientSslTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg(in["client_ssl_template_cfg"].([]interface{}))
		ret.ConnReuseTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg(in["conn_reuse_template_cfg"].([]interface{}))
		ret.FastTcpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg(in["fast_tcp_template_cfg"].([]interface{}))
		ret.FastUdpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg(in["fast_udp_template_cfg"].([]interface{}))
		ret.FixTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFixTemplateCfg(in["fix_template_cfg"].([]interface{}))
		ret.HttpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg(in["http_template_cfg"].([]interface{}))
		ret.LinkCostTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg(in["link_cost_template_cfg"].([]interface{}))
		ret.PbslbEntryCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg(in["pbslb_entry_cfg"].([]interface{}))
		ret.PersistCookieTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg(in["persist_cookie_template_cfg"].([]interface{}))
		ret.PersistSrcipTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg(in["persist_srcip_template_cfg"].([]interface{}))
		ret.ServerSslTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg(in["server_ssl_template_cfg"].([]interface{}))
		ret.ProxyTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg(in["proxy_template_cfg"].([]interface{}))
		ret.StreamTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg(in["stream_template_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceMax = in["gslb_device_max"].(int)
		ret.GslbDeviceMinGuarantee = in["gslb_device_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbGeoLocationMax = in["gslb_geo_location_max"].(int)
		ret.GslbGeoLocationMinGuarantee = in["gslb_geo_location_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbIpListCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbIpListMax = in["gslb_ip_list_max"].(int)
		ret.GslbIpListMinGuarantee = in["gslb_ip_list_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbPolicyMax = in["gslb_policy_max"].(int)
		ret.GslbPolicyMinGuarantee = in["gslb_policy_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceMax = in["gslb_service_max"].(int)
		ret.GslbServiceMinGuarantee = in["gslb_service_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceIpMax = in["gslb_service_ip_max"].(int)
		ret.GslbServiceIpMinGuarantee = in["gslb_service_ip_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServicePortMax = in["gslb_service_port_max"].(int)
		ret.GslbServicePortMinGuarantee = in["gslb_service_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbSiteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSiteMax = in["gslb_site_max"].(int)
		ret.GslbSiteMinGuarantee = in["gslb_site_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSvcGroupMax = in["gslb_svc_group_max"].(int)
		ret.GslbSvcGroupMinGuarantee = in["gslb_svc_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbTemplateMax = in["gslb_template_max"].(int)
		ret.GslbTemplateMinGuarantee = in["gslb_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbZoneCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbZoneMax = in["gslb_zone_max"].(int)
		ret.GslbZoneMinGuarantee = in["gslb_zone_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HealthMonitorMax = in["health_monitor_max"].(int)
		ret.HealthMonitorMinGuarantee = in["health_monitor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesRealPortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesRealPortCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesRealPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealPortMax = in["real_port_max"].(int)
		ret.RealPortMinGuarantee = in["real_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesRealServerCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesRealServerCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesRealServerCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealServerMax = in["real_server_max"].(int)
		ret.RealServerMinGuarantee = in["real_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesServiceGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceGroupMax = in["service_group_max"].(int)
		ret.ServiceGroupMinGuarantee = in["service_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesVirtualServerCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualServerMax = in["virtual_server_max"].(int)
		ret.VirtualServerMinGuarantee = in["virtual_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesVirtualPortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualPortMax = in["virtual_port_max"].(int)
		ret.VirtualPortMinGuarantee = in["virtual_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheTemplateMax = in["cache_template_max"].(int)
		ret.CacheTemplateMinGuarantee = in["cache_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSslTemplateMax = in["client_ssl_template_max"].(int)
		ret.ClientSslTemplateMinGuarantee = in["client_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnReuseTemplateMax = in["conn_reuse_template_max"].(int)
		ret.ConnReuseTemplateMinGuarantee = in["conn_reuse_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastTcpTemplateMax = in["fast_tcp_template_max"].(int)
		ret.FastTcpTemplateMinGuarantee = in["fast_tcp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastUdpTemplateMax = in["fast_udp_template_max"].(int)
		ret.FastUdpTemplateMinGuarantee = in["fast_udp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFixTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixTemplateMax = in["fix_template_max"].(int)
		ret.FixTemplateMinGuarantee = in["fix_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpTemplateMax = in["http_template_max"].(int)
		ret.HttpTemplateMinGuarantee = in["http_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LinkCostTemplateMax = in["link_cost_template_max"].(int)
		ret.LinkCostTemplateMinGuarantee = in["link_cost_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PbslbEntryMax = in["pbslb_entry_max"].(int)
		ret.PbslbEntryMinGuarantee = in["pbslb_entry_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistCookieTemplateMax = in["persist_cookie_template_max"].(int)
		ret.PersistCookieTemplateMinGuarantee = in["persist_cookie_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistSrcipTemplateMax = in["persist_srcip_template_max"].(int)
		ret.PersistSrcipTemplateMinGuarantee = in["persist_srcip_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerSslTemplateMax = in["server_ssl_template_max"].(int)
		ret.ServerSslTemplateMinGuarantee = in["server_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyTemplateMax = in["proxy_template_max"].(int)
		ret.ProxyTemplateMinGuarantee = in["proxy_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamTemplateMax = in["stream_template_max"].(int)
		ret.StreamTemplateMinGuarantee = in["stream_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResources(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResources {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg(in["static_ipv4_route_cfg"].([]interface{}))
		ret.StaticIpv6RouteCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg(in["static_ipv6_route_cfg"].([]interface{}))
		ret.Ipv4AclLineCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg(in["ipv4_acl_line_cfg"].([]interface{}))
		ret.Ipv6AclLineCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg(in["ipv6_acl_line_cfg"].([]interface{}))
		ret.StaticArpCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg(in["static_arp_cfg"].([]interface{}))
		ret.StaticNeighborCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg(in["static_neighbor_cfg"].([]interface{}))
		ret.StaticMacCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg(in["static_mac_cfg"].([]interface{}))
		ret.ObjectGroupCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg(in["object_group_cfg"].([]interface{}))
		ret.ObjectGroupClauseCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg(in["object_group_clause_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteMax = in["static_ipv4_route_max"].(int)
		ret.StaticIpv4RouteMinGuarantee = in["static_ipv4_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv6RouteMax = in["static_ipv6_route_max"].(int)
		ret.StaticIpv6RouteMinGuarantee = in["static_ipv6_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4AclLineMax = in["ipv4_acl_line_max"].(int)
		ret.Ipv4AclLineMinGuarantee = in["ipv4_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6AclLineMax = in["ipv6_acl_line_max"].(int)
		ret.Ipv6AclLineMinGuarantee = in["ipv6_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticArpMax = in["static_arp_max"].(int)
		ret.StaticArpMinGuarantee = in["static_arp_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNeighborMax = in["static_neighbor_max"].(int)
		ret.StaticNeighborMinGuarantee = in["static_neighbor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticMacMax = in["static_mac_max"].(int)
		ret.StaticMacMinGuarantee = in["static_mac_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupMax = in["object_group_max"].(int)
		ret.ObjectGroupMinGuarantee = in["object_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupClauseMax = in["object_group_clause_max"].(int)
		ret.ObjectGroupClauseMinGuarantee = in["object_group_clause_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResources(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResources {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesBwLimitCfg(in["bw_limit_cfg"].([]interface{}))
		ret.ConcurrentSessionLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg(in["concurrent_session_limit_cfg"].([]interface{}))
		ret.L4SessionLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg(in["l4_session_limit_cfg"].([]interface{}))
		ret.L4cpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg(in["l4cps_limit_cfg"].([]interface{}))
		ret.L7cpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg(in["l7cps_limit_cfg"].([]interface{}))
		ret.NatcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg(in["natcps_limit_cfg"].([]interface{}))
		ret.FwcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg(in["fwcps_limit_cfg"].([]interface{}))
		ret.SslThroughputLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg(in["ssl_throughput_limit_cfg"].([]interface{}))
		ret.SslcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg(in["sslcps_limit_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesBwLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitMax = in["bw_limit_max"].(int)
		ret.BwLimitWatermarkDisable = in["bw_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentSessionLimitMax = in["concurrent_session_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4SessionLimitMax = in["l4_session_limit_max"].(string)
		ret.L4SessionLimitMinGuarantee = in["l4_session_limit_min_guarantee"].(string)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4cpsLimitMax = in["l4cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L7cpsLimitMax = in["l7cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatcpsLimitMax = in["natcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FwcpsLimitMax = in["fwcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslThroughputLimitMax = in["ssl_throughput_limit_max"].(int)
		ret.SslThroughputLimitWatermarkDisable = in["ssl_throughput_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslcpsLimitMax = in["sslcps_limit_max"].(int)
	}
	return ret
}

func dataToEndpointSystemResourceAccounting(d *schema.ResourceData) edpt.SystemResourceAccounting {
	var ret edpt.SystemResourceAccounting
	ret.Inst.TemplateList = getSliceSystemResourceAccountingTemplateList(d.Get("template_list").([]interface{}))
	//omit uuid
	return ret
}
