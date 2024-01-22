package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sflow_global_oper`: Operational Status for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceSflowGlobalOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"if_stats_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"if_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"if_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_sample_records": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"counter_sample_records": {
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

func resourceSflowGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SflowGlobalOperOper := setObjectSflowGlobalOperOper(res)
		d.Set("oper", SflowGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSflowGlobalOperOper(ret edpt.DataSflowGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"if_stats_list": setSliceSflowGlobalOperOperIfStatsList(ret.DtSflowGlobalOper.Oper.IfStatsList),
		},
	}
}

func setSliceSflowGlobalOperOperIfStatsList(d []edpt.SflowGlobalOperOperIfStatsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["if_type"] = item.IfType
		in["if_num"] = item.IfNum
		in["packet_sample_records"] = item.PacketSampleRecords
		in["counter_sample_records"] = item.CounterSampleRecords
		result = append(result, in)
	}
	return result
}

func getObjectSflowGlobalOperOper(d []interface{}) edpt.SflowGlobalOperOper {

	count1 := len(d)
	var ret edpt.SflowGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IfStatsList = getSliceSflowGlobalOperOperIfStatsList(in["if_stats_list"].([]interface{}))
	}
	return ret
}

func getSliceSflowGlobalOperOperIfStatsList(d []interface{}) []edpt.SflowGlobalOperOperIfStatsList {

	count1 := len(d)
	ret := make([]edpt.SflowGlobalOperOperIfStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowGlobalOperOperIfStatsList
		oi.IfType = in["if_type"].(string)
		oi.IfNum = in["if_num"].(int)
		oi.PacketSampleRecords = in["packet_sample_records"].(int)
		oi.CounterSampleRecords = in["counter_sample_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSflowGlobalOper(d *schema.ResourceData) edpt.SflowGlobalOper {
	var ret edpt.SflowGlobalOper

	ret.Oper = getObjectSflowGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
