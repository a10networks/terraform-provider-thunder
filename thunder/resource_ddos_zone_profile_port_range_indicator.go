package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfilePortRangeIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_port_range_indicator`: DDoS zone profile indicator config\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfilePortRangeIndicatorCreate,
		UpdateContext: resourceDdosZoneProfilePortRangeIndicatorUpdate,
		ReadContext:   resourceDdosZoneProfilePortRangeIndicatorRead,
		DeleteContext: resourceDdosZoneProfilePortRangeIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"indicator_name": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': pkt-rate; 'pkt-drop-rate': pkt-drop-rate; 'bit-rate': bit-rate; 'pkt-drop-ratio': pkt-drop-ratio; 'bytes-to-bytes-from-ratio': bytes-to-bytes-from-ratio; 'concurrent-conns': concurrent-conns; 'conn-miss-rate': conn-miss-rate; 'syn-rate': syn-rate; 'fin-rate': fin-rate; 'rst-rate': rst-rate; 'small-window-ack-rate': small-window-ack-rate; 'empty-ack-rate': empty-ack-rate; 'small-payload-rate': small-payload-rate; 'syn-fin-ratio': syn-fin-ratio; 'cpu-utilization': cpu-utilization; 'interface-utilization': interface-utilization;",
			},
			"src_threshold_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
						},
						"src_threshold_large_num": {
							Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
						},
						"src_threshold_str": {
							Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
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
			"zone_threshold_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
						},
						"zone_threshold_large_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
						},
						"zone_threshold_str": {
							Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"profile_name": {
				Type: schema.TypeString, Required: true, Description: "ProfileName",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
		},
	}
}
func resourceDdosZoneProfilePortRangeIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRangeIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortRangeIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfilePortRangeIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRangeIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortRangeIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfilePortRangeIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRangeIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfilePortRangeIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRangeIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneProfilePortRangeIndicatorSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortRangeIndicatorSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortRangeIndicatorSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfilePortRangeIndicatorZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortRangeIndicatorZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortRangeIndicatorZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfilePortRangeIndicator(d *schema.ResourceData) edpt.DdosZoneProfilePortRangeIndicator {
	var ret edpt.DdosZoneProfilePortRangeIndicator
	ret.Inst.IndicatorName = d.Get("indicator_name").(string)
	ret.Inst.SrcThresholdCfg = getObjectDdosZoneProfilePortRangeIndicatorSrcThresholdCfg(d.Get("src_threshold_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdCfg = getObjectDdosZoneProfilePortRangeIndicatorZoneThresholdCfg(d.Get("zone_threshold_cfg").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	return ret
}
