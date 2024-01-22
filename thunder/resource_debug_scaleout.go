package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugScaleout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_scaleout`: Debug scaleout\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugScaleoutCreate,
		UpdateContext: resourceDebugScaleoutUpdate,
		ReadContext:   resourceDebugScaleoutRead,
		DeleteContext: resourceDebugScaleoutDelete,

		Schema: map[string]*schema.Schema{
			"config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for scaleout config change",
			},
			"debug_level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-3)",
			},
			"event": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for scaleout events",
			},
			"packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for scaleout packet flow",
			},
			"session_sync": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for scaleout session sync events",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugScaleoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugScaleoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugScaleout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugScaleoutRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugScaleoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugScaleoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugScaleout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugScaleoutRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugScaleoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugScaleoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugScaleout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugScaleoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugScaleoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugScaleout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugScaleout(d *schema.ResourceData) edpt.DebugScaleout {
	var ret edpt.DebugScaleout
	ret.Inst.Config = d.Get("config").(int)
	ret.Inst.DebugLevel = d.Get("debug_level").(int)
	ret.Inst.Event = d.Get("event").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	ret.Inst.SessionSync = d.Get("session_sync").(int)
	//omit uuid
	return ret
}
