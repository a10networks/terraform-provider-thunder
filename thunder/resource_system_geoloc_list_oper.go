package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geoloc_list_oper`: Operational Status for the object geoloc-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocListOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of Geolocation list",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geoloc_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"geoloc_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSystemGeolocListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocListOperOper := setObjectSystemGeolocListOperOper(res)
		d.Set("oper", SystemGeolocListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocListOperOper(ret edpt.DataSystemGeolocListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geoloc_list": setSliceSystemGeolocListOperOperGeolocList(ret.DtSystemGeolocListOper.Oper.GeolocList),
		},
	}
}

func setSliceSystemGeolocListOperOperGeolocList(d []edpt.SystemGeolocListOperOperGeolocList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["geoloc_name"] = item.GeolocName
		in["hit_count"] = item.HitCount
		in["active"] = item.Active
		result = append(result, in)
	}
	return result
}

func getObjectSystemGeolocListOperOper(d []interface{}) edpt.SystemGeolocListOperOper {

	count1 := len(d)
	var ret edpt.SystemGeolocListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeolocList = getSliceSystemGeolocListOperOperGeolocList(in["geoloc_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeolocListOperOperGeolocList(d []interface{}) []edpt.SystemGeolocListOperOperGeolocList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListOperOperGeolocList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListOperOperGeolocList
		oi.Type = in["type"].(string)
		oi.GeolocName = in["geoloc_name"].(string)
		oi.HitCount = in["hit_count"].(int)
		oi.Active = in["active"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocListOper(d *schema.ResourceData) edpt.SystemGeolocListOper {
	var ret edpt.SystemGeolocListOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectSystemGeolocListOperOper(d.Get("oper").([]interface{}))
	return ret
}
