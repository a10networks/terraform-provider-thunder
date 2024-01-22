package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangePatternRecognition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_range_pattern_recognition`: Enable Pattern Recognition\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortRangePatternRecognitionCreate,
		UpdateContext: resourceDdosDstZonePortRangePatternRecognitionUpdate,
		ReadContext:   resourceDdosDstZonePortRangePatternRecognitionRead,
		DeleteContext: resourceDdosDstZonePortRangePatternRecognitionDelete,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type: schema.TypeString, Required: true, Description: "'heuristic': heuristic algorithm;",
			},
			"app_payload_offset": {
				Type: schema.TypeInt, Optional: true, Description: "Set offset of the payload, default 0",
			},
			"capture_traffic": {
				Type: schema.TypeString, Optional: true, Description: "'all': Capture all packets; 'dropped': Capture dropped packets (default);",
			},
			"filter_inactive_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Extracted filter inactive threshold",
			},
			"filter_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Extracted filter threshold",
			},
			"mode": {
				Type: schema.TypeString, Optional: true, Description: "'capture-never-expire': War-time capture without rate exceeding and never expires; 'manual': Manual mode;",
			},
			"sensitivity": {
				Type: schema.TypeString, Optional: true, Description: "'high': High Sensitivity; 'medium': Medium Sensitivity; 'low': Low Sensitivity;",
			},
			"triggered_by": {
				Type: schema.TypeString, Optional: true, Description: "'zone-escalation': Zone escalation trigger pattern recognition; 'packet-rate-exceeds': Packet rate limit exceeds trigger pattern recognition (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}
func resourceDdosDstZonePortRangePatternRecognitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangePatternRecognitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangePatternRecognition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangePatternRecognitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortRangePatternRecognitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangePatternRecognitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangePatternRecognition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangePatternRecognitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortRangePatternRecognitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangePatternRecognitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangePatternRecognition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortRangePatternRecognitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangePatternRecognitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangePatternRecognition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZonePortRangePatternRecognition(d *schema.ResourceData) edpt.DdosDstZonePortRangePatternRecognition {
	var ret edpt.DdosDstZonePortRangePatternRecognition
	ret.Inst.Algorithm = d.Get("algorithm").(string)
	ret.Inst.AppPayloadOffset = d.Get("app_payload_offset").(int)
	ret.Inst.CaptureTraffic = d.Get("capture_traffic").(string)
	ret.Inst.FilterInactiveThreshold = d.Get("filter_inactive_threshold").(int)
	ret.Inst.FilterThreshold = d.Get("filter_threshold").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	ret.Inst.TriggeredBy = d.Get("triggered_by").(string)
	//omit uuid
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
