package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_automatic_update_info_oper`: Operational Status for the object info\n\n__PLACEHOLDER__",
		ReadContext: resourceAutomaticUpdateInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"feature_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"schedule": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"last_update": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"next_check": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceAutomaticUpdateInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AutomaticUpdateInfoOperOper := setObjectAutomaticUpdateInfoOperOper(res)
		d.Set("oper", AutomaticUpdateInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAutomaticUpdateInfoOperOper(ret edpt.DataAutomaticUpdateInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"feature_list": setSliceAutomaticUpdateInfoOperOperFeatureList(ret.DtAutomaticUpdateInfoOper.Oper.FeatureList),
		},
	}
}

func setSliceAutomaticUpdateInfoOperOperFeatureList(d []edpt.AutomaticUpdateInfoOperOperFeatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["feature_name"] = item.FeatureName
		in["version"] = item.Version
		in["schedule"] = item.Schedule
		in["time"] = item.Time
		in["last_update"] = item.LastUpdate
		in["next_check"] = item.NextCheck
		result = append(result, in)
	}
	return result
}

func getObjectAutomaticUpdateInfoOperOper(d []interface{}) edpt.AutomaticUpdateInfoOperOper {

	count1 := len(d)
	var ret edpt.AutomaticUpdateInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FeatureList = getSliceAutomaticUpdateInfoOperOperFeatureList(in["feature_list"].([]interface{}))
	}
	return ret
}

func getSliceAutomaticUpdateInfoOperOperFeatureList(d []interface{}) []edpt.AutomaticUpdateInfoOperOperFeatureList {

	count1 := len(d)
	ret := make([]edpt.AutomaticUpdateInfoOperOperFeatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AutomaticUpdateInfoOperOperFeatureList
		oi.FeatureName = in["feature_name"].(string)
		oi.Version = in["version"].(string)
		oi.Schedule = in["schedule"].(string)
		oi.Time = in["time"].(string)
		oi.LastUpdate = in["last_update"].(string)
		oi.NextCheck = in["next_check"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAutomaticUpdateInfoOper(d *schema.ResourceData) edpt.AutomaticUpdateInfoOper {
	var ret edpt.AutomaticUpdateInfoOper

	ret.Oper = getObjectAutomaticUpdateInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
