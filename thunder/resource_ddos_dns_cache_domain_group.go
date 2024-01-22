package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheDomainGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_domain_group`: DNS Cache Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheDomainGroupCreate,
		UpdateContext: resourceDdosDnsCacheDomainGroupUpdate,
		ReadContext:   resourceDdosDnsCacheDomainGroupRead,
		DeleteContext: resourceDdosDnsCacheDomainGroupDelete,

		Schema: map[string]*schema.Schema{
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS domain group",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDnsCacheDomainGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheDomainGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheDomainGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheDomainGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheDomainGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheDomainGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDnsCacheDomainGroupDomainListPolicyList(d []interface{}) []edpt.DdosDnsCacheDomainGroupDomainListPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheDomainGroupDomainListPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheDomainGroupDomainListPolicyList
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
		oi.PacketCapturing = getObjectDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing(in["packet_capturing"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing(d []interface{}) edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing {

	count1 := len(d)
	var ret edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RootZoneList = getSliceDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList(in["root_zone_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList(d []interface{}) []edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheDomainGroup(d *schema.ResourceData) edpt.DdosDnsCacheDomainGroup {
	var ret edpt.DdosDnsCacheDomainGroup
	ret.Inst.DomainListPolicyList = getSliceDdosDnsCacheDomainGroupDomainListPolicyList(d.Get("domain_list_policy_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
