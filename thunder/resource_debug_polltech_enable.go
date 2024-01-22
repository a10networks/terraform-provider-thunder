package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugPolltechEnable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_polltech_enable`: Enable polltech switch\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugPolltechEnableCreate,
		UpdateContext: resourceDebugPolltechEnableUpdate,
		ReadContext:   resourceDebugPolltechEnableRead,
		DeleteContext: resourceDebugPolltechEnableDelete,

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
func resourceDebugPolltechEnableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPolltechEnableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPolltechEnable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPolltechEnableRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugPolltechEnableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPolltechEnableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPolltechEnable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPolltechEnableRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugPolltechEnableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPolltechEnableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPolltechEnable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugPolltechEnableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPolltechEnableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPolltechEnable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugPolltechEnable(d *schema.ResourceData) edpt.DebugPolltechEnable {
	var ret edpt.DebugPolltechEnable
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
