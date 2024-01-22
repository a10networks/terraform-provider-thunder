package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_dns`: DNS related policy\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyDnsCreate,
		UpdateContext: resourceGslbPolicyDnsUpdate,
		ReadContext:   resourceGslbPolicyDnsRead,
		DeleteContext: resourceGslbPolicyDnsDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply DNS action for service",
			},
			"action_type": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop query; 'reject': Send refuse response; 'ignore': Send empty response;",
			},
			"active_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only keep active servers",
			},
			"active_only_fail_safe": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Continue if no candidate",
			},
			"aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify aging-time, default is TTL in DNS record, unit: second (Aging time, default 0 means using TTL in DNS record as aging time)",
			},
			"backup_alias": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return alias name when fail",
			},
			"backup_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return fallback server when fail",
			},
			"block_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Action",
			},
			"block_type": {
				Type: schema.TypeString, Optional: true, Description: "",
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
			"cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cache DNS Server response",
			},
			"cname_detect": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Apply GSLB for DNS Server response when service is Canonical Name (CNAME)",
			},
			"delegation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Zone Delegation",
			},
			"dns_addition_mx": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append MX Records in Addition Section",
			},
			"dns_auto_map": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Automatically build DNS Infrastructure",
			},
			"dynamic_preference": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Make dynamically change the preference",
			},
			"dynamic_weight": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "dynamically change the weight",
			},
			"external_ip": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Return DNS response with external IP address",
			},
			"external_soa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return DNS response with external SOA Record",
			},
			"geoloc_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply DNS action by geo-location",
			},
			"geoloc_alias": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return alias name by geo-location",
			},
			"geoloc_policy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply different policy by geo-location",
			},
			"hint": {
				Type: schema.TypeString, Optional: true, Default: "addition", Description: "'none': None; 'answer': Append Hint Records in DNS Answer Section; 'addition': Append Hint Records in DNS Addition Section;",
			},
			"ip_replace": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace DNS Server Response with GSLB Service-IPs",
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
			"logging": {
				Type: schema.TypeString, Optional: true, Description: "'none': None; 'query': DNS Query; 'response': DNS Response; 'both': Both DNS Query and Response;",
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
			"selected_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only keep selected servers",
			},
			"selected_only_value": {
				Type: schema.TypeInt, Optional: true, Description: "Answer Number",
			},
			"server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB as DNS server mode",
			},
			"server_addition_mx": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append MX Records in Addition Section",
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
			"server_auto_ns": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide A-Records for NS-Records automatically",
			},
			"server_auto_ptr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide PTR Records automatically",
			},
			"server_caa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide CAA Records",
			},
			"server_cname": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide CNAME Records",
			},
			"server_custom": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide Custom Records",
			},
			"server_full_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append All A Records in Authoritative Section",
			},
			"server_mode_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only run GSLB as DNS server mode",
			},
			"server_mx": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide MX Records",
			},
			"server_naptr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide NAPTR Records",
			},
			"server_ns": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide NS Records",
			},
			"server_ns_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append All NS Records in Authoritative Section",
			},
			"server_ptr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide PTR Records",
			},
			"server_sec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide DNSSEC support",
			},
			"server_srv": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide SRV Records",
			},
			"server_txt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Provide TXT Records",
			},
			"sticky": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Make DNS Record sticky for certain time",
			},
			"sticky_aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify aging-time, unit: min, default is 5 (Aging time)",
			},
			"sticky_ipv6_mask": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Specify IPv6 mask length, default is 128",
			},
			"sticky_mask": {
				Type: schema.TypeString, Optional: true, Default: "/32", Description: "Specify IP mask, default is /32",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_owner_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only run GSLB as DNS server mode with zone ownership",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbPolicyDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbPolicyDnsBlockValue(d []interface{}) []edpt.GslbPolicyDnsBlockValue {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyDnsBlockValue, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyDnsBlockValue
		oi.BlockValue = in["block_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyDnsIpv6(d []interface{}) []edpt.GslbPolicyDnsIpv6 {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyDnsIpv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyDnsIpv6
		oi.DnsIpv6Option = in["dns_ipv6_option"].(string)
		oi.DnsIpv6MappingType = in["dns_ipv6_mapping_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyDnsProxyBlockPortRangeList(d []interface{}) []edpt.GslbPolicyDnsProxyBlockPortRangeList {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyDnsProxyBlockPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyDnsProxyBlockPortRangeList
		oi.ProxyBlockRangeFrom = in["proxy_block_range_from"].(int)
		oi.ProxyBlockRangeTo = in["proxy_block_range_to"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbPolicyDns(d *schema.ResourceData) edpt.GslbPolicyDns {
	var ret edpt.GslbPolicyDns
	ret.Inst.Action = d.Get("action").(int)
	ret.Inst.ActionType = d.Get("action_type").(string)
	ret.Inst.ActiveOnly = d.Get("active_only").(int)
	ret.Inst.ActiveOnlyFailSafe = d.Get("active_only_fail_safe").(int)
	ret.Inst.AgingTime = d.Get("aging_time").(int)
	ret.Inst.BackupAlias = d.Get("backup_alias").(int)
	ret.Inst.BackupServer = d.Get("backup_server").(int)
	ret.Inst.BlockAction = d.Get("block_action").(int)
	ret.Inst.BlockType = d.Get("block_type").(string)
	ret.Inst.BlockValue = getSliceGslbPolicyDnsBlockValue(d.Get("block_value").([]interface{}))
	ret.Inst.Cache = d.Get("cache").(int)
	ret.Inst.CnameDetect = d.Get("cname_detect").(int)
	ret.Inst.Delegation = d.Get("delegation").(int)
	ret.Inst.DnsAdditionMx = d.Get("dns_addition_mx").(int)
	ret.Inst.DnsAutoMap = d.Get("dns_auto_map").(int)
	ret.Inst.DynamicPreference = d.Get("dynamic_preference").(int)
	ret.Inst.DynamicWeight = d.Get("dynamic_weight").(int)
	ret.Inst.ExternalIp = d.Get("external_ip").(int)
	ret.Inst.ExternalSoa = d.Get("external_soa").(int)
	ret.Inst.GeolocAction = d.Get("geoloc_action").(int)
	ret.Inst.GeolocAlias = d.Get("geoloc_alias").(int)
	ret.Inst.GeolocPolicy = d.Get("geoloc_policy").(int)
	ret.Inst.Hint = d.Get("hint").(string)
	ret.Inst.IpReplace = d.Get("ip_replace").(int)
	ret.Inst.Ipv6 = getSliceGslbPolicyDnsIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.ProxyBlockPortRangeList = getSliceGslbPolicyDnsProxyBlockPortRangeList(d.Get("proxy_block_port_range_list").([]interface{}))
	ret.Inst.SelectedOnly = d.Get("selected_only").(int)
	ret.Inst.SelectedOnlyValue = d.Get("selected_only_value").(int)
	ret.Inst.Server = d.Get("server").(int)
	ret.Inst.ServerAdditionMx = d.Get("server_addition_mx").(int)
	ret.Inst.ServerAny = d.Get("server_any").(int)
	ret.Inst.ServerAnyWithMetric = d.Get("server_any_with_metric").(int)
	ret.Inst.ServerAuthoritative = d.Get("server_authoritative").(int)
	ret.Inst.ServerAutoNs = d.Get("server_auto_ns").(int)
	ret.Inst.ServerAutoPtr = d.Get("server_auto_ptr").(int)
	ret.Inst.ServerCaa = d.Get("server_caa").(int)
	ret.Inst.ServerCname = d.Get("server_cname").(int)
	ret.Inst.ServerCustom = d.Get("server_custom").(int)
	ret.Inst.ServerFullList = d.Get("server_full_list").(int)
	ret.Inst.ServerModeOnly = d.Get("server_mode_only").(int)
	ret.Inst.ServerMx = d.Get("server_mx").(int)
	ret.Inst.ServerNaptr = d.Get("server_naptr").(int)
	ret.Inst.ServerNs = d.Get("server_ns").(int)
	ret.Inst.ServerNsList = d.Get("server_ns_list").(int)
	ret.Inst.ServerPtr = d.Get("server_ptr").(int)
	ret.Inst.ServerSec = d.Get("server_sec").(int)
	ret.Inst.ServerSrv = d.Get("server_srv").(int)
	ret.Inst.ServerTxt = d.Get("server_txt").(int)
	ret.Inst.Sticky = d.Get("sticky").(int)
	ret.Inst.StickyAgingTime = d.Get("sticky_aging_time").(int)
	ret.Inst.StickyIpv6Mask = d.Get("sticky_ipv6_mask").(int)
	ret.Inst.StickyMask = d.Get("sticky_mask").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Ttl = d.Get("ttl").(int)
	ret.Inst.UseServerTtl = d.Get("use_server_ttl").(int)
	//omit uuid
	ret.Inst.ZoneOwnerMode = d.Get("zone_owner_mode").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
