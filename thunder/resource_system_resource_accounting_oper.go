package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceAccountingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_resource_accounting_oper`: Operational Status for the object resource-accounting\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemResourceAccountingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"partition_resource": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"res_type": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"resource_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"resources": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"resource_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"current": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"minimum": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"maximum": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"utilization": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"max_exceed": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"threshold_exceed": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"average": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"peak": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"cap": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"cap_utilization": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
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

func resourceSystemResourceAccountingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemResourceAccountingOperOper := setObjectSystemResourceAccountingOperOper(res)
		d.Set("oper", SystemResourceAccountingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemResourceAccountingOperOper(ret edpt.DataSystemResourceAccountingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scope":              ret.DtSystemResourceAccountingOper.Oper.Scope,
			"partition_resource": setSliceSystemResourceAccountingOperOperPartitionResource(ret.DtSystemResourceAccountingOper.Oper.PartitionResource),
		},
	}
}

func setSliceSystemResourceAccountingOperOperPartitionResource(d []edpt.SystemResourceAccountingOperOperPartitionResource) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["partition_name"] = item.PartitionName
		in["res_type"] = setSliceSystemResourceAccountingOperOperPartitionResourceResType(item.ResType)
		result = append(result, in)
	}
	return result
}

func setSliceSystemResourceAccountingOperOperPartitionResourceResType(d []edpt.SystemResourceAccountingOperOperPartitionResourceResType) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["resource_type"] = item.ResourceType
		in["resources"] = setSliceSystemResourceAccountingOperOperPartitionResourceResTypeResources(item.Resources)
		result = append(result, in)
	}
	return result
}

func setSliceSystemResourceAccountingOperOperPartitionResourceResTypeResources(d []edpt.SystemResourceAccountingOperOperPartitionResourceResTypeResources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["resource_name"] = item.ResourceName
		in["current"] = item.Current
		in["minimum"] = item.Minimum
		in["maximum"] = item.Maximum
		in["utilization"] = item.Utilization
		in["max_exceed"] = item.MaxExceed
		in["threshold_exceed"] = item.ThresholdExceed
		in["average"] = item.Average
		in["peak"] = item.Peak
		in["cap"] = item.Cap
		in["cap_utilization"] = item.CapUtilization
		result = append(result, in)
	}
	return result
}

func getObjectSystemResourceAccountingOperOper(d []interface{}) edpt.SystemResourceAccountingOperOper {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Scope = in["scope"].(string)
		ret.PartitionResource = getSliceSystemResourceAccountingOperOperPartitionResource(in["partition_resource"].([]interface{}))
	}
	return ret
}

func getSliceSystemResourceAccountingOperOperPartitionResource(d []interface{}) []edpt.SystemResourceAccountingOperOperPartitionResource {

	count1 := len(d)
	ret := make([]edpt.SystemResourceAccountingOperOperPartitionResource, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemResourceAccountingOperOperPartitionResource
		oi.PartitionName = in["partition_name"].(string)
		oi.ResType = getSliceSystemResourceAccountingOperOperPartitionResourceResType(in["res_type"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemResourceAccountingOperOperPartitionResourceResType(d []interface{}) []edpt.SystemResourceAccountingOperOperPartitionResourceResType {

	count1 := len(d)
	ret := make([]edpt.SystemResourceAccountingOperOperPartitionResourceResType, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemResourceAccountingOperOperPartitionResourceResType
		oi.ResourceType = in["resource_type"].(string)
		oi.Resources = getSliceSystemResourceAccountingOperOperPartitionResourceResTypeResources(in["resources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemResourceAccountingOperOperPartitionResourceResTypeResources(d []interface{}) []edpt.SystemResourceAccountingOperOperPartitionResourceResTypeResources {

	count1 := len(d)
	ret := make([]edpt.SystemResourceAccountingOperOperPartitionResourceResTypeResources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemResourceAccountingOperOperPartitionResourceResTypeResources
		oi.ResourceName = in["resource_name"].(string)
		oi.Current = in["current"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Utilization = in["utilization"].(string)
		oi.MaxExceed = in["max_exceed"].(string)
		oi.ThresholdExceed = in["threshold_exceed"].(string)
		oi.Average = in["average"].(string)
		oi.Peak = in["peak"].(string)
		oi.Cap = in["cap"].(string)
		oi.CapUtilization = in["cap_utilization"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemResourceAccountingOper(d *schema.ResourceData) edpt.SystemResourceAccountingOper {
	var ret edpt.SystemResourceAccountingOper

	ret.Oper = getObjectSystemResourceAccountingOperOper(d.Get("oper").([]interface{}))
	return ret
}
