package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutAppsSeparateV4V6TrafficMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_apps_separate_v4_v6_traffic_map`: Separates traffic maps for IPv4 and IPv6\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutAppsSeparateV4V6TrafficMapCreate,
		UpdateContext: resourceScaleoutAppsSeparateV4V6TrafficMapUpdate,
		ReadContext:   resourceScaleoutAppsSeparateV4V6TrafficMapRead,
		DeleteContext: resourceScaleoutAppsSeparateV4V6TrafficMapDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Separates traffic maps for IPv4 and IPv6",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutAppsSeparateV4V6TrafficMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSeparateV4V6TrafficMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSeparateV4V6TrafficMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutAppsSeparateV4V6TrafficMapRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutAppsSeparateV4V6TrafficMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSeparateV4V6TrafficMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSeparateV4V6TrafficMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutAppsSeparateV4V6TrafficMapRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutAppsSeparateV4V6TrafficMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSeparateV4V6TrafficMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSeparateV4V6TrafficMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutAppsSeparateV4V6TrafficMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutAppsSeparateV4V6TrafficMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutAppsSeparateV4V6TrafficMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutAppsSeparateV4V6TrafficMap(d *schema.ResourceData) edpt.ScaleoutAppsSeparateV4V6TrafficMap {
	var ret edpt.ScaleoutAppsSeparateV4V6TrafficMap
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
