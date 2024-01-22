package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServicePatternRecognition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_pattern_recognition`: Enable Pattern Recognition\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServicePatternRecognitionCreate,
		UpdateContext: resourceDdosDstZonePortZoneServicePatternRecognitionUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServicePatternRecognitionRead,
		DeleteContext: resourceDdosDstZonePortZoneServicePatternRecognitionDelete,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type: schema.TypeString, Required: true, Description: "'heuristic': heuristic algorithm;",
			},
			"app_payload_offset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set offset of the payload",
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServicePatternRecognitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServicePatternRecognitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServicePatternRecognition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServicePatternRecognitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServicePatternRecognitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServicePatternRecognitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServicePatternRecognition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServicePatternRecognitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServicePatternRecognitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServicePatternRecognitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServicePatternRecognition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServicePatternRecognitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServicePatternRecognitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServicePatternRecognition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZonePortZoneServicePatternRecognition(d *schema.ResourceData) edpt.DdosDstZonePortZoneServicePatternRecognition {
	var ret edpt.DdosDstZonePortZoneServicePatternRecognition
	ret.Inst.Algorithm = d.Get("algorithm").(string)
	ret.Inst.AppPayloadOffset = d.Get("app_payload_offset").(int)
	ret.Inst.CaptureTraffic = d.Get("capture_traffic").(string)
	ret.Inst.FilterInactiveThreshold = d.Get("filter_inactive_threshold").(int)
	ret.Inst.FilterThreshold = d.Get("filter_threshold").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	ret.Inst.TriggeredBy = d.Get("triggered_by").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
