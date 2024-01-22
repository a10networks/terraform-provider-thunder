package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_sharded_domain_group_sharded_domain_list_policy`: DNS Sharded Domain List\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyCreate,
		UpdateContext: resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyUpdate,
		ReadContext:   resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyRead,
		DeleteContext: resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyDelete,

		Schema: map[string]*schema.Schema{
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
				Type: schema.TypeString, Required: true, Description: "DNS sharded domain list policy",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing148(d []interface{}) edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing148 {

	count1 := len(d)
	var ret edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing148
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RootZoneList = getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149(in["root_zone_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149 {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicy(d *schema.ResourceData) edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicy {
	var ret edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicy
	ret.Inst.ClientIpv4 = d.Get("client_ipv4").(string)
	ret.Inst.ClientIpv6 = d.Get("client_ipv6").(string)
	ret.Inst.Force = d.Get("force").(int)
	ret.Inst.ManualRefresh = d.Get("manual_refresh").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCapturing = getObjectDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing148(d.Get("packet_capturing").([]interface{}))
	ret.Inst.RefreshIntervalHours = d.Get("refresh_interval_hours").(int)
	ret.Inst.ServerIpv4 = d.Get("server_ipv4").(string)
	ret.Inst.ServerIpv6 = d.Get("server_ipv6").(string)
	ret.Inst.ServerV4Port = d.Get("server_v4_port").(int)
	ret.Inst.ServerV6Port = d.Get("server_v6_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
