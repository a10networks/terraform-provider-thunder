package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugSystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_system`: System module parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugSystemCreate,
		UpdateContext: resourceDebugSystemUpdate,
		ReadContext:   resourceDebugSystemRead,
		DeleteContext: resourceDebugSystemDelete,

		Schema: map[string]*schema.Schema{
			"debug_system_enum": {
				Type: schema.TypeString, Optional: true, Description: "'all': System all debug switch; 'aaa': System debug aaa switch; 'import-export': System debug import or export switch; 'ssl': System debug ssl switch;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSystemCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSystem(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSystemRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSystemUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSystem(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSystemRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSystemDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSystem(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSystemRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSystem(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugSystem(d *schema.ResourceData) edpt.DebugSystem {
	var ret edpt.DebugSystem
	ret.Inst.Debug_system_enum = d.Get("debug_system_enum").(string)
	//omit uuid
	return ret
}
