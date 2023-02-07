package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZone() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone`: Specify the DNS zone name for which global SLB is provided\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneCreate,
		UpdateContext: resourceGslbZoneUpdate,
		ReadContext:   resourceGslbZoneRead,
		DeleteContext: resourceGslbZoneDelete,
		Schema: map[string]*schema.Schema{
			"disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Disable all services in the GSLB zone",
			},
			"dns_mx_record_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mx_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify Domain Name",
						},
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify Priority",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify TTL",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "uuid of the object",
						},
						"sampling_enable": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_ns_record_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ns_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify Domain Name",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify TTL",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "uuid of the object",
						},
						"sampling_enable": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_soa_record": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"soa_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "DNS Server Name",
						},
						"mail": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Mailbox",
						},
						"expire": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1209600,
							Description: "Specify Expire Time Interval, default is 1209600",
						},
						"refresh": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3600,
							Description: "Specify Refresh Time Interval, default is 3600",
						},
						"retry": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     900,
							Description: "Specify Retry Time Interval, default is 900",
						},
						"serial": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify Serial Number, default is Current Time (Time Interval)",
						},
						"soa_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify Negative caching TTL, default is Zone TTL",
						},
						"external": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify External SOA Record (DNS Server Name)",
						},
						"ex_mail": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Mailbox",
						},
						"ex_expire": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1209600,
							Description: "Specify Expire Time Interval, default is 1209600",
						},
						"ex_refresh": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3600,
							Description: "Specify Refresh Time Interval, default is 3600",
						},
						"ex_retry": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     900,
							Description: "Specify Retry Time Interval, default is 900",
						},
						"ex_serial": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify Serial Number, default is Current Time (Time Interval)",
						},
						"ex_soa_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify Negative caching TTL, default is Zone TTL",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specify the name for the DNS zone",
			},
			"policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Specify the policy for this zone (Specify policy name)",
			},
			"sampling_enable": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'all': all; 'received-query': Total Number of DNS queries received for the zone; 'sent-response': Total Number of DNS replies sent to clients for the zone; 'proxy-mode-response': Total Number of DNS replies sent to clients by the ACOS device as a DNS proxy for the zone; 'cache-mode-response': Total Number of cached DNS replies sent to clients by the ACOS device for the zone. (This statistic applies only if the DNS cac; 'server-mode-response': Total Number of DNS replies sent to clients by the ACOS device as a DNS server for the zone. (This statistic applies only if th; 'sticky-mode-response': Total Number of DNS replies sent to clients by the ACOS device to keep the clients on the same site. (This statistic applies on; 'backup-mode-response': Total Number of DNS replies sent to clients by the ACOS device in backup mode;",
						},
					},
				},
			},
			"service_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number of the service",
						},
						"service_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify the service name for the zone, * for wildcard",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'drop': Drop query; 'forward': Forward packet; 'ignore': Send empty response; 'reject': Send refuse response;",
						},
						"forward_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'both': Forward both query and response; 'query': Forward query; 'response': Forward response;",
						},
						"disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Disable",
						},
						"health_check_gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "enable",
							Description: "'enable': Enable Gateway Status Check; 'disable': Disable Gateway Status Check;",
						},
						"health_check_port": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"health_check_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Check Related Port Status (Port Number)",
									},
								},
							},
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify policy for this service (Specify policy name)",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "uuid of the object",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Customized tag",
						},
						"sampling_enable": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'all': all; 'received-query': Number of DNS queries received for the service; 'sent-response': Number of DNS replies sent to clients for the service; 'proxy-mode-response': Number of DNS replies sent to clients by the ACOS device as a DNS proxy for the service; 'cache-mode-response': Number of cached DNS replies sent to clients by the ACOS device for the service. (This statistic applies only if the DNS cache; 'server-mode-response': Number of DNS replies sent to clients by the ACOS device as a DNS server for the service. (This statistic applies only if the D; 'sticky-mode-response': Number of DNS replies sent to clients by the ACOS device to keep the clients on the same site. (This statistic applies only if; 'backup-mode-response': help Number of DNS replies sent to clients by the ACOS device in backup mode;",
									},
								},
							},
						},
						"dns_a_record": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_a_record_srv_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"svrname": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specify name",
												},
												"no_resp": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Don't use this Service-IP as DNS response",
												},
												"as_backup": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "As backup when fail",
												},
												"weight": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specify weight for Service-IP (Weight value)",
												},
												"ttl": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Specify TTL for Service-IP",
												},
												"as_replace": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Return this Service-IP when enable ip-replace",
												},
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Disable this Service-IP",
												},
												"static": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Return this Service-IP in DNS server mode",
												},
												"admin_ip": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specify admin priority of Service-IP (Specify the priority)",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "uuid of the object",
												},
											},
										},
									},
									"dns_a_record_ipv4_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_a_record_ip": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specify IP address",
												},
												"no_resp": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Don't use this Service-IP as DNS response",
												},
												"as_backup": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "As backup when fail",
												},
												"weight": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specify weight for Service-IP (Weight value)",
												},
												"ttl": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Specify TTL for Service-IP",
												},
												"as_replace": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Return this Service-IP when enable ip-replace",
												},
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Disable this Service-IP",
												},
												"static": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Return this Service-IP in DNS server mode",
												},
												"admin_ip": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specify admin priority of Service-IP (Specify the priority)",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "uuid of the object",
												},
											},
										},
									},
									"dns_a_record_ipv6_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_a_record_ipv6": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "IPV6 address",
												},
												"no_resp": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Don't use this Service-IP as DNS response",
												},
												"as_backup": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "As backup when fail",
												},
												"weight": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specify weight for Service-IP (Weight value)",
												},
												"ttl": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Specify TTL for Service-IP",
												},
												"as_replace": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Return this Service-IP when enable ip-replace",
												},
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Disable this Service-IP",
												},
												"static": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Return this Service-IP in DNS server mode",
												},
												"admin_ip": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specify admin priority of Service-IP (Specify the priority)",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"dns_cname_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alias_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify the alias name",
									},
									"admin_preference": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     100,
										Description: "Specify Administrative Preference, default is 100",
									},
									"weight": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     1,
										Description: "Specify Weight, default is 1",
									},
									"as_backup": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "As backup when fail",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'cname-hits': Number of times the CNAME has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_mx_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mx_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify Domain Name",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify Priority",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Specify TTL",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of times the record has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_ns_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ns_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify Domain Name",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Specify TTL",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of times the record has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_ptr_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ptr_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify Domain Name",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Specify TTL",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of times the record has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_srv_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"srv_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify Domain Name",
									},
									"port": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specify Port (Port Number)",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify Priority",
									},
									"weight": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     10,
										Description: "Specify Weight, default is 10",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Specify TTL",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of times the record has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_naptr_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"naptr_target": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify the replacement or regular expression",
									},
									"service_proto": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify Service and Protocol",
									},
									"flag": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify the flag (e.g., a, s). Default is empty flag",
									},
									"order": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify Order",
									},
									"preference": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify Preference",
									},
									"regexp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Return the regular expression",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Specify TTL",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'naptr-hits': Number of times the NAPTR has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_txt_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"record_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify the Object Name for TXT Data",
									},
									"txt_data": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specify TXT Data",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Specify TTL",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of times the record has been used;",
												},
											},
										},
									},
								},
							},
						},
						"dns_record_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specify DNS Type",
									},
									"data": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specify DNS Data",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
								},
							},
						},
						"geo_location_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specify the geo-location",
									},
									"alias": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alias": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Send CNAME response for this geo-location (Specify a CNAME record)",
												},
											},
										},
									},
									"action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Action for this geo-location",
									},
									"action_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'allow': Allow query from this geo-location; 'drop': Drop query from this geo-location; 'forward': Forward packet for this geo-location; 'ignore': Send empty response to this geo-location; 'reject': Send refuse response to this geo-location;",
									},
									"forward_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'both': Forward both query and response; 'query': Forward query from this geo-location; 'response': Forward response to this geo-location;",
									},
									"policy": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Policy for this geo-location (Specify the policy name)",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"template": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dnssec": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify DNSSEC template (Specify template name)",
						},
					},
				},
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "Specify the zone ttl value (TTL value, unit: second, default is 10)",
			},
			"use_server_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Use DNS Server Response TTL value in GSLB Proxy mode",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customized tag",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZone(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZone(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}
func resourceGslbZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZone(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZone(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}
func getSliceGslbZoneDnsMxRecordList(d []interface{}) []edpt.GslbZoneDnsMxRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneDnsMxRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Priority = in["priority"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneDnsMxRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneDnsMxRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneDnsMxRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneDnsMxRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneDnsMxRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneDnsNsRecordList(d []interface{}) []edpt.GslbZoneDnsNsRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneDnsNsRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneDnsNsRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneDnsNsRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneDnsNsRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneDnsNsRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneDnsNsRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getObjectGslbZoneDnsSoaRecord(d []interface{}) edpt.GslbZoneDnsSoaRecord {
	count := len(d)
	var ret edpt.GslbZoneDnsSoaRecord
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SoaName = in["soa_name"].(string)
		ret.Mail = in["mail"].(string)
		ret.Expire = in["expire"].(int)
		ret.Refresh = in["refresh"].(int)
		ret.Retry = in["retry"].(int)
		ret.Serial = in["serial"].(int)
		ret.SoaTtl = in["soa_ttl"].(int)
		ret.External = in["external"].(string)
		ret.ExMail = in["ex_mail"].(string)
		ret.ExExpire = in["ex_expire"].(int)
		ret.ExRefresh = in["ex_refresh"].(int)
		ret.ExRetry = in["ex_retry"].(int)
		ret.ExSerial = in["ex_serial"].(int)
		ret.ExSoaTtl = in["ex_soa_ttl"].(int)
	}
	return ret
}
func getSliceGslbZoneSamplingEnable(d []interface{}) []edpt.GslbZoneSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceList(d []interface{}) []edpt.GslbZoneServiceList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceList
		oi.ServicePort = in["service_port"].(int)
		oi.ServiceName = in["service_name"].(string)
		oi.Action = in["action"].(string)
		oi.ForwardType = in["forward_type"].(string)
		oi.Disable = in["disable"].(int)
		oi.HealthCheckGateway = in["health_check_gateway"].(string)
		oi.HealthCheckPort = getSliceGslbZoneServiceListHealthCheckPort(in["health_check_port"].([]interface{}))
		oi.Policy = in["policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceGslbZoneServiceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.DnsARecord = getObjectGslbZoneServiceListDnsARecord(in["dns_a_record"].([]interface{}))
		oi.DnsCnameRecordList = getSliceGslbZoneServiceListDnsCnameRecordList(in["dns_cname_record_list"].([]interface{}))
		oi.DnsMxRecordList = getSliceGslbZoneServiceListDnsMxRecordList(in["dns_mx_record_list"].([]interface{}))
		oi.DnsNsRecordList = getSliceGslbZoneServiceListDnsNsRecordList(in["dns_ns_record_list"].([]interface{}))
		oi.DnsPtrRecordList = getSliceGslbZoneServiceListDnsPtrRecordList(in["dns_ptr_record_list"].([]interface{}))
		oi.DnsSrvRecordList = getSliceGslbZoneServiceListDnsSrvRecordList(in["dns_srv_record_list"].([]interface{}))
		oi.DnsNaptrRecordList = getSliceGslbZoneServiceListDnsNaptrRecordList(in["dns_naptr_record_list"].([]interface{}))
		oi.DnsTxtRecordList = getSliceGslbZoneServiceListDnsTxtRecordList(in["dns_txt_record_list"].([]interface{}))
		oi.DnsRecordList = getSliceGslbZoneServiceListDnsRecordList(in["dns_record_list"].([]interface{}))
		oi.GeoLocationList = getSliceGslbZoneServiceListGeoLocationList(in["geo_location_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListHealthCheckPort(d []interface{}) []edpt.GslbZoneServiceListHealthCheckPort {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListHealthCheckPort, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListHealthCheckPort
		oi.HealthCheckPort = in["health_check_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getObjectGslbZoneServiceListDnsARecord(d []interface{}) edpt.GslbZoneServiceListDnsARecord {
	count := len(d)
	var ret edpt.GslbZoneServiceListDnsARecord
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsARecordSrvList = getSliceGslbZoneServiceListDnsARecordDnsARecordSrvList(in["dns_a_record_srv_list"].([]interface{}))
		ret.DnsARecordIpv4List = getSliceGslbZoneServiceListDnsARecordDnsARecordIpv4List(in["dns_a_record_ipv4_list"].([]interface{}))
		ret.DnsARecordIpv6List = getSliceGslbZoneServiceListDnsARecordDnsARecordIpv6List(in["dns_a_record_ipv6_list"].([]interface{}))
	}
	return ret
}
func getSliceGslbZoneServiceListDnsARecordDnsARecordSrvList(d []interface{}) []edpt.GslbZoneServiceListDnsARecordDnsARecordSrvList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsARecordDnsARecordSrvList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsARecordDnsARecordSrvList
		oi.Svrname = in["svrname"].(string)
		oi.NoResp = in["no_resp"].(int)
		oi.AsBackup = in["as_backup"].(int)
		oi.Weight = in["weight"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.AsReplace = in["as_replace"].(int)
		oi.Disable = in["disable"].(int)
		oi.Static = in["static"].(int)
		oi.AdminIp = in["admin_ip"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsARecordDnsARecordIpv4List(d []interface{}) []edpt.GslbZoneServiceListDnsARecordDnsARecordIpv4List {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsARecordDnsARecordIpv4List, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsARecordDnsARecordIpv4List
		oi.DnsARecordIp = in["dns_a_record_ip"].(string)
		oi.NoResp = in["no_resp"].(int)
		oi.AsBackup = in["as_backup"].(int)
		oi.Weight = in["weight"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.AsReplace = in["as_replace"].(int)
		oi.Disable = in["disable"].(int)
		oi.Static = in["static"].(int)
		oi.AdminIp = in["admin_ip"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsARecordDnsARecordIpv6List(d []interface{}) []edpt.GslbZoneServiceListDnsARecordDnsARecordIpv6List {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsARecordDnsARecordIpv6List, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsARecordDnsARecordIpv6List
		oi.DnsARecordIpv6 = in["dns_a_record_ipv6"].(string)
		oi.NoResp = in["no_resp"].(int)
		oi.AsBackup = in["as_backup"].(int)
		oi.Weight = in["weight"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.AsReplace = in["as_replace"].(int)
		oi.Disable = in["disable"].(int)
		oi.Static = in["static"].(int)
		oi.AdminIp = in["admin_ip"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsCnameRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsCnameRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsCnameRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsCnameRecordList
		oi.AliasName = in["alias_name"].(string)
		oi.AdminPreference = in["admin_preference"].(int)
		oi.Weight = in["weight"].(int)
		oi.AsBackup = in["as_backup"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsCnameRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsCnameRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsCnameRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsCnameRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsCnameRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsMxRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsMxRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsMxRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Priority = in["priority"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsMxRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsMxRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsMxRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsMxRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsMxRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsNsRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsNsRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsNsRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsNsRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsNsRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsNsRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsNsRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsNsRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsPtrRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsPtrRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsPtrRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsPtrRecordList
		oi.PtrName = in["ptr_name"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsPtrRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsPtrRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsPtrRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsPtrRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsPtrRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsSrvRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsSrvRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsSrvRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsSrvRecordList
		oi.SrvName = in["srv_name"].(string)
		oi.Port = in["port"].(int)
		oi.Priority = in["priority"].(int)
		oi.Weight = in["weight"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsSrvRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsSrvRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsSrvRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsSrvRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsSrvRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsNaptrRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsNaptrRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsNaptrRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsNaptrRecordList
		oi.NaptrTarget = in["naptr_target"].(string)
		oi.ServiceProto = in["service_proto"].(string)
		oi.Flag = in["flag"].(string)
		oi.Order = in["order"].(int)
		oi.Preference = in["preference"].(int)
		oi.Regexp = in["regexp"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsNaptrRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsNaptrRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsNaptrRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsNaptrRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsNaptrRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsTxtRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsTxtRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsTxtRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsTxtRecordList
		oi.RecordName = in["record_name"].(string)
		oi.TxtData = in["txt_data"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceListDnsTxtRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsTxtRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceListDnsTxtRecordListSamplingEnable {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsTxtRecordListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsTxtRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListDnsRecordList(d []interface{}) []edpt.GslbZoneServiceListDnsRecordList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListDnsRecordList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListDnsRecordList
		oi.Type = in["type"].(int)
		oi.Data = in["data"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListGeoLocationList(d []interface{}) []edpt.GslbZoneServiceListGeoLocationList {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListGeoLocationList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListGeoLocationList
		oi.GeoName = in["geo_name"].(string)
		oi.Alias = getSliceGslbZoneServiceListGeoLocationListAlias(in["alias"].([]interface{}))
		oi.Action = in["action"].(int)
		oi.ActionType = in["action_type"].(string)
		oi.ForwardType = in["forward_type"].(string)
		oi.Policy = in["policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getSliceGslbZoneServiceListGeoLocationListAlias(d []interface{}) []edpt.GslbZoneServiceListGeoLocationListAlias {
	count := len(d)
	ret := make([]edpt.GslbZoneServiceListGeoLocationListAlias, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceListGeoLocationListAlias
		oi.Alias = in["alias"].(string)
		ret = append(ret, oi)
	}
	return ret
}
func getObjectGslbZoneTemplate(d []interface{}) edpt.GslbZoneTemplate {
	count := len(d)
	var ret edpt.GslbZoneTemplate
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Dnssec = in["dnssec"].(string)
	}
	return ret
}
func dataToEndpointGslbZone(d *schema.ResourceData) edpt.GslbZone {
	var ret edpt.GslbZone
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DnsMxRecordList = getSliceGslbZoneDnsMxRecordList(d.Get("dns_mx_record_list").([]interface{}))
	ret.Inst.DnsNsRecordList = getSliceGslbZoneDnsNsRecordList(d.Get("dns_ns_record_list").([]interface{}))
	ret.Inst.DnsSoaRecord = getObjectGslbZoneDnsSoaRecord(d.Get("dns_soa_record").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Policy = d.Get("policy").(string)
	ret.Inst.SamplingEnable = getSliceGslbZoneSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceList = getSliceGslbZoneServiceList(d.Get("service_list").([]interface{}))
	ret.Inst.Template = getObjectGslbZoneTemplate(d.Get("template").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	ret.Inst.UseServerTtl = d.Get("use_server_ttl").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
