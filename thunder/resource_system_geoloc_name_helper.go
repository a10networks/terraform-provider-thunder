package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocNameHelper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_geoloc_name_helper`: Geolocation name helper\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGeolocNameHelperCreate,
		UpdateContext: resourceSystemGeolocNameHelperUpdate,
		ReadContext:   resourceSystemGeolocNameHelperRead,
		DeleteContext: resourceSystemGeolocNameHelperDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'place-holder': place-holder;",
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
func resourceSystemGeolocNameHelperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocNameHelperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocNameHelper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocNameHelperRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGeolocNameHelperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocNameHelperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocNameHelper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocNameHelperRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGeolocNameHelperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocNameHelperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocNameHelper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGeolocNameHelperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocNameHelperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocNameHelper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemGeolocNameHelperSamplingEnable(d []interface{}) []edpt.SystemGeolocNameHelperSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocNameHelperSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocNameHelperSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocNameHelper(d *schema.ResourceData) edpt.SystemGeolocNameHelper {
	var ret edpt.SystemGeolocNameHelper
	ret.Inst.SamplingEnable = getSliceSystemGeolocNameHelperSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
