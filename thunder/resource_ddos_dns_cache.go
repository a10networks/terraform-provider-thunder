package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache`: DNS Cache Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheCreate,
		UpdateContext: resourceDdosDnsCacheUpdate,
		ReadContext:   resourceDdosDnsCacheRead,
		DeleteContext: resourceDdosDnsCacheDelete,

		Schema: map[string]*schema.Schema{
			"any_query_action_str": {
				Type: schema.TypeString, Optional: true, Default: "respond-refuse", Description: "'respond-refuse': Send refuse response (default); 'respond-empty': Send empty response; 'drop': Drop the request;",
			},
			"default_serving_action": {
				Type: schema.TypeString, Optional: true, Default: "serve-from-cache", Description: "'serve-from-cache': Serve DNS records; 'forward': Forward to DNS server; 'drop': Drop the request;",
			},
			"domain_group": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "DNS domain group",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"domain_list_policy_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "DNS domain list policy",
									},
									"server_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "Master ipv4 address",
									},
									"server_v4_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port number (default 53)",
									},
									"client_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "Client ipv4 address",
									},
									"server_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Master ipv6 address",
									},
									"server_v6_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port number (default 53)",
									},
									"client_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Client ipv6 address",
									},
									"refresh_interval_hours": {
										Type: schema.TypeInt, Optional: true, Default: 4, Description: "Zone transfer refresh rate in hours (Default 4). 0 means no refresh",
									},
									"ttl_override": {
										Type: schema.TypeInt, Optional: true, Description: "Override the TTL value for zone transfer",
									},
									"respond_with_authority": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond with authority section for all requests under this list",
									},
									"oversize_answer_response": {
										Type: schema.TypeString, Optional: true, Default: "set-truncate-bit", Description: "'set-truncate-bit': Set the TC bit for oversize answer(default); 'disable-truncate-bit': Do not set TC bit for oversize answer;",
									},
									"resolve_cname_record": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always try to resolve domain in CNAME record answer section",
									},
									"manual_refresh": {
										Type: schema.TypeString, Optional: true, Description: "Manually refresh the particular zone",
									},
									"force": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force update even the serial is the same",
									},
									"cache_all_records": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "cache all fqdn records including uncommon types",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"packet_capturing": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"root_zone_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"root_zone": {
																Type: schema.TypeString, Optional: true, Description: "Specify root zone to be captured",
															},
															"capture_config": {
																Type: schema.TypeString, Optional: true, Description: "Capture-config name",
															},
															"capture_mode": {
																Type: schema.TypeString, Optional: true, Description: "'regular': Capture packet anyway; 'capture-on-failure': Capture packet if last XFR was failed;",
															},
														},
													},
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
					},
				},
			},
			"fqdn_manual_override_action_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn_name": {
							Type: schema.TypeString, Required: true, Description: "Specify fqdn name",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'default': Default; 'forward': Forward to DNS server; 'drop': Drop the request; 'serve-from-cache': Serve DNS records;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Cache Instance Name",
			},
			"neg_cache_action_follow_q_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Negative cached response queries counted toward query-rate-threshold",
			},
			"non_authoritative_zone_query_action_str": {
				Type: schema.TypeString, Optional: true, Default: "respond-refuse", Description: "'default': Default action: respond-refuse; 'forward': Forward to DNS server; 'respond-refuse': Send refuse response; 'drop': Drop the request;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-cached-fqdn': total-cached-fqdn; 'total-cached-records': total-cached-records; 'fqdn-a': fqdn-a; 'fqdn-aaaa': fqdn-aaaa; 'fqdn-cname': fqdn-cname; 'fqdn-ns': fqdn-ns; 'fqdn-mx': fqdn-mx; 'fqdn-soa': fqdn-soa; 'fqdn-srv': fqdn-srv; 'fqdn-txt': fqdn-txt; 'fqdn-ptr': fqdn-ptr; 'fqdn-other': fqdn-other; 'fqdn-wildcard': fqdn-wildcard; 'fqdn-delegation': fqdn-delegation; 'shard-size': shard-size; 'resp-ext-size': resp-ext-size; 'a-record': a-record; 'aaaa-record': aaaa-record; 'cname-record': cname-record; 'ns-record': ns-record; 'mx-record': mx-record; 'soa-record': soa-record; 'srv-record': srv-record; 'txt-record': txt-record; 'ptr-record': ptr-record; 'other-record': other-record; 'fqdn-in-shard-filter': fqdn-in-shard-filter;",
						},
					},
				},
			},
			"sharded_domain_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "DNS sharded domain group",
						},
						"match_action": {
							Type: schema.TypeString, Optional: true, Default: "forward", Description: "'forward': Forward query to server (default); 'tunnel-encap': Encapsulate the query and send on a tunnel;",
						},
						"encap_template": {
							Type: schema.TypeString, Optional: true, Description: "DDOS encap template to sepcify the tunnel endpoint",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sharded_domain_list_policy_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "DNS sharded domain list policy",
									},
									"server_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "Master ipv4 address",
									},
									"server_v4_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port number (default 53)",
									},
									"client_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "Client ipv4 address",
									},
									"server_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Master ipv6 address",
									},
									"server_v6_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port number (default 53)",
									},
									"client_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Client ipv6 address",
									},
									"refresh_interval_hours": {
										Type: schema.TypeInt, Optional: true, Default: 4, Description: "Zone transfer refresh rate in hours (Default 4). 0 means no refresh",
									},
									"manual_refresh": {
										Type: schema.TypeString, Optional: true, Description: "Manually refresh the particular zone",
									},
									"force": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force update even the serial is the same",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"packet_capturing": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"root_zone_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"root_zone": {
																Type: schema.TypeString, Optional: true, Description: "Specify root zone to be captured",
															},
															"capture_config": {
																Type: schema.TypeString, Optional: true, Description: "Capture-config name",
															},
															"capture_mode": {
																Type: schema.TypeString, Optional: true, Description: "'regular': Capture packet anyway; 'capture-on-failure': Capture packet if last XFR was failed;",
															},
														},
													},
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
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_domain_lookup_miss_action": {
				Type: schema.TypeString, Optional: true, Default: "respond-nxdomain", Description: "'respond-nxdomain': Send NxDomain response; 'drop': Drop the request;",
			},
			"zone_manual_override_action_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_name": {
							Type: schema.TypeString, Required: true, Description: "Specify zone name",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'default': Default; 'forward': Forward to DNS server; 'drop': Drop the request; 'serve-from-cache': Serve DNS records;",
						},
					},
				},
			},
			"zone_transfer": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceDdosDnsCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDnsCacheDomainGroup150(d []interface{}) edpt.DdosDnsCacheDomainGroup150 {

	count1 := len(d)
	var ret edpt.DdosDnsCacheDomainGroup150
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
		ret.DomainListPolicyList = getSliceDdosDnsCacheDomainGroupDomainListPolicyList151(in["domain_list_policy_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDnsCacheDomainGroupDomainListPolicyList151(d []interface{}) []edpt.DdosDnsCacheDomainGroupDomainListPolicyList151 {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheDomainGroupDomainListPolicyList151, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheDomainGroupDomainListPolicyList151
		oi.Name = in["name"].(string)
		oi.ServerIpv4 = in["server_ipv4"].(string)
		oi.ServerV4Port = in["server_v4_port"].(int)
		oi.ClientIpv4 = in["client_ipv4"].(string)
		oi.ServerIpv6 = in["server_ipv6"].(string)
		oi.ServerV6Port = in["server_v6_port"].(int)
		oi.ClientIpv6 = in["client_ipv6"].(string)
		oi.RefreshIntervalHours = in["refresh_interval_hours"].(int)
		oi.TtlOverride = in["ttl_override"].(int)
		oi.RespondWithAuthority = in["respond_with_authority"].(int)
		oi.OversizeAnswerResponse = in["oversize_answer_response"].(string)
		oi.ResolveCnameRecord = in["resolve_cname_record"].(int)
		oi.ManualRefresh = in["manual_refresh"].(string)
		oi.Force = in["force"].(int)
		oi.CacheAllRecords = in["cache_all_records"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PacketCapturing = getObjectDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing152(in["packet_capturing"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing152(d []interface{}) edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing152 {

	count1 := len(d)
	var ret edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing152
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RootZoneList = getSliceDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList153(in["root_zone_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList153(d []interface{}) []edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList153 {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList153, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList153
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheFqdnManualOverrideActionList(d []interface{}) []edpt.DdosDnsCacheFqdnManualOverrideActionList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheFqdnManualOverrideActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheFqdnManualOverrideActionList
		oi.FqdnName = in["fqdn_name"].(string)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheSamplingEnable(d []interface{}) []edpt.DdosDnsCacheSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheShardedDomainGroupList(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupList
		oi.Name = in["name"].(string)
		oi.MatchAction = in["match_action"].(string)
		oi.EncapTemplate = in["encap_template"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.ShardedDomainListPolicyList = getSliceDdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList(in["sharded_domain_list_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList
		oi.Name = in["name"].(string)
		oi.ServerIpv4 = in["server_ipv4"].(string)
		oi.ServerV4Port = in["server_v4_port"].(int)
		oi.ClientIpv4 = in["client_ipv4"].(string)
		oi.ServerIpv6 = in["server_ipv6"].(string)
		oi.ServerV6Port = in["server_v6_port"].(int)
		oi.ClientIpv6 = in["client_ipv6"].(string)
		oi.RefreshIntervalHours = in["refresh_interval_hours"].(int)
		oi.ManualRefresh = in["manual_refresh"].(string)
		oi.Force = in["force"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PacketCapturing = getObjectDdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturing(in["packet_capturing"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturing(d []interface{}) edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturing {

	count1 := len(d)
	var ret edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturing
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RootZoneList = getSliceDdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList(in["root_zone_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheZoneManualOverrideActionList(d []interface{}) []edpt.DdosDnsCacheZoneManualOverrideActionList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheZoneManualOverrideActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheZoneManualOverrideActionList
		oi.ZoneName = in["zone_name"].(string)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDnsCacheZoneTransfer154(d []interface{}) edpt.DdosDnsCacheZoneTransfer154 {

	var ret edpt.DdosDnsCacheZoneTransfer154
	return ret
}

func dataToEndpointDdosDnsCache(d *schema.ResourceData) edpt.DdosDnsCache {
	var ret edpt.DdosDnsCache
	ret.Inst.AnyQueryActionStr = d.Get("any_query_action_str").(string)
	ret.Inst.DefaultServingAction = d.Get("default_serving_action").(string)
	ret.Inst.DomainGroup = getObjectDdosDnsCacheDomainGroup150(d.Get("domain_group").([]interface{}))
	ret.Inst.FqdnManualOverrideActionList = getSliceDdosDnsCacheFqdnManualOverrideActionList(d.Get("fqdn_manual_override_action_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NegCacheActionFollowQRate = d.Get("neg_cache_action_follow_q_rate").(int)
	ret.Inst.NonAuthoritativeZoneQueryActionStr = d.Get("non_authoritative_zone_query_action_str").(string)
	ret.Inst.SamplingEnable = getSliceDdosDnsCacheSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ShardedDomainGroupList = getSliceDdosDnsCacheShardedDomainGroupList(d.Get("sharded_domain_group_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneDomainLookupMissAction = d.Get("zone_domain_lookup_miss_action").(string)
	ret.Inst.ZoneManualOverrideActionList = getSliceDdosDnsCacheZoneManualOverrideActionList(d.Get("zone_manual_override_action_list").([]interface{}))
	ret.Inst.ZoneTransfer = getObjectDdosDnsCacheZoneTransfer154(d.Get("zone_transfer").([]interface{}))
	return ret
}
