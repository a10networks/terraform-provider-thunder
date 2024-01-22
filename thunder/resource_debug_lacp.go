package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugLacp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_lacp`: lacp - LACP commands\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugLacpCreate,
		UpdateContext: resourceDebugLacpUpdate,
		ReadContext:   resourceDebugLacpRead,
		DeleteContext: resourceDebugLacpDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "all - turn on all debugging",
			},
			"cli": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "cli - echo commands to console",
			},
			"detail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "detail - echo timer start/stop to console",
			},
			"event": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "event - echo events to console",
			},
			"ha": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ha - echo High availability events to console",
			},
			"packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet - echo packet contents to console",
			},
			"sync": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "sync - echo synchronization to console",
			},
			"timer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "timer - echo timer expiry to console",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugLacpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLacpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLacp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLacpRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugLacpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLacpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLacp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLacpRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugLacpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLacpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLacp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugLacpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLacpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLacp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugLacp(d *schema.ResourceData) edpt.DebugLacp {
	var ret edpt.DebugLacp
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Cli = d.Get("cli").(int)
	ret.Inst.Detail = d.Get("detail").(int)
	ret.Inst.Event = d.Get("event").(int)
	ret.Inst.Ha = d.Get("ha").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	ret.Inst.Sync = d.Get("sync").(int)
	ret.Inst.Timer = d.Get("timer").(int)
	//omit uuid
	return ret
}
