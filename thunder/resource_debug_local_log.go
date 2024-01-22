package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugLocalLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_local_log`: Debug LOCAL log\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugLocalLogCreate,
		UpdateContext: resourceDebugLocalLogUpdate,
		ReadContext:   resourceDebugLocalLogRead,
		DeleteContext: resourceDebugLocalLogDelete,

		Schema: map[string]*schema.Schema{
			"group": {
				Type: schema.TypeString, Optional: true, Description: "'query': Debug local log query (axapi).; 'db-mgr': Debug local log database management.; 'db-write': Debug local log insert local log into database.; 'queue': Debug local log generate and inqueue/dequeue.;",
			},
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-3. 1:error, 2:notice, 3:Info)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugLocalLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLocalLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLocalLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLocalLogRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugLocalLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLocalLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLocalLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLocalLogRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugLocalLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLocalLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLocalLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugLocalLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLocalLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLocalLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugLocalLog(d *schema.ResourceData) edpt.DebugLocalLog {
	var ret edpt.DebugLocalLog
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
