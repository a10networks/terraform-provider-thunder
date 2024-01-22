package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRangePatternRecognition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_port_range_pattern_recognition`: Enable Pattern Recognition\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryPortRangePatternRecognitionCreate,
		UpdateContext: resourceDdosDstEntryPortRangePatternRecognitionUpdate,
		ReadContext:   resourceDdosDstEntryPortRangePatternRecognitionRead,
		DeleteContext: resourceDdosDstEntryPortRangePatternRecognitionDelete,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type: schema.TypeString, Required: true, Description: "'heuristic': heuristic algorithm;",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntryPortRangePatternRecognitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangePatternRecognitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangePatternRecognition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortRangePatternRecognitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryPortRangePatternRecognitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangePatternRecognitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangePatternRecognition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortRangePatternRecognitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryPortRangePatternRecognitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangePatternRecognitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangePatternRecognition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryPortRangePatternRecognitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangePatternRecognitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangePatternRecognition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstEntryPortRangePatternRecognition(d *schema.ResourceData) edpt.DdosDstEntryPortRangePatternRecognition {
	var ret edpt.DdosDstEntryPortRangePatternRecognition
	ret.Inst.Algorithm = d.Get("algorithm").(string)
	ret.Inst.FilterInactiveThreshold = d.Get("filter_inactive_threshold").(int)
	ret.Inst.FilterThreshold = d.Get("filter_threshold").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	//omit uuid
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
