package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemLro() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_lro`: Large Receive Offload\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemLroCreate,
		UpdateContext: resourceSystemLroUpdate,
		ReadContext:   resourceSystemLroRead,
		DeleteContext: resourceSystemLroDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Large Receive Offload",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Large Receive Offload",
			},
		},
	}
}
func resourceSystemLroCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLroCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLro(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemLroRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemLroUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLroUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLro(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemLroRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemLroDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLroDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLro(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemLroRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLroRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLro(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemLro(d *schema.ResourceData) edpt.SystemLro {
	var ret edpt.SystemLro
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	return ret
}
