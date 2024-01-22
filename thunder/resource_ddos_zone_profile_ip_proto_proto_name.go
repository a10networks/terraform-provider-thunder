package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfileIpProtoProtoName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_ip_proto_proto_name`: DDOS IP protocol profile configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfileIpProtoProtoNameCreate,
		UpdateContext: resourceDdosZoneProfileIpProtoProtoNameUpdate,
		ReadContext:   resourceDdosZoneProfileIpProtoProtoNameRead,
		DeleteContext: resourceDdosZoneProfileIpProtoProtoNameDelete,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6; 'gre': ip-proto gre; 'ipv4-encap': ip-proto IPv4 Encapsulation; 'ipv6-encap': ip-proto IPv6 Encapsulation;",
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
func resourceDdosZoneProfileIpProtoProtoNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNameRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileIpProtoProtoNameRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfileIpProtoProtoNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfileIpProtoProtoNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileIpProtoProtoNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfileIpProtoProtoName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneProfileIpProtoProtoNameIndicatorList(d []interface{}) []edpt.DdosZoneProfileIpProtoProtoNameIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfileIpProtoProtoNameIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfileIpProtoProtoNameIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNameIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNameIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNameIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNameIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNameIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNameIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNameIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNameIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfileIpProtoProtoName(d *schema.ResourceData) edpt.DdosZoneProfileIpProtoProtoName {
	var ret edpt.DdosZoneProfileIpProtoProtoName
	ret.Inst.IndicatorList = getSliceDdosZoneProfileIpProtoProtoNameIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	//omit uuid
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
