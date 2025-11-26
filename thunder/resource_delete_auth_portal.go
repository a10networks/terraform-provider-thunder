package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteAuthPortal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_auth_portal`: Portal file for http authentication\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteAuthPortalCreate,
		UpdateContext: resourceDeleteAuthPortalUpdate,
		ReadContext:   resourceDeleteAuthPortalRead,
		DeleteContext: resourceDeleteAuthPortalDelete,

		Schema: map[string]*schema.Schema{
			"portal_name": {
				Type: schema.TypeString, Optional: true, Description: "Local File Name",
			},
		},
	}
}
func resourceDeleteAuthPortalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthPortalRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteAuthPortalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthPortalRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteAuthPortalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteAuthPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteAuthPortal(d *schema.ResourceData) edpt.DeleteAuthPortal {
	var ret edpt.DeleteAuthPortal
	ret.Inst.PortalName = d.Get("portal_name").(string)
	return ret
}
