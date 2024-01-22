package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service`: Service information for the GSLB zone\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceCreate,
		UpdateContext: resourceGslbZoneServiceUpdate,
		ReadContext:   resourceGslbZoneServiceRead,
		DeleteContext: resourceGslbZoneServiceDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop query; 'forward': Forward packet; 'ignore': Send empty response; 'reject': Send refuse response;",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable",
			},
			"dns_a_record": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_a_record_srv_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"svrname": {
										Type: schema.TypeString, Required: true, Description: "Specify name",
									},
									"no_resp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't use this Service-IP as DNS response",
									},
									"as_backup": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "As backup when fail",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "Specify weight for Service-IP (Weight value)",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL for Service-IP",
									},
									"as_replace": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP when enable ip-replace",
									},
									"disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this Service-IP",
									},
									"static": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP in DNS server mode",
									},
									"admin_ip": {
										Type: schema.TypeInt, Optional: true, Description: "Specify admin priority of Service-IP (Specify the priority)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"dns_a_record_ipv4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_a_record_ip": {
										Type: schema.TypeString, Required: true, Description: "Specify IP address",
									},
									"no_resp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't use this Service-IP as DNS response",
									},
									"as_backup": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "As backup when fail",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "Specify weight for Service-IP (Weight value)",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL for Service-IP",
									},
									"as_replace": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP when enable ip-replace",
									},
									"disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this Service-IP",
									},
									"static": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP in DNS server mode",
									},
									"admin_ip": {
										Type: schema.TypeInt, Optional: true, Description: "Specify admin priority of Service-IP (Specify the priority)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"dns_a_record_ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_a_record_ipv6": {
										Type: schema.TypeString, Required: true, Description: "IPV6 address",
									},
									"no_resp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't use this Service-IP as DNS response",
									},
									"as_backup": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "As backup when fail",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "Specify weight for Service-IP (Weight value)",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL for Service-IP",
									},
									"as_replace": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP when enable ip-replace",
									},
									"disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this Service-IP",
									},
									"static": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP in DNS server mode",
									},
									"admin_ip": {
										Type: schema.TypeInt, Optional: true, Description: "Specify admin priority of Service-IP (Specify the priority)",
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
			"dns_caa_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical_flag": {
							Type: schema.TypeInt, Required: true, Description: "Issuer Critical Flag",
						},
						"property_tag": {
							Type: schema.TypeString, Required: true, Description: "Specify other property tags, only allowed lowercase alphanumeric",
						},
						"rdata": {
							Type: schema.TypeString, Required: true, Description: "Specify the Issuer Domain Name or a URL",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the CAA has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_cname_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the alias name",
						},
						"admin_preference": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Specify Administrative Preference, default is 100",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify Weight, default is 1",
						},
						"as_backup": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "As backup when fail",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'cname-hits': Number of times the CNAME has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_mx_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mx_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Priority",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_naptr_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"naptr_target": {
							Type: schema.TypeString, Required: true, Description: "Specify the replacement or regular expression",
						},
						"service_proto": {
							Type: schema.TypeString, Required: true, Description: "Specify Service and Protocol",
						},
						"flag": {
							Type: schema.TypeString, Required: true, Description: "Specify the flag (e.g., a, s). Default is empty flag",
						},
						"order": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Order",
						},
						"preference": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Preference",
						},
						"regexp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return the regular expression",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'naptr-hits': Number of times the NAPTR has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_ns_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ns_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_ptr_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ptr_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeInt, Required: true, Description: "Specify DNS Type",
						},
						"data": {
							Type: schema.TypeString, Optional: true, Description: "Specify DNS Data",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dns_srv_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"srv_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Specify Port (Port Number)",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Priority",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify Weight, default is 10",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"dns_txt_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"record_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the Object Name for TXT Data",
						},
						"txt_data": {
							Type: schema.TypeString, Optional: true, Description: "Specify TXT Data",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the record has been used;",
									},
								},
							},
						},
					},
				},
			},
			"forward_type": {
				Type: schema.TypeString, Optional: true, Description: "'both': Forward both query and response; 'query': Forward query; 'response': Forward response;",
			},
			"geo_location_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the geo-location",
						},
						"alias": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alias": {
										Type: schema.TypeString, Optional: true, Description: "Send CNAME response for this geo-location (Specify a CNAME record)",
									},
								},
							},
						},
						"action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action for this geo-location",
						},
						"action_type": {
							Type: schema.TypeString, Optional: true, Description: "'allow': Allow query from this geo-location; 'drop': Drop query from this geo-location; 'forward': Forward packet for this geo-location; 'ignore': Send empty response to this geo-location; 'reject': Send refuse response to this geo-location;",
						},
						"forward_type": {
							Type: schema.TypeString, Optional: true, Description: "'both': Forward both query and response; 'query': Forward query from this geo-location; 'response': Forward response to this geo-location;",
						},
						"policy": {
							Type: schema.TypeString, Optional: true, Description: "Policy for this geo-location (Specify the policy name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"health_check_gateway": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Gateway Status Check; 'disable': Disable Gateway Status Check;",
			},
			"health_check_port": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check_port": {
							Type: schema.TypeInt, Optional: true, Description: "Check Related Port Status (Port Number)",
						},
					},
				},
			},
			"policy": {
				Type: schema.TypeString, Optional: true, Description: "Specify policy for this service (Specify policy name)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'received-query': Number of DNS queries received for the service; 'sent-response': Number of DNS replies sent to clients for the service; 'proxy-mode-response': Number of DNS replies sent to clients by the ACOS device as a DNS proxy for the service; 'cache-mode-response': Number of cached DNS replies sent to clients by the ACOS device for the service. (This statistic applies only if the DNS cache; 'server-mode-response': Number of DNS replies sent to clients by the ACOS device as a DNS server for the service. (This statistic applies only if the D; 'sticky-mode-response': Number of DNS replies sent to clients by the ACOS device to keep the clients on the same site. (This statistic applies only if; 'backup-mode-response': help Number of DNS replies sent to clients by the ACOS device in backup mode;",
						},
					},
				},
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the service name for the zone, * for wildcard",
			},
			"service_port": {
				Type: schema.TypeInt, Required: true, Description: "Port number of the service",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbZoneServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGslbZoneServiceDnsARecord399(d []interface{}) edpt.GslbZoneServiceDnsARecord399 {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsARecord399
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsARecordSrvList = getSliceGslbZoneServiceDnsARecordDnsARecordSrvList(in["dns_a_record_srv_list"].([]interface{}))
		ret.DnsARecordIpv4List = getSliceGslbZoneServiceDnsARecordDnsARecordIpv4List(in["dns_a_record_ipv4_list"].([]interface{}))
		ret.DnsARecordIpv6List = getSliceGslbZoneServiceDnsARecordDnsARecordIpv6List(in["dns_a_record_ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbZoneServiceDnsARecordDnsARecordSrvList(d []interface{}) []edpt.GslbZoneServiceDnsARecordDnsARecordSrvList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsARecordDnsARecordSrvList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsARecordDnsARecordSrvList
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

func getSliceGslbZoneServiceDnsARecordDnsARecordIpv4List(d []interface{}) []edpt.GslbZoneServiceDnsARecordDnsARecordIpv4List {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsARecordDnsARecordIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsARecordDnsARecordIpv4List
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

func getSliceGslbZoneServiceDnsARecordDnsARecordIpv6List(d []interface{}) []edpt.GslbZoneServiceDnsARecordDnsARecordIpv6List {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsARecordDnsARecordIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsARecordDnsARecordIpv6List
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

func getSliceGslbZoneServiceDnsCaaRecordList(d []interface{}) []edpt.GslbZoneServiceDnsCaaRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsCaaRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsCaaRecordList
		oi.CriticalFlag = in["critical_flag"].(int)
		oi.PropertyTag = in["property_tag"].(string)
		oi.Rdata = in["rdata"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsCaaRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsCaaRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsCaaRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsCaaRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsCaaRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsCnameRecordList(d []interface{}) []edpt.GslbZoneServiceDnsCnameRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsCnameRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsCnameRecordList
		oi.AliasName = in["alias_name"].(string)
		oi.AdminPreference = in["admin_preference"].(int)
		oi.Weight = in["weight"].(int)
		oi.AsBackup = in["as_backup"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsCnameRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsCnameRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsCnameRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsCnameRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsCnameRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsMxRecordList(d []interface{}) []edpt.GslbZoneServiceDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Priority = in["priority"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsMxRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsMxRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsMxRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsMxRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsMxRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsNaptrRecordList(d []interface{}) []edpt.GslbZoneServiceDnsNaptrRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsNaptrRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsNaptrRecordList
		oi.NaptrTarget = in["naptr_target"].(string)
		oi.ServiceProto = in["service_proto"].(string)
		oi.Flag = in["flag"].(string)
		oi.Order = in["order"].(int)
		oi.Preference = in["preference"].(int)
		oi.Regexp = in["regexp"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsNaptrRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsNaptrRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsNaptrRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsNaptrRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsNaptrRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsNsRecordList(d []interface{}) []edpt.GslbZoneServiceDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsNsRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsNsRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsNsRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsNsRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsNsRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsPtrRecordList(d []interface{}) []edpt.GslbZoneServiceDnsPtrRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsPtrRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsPtrRecordList
		oi.PtrName = in["ptr_name"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsPtrRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsPtrRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsPtrRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsPtrRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsPtrRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsRecordList(d []interface{}) []edpt.GslbZoneServiceDnsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsRecordList
		oi.Type = in["type"].(int)
		oi.Data = in["data"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsSrvRecordList(d []interface{}) []edpt.GslbZoneServiceDnsSrvRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsSrvRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsSrvRecordList
		oi.SrvName = in["srv_name"].(string)
		oi.Port = in["port"].(int)
		oi.Priority = in["priority"].(int)
		oi.Weight = in["weight"].(int)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsSrvRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsSrvRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsSrvRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsSrvRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsSrvRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsTxtRecordList(d []interface{}) []edpt.GslbZoneServiceDnsTxtRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsTxtRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsTxtRecordList
		oi.RecordName = in["record_name"].(string)
		oi.TxtData = in["txt_data"].(string)
		oi.Ttl = in["ttl"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceGslbZoneServiceDnsTxtRecordListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceDnsTxtRecordListSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsTxtRecordListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsTxtRecordListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsTxtRecordListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceGeoLocationList(d []interface{}) []edpt.GslbZoneServiceGeoLocationList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceGeoLocationList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceGeoLocationList
		oi.GeoName = in["geo_name"].(string)
		oi.Alias = getSliceGslbZoneServiceGeoLocationListAlias(in["alias"].([]interface{}))
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

func getSliceGslbZoneServiceGeoLocationListAlias(d []interface{}) []edpt.GslbZoneServiceGeoLocationListAlias {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceGeoLocationListAlias, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceGeoLocationListAlias
		oi.Alias = in["alias"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceHealthCheckPort(d []interface{}) []edpt.GslbZoneServiceHealthCheckPort {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceHealthCheckPort, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceHealthCheckPort
		oi.HealthCheckPort = in["health_check_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneServiceSamplingEnable(d []interface{}) []edpt.GslbZoneServiceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneService(d *schema.ResourceData) edpt.GslbZoneService {
	var ret edpt.GslbZoneService
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DnsARecord = getObjectGslbZoneServiceDnsARecord399(d.Get("dns_a_record").([]interface{}))
	ret.Inst.DnsCaaRecordList = getSliceGslbZoneServiceDnsCaaRecordList(d.Get("dns_caa_record_list").([]interface{}))
	ret.Inst.DnsCnameRecordList = getSliceGslbZoneServiceDnsCnameRecordList(d.Get("dns_cname_record_list").([]interface{}))
	ret.Inst.DnsMxRecordList = getSliceGslbZoneServiceDnsMxRecordList(d.Get("dns_mx_record_list").([]interface{}))
	ret.Inst.DnsNaptrRecordList = getSliceGslbZoneServiceDnsNaptrRecordList(d.Get("dns_naptr_record_list").([]interface{}))
	ret.Inst.DnsNsRecordList = getSliceGslbZoneServiceDnsNsRecordList(d.Get("dns_ns_record_list").([]interface{}))
	ret.Inst.DnsPtrRecordList = getSliceGslbZoneServiceDnsPtrRecordList(d.Get("dns_ptr_record_list").([]interface{}))
	ret.Inst.DnsRecordList = getSliceGslbZoneServiceDnsRecordList(d.Get("dns_record_list").([]interface{}))
	ret.Inst.DnsSrvRecordList = getSliceGslbZoneServiceDnsSrvRecordList(d.Get("dns_srv_record_list").([]interface{}))
	ret.Inst.DnsTxtRecordList = getSliceGslbZoneServiceDnsTxtRecordList(d.Get("dns_txt_record_list").([]interface{}))
	ret.Inst.ForwardType = d.Get("forward_type").(string)
	ret.Inst.GeoLocationList = getSliceGslbZoneServiceGeoLocationList(d.Get("geo_location_list").([]interface{}))
	ret.Inst.HealthCheckGateway = d.Get("health_check_gateway").(string)
	ret.Inst.HealthCheckPort = getSliceGslbZoneServiceHealthCheckPort(d.Get("health_check_port").([]interface{}))
	ret.Inst.Policy = d.Get("policy").(string)
	ret.Inst.SamplingEnable = getSliceGslbZoneServiceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
