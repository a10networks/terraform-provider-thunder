package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6BindingTableValidate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_binding_table_validate`: Check for errors in LW-4over6 Binding Table\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6BindingTableValidateCreate,
		UpdateContext: resourceCgnv6Lw4o6BindingTableValidateUpdate,
		ReadContext:   resourceCgnv6Lw4o6BindingTableValidateRead,
		DeleteContext: resourceCgnv6Lw4o6BindingTableValidateDelete,

		Schema: map[string]*schema.Schema{
			"binding_name": {
				Type: schema.TypeString, Optional: true, Description: "LW-4over6 Binding Table Name",
			},
		},
	}
}
func resourceCgnv6Lw4o6BindingTableValidateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableValidateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableValidate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableValidateRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableValidateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableValidateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableValidate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableValidateRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6BindingTableValidateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableValidateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableValidate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableValidateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableValidateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableValidate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Lw4o6BindingTableValidate(d *schema.ResourceData) edpt.Cgnv6Lw4o6BindingTableValidate {
	var ret edpt.Cgnv6Lw4o6BindingTableValidate
	ret.Inst.BindingName = d.Get("binding_name").(string)
	return ret
}
