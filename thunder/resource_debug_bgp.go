package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugBgp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_bgp`: Debug Border Gateway Protocol (BGP)\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugBgpCreate,
		UpdateContext: resourceDebugBgpUpdate,
		ReadContext:   resourceDebugBgpRead,
		DeleteContext: resourceDebugBgpDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "all debugging",
			},
			"bfd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
			},
			"dampening": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP Dampening",
			},
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP events",
			},
			"filters": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP filters",
			},
			"fsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP Finite State Machine",
			},
			"in": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inbound updates",
			},
			"keepalives": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP keepalives",
			},
			"nht": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NHT message",
			},
			"nsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM message",
			},
			"out": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outbound updates",
			},
			"updates": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP updates",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBgpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBgp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugBgpRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBgpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBgp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugBgpRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBgpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBgp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBgpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBgp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugBgp(d *schema.ResourceData) edpt.DebugBgp {
	var ret edpt.DebugBgp
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.Dampening = d.Get("dampening").(int)
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Filters = d.Get("filters").(int)
	ret.Inst.Fsm = d.Get("fsm").(int)
	ret.Inst.In = d.Get("in").(int)
	ret.Inst.Keepalives = d.Get("keepalives").(int)
	ret.Inst.Nht = d.Get("nht").(int)
	ret.Inst.Nsm = d.Get("nsm").(int)
	ret.Inst.Out = d.Get("out").(int)
	ret.Inst.Updates = d.Get("updates").(int)
	//omit uuid
	return ret
}
