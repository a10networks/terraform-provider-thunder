package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosGeoLocationDbOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_geo_location_db_oper`: Operational Status for the object db\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosGeoLocationDbOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_list": {
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
									"subcnt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_geo_location": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosGeoLocationDbOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosGeoLocationDbOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosGeoLocationDbOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosGeoLocationDbOperOper := setObjectDdosGeoLocationDbOperOper(res)
		d.Set("oper", DdosGeoLocationDbOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosGeoLocationDbOperOper(ret edpt.DataDdosGeoLocationDbOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geo_location_list":  setSliceDdosGeoLocationDbOperOperGeoLocationList(ret.DtDdosGeoLocationDbOper.Oper.GeoLocationList),
			"total_geo_location": ret.DtDdosGeoLocationDbOper.Oper.TotalGeoLocation,
			"filter":             ret.DtDdosGeoLocationDbOper.Oper.Filter,
			"ip":                 ret.DtDdosGeoLocationDbOper.Oper.Ip,
			"ipv6":               ret.DtDdosGeoLocationDbOper.Oper.Ipv6,
		},
	}
}

func setSliceDdosGeoLocationDbOperOperGeoLocationList(d []edpt.DdosGeoLocationDbOperOperGeoLocationList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["from"] = item.From
		in["tomask"] = item.Tomask
		in["subcnt"] = item.Subcnt
		in["type"] = item.Type
		result = append(result, in)
	}
	return result
}

func getObjectDdosGeoLocationDbOperOper(d []interface{}) edpt.DdosGeoLocationDbOperOper {

	count1 := len(d)
	var ret edpt.DdosGeoLocationDbOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeoLocationList = getSliceDdosGeoLocationDbOperOperGeoLocationList(in["geo_location_list"].([]interface{}))
		ret.TotalGeoLocation = in["total_geo_location"].(int)
		ret.Filter = in["filter"].(string)
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
	}
	return ret
}

func getSliceDdosGeoLocationDbOperOperGeoLocationList(d []interface{}) []edpt.DdosGeoLocationDbOperOperGeoLocationList {

	count1 := len(d)
	ret := make([]edpt.DdosGeoLocationDbOperOperGeoLocationList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosGeoLocationDbOperOperGeoLocationList
		oi.Name = in["name"].(string)
		oi.From = in["from"].(string)
		oi.Tomask = in["tomask"].(string)
		oi.Subcnt = in["subcnt"].(int)
		oi.Type = in["type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosGeoLocationDbOper(d *schema.ResourceData) edpt.DdosGeoLocationDbOper {
	var ret edpt.DdosGeoLocationDbOper

	ret.Oper = getObjectDdosGeoLocationDbOperOper(d.Get("oper").([]interface{}))
	return ret
}
