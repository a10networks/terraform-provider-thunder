package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugDumpthread() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_dumpthread`: Generate system threads info\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugDumpthreadCreate,
		UpdateContext: resourceDebugDumpthreadUpdate,
		ReadContext:   resourceDebugDumpthreadRead,
		DeleteContext: resourceDebugDumpthreadDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugDumpthreadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDumpthreadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDumpthread(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDumpthreadRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugDumpthreadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDumpthreadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDumpthread(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDumpthreadRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugDumpthreadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDumpthreadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDumpthread(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugDumpthreadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDumpthreadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDumpthread(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugDumpthread(d *schema.ResourceData) edpt.DebugDumpthread {
	var ret edpt.DebugDumpthread
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
