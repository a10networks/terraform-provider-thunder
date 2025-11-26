package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTopkDestinationsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_topk_destinations_oper`: Operational Status for the object topk-destinations\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectTopkDestinationsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"indicators": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicator_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"indicator_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"destinations": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"topk_type": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reset_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}

func resourceDdosNetworkObjectTopkDestinationsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTopkDestinationsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTopkDestinationsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectTopkDestinationsOperOper := setObjectDdosNetworkObjectTopkDestinationsOperOper(res)
		d.Set("oper", DdosNetworkObjectTopkDestinationsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectTopkDestinationsOperOper(ret edpt.DataDdosNetworkObjectTopkDestinationsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators": setSliceDdosNetworkObjectTopkDestinationsOperOperIndicators(ret.DtDdosNetworkObjectTopkDestinationsOper.Oper.Indicators),
			"topk_type":  ret.DtDdosNetworkObjectTopkDestinationsOper.Oper.TopkType,
			"reset_time": ret.DtDdosNetworkObjectTopkDestinationsOper.Oper.ResetTime,
		},
	}
}

func setSliceDdosNetworkObjectTopkDestinationsOperOperIndicators(d []edpt.DdosNetworkObjectTopkDestinationsOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations(d []edpt.DdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func getObjectDdosNetworkObjectTopkDestinationsOperOper(d []interface{}) edpt.DdosNetworkObjectTopkDestinationsOperOper {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTopkDestinationsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosNetworkObjectTopkDestinationsOperOperIndicators(in["indicators"].([]interface{}))
		ret.TopkType = in["topk_type"].(int)
		ret.ResetTime = in["reset_time"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectTopkDestinationsOperOperIndicators(d []interface{}) []edpt.DdosNetworkObjectTopkDestinationsOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectTopkDestinationsOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectTopkDestinationsOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations(d []interface{}) []edpt.DdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectTopkDestinationsOper(d *schema.ResourceData) edpt.DdosNetworkObjectTopkDestinationsOper {
	var ret edpt.DdosNetworkObjectTopkDestinationsOper

	ret.Oper = getObjectDdosNetworkObjectTopkDestinationsOperOper(d.Get("oper").([]interface{}))

	ret.ObjectName = d.Get("object_name").(string)
	return ret
}
