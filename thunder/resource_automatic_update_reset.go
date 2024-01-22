package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateReset() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_automatic_update_reset`: Manually reset to default software version\n\n__PLACEHOLDER__",
		CreateContext: resourceAutomaticUpdateResetCreate,
		UpdateContext: resourceAutomaticUpdateResetUpdate,
		ReadContext:   resourceAutomaticUpdateResetRead,
		DeleteContext: resourceAutomaticUpdateResetDelete,

		Schema: map[string]*schema.Schema{
			"feature_name": {
				Type: schema.TypeString, Optional: true, Description: "'app-fw': Application Firewall; 'ca-bundle': CA Certificate Bundle;",
			},
		},
	}
}
func resourceAutomaticUpdateResetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateResetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateReset(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateResetRead(ctx, d, meta)
	}
	return diags
}

func resourceAutomaticUpdateResetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateResetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateReset(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateResetRead(ctx, d, meta)
	}
	return diags
}
func resourceAutomaticUpdateResetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateResetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateReset(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAutomaticUpdateResetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateResetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateReset(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAutomaticUpdateReset(d *schema.ResourceData) edpt.AutomaticUpdateReset {
	var ret edpt.AutomaticUpdateReset
	ret.Inst.FeatureName = d.Get("feature_name").(string)
	return ret
}
