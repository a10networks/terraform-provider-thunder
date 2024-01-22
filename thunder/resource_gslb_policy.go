package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy`: Policy for GSLB zone, service or geo-location\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyCreate,
		UpdateContext: resourceGslbPolicyUpdate,
		ReadContext:   resourceGslbPolicyRead,
		DeleteContext: resourceGslbPolicyDelete,

		Schema: map[string]*schema.Schema{
			"active_rdt": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the active rdt",
						},
						"single_shot": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Single Shot RDT",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify timeout if round-delay-time samples are not ready (Specify timeout, unit:sec,default is 3)",
						},
						"skip": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Skip query if round-delay-time samples are not ready (Specify maximum skip count,default is 3)",
						},
						"keep_tracking": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep tracking client even round-delay-time samples are ready",
						},
						"ignore_id": {
							Type: schema.TypeInt, Optional: true, Description: "Ignore IP Address specified in IP List by ID",
						},
						"samples": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify samples number for round-delay-time (Number of samples,default is 5)",
						},
						"tolerance": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "The difference percentage between the round-delay-time, default is 10 (Tolerance)",
						},
						"difference": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "The difference between the round-delay-time, default is 0",
						},
						"limit": {
							Type: schema.TypeInt, Optional: true, Default: 16383, Description: "Limit of allowed RDT, default is 16383 (Limit, unit: millisecond)",
						},
						"fail_break": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when no valid RDT",
						},
						"controller": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Active round-delay-time by controller",
						},
						"proto_rdt_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the round-delay-time to the controller",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"active_servers_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Select Service-IP with the highest number of active servers",
			},
			"active_servers_fail_break": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when no active server",
			},
			"admin_ip_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable admin ip",
			},
			"admin_ip_top_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return highest priority server only",
			},
			"admin_preference": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select Service-IP for the device having maximum admin preference",
			},
			"alias_admin_preference": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select alias name having maximum admin preference",
			},
			"amount_first": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select record based on the amount of available service-ip",
			},
			"auto_map": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 300, Description: "Specify Auto Map TTL (TTL, default is 300)",
						},
						"module_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Disable Auto Map Module",
						},
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All modules",
						},
						"module_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"bw_cost_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bw cost",
			},
			"bw_cost_fail_break": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when exceed limit",
			},
			"capacity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capacity_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable capacity",
						},
						"threshold": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "Specify capacity threshold, default is 90",
						},
						"capacity_fail_break": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when exceed threshold",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"connection_load": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_load_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable connection-load",
						},
						"connection_load_fail_break": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when exceed limit",
						},
						"connection_load_samples": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify samples for connection-load (Number of samples used to calculate the connection load, default is 5)",
						},
						"connection_load_interval": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Interval between two samples, Unit: second (Interval value,default is 5)",
						},
						"limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit of maxinum connection load, default is unlimited",
						},
						"connection_load_limit": {
							Type: schema.TypeInt, Optional: true, Description: "The value of the connection-load limit, default is unlimited",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dns": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply DNS action for service",
						},
						"active_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only keep active servers",
						},
						"active_only_fail_safe": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Continue if no candidate",
						},
						"dns_addition_mx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append MX Records in Addition Section",
						},
						"dns_auto_map": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Automatically build DNS Infrastructure",
						},
						"backup_alias": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return alias name when fail",
						},
						"backup_server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return fallback server when fail",
						},
						"external_ip": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Return DNS response with external IP address",
						},
						"external_soa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return DNS response with external SOA Record",
						},
						"cname_detect": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Apply GSLB for DNS Server response when service is Canonical Name (CNAME)",
						},
						"ip_replace": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace DNS Server Response with GSLB Service-IPs",
						},
						"geoloc_alias": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return alias name by geo-location",
						},
						"geoloc_action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply DNS action by geo-location",
						},
						"geoloc_policy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply different policy by geo-location",
						},
						"selected_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only keep selected servers",
						},
						"selected_only_value": {
							Type: schema.TypeInt, Optional: true, Description: "Answer Number",
						},
						"cache": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cache DNS Server response",
						},
						"aging_time": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify aging-time, default is TTL in DNS record, unit: second (Aging time, default 0 means using TTL in DNS record as aging time)",
						},
						"delegation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Zone Delegation",
						},
						"hint": {
							Type: schema.TypeString, Optional: true, Default: "addition", Description: "'none': None; 'answer': Append Hint Records in DNS Answer Section; 'addition': Append Hint Records in DNS Addition Section;",
						},
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "'none': None; 'query': DNS Query; 'response': DNS Response; 'both': Both DNS Query and Response;",
						},
						"template": {
							Type: schema.TypeString, Optional: true, Description: "Logging template (Logging Template Name)",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the TTL value contained in DNS record (TTL value, unit: second, default is 10)",
						},
						"use_server_ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DNS Server Response TTL value in GSLB Proxy mode",
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB as DNS server mode",
						},
						"server_srv": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide SRV Records",
						},
						"server_mx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide MX Records",
						},
						"server_naptr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide NAPTR Records",
						},
						"server_addition_mx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append MX Records in Addition Section",
						},
						"server_ns": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide NS Records",
						},
						"server_auto_ns": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide A-Records for NS-Records automatically",
						},
						"server_ptr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide PTR Records",
						},
						"server_auto_ptr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide PTR Records automatically",
						},
						"server_txt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide TXT Records",
						},
						"server_custom": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide Custom Records",
						},
						"server_any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide All Records",
						},
						"server_any_with_metric": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide All Records with GSLB Metrics applied to A/AAAA Records",
						},
						"server_authoritative": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "As authoritative server",
						},
						"server_sec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide DNSSEC support",
						},
						"server_ns_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append All NS Records in Authoritative Section",
						},
						"server_full_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append All A Records in Authoritative Section",
						},
						"server_mode_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only run GSLB as DNS server mode",
						},
						"zone_owner_mode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only run GSLB as DNS server mode with zone ownership",
						},
						"server_cname": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide CNAME Records",
						},
						"server_caa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide CAA Records",
						},
						"ipv6": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_ipv6_option": {
										Type: schema.TypeString, Optional: true, Description: "'mix': Return both AAAA Record and A Record; 'smart': Return AAAA Record by DNS Query Type; 'mapping': Map A Record to AAAA Record;",
									},
									"dns_ipv6_mapping_type": {
										Type: schema.TypeString, Optional: true, Description: "'addition': Append Mapped Record in DNS Addition Section; 'answer': Append Mapped Record in DNS Answer Section; 'exclusive': Only return AAAA Record; 'replace': Replace Record with Mapped Record;",
									},
								},
							},
						},
						"block_action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Action",
						},
						"action_type": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop query; 'reject': Send refuse response; 'ignore': Send empty response;",
						},
						"proxy_block_port_range_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proxy_block_range_from": {
										Type: schema.TypeInt, Optional: true, Description: "Specify Type Range (From)",
									},
									"proxy_block_range_to": {
										Type: schema.TypeInt, Optional: true, Description: "To",
									},
								},
							},
						},
						"block_value": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"block_value": {
										Type: schema.TypeInt, Optional: true, Description: "Specify Type Number",
									},
								},
							},
						},
						"block_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sticky": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Make DNS Record sticky for certain time",
						},
						"sticky_mask": {
							Type: schema.TypeString, Optional: true, Default: "/32", Description: "Specify IP mask, default is /32",
						},
						"sticky_ipv6_mask": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "Specify IPv6 mask length, default is 128",
						},
						"sticky_aging_time": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify aging-time, unit: min, default is 5 (Aging time)",
						},
						"dynamic_preference": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Make dynamically change the preference",
						},
						"dynamic_weight": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "dynamically change the weight",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"edns": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_subnet_geographic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use client subnet for geo-location",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"geo_location_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify geo-location name, section range is (1-15)",
						},
						"ip_multiple_fields": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_sub": {
										Type: schema.TypeString, Optional: true, Description: "Specify IP information",
									},
									"ip_mask_sub": {
										Type: schema.TypeString, Optional: true, Description: "Specify IP/mask format (Specify IP address mask)",
									},
									"ip_addr2_sub": {
										Type: schema.TypeString, Optional: true, Description: "Specify IP address range",
									},
								},
							},
						},
						"ipv6_multiple_fields": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_sub": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 information",
									},
									"ipv6_mask_sub": {
										Type: schema.TypeInt, Optional: true, Description: "Specify IPv6/mask format (Specify IP address mask)",
									},
									"ipv6_addr2_sub": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address range",
									},
								},
							},
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
			"geo_location_match": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"overlap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable overlap mode to do longest match",
						},
						"geo_type_overlap": {
							Type: schema.TypeString, Optional: true, Description: "'global': Global Geo-location; 'policy': Policy Geo-location;",
						},
						"match_first": {
							Type: schema.TypeString, Optional: true, Default: "global", Description: "'global': Global Geo-location; 'policy': Policy Geo-location;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"geographic": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Select Service-IP by geographic",
			},
			"health_check": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Select Service-IP by health status",
			},
			"health_check_preference_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check health preference",
			},
			"health_preference_top": {
				Type: schema.TypeInt, Optional: true, Description: "Only keep top n",
			},
			"ip_list": {
				Type: schema.TypeString, Optional: true, Description: "Specify IP List (IP List Name)",
			},
			"least_response": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Least response selection",
			},
			"metric_fail_break": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break if no valid Service-IP",
			},
			"metric_force_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always check Service-IP for all enabled metrics",
			},
			"metric_order": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify order of metric",
			},
			"metric_type": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify policy name",
			},
			"num_session_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Select Service-IP for device having maximum number of available sessions",
			},
			"num_session_tolerance": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "The difference between the available sessions, default is 10 (Tolerance)",
			},
			"ordered_ip_top_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return highest priority server only",
			},
			"round_robin": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Round robin selection, enabled by default",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weighted_alias": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select alias name by weighted preference",
			},
			"weighted_ip_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Select Service-IP by weighted preference",
			},
			"weighted_ip_total_hits": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Weighted by total hits",
			},
			"weighted_site_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Select Service-IP by weighted site preference",
			},
			"weighted_site_total_hits": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Weighted by total hits",
			},
		},
	}
}
func resourceGslbPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGslbPolicyActiveRdt384(d []interface{}) edpt.GslbPolicyActiveRdt384 {

	count1 := len(d)
	var ret edpt.GslbPolicyActiveRdt384
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.SingleShot = in["single_shot"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Skip = in["skip"].(int)
		ret.KeepTracking = in["keep_tracking"].(int)
		ret.IgnoreId = in["ignore_id"].(int)
		ret.Samples = in["samples"].(int)
		ret.Tolerance = in["tolerance"].(int)
		ret.Difference = in["difference"].(int)
		ret.Limit = in["limit"].(int)
		ret.FailBreak = in["fail_break"].(int)
		ret.Controller = in["controller"].(int)
		ret.ProtoRdtEnable = in["proto_rdt_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectGslbPolicyAutoMap385(d []interface{}) edpt.GslbPolicyAutoMap385 {

	count1 := len(d)
	var ret edpt.GslbPolicyAutoMap385
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ttl = in["ttl"].(int)
		ret.ModuleDisable = in["module_disable"].(int)
		ret.All = in["all"].(int)
		ret.ModuleType = in["module_type"].(string)
		//omit uuid
	}
	return ret
}

func getObjectGslbPolicyCapacity386(d []interface{}) edpt.GslbPolicyCapacity386 {

	count1 := len(d)
	var ret edpt.GslbPolicyCapacity386
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CapacityEnable = in["capacity_enable"].(int)
		ret.Threshold = in["threshold"].(int)
		ret.CapacityFailBreak = in["capacity_fail_break"].(int)
		//omit uuid
	}
	return ret
}

func getObjectGslbPolicyConnectionLoad387(d []interface{}) edpt.GslbPolicyConnectionLoad387 {

	count1 := len(d)
	var ret edpt.GslbPolicyConnectionLoad387
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectionLoadEnable = in["connection_load_enable"].(int)
		ret.ConnectionLoadFailBreak = in["connection_load_fail_break"].(int)
		ret.ConnectionLoadSamples = in["connection_load_samples"].(int)
		ret.ConnectionLoadInterval = in["connection_load_interval"].(int)
		ret.Limit = in["limit"].(int)
		ret.ConnectionLoadLimit = in["connection_load_limit"].(int)
		//omit uuid
	}
	return ret
}

func getObjectGslbPolicyDns388(d []interface{}) edpt.GslbPolicyDns388 {

	count1 := len(d)
	var ret edpt.GslbPolicyDns388
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(int)
		ret.ActiveOnly = in["active_only"].(int)
		ret.ActiveOnlyFailSafe = in["active_only_fail_safe"].(int)
		ret.DnsAdditionMx = in["dns_addition_mx"].(int)
		ret.DnsAutoMap = in["dns_auto_map"].(int)
		ret.BackupAlias = in["backup_alias"].(int)
		ret.BackupServer = in["backup_server"].(int)
		ret.ExternalIp = in["external_ip"].(int)
		ret.ExternalSoa = in["external_soa"].(int)
		ret.CnameDetect = in["cname_detect"].(int)
		ret.IpReplace = in["ip_replace"].(int)
		ret.GeolocAlias = in["geoloc_alias"].(int)
		ret.GeolocAction = in["geoloc_action"].(int)
		ret.GeolocPolicy = in["geoloc_policy"].(int)
		ret.SelectedOnly = in["selected_only"].(int)
		ret.SelectedOnlyValue = in["selected_only_value"].(int)
		ret.Cache = in["cache"].(int)
		ret.AgingTime = in["aging_time"].(int)
		ret.Delegation = in["delegation"].(int)
		ret.Hint = in["hint"].(string)
		ret.Logging = in["logging"].(string)
		ret.Template = in["template"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.UseServerTtl = in["use_server_ttl"].(int)
		ret.Server = in["server"].(int)
		ret.ServerSrv = in["server_srv"].(int)
		ret.ServerMx = in["server_mx"].(int)
		ret.ServerNaptr = in["server_naptr"].(int)
		ret.ServerAdditionMx = in["server_addition_mx"].(int)
		ret.ServerNs = in["server_ns"].(int)
		ret.ServerAutoNs = in["server_auto_ns"].(int)
		ret.ServerPtr = in["server_ptr"].(int)
		ret.ServerAutoPtr = in["server_auto_ptr"].(int)
		ret.ServerTxt = in["server_txt"].(int)
		ret.ServerCustom = in["server_custom"].(int)
		ret.ServerAny = in["server_any"].(int)
		ret.ServerAnyWithMetric = in["server_any_with_metric"].(int)
		ret.ServerAuthoritative = in["server_authoritative"].(int)
		ret.ServerSec = in["server_sec"].(int)
		ret.ServerNsList = in["server_ns_list"].(int)
		ret.ServerFullList = in["server_full_list"].(int)
		ret.ServerModeOnly = in["server_mode_only"].(int)
		ret.ZoneOwnerMode = in["zone_owner_mode"].(int)
		ret.ServerCname = in["server_cname"].(int)
		ret.ServerCaa = in["server_caa"].(int)
		ret.Ipv6 = getSliceGslbPolicyDnsIpv6389(in["ipv6"].([]interface{}))
		ret.BlockAction = in["block_action"].(int)
		ret.ActionType = in["action_type"].(string)
		ret.ProxyBlockPortRangeList = getSliceGslbPolicyDnsProxyBlockPortRangeList390(in["proxy_block_port_range_list"].([]interface{}))
		ret.BlockValue = getSliceGslbPolicyDnsBlockValue391(in["block_value"].([]interface{}))
		ret.BlockType = in["block_type"].(string)
		ret.Sticky = in["sticky"].(int)
		ret.StickyMask = in["sticky_mask"].(string)
		ret.StickyIpv6Mask = in["sticky_ipv6_mask"].(int)
		ret.StickyAgingTime = in["sticky_aging_time"].(int)
		ret.DynamicPreference = in["dynamic_preference"].(int)
		ret.DynamicWeight = in["dynamic_weight"].(int)
		//omit uuid
	}
	return ret
}

func getSliceGslbPolicyDnsIpv6389(d []interface{}) []edpt.GslbPolicyDnsIpv6389 {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyDnsIpv6389, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyDnsIpv6389
		oi.DnsIpv6Option = in["dns_ipv6_option"].(string)
		oi.DnsIpv6MappingType = in["dns_ipv6_mapping_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyDnsProxyBlockPortRangeList390(d []interface{}) []edpt.GslbPolicyDnsProxyBlockPortRangeList390 {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyDnsProxyBlockPortRangeList390, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyDnsProxyBlockPortRangeList390
		oi.ProxyBlockRangeFrom = in["proxy_block_range_from"].(int)
		oi.ProxyBlockRangeTo = in["proxy_block_range_to"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyDnsBlockValue391(d []interface{}) []edpt.GslbPolicyDnsBlockValue391 {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyDnsBlockValue391, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyDnsBlockValue391
		oi.BlockValue = in["block_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbPolicyEdns392(d []interface{}) edpt.GslbPolicyEdns392 {

	count1 := len(d)
	var ret edpt.GslbPolicyEdns392
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSubnetGeographic = in["client_subnet_geographic"].(int)
		//omit uuid
	}
	return ret
}

func getSliceGslbPolicyGeoLocationList(d []interface{}) []edpt.GslbPolicyGeoLocationList {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyGeoLocationList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyGeoLocationList
		oi.Name = in["name"].(string)
		oi.IpMultipleFields = getSliceGslbPolicyGeoLocationListIpMultipleFields(in["ip_multiple_fields"].([]interface{}))
		oi.Ipv6MultipleFields = getSliceGslbPolicyGeoLocationListIpv6MultipleFields(in["ipv6_multiple_fields"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyGeoLocationListIpMultipleFields(d []interface{}) []edpt.GslbPolicyGeoLocationListIpMultipleFields {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyGeoLocationListIpMultipleFields, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyGeoLocationListIpMultipleFields
		oi.IpSub = in["ip_sub"].(string)
		oi.IpMaskSub = in["ip_mask_sub"].(string)
		oi.IpAddr2Sub = in["ip_addr2_sub"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyGeoLocationListIpv6MultipleFields(d []interface{}) []edpt.GslbPolicyGeoLocationListIpv6MultipleFields {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyGeoLocationListIpv6MultipleFields, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyGeoLocationListIpv6MultipleFields
		oi.Ipv6Sub = in["ipv6_sub"].(string)
		oi.Ipv6MaskSub = in["ipv6_mask_sub"].(int)
		oi.Ipv6Addr2Sub = in["ipv6_addr2_sub"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbPolicyGeoLocationMatch393(d []interface{}) edpt.GslbPolicyGeoLocationMatch393 {

	count1 := len(d)
	var ret edpt.GslbPolicyGeoLocationMatch393
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Overlap = in["overlap"].(int)
		ret.GeoTypeOverlap = in["geo_type_overlap"].(string)
		ret.MatchFirst = in["match_first"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointGslbPolicy(d *schema.ResourceData) edpt.GslbPolicy {
	var ret edpt.GslbPolicy
	ret.Inst.ActiveRdt = getObjectGslbPolicyActiveRdt384(d.Get("active_rdt").([]interface{}))
	ret.Inst.ActiveServersEnable = d.Get("active_servers_enable").(int)
	ret.Inst.ActiveServersFailBreak = d.Get("active_servers_fail_break").(int)
	ret.Inst.AdminIpEnable = d.Get("admin_ip_enable").(int)
	ret.Inst.AdminIpTopOnly = d.Get("admin_ip_top_only").(int)
	ret.Inst.AdminPreference = d.Get("admin_preference").(int)
	ret.Inst.AliasAdminPreference = d.Get("alias_admin_preference").(int)
	ret.Inst.AmountFirst = d.Get("amount_first").(int)
	ret.Inst.AutoMap = getObjectGslbPolicyAutoMap385(d.Get("auto_map").([]interface{}))
	ret.Inst.BwCostEnable = d.Get("bw_cost_enable").(int)
	ret.Inst.BwCostFailBreak = d.Get("bw_cost_fail_break").(int)
	ret.Inst.Capacity = getObjectGslbPolicyCapacity386(d.Get("capacity").([]interface{}))
	ret.Inst.ConnectionLoad = getObjectGslbPolicyConnectionLoad387(d.Get("connection_load").([]interface{}))
	ret.Inst.Dns = getObjectGslbPolicyDns388(d.Get("dns").([]interface{}))
	ret.Inst.Edns = getObjectGslbPolicyEdns392(d.Get("edns").([]interface{}))
	ret.Inst.GeoLocationList = getSliceGslbPolicyGeoLocationList(d.Get("geo_location_list").([]interface{}))
	ret.Inst.GeoLocationMatch = getObjectGslbPolicyGeoLocationMatch393(d.Get("geo_location_match").([]interface{}))
	ret.Inst.Geographic = d.Get("geographic").(int)
	ret.Inst.HealthCheck = d.Get("health_check").(int)
	ret.Inst.HealthCheckPreferenceEnable = d.Get("health_check_preference_enable").(int)
	ret.Inst.HealthPreferenceTop = d.Get("health_preference_top").(int)
	ret.Inst.IpList = d.Get("ip_list").(string)
	ret.Inst.LeastResponse = d.Get("least_response").(int)
	ret.Inst.MetricFailBreak = d.Get("metric_fail_break").(int)
	ret.Inst.MetricForceCheck = d.Get("metric_force_check").(int)
	ret.Inst.MetricOrder = d.Get("metric_order").(int)
	ret.Inst.MetricType = d.Get("metric_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NumSessionEnable = d.Get("num_session_enable").(int)
	ret.Inst.NumSessionTolerance = d.Get("num_session_tolerance").(int)
	ret.Inst.OrderedIpTopOnly = d.Get("ordered_ip_top_only").(int)
	ret.Inst.RoundRobin = d.Get("round_robin").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.WeightedAlias = d.Get("weighted_alias").(int)
	ret.Inst.WeightedIpEnable = d.Get("weighted_ip_enable").(int)
	ret.Inst.WeightedIpTotalHits = d.Get("weighted_ip_total_hits").(int)
	ret.Inst.WeightedSiteEnable = d.Get("weighted_site_enable").(int)
	ret.Inst.WeightedSiteTotalHits = d.Get("weighted_site_total_hits").(int)
	return ret
}
