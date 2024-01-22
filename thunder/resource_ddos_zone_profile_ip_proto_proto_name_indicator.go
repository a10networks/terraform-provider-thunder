package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfileIpProtoProtoNameIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_ip_proto_proto_name_indicator`: DDoS zone profile indicator config\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfileIpProtoProtoNameIndicatorCreate,
		UpdateContext: resourceDdosZoneProfileIpProtoProtoNameIndicatorUpdate,
		ReadContext:   resourceDdosZoneProfileIpProtoProtoNameIndicatorRead,
		DeleteContext: resourceDdosZoneProfileIpProtoProtoNameIndicatorDelete,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"profile_name": {
				Type: schema.TypeString, Required: true, Description: "ProfileName",
			},
		},
	}
}
func resourceDdosZoneProfileIpProtoProtoNameIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNameIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNameIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNameIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNameIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNameIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfileIpProtoProtoNameIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNameIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNameIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNameIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneProfileIpProtoProtoNameIndicatorSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNameIndicatorSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNameIndicatorSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNameIndicatorZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNameIndicatorZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNameIndicatorZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfileIpProtoProtoNameIndicator(d *schema.ResourceData) edpt.DdosZoneProfileIpProtoProtoNameIndicator {
	var ret edpt.DdosZoneProfileIpProtoProtoNameIndicator
	ret.Inst.IndicatorName = d.Get("indicator_name").(string)
	ret.Inst.SrcThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNameIndicatorSrcThresholdCfg(d.Get("src_threshold_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNameIndicatorZoneThresholdCfg(d.Get("zone_threshold_cfg").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
