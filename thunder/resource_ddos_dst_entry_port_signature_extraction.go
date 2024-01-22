package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortSignatureExtraction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_port_signature_extraction`: Enable Signature Extraction\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryPortSignatureExtractionCreate,
		UpdateContext: resourceDdosDstEntryPortSignatureExtractionUpdate,
		ReadContext:   resourceDdosDstEntryPortSignatureExtractionRead,
		DeleteContext: resourceDdosDstEntryPortSignatureExtractionDelete,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type: schema.TypeString, Required: true, Description: "'heuristic': heuristic algorithm;",
			},
			"manual_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable manual mode",
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
func resourceDdosDstEntryPortSignatureExtractionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortSignatureExtractionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortSignatureExtraction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortSignatureExtractionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryPortSignatureExtractionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortSignatureExtractionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortSignatureExtraction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortSignatureExtractionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryPortSignatureExtractionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortSignatureExtractionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortSignatureExtraction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryPortSignatureExtractionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortSignatureExtractionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortSignatureExtraction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstEntryPortSignatureExtraction(d *schema.ResourceData) edpt.DdosDstEntryPortSignatureExtraction {
	var ret edpt.DdosDstEntryPortSignatureExtraction
	ret.Inst.Algorithm = d.Get("algorithm").(string)
	ret.Inst.ManualMode = d.Get("manual_mode").(int)
	//omit uuid
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
