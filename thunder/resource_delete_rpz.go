package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteRpz() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_rpz`: Delete Response Policy Zone file\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteRpzCreate,
		UpdateContext: resourceDeleteRpzUpdate,
		ReadContext:   resourceDeleteRpzRead,
		DeleteContext: resourceDeleteRpzDelete,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Response Policy Zone File",
			},
		},
	}
}
func resourceDeleteRpzCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteRpzCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteRpz(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteRpzRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteRpzUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteRpzUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteRpz(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteRpzRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteRpzDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteRpzDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteRpz(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteRpzRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteRpzRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteRpz(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteRpz(d *schema.ResourceData) edpt.DeleteRpz {
	var ret edpt.DeleteRpz
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
