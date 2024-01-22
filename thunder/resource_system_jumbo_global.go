package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemJumboGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_jumbo_global`: To enable/disable jumbo\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemJumboGlobalCreate,
		UpdateContext: resourceSystemJumboGlobalUpdate,
		ReadContext:   resourceSystemJumboGlobalRead,
		DeleteContext: resourceSystemJumboGlobalDelete,

		Schema: map[string]*schema.Schema{
			"enable_jumbo": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable jumbo frame",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemJumboGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJumboGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJumboGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemJumboGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemJumboGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJumboGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJumboGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemJumboGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemJumboGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJumboGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJumboGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemJumboGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJumboGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJumboGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemJumboGlobal(d *schema.ResourceData) edpt.SystemJumboGlobal {
	var ret edpt.SystemJumboGlobal
	ret.Inst.EnableJumbo = d.Get("enable_jumbo").(int)
	//omit uuid
	return ret
}
