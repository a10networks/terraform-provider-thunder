package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_port_range`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcPortRangeCreate,
		UpdateContext: resourceDdosDstZoneSrcPortRangeUpdate,
		ReadContext:   resourceDdosDstZoneSrcPortRangeRead,
		DeleteContext: resourceDdosDstZoneSrcPortRangeDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capture_config_name": {
							Type: schema.TypeString, Optional: true, Description: "Capture-config name",
						},
						"capture_config_mode": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Apply capture-config to dropped packets; 'forward': Apply capture-config to forwarded packets; 'all': Apply capture-config to both dropped and forwarded packets;",
						},
					},
				},
			},
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
										Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the port-range",
									},
									"zone_threshold_large_num": {
										Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the port-range",
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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"src_port_range_end": {
				Type: schema.TypeInt, Required: true, Description: "Src Port-Range End Port Number",
			},
			"src_port_range_start": {
				Type: schema.TypeInt, Required: true, Description: "Src Port-Range Start Port Number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosDstZoneSrcPortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcPortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcPortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcPortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneSrcPortRangeCaptureConfig(d []interface{}) edpt.DdosDstZoneSrcPortRangeCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangeCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortRangeGlidCfg(d []interface{}) edpt.DdosDstZoneSrcPortRangeGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangeGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortRangeLevelList(d []interface{}) []edpt.DdosDstZoneSrcPortRangeLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangeLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangeLevelList
		oi.LevelNum = in["level_num"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneSrcPortRangeLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortRangeLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortRangeLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangeLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangeLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortRangePortInd241(d []interface{}) edpt.DdosDstZoneSrcPortRangePortInd241 {

	var ret edpt.DdosDstZoneSrcPortRangePortInd241
	return ret
}

func getObjectDdosDstZoneSrcPortRangeZoneTemplate(d []interface{}) edpt.DdosDstZoneSrcPortRangeZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangeZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortRange(d *schema.ResourceData) edpt.DdosDstZoneSrcPortRange {
	var ret edpt.DdosDstZoneSrcPortRange
	ret.Inst.CaptureConfig = getObjectDdosDstZoneSrcPortRangeCaptureConfig(d.Get("capture_config").([]interface{}))
	ret.Inst.DefaultActionList = d.Get("default_action_list").(string)
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZoneSrcPortRangeGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.LevelList = getSliceDdosDstZoneSrcPortRangeLevelList(d.Get("level_list").([]interface{}))
	ret.Inst.PortInd = getObjectDdosDstZoneSrcPortRangePortInd241(d.Get("port_ind").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.SrcPortRangeEnd = d.Get("src_port_range_end").(int)
	ret.Inst.SrcPortRangeStart = d.Get("src_port_range_start").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneSrcPortRangeZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
