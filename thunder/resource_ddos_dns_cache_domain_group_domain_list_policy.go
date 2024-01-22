package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheDomainGroupDomainListPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_domain_group_domain_list_policy`: DNS Cache Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheDomainGroupDomainListPolicyCreate,
		UpdateContext: resourceDdosDnsCacheDomainGroupDomainListPolicyUpdate,
		ReadContext:   resourceDdosDnsCacheDomainGroupDomainListPolicyRead,
		DeleteContext: resourceDdosDnsCacheDomainGroupDomainListPolicyDelete,

		Schema: map[string]*schema.Schema{
			"cache_all_records": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "cache all fqdn records including uncommon types",
			},
			"client_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Client ipv4 address",
			},
			"client_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Client ipv6 address",
			},
			"force": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force update even the serial is the same",
			},
			"manual_refresh": {
				Type: schema.TypeString, Optional: true, Description: "Manually refresh the particular zone",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS domain list policy",
			},
			"oversize_answer_response": {
				Type: schema.TypeString, Optional: true, Default: "set-truncate-bit", Description: "'set-truncate-bit': Set the TC bit for oversize answer(default); 'disable-truncate-bit': Do not set TC bit for oversize answer;",
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
			"refresh_interval_hours": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Zone transfer refresh rate in hours (Default 4). 0 means no refresh",
			},
			"resolve_cname_record": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always try to resolve domain in CNAME record answer section",
			},
			"respond_with_authority": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond with authority section for all requests under this list",
			},
			"server_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Master ipv4 address",
			},
			"server_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Master ipv6 address",
			},
			"server_v4_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port number (default 53)",
			},
			"server_v6_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port number (default 53)",
			},
			"ttl_override": {
				Type: schema.TypeInt, Optional: true, Description: "Override the TTL value for zone transfer",
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
func resourceDdosDnsCacheDomainGroupDomainListPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheDomainGroupDomainListPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheDomainGroupDomainListPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheDomainGroupDomainListPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheDomainGroupDomainListPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheDomainGroupDomainListPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing146(d []interface{}) edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing146 {

	count1 := len(d)
	var ret edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing146
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RootZoneList = getSliceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147(in["root_zone_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147(d []interface{}) []edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147 {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheDomainGroupDomainListPolicy(d *schema.ResourceData) edpt.DdosDnsCacheDomainGroupDomainListPolicy {
	var ret edpt.DdosDnsCacheDomainGroupDomainListPolicy
	ret.Inst.CacheAllRecords = d.Get("cache_all_records").(int)
	ret.Inst.ClientIpv4 = d.Get("client_ipv4").(string)
	ret.Inst.ClientIpv6 = d.Get("client_ipv6").(string)
	ret.Inst.Force = d.Get("force").(int)
	ret.Inst.ManualRefresh = d.Get("manual_refresh").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OversizeAnswerResponse = d.Get("oversize_answer_response").(string)
	ret.Inst.PacketCapturing = getObjectDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing146(d.Get("packet_capturing").([]interface{}))
	ret.Inst.RefreshIntervalHours = d.Get("refresh_interval_hours").(int)
	ret.Inst.ResolveCnameRecord = d.Get("resolve_cname_record").(int)
	ret.Inst.RespondWithAuthority = d.Get("respond_with_authority").(int)
	ret.Inst.ServerIpv4 = d.Get("server_ipv4").(string)
	ret.Inst.ServerIpv6 = d.Get("server_ipv6").(string)
	ret.Inst.ServerV4Port = d.Get("server_v4_port").(int)
	ret.Inst.ServerV6Port = d.Get("server_v6_port").(int)
	ret.Inst.TtlOverride = d.Get("ttl_override").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
