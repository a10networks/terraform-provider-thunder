package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocNameHelperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geoloc_name_helper_oper`: Operational Status for the object geoloc-name-helper\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocNameHelperOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geoloc_candidate_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geoloc_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"has_subregion": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"geoloc": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemGeolocNameHelperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocNameHelperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocNameHelperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocNameHelperOperOper := setObjectSystemGeolocNameHelperOperOper(res)
		d.Set("oper", SystemGeolocNameHelperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocNameHelperOperOper(ret edpt.DataSystemGeolocNameHelperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geoloc_candidate_list": setSliceSystemGeolocNameHelperOperOperGeolocCandidateList(ret.DtSystemGeolocNameHelperOper.Oper.GeolocCandidateList),
			"geoloc":                ret.DtSystemGeolocNameHelperOper.Oper.Geoloc,
		},
	}
}

func setSliceSystemGeolocNameHelperOperOperGeolocCandidateList(d []edpt.SystemGeolocNameHelperOperOperGeolocCandidateList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["geoloc_name"] = item.GeolocName
		in["has_subregion"] = item.HasSubregion
		result = append(result, in)
	}
	return result
}

func getObjectSystemGeolocNameHelperOperOper(d []interface{}) edpt.SystemGeolocNameHelperOperOper {

	count1 := len(d)
	var ret edpt.SystemGeolocNameHelperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeolocCandidateList = getSliceSystemGeolocNameHelperOperOperGeolocCandidateList(in["geoloc_candidate_list"].([]interface{}))
		ret.Geoloc = in["geoloc"].(string)
	}
	return ret
}

func getSliceSystemGeolocNameHelperOperOperGeolocCandidateList(d []interface{}) []edpt.SystemGeolocNameHelperOperOperGeolocCandidateList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocNameHelperOperOperGeolocCandidateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocNameHelperOperOperGeolocCandidateList
		oi.GeolocName = in["geoloc_name"].(string)
		oi.HasSubregion = in["has_subregion"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocNameHelperOper(d *schema.ResourceData) edpt.SystemGeolocNameHelperOper {
	var ret edpt.SystemGeolocNameHelperOper

	ret.Oper = getObjectSystemGeolocNameHelperOperOper(d.Get("oper").([]interface{}))
	return ret
}
