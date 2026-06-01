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
			"separate_v4_v6_traffic_map": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Separates traffic maps for IPv4 and IPv6",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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

func getObjectScaleoutAppsSeparateV4V6TrafficMap1410(d []interface{}) edpt.ScaleoutAppsSeparateV4V6TrafficMap1410 {

	count1 := len(d)
	var ret edpt.ScaleoutAppsSeparateV4V6TrafficMap1410
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutAppsSkipMacOverwrite1411(d []interface{}) edpt.ScaleoutAppsSkipMacOverwrite1411 {

	count1 := len(d)
	var ret edpt.ScaleoutAppsSkipMacOverwrite1411
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
	ret.Inst.SeparateV4V6TrafficMap = getObjectScaleoutAppsSeparateV4V6TrafficMap1410(d.Get("separate_v4_v6_traffic_map").([]interface{}))
	ret.Inst.SkipMacOverwrite = getObjectScaleoutAppsSkipMacOverwrite1411(d.Get("skip_mac_overwrite").([]interface{}))
	//omit uuid
	return ret
}
