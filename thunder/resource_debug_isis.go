package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_isis`: Intermediate System - Intermediate System (IS-IS)\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIsisCreate,
		UpdateContext: resourceDebugIsisUpdate,
		ReadContext:   resourceDebugIsisRead,
		DeleteContext: resourceDebugIsisDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all debugging",
			},
			"bfd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
			},
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS Events",
			},
			"ifsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS Interface Finite State Machine",
			},
			"lsp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS Link State PDU",
			},
			"nfsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS Neighbor Finite State Machine",
			},
			"nsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS NSM information",
			},
			"pdu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS Protocol Data Unit",
			},
			"spf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IS-IS SPF Calculation",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIsis(d *schema.ResourceData) edpt.DebugIsis {
	var ret edpt.DebugIsis
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Ifsm = d.Get("ifsm").(int)
	ret.Inst.Lsp = d.Get("lsp").(int)
	ret.Inst.Nfsm = d.Get("nfsm").(int)
	ret.Inst.Nsm = d.Get("nsm").(int)
	ret.Inst.Pdu = d.Get("pdu").(int)
	ret.Inst.Spf = d.Get("spf").(int)
	//omit uuid
	return ret
}
