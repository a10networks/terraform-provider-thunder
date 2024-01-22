package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_port_zone_src_port`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcPortZoneSrcPortCreate,
		UpdateContext: resourceDdosDstZoneSrcPortZoneSrcPortUpdate,
		ReadContext:   resourceDdosDstZoneSrcPortZoneSrcPortRead,
		DeleteContext: resourceDdosDstZoneSrcPortZoneSrcPortDelete,

		Schema: map[string]*schema.Schema{
			"default_action_list": {
				Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
			},
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
						},
					},
				},
			},
			"level_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level_num": {
							Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"indicator_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'bit-rate': rate of incoming bits;",
									},
									"zone_threshold_num": {
										Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
									},
									"zone_threshold_large_num": {
										Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
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
					},
				},
			},
			"outbound_src_tracking": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
			},
			"port_ind": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Source Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-udp': DNS-UDP Port; 'dns-tcp': DNS-TCP Port; 'udp': UDP port; 'tcp': TCP Port;",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_udp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS udp src template",
						},
						"src_tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS tcp src template",
						},
						"src_dns": {
							Type: schema.TypeString, Optional: true, Description: "DDOS dns src template",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneSrcPortZoneSrcPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcPortZoneSrcPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneSrcPortZoneSrcPortGlidCfg(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortLevelList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortLevelList
		oi.LevelNum = in["level_num"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortPortInd243(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortPortInd243 {

	var ret edpt.DdosDstZoneSrcPortZoneSrcPortPortInd243
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortZoneTemplate(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
		ret.SrcDns = in["src_dns"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPort(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPort {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPort
	ret.Inst.DefaultActionList = d.Get("default_action_list").(string)
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZoneSrcPortZoneSrcPortGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.LevelList = getSliceDdosDstZoneSrcPortZoneSrcPortLevelList(d.Get("level_list").([]interface{}))
	ret.Inst.OutboundSrcTracking = d.Get("outbound_src_tracking").(string)
	ret.Inst.PortInd = getObjectDdosDstZoneSrcPortZoneSrcPortPortInd243(d.Get("port_ind").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneSrcPortZoneSrcPortZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
