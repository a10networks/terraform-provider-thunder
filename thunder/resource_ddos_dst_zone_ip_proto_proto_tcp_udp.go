package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoTcpUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_tcp_udp`: DDOS IP protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoTcpUdpCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoTcpUdpUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoTcpUdpRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoTcpUdpDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for this ip-proto",
			},
			"drop_frag_pkt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
			},
			"glid_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
						},
						"glid_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
						},
						"per_addr_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
						},
						"action_list": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
					},
				},
			},
			"ip_filtering_policy": {
				Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
			},
			"ip_filtering_policy_oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': ip-proto tcp; 'udp': ip-proto udp;",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoTcpUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoTcpUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoTcpUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoTcpUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoTcpUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoTcpUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoTcpUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoTcpUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoTcpUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoTcpUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoTcpUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoTcpUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoTcpUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoTcpUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpGlidCfg(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
		ret.ActionList = in["action_list"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOper211(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOper211 {

	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOper211
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoTcpUdp(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoTcpUdp {
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdp
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZoneIpProtoProtoTcpUdpGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOper211(d.Get("ip_filtering_policy_oper").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
