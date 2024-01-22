package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateCheckNow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_automatic_update_check_now`: Manually check for update now\n\n__PLACEHOLDER__",
		CreateContext: resourceAutomaticUpdateCheckNowCreate,
		UpdateContext: resourceAutomaticUpdateCheckNowUpdate,
		ReadContext:   resourceAutomaticUpdateCheckNowRead,
		DeleteContext: resourceAutomaticUpdateCheckNowDelete,

		Schema: map[string]*schema.Schema{
			"feature_name": {
				Type: schema.TypeString, Optional: true, Description: "'app-fw': Application Firewall; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;",
			},
			"from_staging_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Get files from GLM Staging storage",
			},
			"prod_ver": {
				Type: schema.TypeString, Optional: true, Description: "update to this specific version, if this option is not configured, update to the latest version",
			},
			"stage_ver": {
				Type: schema.TypeString, Optional: true, Description: "update this specific version",
			},
		},
	}
}
func resourceAutomaticUpdateCheckNowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateCheckNowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateCheckNow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateCheckNowRead(ctx, d, meta)
	}
	return diags
}

func resourceAutomaticUpdateCheckNowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateCheckNowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateCheckNow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateCheckNowRead(ctx, d, meta)
	}
	return diags
}
func resourceAutomaticUpdateCheckNowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateCheckNowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateCheckNow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAutomaticUpdateCheckNowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateCheckNowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateCheckNow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAutomaticUpdateCheckNow(d *schema.ResourceData) edpt.AutomaticUpdateCheckNow {
	var ret edpt.AutomaticUpdateCheckNow
	ret.Inst.FeatureName = d.Get("feature_name").(string)
	ret.Inst.FromStagingServer = d.Get("from_staging_server").(int)
	ret.Inst.ProdVer = d.Get("prod_ver").(string)
	ret.Inst.StageVer = d.Get("stage_ver").(string)
	return ret
}
