package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPortOther() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_port_zone_src_port_other`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherCreate,
		UpdateContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherUpdate,
		ReadContext:   resourceDdosDstZoneSrcPortZoneSrcPortOtherRead,
		DeleteContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherDelete,

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
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "'other': other;",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
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
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneSrcPortZoneSrcPortOtherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOther(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortOtherRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortOtherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOther(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortOtherRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcPortZoneSrcPortOtherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOther(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortOtherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOther(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherGlidCfg(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherLevelList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelList
		oi.LevelNum = in["level_num"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherPortInd242(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortInd242 {

	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortInd242
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherZoneTemplate(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPortOther(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPortOther {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOther
	ret.Inst.DefaultActionList = d.Get("default_action_list").(string)
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZoneSrcPortZoneSrcPortOtherGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.LevelList = getSliceDdosDstZoneSrcPortZoneSrcPortOtherLevelList(d.Get("level_list").([]interface{}))
	ret.Inst.PortInd = getObjectDdosDstZoneSrcPortZoneSrcPortOtherPortInd242(d.Get("port_ind").([]interface{}))
	ret.Inst.PortOther = d.Get("port_other").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneSrcPortZoneSrcPortOtherZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
