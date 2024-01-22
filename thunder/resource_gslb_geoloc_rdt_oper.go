package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbGeolocRdtOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_geoloc_rdt_oper`: Operational Status for the object geoloc-rdt\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbGeolocRdtOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geoloc_rdt_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gl_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"site_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rdt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_rdt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geo_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"site_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"active_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceGslbGeolocRdtOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGeolocRdtOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGeolocRdtOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbGeolocRdtOperOper := setObjectGslbGeolocRdtOperOper(res)
		d.Set("oper", GslbGeolocRdtOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbGeolocRdtOperOper(ret edpt.DataGslbGeolocRdtOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geoloc_rdt_list": setSliceGslbGeolocRdtOperOperGeolocRdtList(ret.DtGslbGeolocRdtOper.Oper.GeolocRdtList),
			"total_rdt":       ret.DtGslbGeolocRdtOper.Oper.TotalRdt,
			"geo_name":        ret.DtGslbGeolocRdtOper.Oper.GeoName,
			"site_name":       ret.DtGslbGeolocRdtOper.Oper.SiteName,
			"active_status":   ret.DtGslbGeolocRdtOper.Oper.ActiveStatus,
		},
	}
}

func setSliceGslbGeolocRdtOperOperGeolocRdtList(d []edpt.GslbGeolocRdtOperOperGeolocRdtList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["gl_name"] = item.GlName
		in["site_name"] = item.SiteName
		in["type"] = item.Type
		in["rdt"] = item.Rdt
		in["age"] = item.Age
		result = append(result, in)
	}
	return result
}

func getObjectGslbGeolocRdtOperOper(d []interface{}) edpt.GslbGeolocRdtOperOper {

	count1 := len(d)
	var ret edpt.GslbGeolocRdtOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeolocRdtList = getSliceGslbGeolocRdtOperOperGeolocRdtList(in["geoloc_rdt_list"].([]interface{}))
		ret.TotalRdt = in["total_rdt"].(int)
		ret.GeoName = in["geo_name"].(string)
		ret.SiteName = in["site_name"].(string)
		ret.ActiveStatus = in["active_status"].(string)
	}
	return ret
}

func getSliceGslbGeolocRdtOperOperGeolocRdtList(d []interface{}) []edpt.GslbGeolocRdtOperOperGeolocRdtList {

	count1 := len(d)
	ret := make([]edpt.GslbGeolocRdtOperOperGeolocRdtList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbGeolocRdtOperOperGeolocRdtList
		oi.GlName = in["gl_name"].(string)
		oi.SiteName = in["site_name"].(string)
		oi.Type = in["type"].(string)
		oi.Rdt = in["rdt"].(int)
		oi.Age = in["age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbGeolocRdtOper(d *schema.ResourceData) edpt.GslbGeolocRdtOper {
	var ret edpt.GslbGeolocRdtOper

	ret.Oper = getObjectGslbGeolocRdtOperOper(d.Get("oper").([]interface{}))
	return ret
}
