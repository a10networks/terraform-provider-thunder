package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfileIpProtoProtoNumberIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_ip_proto_proto_number_indicator`: DDoS zone profile indicator config\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfileIpProtoProtoNumberIndicatorCreate,
		UpdateContext: resourceDdosZoneProfileIpProtoProtoNumberIndicatorUpdate,
		ReadContext:   resourceDdosZoneProfileIpProtoProtoNumberIndicatorRead,
		DeleteContext: resourceDdosZoneProfileIpProtoProtoNumberIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"indicator_name": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': pkt-rate; 'pkt-drop-rate': pkt-drop-rate; 'bit-rate': bit-rate; 'pkt-drop-ratio': pkt-drop-ratio; 'bytes-to-bytes-from-ratio': bytes-to-bytes-from-ratio; 'frag-rate': frag-rate; 'cpu-utilization': cpu-utilization; 'interface-utilization': interface-utilization;",
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
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
			"profile_name": {
				Type: schema.TypeString, Required: true, Description: "ProfileName",
			},
		},
	}
}
func resourceDdosZoneProfileIpProtoProtoNumberIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumberIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNumberIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNumberIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumberIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNumberIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfileIpProtoProtoNumberIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumberIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNumberIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumberIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneProfileIpProtoProtoNumberIndicatorSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNumberIndicatorSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNumberIndicatorSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNumberIndicatorZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNumberIndicatorZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNumberIndicatorZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfileIpProtoProtoNumberIndicator(d *schema.ResourceData) edpt.DdosZoneProfileIpProtoProtoNumberIndicator {
	var ret edpt.DdosZoneProfileIpProtoProtoNumberIndicator
	ret.Inst.IndicatorName = d.Get("indicator_name").(string)
	ret.Inst.SrcThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNumberIndicatorSrcThresholdCfg(d.Get("src_threshold_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNumberIndicatorZoneThresholdCfg(d.Get("zone_threshold_cfg").([]interface{}))
	ret.Inst.ProtocolNum = d.Get("protocol_num").(string)
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
