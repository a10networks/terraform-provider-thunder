package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnTemplGtpPlcyTopnNode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_templ_gtp_plcy_topn_node`: Activate templ-gtp-plcy-topn template for template.gtp-policy\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnTemplGtpPlcyTopnNodeCreate,
		UpdateContext: resourceVisibilityTopnTemplGtpPlcyTopnNodeUpdate,
		ReadContext:   resourceVisibilityTopnTemplGtpPlcyTopnNodeRead,
		DeleteContext: resourceVisibilityTopnTemplGtpPlcyTopnNodeDelete,

		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeString, Optional: true, Description: "Name of the templated to be activated",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityTopnTemplGtpPlcyTopnNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnNodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnNode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnTemplGtpPlcyTopnNodeRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnTemplGtpPlcyTopnNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnNodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnNode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnTemplGtpPlcyTopnNodeRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnTemplGtpPlcyTopnNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnNodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnNode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnTemplGtpPlcyTopnNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnNodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnNode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnTemplGtpPlcyTopnNode(d *schema.ResourceData) edpt.VisibilityTopnTemplGtpPlcyTopnNode {
	var ret edpt.VisibilityTopnTemplGtpPlcyTopnNode
	ret.Inst.Activate = d.Get("activate").(string)
	//omit uuid
	return ret
}
