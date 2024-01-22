package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateRevert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_automatic_update_revert`: Manually revert to previous software version\n\n__PLACEHOLDER__",
		CreateContext: resourceAutomaticUpdateRevertCreate,
		UpdateContext: resourceAutomaticUpdateRevertUpdate,
		ReadContext:   resourceAutomaticUpdateRevertRead,
		DeleteContext: resourceAutomaticUpdateRevertDelete,

		Schema: map[string]*schema.Schema{
			"feature_name": {
				Type: schema.TypeString, Optional: true, Description: "'app-fw': Application Firewall; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;",
			},
		},
	}
}
func resourceAutomaticUpdateRevertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateRevertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateRevert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateRevertRead(ctx, d, meta)
	}
	return diags
}

func resourceAutomaticUpdateRevertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateRevertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateRevert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateRevertRead(ctx, d, meta)
	}
	return diags
}
func resourceAutomaticUpdateRevertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateRevertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateRevert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAutomaticUpdateRevertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateRevertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateRevert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAutomaticUpdateRevert(d *schema.ResourceData) edpt.AutomaticUpdateRevert {
	var ret edpt.AutomaticUpdateRevert
	ret.Inst.FeatureName = d.Get("feature_name").(string)
	return ret
}
