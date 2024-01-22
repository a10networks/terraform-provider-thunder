package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutApps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_apps`: Enable Scaleout for apps\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutAppsCreate,
		UpdateContext: resourceScaleoutAppsUpdate,
		ReadContext:   resourceScaleoutAppsRead,
		DeleteContext: resourceScaleoutAppsDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Scaleout for apps",
			},
			"skip_mac_overwrite": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skips overwriting dest MAC of flooded packets on Active node",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutAppsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutApps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutAppsRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutAppsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutApps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutAppsRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutAppsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutApps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutAppsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutApps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectScaleoutAppsSkipMacOverwrite1324(d []interface{}) edpt.ScaleoutAppsSkipMacOverwrite1324 {

	count1 := len(d)
	var ret edpt.ScaleoutAppsSkipMacOverwrite1324
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointScaleoutApps(d *schema.ResourceData) edpt.ScaleoutApps {
	var ret edpt.ScaleoutApps
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.SkipMacOverwrite = getObjectScaleoutAppsSkipMacOverwrite1324(d.Get("skip_mac_overwrite").([]interface{}))
	//omit uuid
	return ret
}
