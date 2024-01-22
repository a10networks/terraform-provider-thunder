package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnGtpApnPrefixTopnNode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_gtp_apn_prefix_topn_node`: Activate gtp-apn-prefix-topn template for fw.gtp.apn-prefix\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnGtpApnPrefixTopnNodeCreate,
		UpdateContext: resourceVisibilityTopnGtpApnPrefixTopnNodeUpdate,
		ReadContext:   resourceVisibilityTopnGtpApnPrefixTopnNodeRead,
		DeleteContext: resourceVisibilityTopnGtpApnPrefixTopnNodeDelete,

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
func resourceVisibilityTopnGtpApnPrefixTopnNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpApnPrefixTopnNodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpApnPrefixTopnNode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnGtpApnPrefixTopnNodeRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnGtpApnPrefixTopnNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpApnPrefixTopnNodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpApnPrefixTopnNode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnGtpApnPrefixTopnNodeRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnGtpApnPrefixTopnNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpApnPrefixTopnNodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpApnPrefixTopnNode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnGtpApnPrefixTopnNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpApnPrefixTopnNodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpApnPrefixTopnNode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnGtpApnPrefixTopnNode(d *schema.ResourceData) edpt.VisibilityTopnGtpApnPrefixTopnNode {
	var ret edpt.VisibilityTopnGtpApnPrefixTopnNode
	ret.Inst.Activate = d.Get("activate").(string)
	//omit uuid
	return ret
}
