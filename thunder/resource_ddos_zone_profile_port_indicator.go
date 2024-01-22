package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfilePortIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_port_indicator`: DDoS zone profile indicator config\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfilePortIndicatorCreate,
		UpdateContext: resourceDdosZoneProfilePortIndicatorUpdate,
		ReadContext:   resourceDdosZoneProfilePortIndicatorRead,
		DeleteContext: resourceDdosZoneProfilePortIndicatorDelete,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"port_protocol": {
				Type: schema.TypeString, Required: true, Description: "PortProtocol",
			},
			"profile_name": {
				Type: schema.TypeString, Required: true, Description: "ProfileName",
			},
		},
	}
}
func resourceDdosZoneProfilePortIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfilePortIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfilePortIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfilePortIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneProfilePortIndicatorSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortIndicatorSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortIndicatorSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfilePortIndicatorZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortIndicatorZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortIndicatorZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfilePortIndicator(d *schema.ResourceData) edpt.DdosZoneProfilePortIndicator {
	var ret edpt.DdosZoneProfilePortIndicator
	ret.Inst.IndicatorName = d.Get("indicator_name").(string)
	ret.Inst.SrcThresholdCfg = getObjectDdosZoneProfilePortIndicatorSrcThresholdCfg(d.Get("src_threshold_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdCfg = getObjectDdosZoneProfilePortIndicatorZoneThresholdCfg(d.Get("zone_threshold_cfg").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.PortProtocol = d.Get("port_protocol").(string)
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
