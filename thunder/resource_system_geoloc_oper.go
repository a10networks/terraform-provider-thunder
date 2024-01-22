package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geoloc_oper`: Operational Status for the object geoloc\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geoloc_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"from": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"tomask": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"last": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"subcnt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_geolocs": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geo_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter1": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pol_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"iprangestrt": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"iprangeend": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6rangestrt": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"depth": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemGeolocOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocOperOper := setObjectSystemGeolocOperOper(res)
		d.Set("oper", SystemGeolocOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocOperOper(ret edpt.DataSystemGeolocOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geoloc_list":   setSliceSystemGeolocOperOperGeolocList(ret.DtSystemGeolocOper.Oper.GeolocList),
			"total_geolocs": ret.DtSystemGeolocOper.Oper.TotalGeolocs,
			"geo_name":      ret.DtSystemGeolocOper.Oper.GeoName,
			"filter1":       ret.DtSystemGeolocOper.Oper.Filter1,
			"filter2":       ret.DtSystemGeolocOper.Oper.Filter2,
			"filter3":       ret.DtSystemGeolocOper.Oper.Filter3,
			"filter4":       ret.DtSystemGeolocOper.Oper.Filter4,
			"pol_name":      ret.DtSystemGeolocOper.Oper.PolName,
			"iprangestrt":   ret.DtSystemGeolocOper.Oper.Iprangestrt,
			"iprangeend":    ret.DtSystemGeolocOper.Oper.Iprangeend,
			"ipv6rangestrt": ret.DtSystemGeolocOper.Oper.Ipv6rangestrt,
			"depth":         ret.DtSystemGeolocOper.Oper.Depth,
		},
	}
}

func setSliceSystemGeolocOperOperGeolocList(d []edpt.SystemGeolocOperOperGeolocList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["from"] = item.From
		in["tomask"] = item.Tomask
		in["last"] = item.Last
		in["hits"] = item.Hits
		in["subcnt"] = item.Subcnt
		in["type"] = item.Type
		result = append(result, in)
	}
	return result
}

func getObjectSystemGeolocOperOper(d []interface{}) edpt.SystemGeolocOperOper {

	count1 := len(d)
	var ret edpt.SystemGeolocOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeolocList = getSliceSystemGeolocOperOperGeolocList(in["geoloc_list"].([]interface{}))
		ret.TotalGeolocs = in["total_geolocs"].(int)
		ret.GeoName = in["geo_name"].(string)
		ret.Filter1 = in["filter1"].(string)
		ret.Filter2 = in["filter2"].(string)
		ret.Filter3 = in["filter3"].(string)
		ret.Filter4 = in["filter4"].(string)
		ret.PolName = in["pol_name"].(string)
		ret.Iprangestrt = in["iprangestrt"].(string)
		ret.Iprangeend = in["iprangeend"].(string)
		ret.Ipv6rangestrt = in["ipv6rangestrt"].(string)
		ret.Depth = in["depth"].(int)
	}
	return ret
}

func getSliceSystemGeolocOperOperGeolocList(d []interface{}) []edpt.SystemGeolocOperOperGeolocList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocOperOperGeolocList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocOperOperGeolocList
		oi.Name = in["name"].(string)
		oi.From = in["from"].(string)
		oi.Tomask = in["tomask"].(string)
		oi.Last = in["last"].(string)
		oi.Hits = in["hits"].(int)
		oi.Subcnt = in["subcnt"].(int)
		oi.Type = in["type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocOper(d *schema.ResourceData) edpt.SystemGeolocOper {
	var ret edpt.SystemGeolocOper

	ret.Oper = getObjectSystemGeolocOperOper(d.Get("oper").([]interface{}))
	return ret
}
