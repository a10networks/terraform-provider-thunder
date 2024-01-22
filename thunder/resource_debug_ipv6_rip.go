package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6Rip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_rip`: Routing Information Protocol (RIP) for IPv6\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6RipCreate,
		UpdateContext: resourceDebugIpv6RipUpdate,
		ReadContext:   resourceDebugIpv6RipRead,
		DeleteContext: resourceDebugIpv6RipDelete,

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
func resourceDebugIpv6RipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6RipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Rip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6RipRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6RipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6RipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Rip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6RipRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6RipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6RipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Rip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6RipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6RipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Rip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6Rip(d *schema.ResourceData) edpt.DebugIpv6Rip {
	var ret edpt.DebugIpv6Rip
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Detail = d.Get("detail").(int)
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Nsm = d.Get("nsm").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	ret.Inst.Recv = d.Get("recv").(int)
	ret.Inst.Send = d.Get("send").(int)
	return ret
}
