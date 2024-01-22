package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCosqStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cosq_stats_oper`: Operational Status for the object cosq-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCosqStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cosq_index": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cosq_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"xaui_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cosq_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cosq_inpkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cosq_outpkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cosq_sharedpkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cosq_ucpeakpkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cosq_uccongdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cosq_congdrop": {
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

func resourceSystemCosqStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCosqStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCosqStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCosqStatsOperOper := setObjectSystemCosqStatsOperOper(res)
		d.Set("oper", SystemCosqStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCosqStatsOperOper(ret edpt.DataSystemCosqStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cosq_index": ret.DtSystemCosqStatsOper.Oper.CosqIndex,
			"cosq_list":  setSliceSystemCosqStatsOperOperCosqList(ret.DtSystemCosqStatsOper.Oper.CosqList),
		},
	}
}

func setSliceSystemCosqStatsOperOperCosqList(d []edpt.SystemCosqStatsOperOperCosqList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ethernet_index"] = item.EthernetIndex
		in["xaui_index"] = item.XauiIndex
		in["interface_name"] = item.InterfaceName
		in["cosq_drop"] = item.CosqDrop
		in["cosq_inpkt"] = item.CosqInpkt
		in["cosq_outpkt"] = item.CosqOutpkt
		in["cosq_sharedpkt"] = item.CosqSharedpkt
		in["cosq_ucpeakpkt"] = item.CosqUcpeakpkt
		in["cosq_uccongdrop"] = item.CosqUccongdrop
		in["cosq_congdrop"] = item.CosqCongdrop
		result = append(result, in)
	}
	return result
}

func getObjectSystemCosqStatsOperOper(d []interface{}) edpt.SystemCosqStatsOperOper {

	count1 := len(d)
	var ret edpt.SystemCosqStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CosqIndex = in["cosq_index"].(int)
		ret.CosqList = getSliceSystemCosqStatsOperOperCosqList(in["cosq_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemCosqStatsOperOperCosqList(d []interface{}) []edpt.SystemCosqStatsOperOperCosqList {

	count1 := len(d)
	ret := make([]edpt.SystemCosqStatsOperOperCosqList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCosqStatsOperOperCosqList
		oi.EthernetIndex = in["ethernet_index"].(int)
		oi.XauiIndex = in["xaui_index"].(int)
		oi.InterfaceName = in["interface_name"].(string)
		oi.CosqDrop = in["cosq_drop"].(int)
		oi.CosqInpkt = in["cosq_inpkt"].(int)
		oi.CosqOutpkt = in["cosq_outpkt"].(int)
		oi.CosqSharedpkt = in["cosq_sharedpkt"].(int)
		oi.CosqUcpeakpkt = in["cosq_ucpeakpkt"].(int)
		oi.CosqUccongdrop = in["cosq_uccongdrop"].(int)
		oi.CosqCongdrop = in["cosq_congdrop"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCosqStatsOper(d *schema.ResourceData) edpt.SystemCosqStatsOper {
	var ret edpt.SystemCosqStatsOper

	ret.Oper = getObjectSystemCosqStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
