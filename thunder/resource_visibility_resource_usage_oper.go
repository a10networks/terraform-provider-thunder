package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"resource_current": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resource_limit": {
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

func resourceVisibilityResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityResourceUsageOperOper := setObjectVisibilityResourceUsageOperOper(res)
		d.Set("oper", VisibilityResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityResourceUsageOperOper(ret edpt.DataVisibilityResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"resource_list": setSliceVisibilityResourceUsageOperOperResourceList(ret.DtVisibilityResourceUsageOper.Oper.ResourceList),
		},
	}
}

func setSliceVisibilityResourceUsageOperOperResourceList(d []edpt.VisibilityResourceUsageOperOperResourceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["resource_name"] = item.ResourceName
		in["resource_current"] = item.ResourceCurrent
		in["resource_limit"] = item.ResourceLimit
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityResourceUsageOperOper(d []interface{}) edpt.VisibilityResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.VisibilityResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ResourceList = getSliceVisibilityResourceUsageOperOperResourceList(in["resource_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityResourceUsageOperOperResourceList(d []interface{}) []edpt.VisibilityResourceUsageOperOperResourceList {

	count1 := len(d)
	ret := make([]edpt.VisibilityResourceUsageOperOperResourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityResourceUsageOperOperResourceList
		oi.ResourceName = in["resource_name"].(string)
		oi.ResourceCurrent = in["resource_current"].(int)
		oi.ResourceLimit = in["resource_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityResourceUsageOper(d *schema.ResourceData) edpt.VisibilityResourceUsageOper {
	var ret edpt.VisibilityResourceUsageOper

	ret.Oper = getObjectVisibilityResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
