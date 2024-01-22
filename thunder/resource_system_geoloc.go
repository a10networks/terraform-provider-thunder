package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeoloc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_geoloc`: \n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGeolocCreate,
		UpdateContext: resourceSystemGeolocUpdate,
		ReadContext:   resourceSystemGeolocRead,
		DeleteContext: resourceSystemGeolocDelete,

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
func resourceSystemGeolocCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoloc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGeolocUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoloc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGeolocDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoloc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGeolocRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoloc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemGeolocSamplingEnable(d []interface{}) []edpt.SystemGeolocSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeoloc(d *schema.ResourceData) edpt.SystemGeoloc {
	var ret edpt.SystemGeoloc
	ret.Inst.SamplingEnable = getSliceSystemGeolocSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
