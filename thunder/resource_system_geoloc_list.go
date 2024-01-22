package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_geoloc_list`: Configure geolocation list\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGeolocListCreate,
		UpdateContext: resourceSystemGeolocListUpdate,
		ReadContext:   resourceSystemGeolocListRead,
		DeleteContext: resourceSystemGeolocListDelete,

		Schema: map[string]*schema.Schema{
			"exclude_geoloc_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_geoloc_name_val": {
							Type: schema.TypeString, Optional: true, Description: "Geolocation name to exclude",
						},
					},
				},
			},
			"include_geoloc_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_geoloc_name_val": {
							Type: schema.TypeString, Optional: true, Description: "Geolocation name to add",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of Geolocation list",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-count': hit-count; 'total-geoloc': total-geoloc; 'total-active': total-active;",
						},
					},
				},
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sharing with other partitions",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemGeolocListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGeolocListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGeolocListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGeolocListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemGeolocListExcludeGeolocNameList(d []interface{}) []edpt.SystemGeolocListExcludeGeolocNameList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListExcludeGeolocNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListExcludeGeolocNameList
		oi.ExcludeGeolocNameVal = in["exclude_geoloc_name_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListIncludeGeolocNameList(d []interface{}) []edpt.SystemGeolocListIncludeGeolocNameList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListIncludeGeolocNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListIncludeGeolocNameList
		oi.IncludeGeolocNameVal = in["include_geoloc_name_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListSamplingEnable(d []interface{}) []edpt.SystemGeolocListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocList(d *schema.ResourceData) edpt.SystemGeolocList {
	var ret edpt.SystemGeolocList
	ret.Inst.ExcludeGeolocNameList = getSliceSystemGeolocListExcludeGeolocNameList(d.Get("exclude_geoloc_name_list").([]interface{}))
	ret.Inst.IncludeGeolocNameList = getSliceSystemGeolocListIncludeGeolocNameList(d.Get("include_geoloc_name_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceSystemGeolocListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
