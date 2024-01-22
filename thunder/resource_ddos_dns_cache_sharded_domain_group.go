package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheShardedDomainGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_sharded_domain_group`: DNS Sharded Domain Group\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheShardedDomainGroupCreate,
		UpdateContext: resourceDdosDnsCacheShardedDomainGroupUpdate,
		ReadContext:   resourceDdosDnsCacheShardedDomainGroupRead,
		DeleteContext: resourceDdosDnsCacheShardedDomainGroupDelete,

		Schema: map[string]*schema.Schema{
			"encap_template": {
				Type: schema.TypeString, Optional: true, Description: "DDOS encap template to sepcify the tunnel endpoint",
			},
			"match_action": {
				Type: schema.TypeString, Optional: true, Default: "forward", Description: "'forward': Forward query to server (default); 'tunnel-encap': Encapsulate the query and send on a tunnel;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS sharded domain group",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDnsCacheShardedDomainGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheShardedDomainGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheShardedDomainGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheShardedDomainGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheShardedDomainGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheShardedDomainGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyList(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyList
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
		oi.PacketCapturing = getObjectDdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturing(in["packet_capturing"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturing(d []interface{}) edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturing {

	count1 := len(d)
	var ret edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturing
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RootZoneList = getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList(in["root_zone_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheShardedDomainGroup(d *schema.ResourceData) edpt.DdosDnsCacheShardedDomainGroup {
	var ret edpt.DdosDnsCacheShardedDomainGroup
	ret.Inst.EncapTemplate = d.Get("encap_template").(string)
	ret.Inst.MatchAction = d.Get("match_action").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ShardedDomainListPolicyList = getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyList(d.Get("sharded_domain_list_policy_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
