package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfileIpProtoProtoNumber() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_ip_proto_proto_number`: DDOS IP protocol profile configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfileIpProtoProtoNumberCreate,
		UpdateContext: resourceDdosZoneProfileIpProtoProtoNumberUpdate,
		ReadContext:   resourceDdosZoneProfileIpProtoProtoNumberRead,
		DeleteContext: resourceDdosZoneProfileIpProtoProtoNumberDelete,

		Schema: map[string]*schema.Schema{
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"protocol_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"profile_name": {
				Type: schema.TypeString, Required: true, Description: "ProfileName",
			},
		},
	}
}
func resourceDdosZoneProfileIpProtoProtoNumberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumber(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNumberRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNumberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumber(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNumberRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfileIpProtoProtoNumberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumber(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNumberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoNumber(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneProfileIpProtoProtoNumberIndicatorList(d []interface{}) []edpt.DdosZoneProfileIpProtoProtoNumberIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfileIpProtoProtoNumberIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfileIpProtoProtoNumberIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNumberIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNumberIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNumberIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNumberIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNumberIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNumberIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNumberIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNumberIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfileIpProtoProtoNumber(d *schema.ResourceData) edpt.DdosZoneProfileIpProtoProtoNumber {
	var ret edpt.DdosZoneProfileIpProtoProtoNumber
	ret.Inst.IndicatorList = getSliceDdosZoneProfileIpProtoProtoNumberIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.ProtocolNum = d.Get("protocol_num").(int)
	//omit uuid
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
