package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteAuthPortalImage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_auth_portal_image`: Image file for default portal\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteAuthPortalImageCreate,
		UpdateContext: resourceDeleteAuthPortalImageUpdate,
		ReadContext:   resourceDeleteAuthPortalImageRead,
		DeleteContext: resourceDeleteAuthPortalImageDelete,

		Schema: map[string]*schema.Schema{
			"image_name": {
				Type: schema.TypeString, Optional: true, Description: "Local File Name",
			},
		},
	}
}
func resourceDeleteAuthPortalImageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalImageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortalImage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthPortalImageRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteAuthPortalImageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalImageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortalImage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthPortalImageRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteAuthPortalImageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalImageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortalImage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteAuthPortalImageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthPortalImageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthPortalImage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteAuthPortalImage(d *schema.ResourceData) edpt.DeleteAuthPortalImage {
	var ret edpt.DeleteAuthPortalImage
	ret.Inst.ImageName = d.Get("image_name").(string)
	return ret
}
