package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_resources": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"res_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"res_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"res_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"dynamic_resources": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"res_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"res_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"res_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"per_resource_limit": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"res_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"res_alloc": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"res_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"hw_resource_limit": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"res_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"res_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"res_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"system_capacity_remaining": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"system_capacity_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"system_capacity_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"system_capacity": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosResourceUsageOperOper := setObjectDdosResourceUsageOperOper(res)
		d.Set("oper", DdosResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosResourceUsageOperOper(ret edpt.DataDdosResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"static_resources":           setSliceDdosResourceUsageOperOperStaticResources(ret.DtDdosResourceUsageOper.Oper.StaticResources),
			"dynamic_resources":          setSliceDdosResourceUsageOperOperDynamicResources(ret.DtDdosResourceUsageOper.Oper.DynamicResources),
			"per_resource_limit":         setSliceDdosResourceUsageOperOperPerResourceLimit(ret.DtDdosResourceUsageOper.Oper.PerResourceLimit),
			"hw_resource_limit":          setSliceDdosResourceUsageOperOperHwResourceLimit(ret.DtDdosResourceUsageOper.Oper.HwResourceLimit),
			"enabled":                    ret.DtDdosResourceUsageOper.Oper.Enabled,
			"system_capacity_remaining":  ret.DtDdosResourceUsageOper.Oper.SystemCapacityRemaining,
			"system_capacity_total":      ret.DtDdosResourceUsageOper.Oper.SystemCapacityTotal,
			"system_capacity_percentage": ret.DtDdosResourceUsageOper.Oper.SystemCapacityPercentage,
			"system_capacity":            ret.DtDdosResourceUsageOper.Oper.SystemCapacity,
		},
	}
}

func setSliceDdosResourceUsageOperOperStaticResources(d []edpt.DdosResourceUsageOperOperStaticResources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["res_name"] = item.ResName
		in["res_alloc"] = item.ResAlloc
		in["res_limit"] = item.ResLimit
		result = append(result, in)
	}
	return result
}

func setSliceDdosResourceUsageOperOperDynamicResources(d []edpt.DdosResourceUsageOperOperDynamicResources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["res_name"] = item.ResName
		in["res_alloc"] = item.ResAlloc
		in["res_limit"] = item.ResLimit
		result = append(result, in)
	}
	return result
}

func setSliceDdosResourceUsageOperOperPerResourceLimit(d []edpt.DdosResourceUsageOperOperPerResourceLimit) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["res_name"] = item.ResName
		in["res_alloc"] = item.ResAlloc
		in["res_limit"] = item.ResLimit
		result = append(result, in)
	}
	return result
}

func setSliceDdosResourceUsageOperOperHwResourceLimit(d []edpt.DdosResourceUsageOperOperHwResourceLimit) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["res_name"] = item.ResName
		in["res_alloc"] = item.ResAlloc
		in["res_limit"] = item.ResLimit
		result = append(result, in)
	}
	return result
}

func getObjectDdosResourceUsageOperOper(d []interface{}) edpt.DdosResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.DdosResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticResources = getSliceDdosResourceUsageOperOperStaticResources(in["static_resources"].([]interface{}))
		ret.DynamicResources = getSliceDdosResourceUsageOperOperDynamicResources(in["dynamic_resources"].([]interface{}))
		ret.PerResourceLimit = getSliceDdosResourceUsageOperOperPerResourceLimit(in["per_resource_limit"].([]interface{}))
		ret.HwResourceLimit = getSliceDdosResourceUsageOperOperHwResourceLimit(in["hw_resource_limit"].([]interface{}))
		ret.Enabled = in["enabled"].(int)
		ret.SystemCapacityRemaining = in["system_capacity_remaining"].(int)
		ret.SystemCapacityTotal = in["system_capacity_total"].(int)
		ret.SystemCapacityPercentage = in["system_capacity_percentage"].(int)
		ret.SystemCapacity = in["system_capacity"].(int)
	}
	return ret
}

func getSliceDdosResourceUsageOperOperStaticResources(d []interface{}) []edpt.DdosResourceUsageOperOperStaticResources {

	count1 := len(d)
	ret := make([]edpt.DdosResourceUsageOperOperStaticResources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosResourceUsageOperOperStaticResources
		oi.ResName = in["res_name"].(string)
		oi.ResAlloc = in["res_alloc"].(int)
		oi.ResLimit = in["res_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosResourceUsageOperOperDynamicResources(d []interface{}) []edpt.DdosResourceUsageOperOperDynamicResources {

	count1 := len(d)
	ret := make([]edpt.DdosResourceUsageOperOperDynamicResources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosResourceUsageOperOperDynamicResources
		oi.ResName = in["res_name"].(string)
		oi.ResAlloc = in["res_alloc"].(int)
		oi.ResLimit = in["res_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosResourceUsageOperOperPerResourceLimit(d []interface{}) []edpt.DdosResourceUsageOperOperPerResourceLimit {

	count1 := len(d)
	ret := make([]edpt.DdosResourceUsageOperOperPerResourceLimit, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosResourceUsageOperOperPerResourceLimit
		oi.ResName = in["res_name"].(string)
		oi.ResAlloc = in["res_alloc"].(string)
		oi.ResLimit = in["res_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosResourceUsageOperOperHwResourceLimit(d []interface{}) []edpt.DdosResourceUsageOperOperHwResourceLimit {

	count1 := len(d)
	ret := make([]edpt.DdosResourceUsageOperOperHwResourceLimit, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosResourceUsageOperOperHwResourceLimit
		oi.ResName = in["res_name"].(string)
		oi.ResAlloc = in["res_alloc"].(int)
		oi.ResLimit = in["res_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosResourceUsageOper(d *schema.ResourceData) edpt.DdosResourceUsageOper {
	var ret edpt.DdosResourceUsageOper

	ret.Oper = getObjectDdosResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
