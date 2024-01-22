package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_rip`: Routing Information Protocol (RIP) for IPv6\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugRipCreate,
		UpdateContext: resourceDebugRipUpdate,
		ReadContext:   resourceDebugRipRead,
		DeleteContext: resourceDebugRipDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all debugging",
			},
			"detail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Detailed information display",
			},
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIPng events",
			},
			"nsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIPng and NSM communication",
			},
			"packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIPng packets",
			},
			"recv": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIPng receive packets",
			},
			"send": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIPng send packets",
			},
		},
	}
}
func resourceDebugRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRipRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRipRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugRip(d *schema.ResourceData) edpt.DebugRip {
	var ret edpt.DebugRip
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Detail = d.Get("detail").(int)
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Nsm = d.Get("nsm").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	ret.Inst.Recv = d.Get("recv").(int)
	ret.Inst.Send = d.Get("send").(int)
	return ret
}
