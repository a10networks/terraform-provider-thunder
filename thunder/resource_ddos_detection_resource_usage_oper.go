package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_detection_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDetectionResourceUsageOperRead,

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
					},
				},
			},
		},
	}
}

func resourceDdosDetectionResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDetectionResourceUsageOperOper := setObjectDdosDetectionResourceUsageOperOper(res)
		d.Set("oper", DdosDetectionResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDetectionResourceUsageOperOper(ret edpt.DataDdosDetectionResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"static_resources":  setSliceDdosDetectionResourceUsageOperOperStaticResources(ret.DtDdosDetectionResourceUsageOper.Oper.StaticResources),
			"dynamic_resources": setSliceDdosDetectionResourceUsageOperOperDynamicResources(ret.DtDdosDetectionResourceUsageOper.Oper.DynamicResources),
		},
	}
}

func setSliceDdosDetectionResourceUsageOperOperStaticResources(d []edpt.DdosDetectionResourceUsageOperOperStaticResources) []map[string]interface{} {
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

func setSliceDdosDetectionResourceUsageOperOperDynamicResources(d []edpt.DdosDetectionResourceUsageOperOperDynamicResources) []map[string]interface{} {
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

func getObjectDdosDetectionResourceUsageOperOper(d []interface{}) edpt.DdosDetectionResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.DdosDetectionResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticResources = getSliceDdosDetectionResourceUsageOperOperStaticResources(in["static_resources"].([]interface{}))
		ret.DynamicResources = getSliceDdosDetectionResourceUsageOperOperDynamicResources(in["dynamic_resources"].([]interface{}))
	}
	return ret
}

func getSliceDdosDetectionResourceUsageOperOperStaticResources(d []interface{}) []edpt.DdosDetectionResourceUsageOperOperStaticResources {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionResourceUsageOperOperStaticResources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionResourceUsageOperOperStaticResources
		oi.ResName = in["res_name"].(string)
		oi.ResAlloc = in["res_alloc"].(int)
		oi.ResLimit = in["res_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDetectionResourceUsageOperOperDynamicResources(d []interface{}) []edpt.DdosDetectionResourceUsageOperOperDynamicResources {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionResourceUsageOperOperDynamicResources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionResourceUsageOperOperDynamicResources
		oi.ResName = in["res_name"].(string)
		oi.ResAlloc = in["res_alloc"].(int)
		oi.ResLimit = in["res_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDetectionResourceUsageOper(d *schema.ResourceData) edpt.DdosDetectionResourceUsageOper {
	var ret edpt.DdosDetectionResourceUsageOper

	ret.Oper = getObjectDdosDetectionResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
