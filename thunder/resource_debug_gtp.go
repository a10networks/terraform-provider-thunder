package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugGtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_gtp`: Debug GTP\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugGtpCreate,
		UpdateContext: resourceDebugGtpUpdate,
		ReadContext:   resourceDebugGtpRead,
		DeleteContext: resourceDebugGtpDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Debug level (Level 1: Errors  Level 2: Minimal_info  Level 3: Detailed_info)",
			},
			"packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug GTP packet",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugGtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugGtpRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugGtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugGtpRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugGtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugGtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugGtp(d *schema.ResourceData) edpt.DebugGtp {
	var ret edpt.DebugGtp
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	//omit uuid
	return ret
}
