package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteDdos() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_ddos`: Delete ddos file\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteDdosCreate,
		UpdateContext: resourceDeleteDdosUpdate,
		ReadContext:   resourceDeleteDdosRead,
		DeleteContext: resourceDeleteDdosDelete,

		Schema: map[string]*schema.Schema{
			"ddos_script": {
				Type: schema.TypeString, Optional: true, Description: "Specify file to be deleted",
			},
		},
	}
}
func resourceDeleteDdosCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDdosCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDdos(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteDdosRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteDdosUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDdosUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDdos(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteDdosRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteDdosDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDdosDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDdos(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteDdosRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDdosRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDdos(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteDdos(d *schema.ResourceData) edpt.DeleteDdos {
	var ret edpt.DeleteDdos
	ret.Inst.DdosScript = d.Get("ddos_script").(string)
	return ret
}
