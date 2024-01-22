package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugSctp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_sctp`: Debug SCTP\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugSctpCreate,
		UpdateContext: resourceDebugSctpUpdate,
		ReadContext:   resourceDebugSctpRead,
		DeleteContext: resourceDebugSctpDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Debug level (Level 1: Errors  Level 2: Minimal_info  Level 3: Detailed_info)",
			},
			"packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug SCTP packet",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugSctpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSctpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSctp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSctpRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugSctpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSctpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSctp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSctpRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugSctpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSctpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSctp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugSctpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSctpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSctp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugSctp(d *schema.ResourceData) edpt.DebugSctp {
	var ret edpt.DebugSctp
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	//omit uuid
	return ret
}
