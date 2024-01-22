package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugRt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_rt`: Routing module parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugRtCreate,
		UpdateContext: resourceDebugRtUpdate,
		ReadContext:   resourceDebugRtRead,
		DeleteContext: resourceDebugRtDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Routing all debug switch",
			},
			"check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ROUTING debug check switch",
			},
			"command": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Routing debug command switch",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugRtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRtRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugRtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRtRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugRtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugRtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugRt(d *schema.ResourceData) edpt.DebugRt {
	var ret edpt.DebugRt
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Check = d.Get("check").(int)
	ret.Inst.Command = d.Get("command").(int)
	//omit uuid
	return ret
}
