package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfilePort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile_port`: Port for DDoS Zone profile\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfilePortCreate,
		UpdateContext: resourceDdosZoneProfilePortUpdate,
		ReadContext:   resourceDdosZoneProfilePortRead,
		DeleteContext: resourceDdosZoneProfilePortDelete,

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
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"port_protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': dns-tcp; 'dns-udp': dns-udp; 'sip-tcp': sip-tcp; 'sip-udp': sip-udp; 'http': http; 'tcp': tcp; 'udp': udp; 'ssl-l4': ssl-l4; 'quic': quic;",
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
func resourceDdosZoneProfilePortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfilePortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfilePortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfilePortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfilePortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfilePortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfilePort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneProfilePortIndicatorList(d []interface{}) []edpt.DdosZoneProfilePortIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfilePortIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfilePortIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfilePortIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfilePortIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfilePortIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfilePortIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfilePort(d *schema.ResourceData) edpt.DdosZoneProfilePort {
	var ret edpt.DdosZoneProfilePort
	ret.Inst.IndicatorList = getSliceDdosZoneProfilePortIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.PortProtocol = d.Get("port_protocol").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	return ret
}
