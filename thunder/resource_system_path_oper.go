package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPathOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_path_oper`: Operational Status for the object path\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPathOperRead,

		Schema: map[string]*schema.Schema{
			"l2hm_path_name": {
				Type: schema.TypeString, Required: true, Description: "Monitor Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l2hm_path": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"endpoint_1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"endpoint_2": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"health_check": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"path_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"apps": {
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

func resourceSystemPathOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPathOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPathOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPathOperOper := setObjectSystemPathOperOper(res)
		d.Set("oper", SystemPathOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPathOperOper(ret edpt.DataSystemPathOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"path_list": setSliceSystemPathOperOperPath_list(ret.DtSystemPathOper.Oper.Path_list),
		},
	}
}

func setSliceSystemPathOperOperPath_list(d []edpt.SystemPathOperOperPath_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["l2hm_path"] = item.L2hm_path
		in["endpoint_1"] = item.Endpoint_1
		in["endpoint_2"] = item.Endpoint_2
		in["health_check"] = item.Health_check
		in["path_state"] = item.Path_state
		in["apps"] = item.Apps
		result = append(result, in)
	}
	return result
}

func getObjectSystemPathOperOper(d []interface{}) edpt.SystemPathOperOper {

	count1 := len(d)
	var ret edpt.SystemPathOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Path_list = getSliceSystemPathOperOperPath_list(in["path_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemPathOperOperPath_list(d []interface{}) []edpt.SystemPathOperOperPath_list {

	count1 := len(d)
	ret := make([]edpt.SystemPathOperOperPath_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPathOperOperPath_list
		oi.L2hm_path = in["l2hm_path"].(string)
		oi.Endpoint_1 = in["endpoint_1"].(string)
		oi.Endpoint_2 = in["endpoint_2"].(string)
		oi.Health_check = in["health_check"].(string)
		oi.Path_state = in["path_state"].(string)
		oi.Apps = in["apps"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemPathOper(d *schema.ResourceData) edpt.SystemPathOper {
	var ret edpt.SystemPathOper

	ret.L2hmPathName = d.Get("l2hm_path_name").(string)

	ret.Oper = getObjectSystemPathOperOper(d.Get("oper").([]interface{}))
	return ret
}
