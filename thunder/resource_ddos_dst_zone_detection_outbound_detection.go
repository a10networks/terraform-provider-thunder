package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionOutboundDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_outbound_detection`: DDOS Outbound Detection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionOutboundDetectionCreate,
		UpdateContext: resourceDdosDstZoneDetectionOutboundDetectionUpdate,
		ReadContext:   resourceDdosDstZoneDetectionOutboundDetectionRead,
		DeleteContext: resourceDdosDstZoneDetectionOutboundDetectionDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': configuration;",
			},
			"discovery_method": {
				Type: schema.TypeString, Optional: true, Description: "'asn': Autonomous Systems number; 'country': Country;",
			},
			"discovery_record": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Maximum number of top locations",
			},
			"enable_top_k": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topk_type": {
							Type: schema.TypeString, Optional: true, Description: "'source-subnet': Topk source subnet;",
						},
						"topk_netmask": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "Subnet mask. The value should be less than or equal to the minimum zone subnet mask + 8 (IPv6 Subnet mask)",
						},
						"topk_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
					},
				},
			},
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets;",
						},
						"tcp_window_size": {
							Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
						},
						"data_packet_size": {
							Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
						},
						"threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
						},
						"threshold_large_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
						},
						"threshold_str": {
							Type: schema.TypeString, Optional: true, Description: "Threshold for each geo-location (Non-zero floating point)",
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
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable outbound detection; 'disable': Disable outbound detection;",
			},
			"topk_source_subnet": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneDetectionOutboundDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionOutboundDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionOutboundDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionOutboundDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionOutboundDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionOutboundDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneDetectionOutboundDetectionEnableTopK(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK
		oi.TopkType = in["topk_type"].(string)
		oi.TopkNetmask = in["topk_netmask"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionIndicatorList(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.ThresholdNum = in["threshold_num"].(int)
		oi.ThresholdLargeNum = in["threshold_large_num"].(int)
		oi.ThresholdStr = in["threshold_str"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnet187(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet187 {

	var ret edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet187
	return ret
}

func dataToEndpointDdosDstZoneDetectionOutboundDetection(d *schema.ResourceData) edpt.DdosDstZoneDetectionOutboundDetection {
	var ret edpt.DdosDstZoneDetectionOutboundDetection
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.DiscoveryMethod = d.Get("discovery_method").(string)
	ret.Inst.DiscoveryRecord = d.Get("discovery_record").(int)
	ret.Inst.EnableTopK = getSliceDdosDstZoneDetectionOutboundDetectionEnableTopK(d.Get("enable_top_k").([]interface{}))
	ret.Inst.IndicatorList = getSliceDdosDstZoneDetectionOutboundDetectionIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.Toggle = d.Get("toggle").(string)
	ret.Inst.TopkSourceSubnet = getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnet187(d.Get("topk_source_subnet").([]interface{}))
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
