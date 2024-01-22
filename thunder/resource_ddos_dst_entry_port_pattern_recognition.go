package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortPatternRecognition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_port_pattern_recognition`: Enable Pattern Recognition\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryPortPatternRecognitionCreate,
		UpdateContext: resourceDdosDstEntryPortPatternRecognitionUpdate,
		ReadContext:   resourceDdosDstEntryPortPatternRecognitionRead,
		DeleteContext: resourceDdosDstEntryPortPatternRecognitionDelete,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntryPortPatternRecognitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortPatternRecognitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortPatternRecognition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortPatternRecognitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryPortPatternRecognitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortPatternRecognitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortPatternRecognition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortPatternRecognitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryPortPatternRecognitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortPatternRecognitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortPatternRecognition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryPortPatternRecognitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortPatternRecognitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortPatternRecognition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstEntryPortPatternRecognition(d *schema.ResourceData) edpt.DdosDstEntryPortPatternRecognition {
	var ret edpt.DdosDstEntryPortPatternRecognition
	ret.Inst.Algorithm = d.Get("algorithm").(string)
	ret.Inst.FilterInactiveThreshold = d.Get("filter_inactive_threshold").(int)
	ret.Inst.FilterThreshold = d.Get("filter_threshold").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	//omit uuid
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
