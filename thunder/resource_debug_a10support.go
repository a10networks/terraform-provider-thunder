package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugA10support() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_a10support`: Enter the password\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugA10supportCreate,
		UpdateContext: resourceDebugA10supportUpdate,
		ReadContext:   resourceDebugA10supportRead,
		DeleteContext: resourceDebugA10supportDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Duration in minutes",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "Password to be set",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugA10supportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugA10supportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugA10support(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugA10supportRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugA10supportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugA10supportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugA10support(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugA10supportRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugA10supportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugA10supportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugA10support(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugA10supportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugA10supportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugA10support(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugA10support(d *schema.ResourceData) edpt.DebugA10support {
	var ret edpt.DebugA10support
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Password = d.Get("password").(string)
	//omit uuid
	return ret
}
