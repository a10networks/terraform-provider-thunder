package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutAppsSkipMacOverwrite() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_apps_skip_mac_overwrite`: Skips overwriting dest MAC of flooded packets on Active node\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutAppsSkipMacOverwriteCreate,
		UpdateContext: resourceScaleoutAppsSkipMacOverwriteUpdate,
		ReadContext:   resourceScaleoutAppsSkipMacOverwriteRead,
		DeleteContext: resourceScaleoutAppsSkipMacOverwriteDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skips overwriting dest MAC of flooded packets on Active node",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutAppsSkipMacOverwriteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSkipMacOverwriteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSkipMacOverwrite(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutAppsSkipMacOverwriteRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutAppsSkipMacOverwriteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSkipMacOverwriteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSkipMacOverwrite(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutAppsSkipMacOverwriteRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutAppsSkipMacOverwriteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSkipMacOverwriteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSkipMacOverwrite(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutAppsSkipMacOverwriteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSkipMacOverwriteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSkipMacOverwrite(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutAppsSkipMacOverwrite(d *schema.ResourceData) edpt.ScaleoutAppsSkipMacOverwrite {
	var ret edpt.ScaleoutAppsSkipMacOverwrite
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
