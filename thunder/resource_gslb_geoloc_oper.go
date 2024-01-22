package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbGeolocOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_geoloc_oper`: Operational Status for the object geoloc\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbGeolocOperRead,

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

func resourceGslbGeolocOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGeolocOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGeolocOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbGeolocOperOper := setObjectGslbGeolocOperOper(res)
		d.Set("oper", GslbGeolocOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbGeolocOperOper(ret edpt.DataGslbGeolocOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geoloc_list":   setSliceGslbGeolocOperOperGeolocList(ret.DtGslbGeolocOper.Oper.GeolocList),
			"total_geolocs": ret.DtGslbGeolocOper.Oper.TotalGeolocs,
			"geo_name":      ret.DtGslbGeolocOper.Oper.GeoName,
			"filter1":       ret.DtGslbGeolocOper.Oper.Filter1,
			"filter2":       ret.DtGslbGeolocOper.Oper.Filter2,
			"filter3":       ret.DtGslbGeolocOper.Oper.Filter3,
			"filter4":       ret.DtGslbGeolocOper.Oper.Filter4,
			"pol_name":      ret.DtGslbGeolocOper.Oper.PolName,
			"iprangestrt":   ret.DtGslbGeolocOper.Oper.Iprangestrt,
			"iprangeend":    ret.DtGslbGeolocOper.Oper.Iprangeend,
			"ipv6rangestrt": ret.DtGslbGeolocOper.Oper.Ipv6rangestrt,
			"depth":         ret.DtGslbGeolocOper.Oper.Depth,
		},
	}
}

func setSliceGslbGeolocOperOperGeolocList(d []edpt.GslbGeolocOperOperGeolocList) []map[string]interface{} {
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

func getObjectGslbGeolocOperOper(d []interface{}) edpt.GslbGeolocOperOper {

	count1 := len(d)
	var ret edpt.GslbGeolocOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeolocList = getSliceGslbGeolocOperOperGeolocList(in["geoloc_list"].([]interface{}))
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

func getSliceGslbGeolocOperOperGeolocList(d []interface{}) []edpt.GslbGeolocOperOperGeolocList {

	count1 := len(d)
	ret := make([]edpt.GslbGeolocOperOperGeolocList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbGeolocOperOperGeolocList
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

func dataToEndpointGslbGeolocOper(d *schema.ResourceData) edpt.GslbGeolocOper {
	var ret edpt.GslbGeolocOper

	ret.Oper = getObjectGslbGeolocOperOper(d.Get("oper").([]interface{}))
	return ret
}
