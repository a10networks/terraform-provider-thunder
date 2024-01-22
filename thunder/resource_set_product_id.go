package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSetProductId() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_set_product_id`: Set product ID\n\n__PLACEHOLDER__",
		CreateContext: resourceSetProductIdCreate,
		UpdateContext: resourceSetProductIdUpdate,
		ReadContext:   resourceSetProductIdRead,
		DeleteContext: resourceSetProductIdDelete,

		Schema: map[string]*schema.Schema{
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "Enter the desired product ID, default is 2",
			},
		},
	}
}
func resourceSetProductIdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSetProductIdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSetProductId(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSetProductIdRead(ctx, d, meta)
	}
	return diags
}

func resourceSetProductIdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSetProductIdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSetProductId(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSetProductIdRead(ctx, d, meta)
	}
	return diags
}
func resourceSetProductIdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSetProductIdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSetProductId(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSetProductIdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSetProductIdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSetProductId(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSetProductId(d *schema.ResourceData) edpt.SetProductId {
	var ret edpt.SetProductId
	ret.Inst.Id1 = d.Get("id1").(int)
	return ret
}
