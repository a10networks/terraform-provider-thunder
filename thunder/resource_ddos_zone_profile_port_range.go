package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfilePortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_port_range`: DDOS Port-Range & Protocol profile configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfilePortRangeCreate,
		UpdateContext: resourceDdosZoneProfilePortRangeUpdate,
		ReadContext:   resourceDdosZoneProfilePortRangeRead,
		DeleteContext: resourceDdosZoneProfilePortRangeDelete,

		Schema: map[string]*schema.Schema{
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
			"port_range_end": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
			},
			"port_range_start": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-tcp': SIP-TCP Port; 'sip-udp': SIP-UDP Port; 'quic': QUIC Port;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosZoneProfilePortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfilePortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfilePortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfilePortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneProfilePortRangeIndicatorList(d []interface{}) []edpt.DdosZoneProfilePortRangeIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfilePortRangeIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfilePortRangeIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfilePortRangeIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfilePortRangeIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfilePortRangeIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortRangeIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortRangeIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfilePortRangeIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortRangeIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortRangeIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfilePortRange(d *schema.ResourceData) edpt.DdosZoneProfilePortRange {
	var ret edpt.DdosZoneProfilePortRange
	ret.Inst.IndicatorList = getSliceDdosZoneProfilePortRangeIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(int)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
