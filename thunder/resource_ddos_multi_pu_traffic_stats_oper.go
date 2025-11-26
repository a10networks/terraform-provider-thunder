package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosMultiPuTrafficStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_multi_pu_traffic_stats_oper`: Operational Status for the object multi-pu-traffic-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosMultiPuTrafficStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pu1_cpu_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pu1_throughput_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pu2_cpu_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pu2_throughput_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pu_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_top": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"pu1_kbit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"pu2_kbit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"pkt_top": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"pu1_pkt": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"pu2_pkt": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"cpu_top": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"pu1_cpu": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"pu2_cpu": {
													Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceDdosMultiPuTrafficStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosMultiPuTrafficStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosMultiPuTrafficStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosMultiPuTrafficStatsOperOper := setObjectDdosMultiPuTrafficStatsOperOper(res)
		d.Set("oper", DdosMultiPuTrafficStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosMultiPuTrafficStatsOperOper(ret edpt.DataDdosMultiPuTrafficStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pu1_cpu_rate":        ret.DtDdosMultiPuTrafficStatsOper.Oper.Pu1CpuRate,
			"pu1_throughput_rate": ret.DtDdosMultiPuTrafficStatsOper.Oper.Pu1ThroughputRate,
			"pu2_cpu_rate":        ret.DtDdosMultiPuTrafficStatsOper.Oper.Pu2CpuRate,
			"pu2_throughput_rate": ret.DtDdosMultiPuTrafficStatsOper.Oper.Pu2ThroughputRate,
			"pu_list":             setSliceDdosMultiPuTrafficStatsOperOperPuList(ret.DtDdosMultiPuTrafficStatsOper.Oper.PuList),
		},
	}
}

func setSliceDdosMultiPuTrafficStatsOperOperPuList(d []edpt.DdosMultiPuTrafficStatsOperOperPuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pu_index"] = item.PuIndex
		in["kbit_top"] = setSliceDdosMultiPuTrafficStatsOperOperPuListKbitTop(item.KbitTop)
		in["pkt_top"] = setSliceDdosMultiPuTrafficStatsOperOperPuListPktTop(item.PktTop)
		in["cpu_top"] = setSliceDdosMultiPuTrafficStatsOperOperPuListCpuTop(item.CpuTop)
		result = append(result, in)
	}
	return result
}

func setSliceDdosMultiPuTrafficStatsOperOperPuListKbitTop(d []edpt.DdosMultiPuTrafficStatsOperOperPuListKbitTop) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["zone_name"] = item.ZoneName
		in["pu1_kbit"] = item.Pu1Kbit
		in["pu2_kbit"] = item.Pu2Kbit
		result = append(result, in)
	}
	return result
}

func setSliceDdosMultiPuTrafficStatsOperOperPuListPktTop(d []edpt.DdosMultiPuTrafficStatsOperOperPuListPktTop) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["zone_name"] = item.ZoneName
		in["pu1_pkt"] = item.Pu1Pkt
		in["pu2_pkt"] = item.Pu2Pkt
		result = append(result, in)
	}
	return result
}

func setSliceDdosMultiPuTrafficStatsOperOperPuListCpuTop(d []edpt.DdosMultiPuTrafficStatsOperOperPuListCpuTop) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["zone_name"] = item.ZoneName
		in["pu1_cpu"] = item.Pu1Cpu
		in["pu2_cpu"] = item.Pu2Cpu
		result = append(result, in)
	}
	return result
}

func getObjectDdosMultiPuTrafficStatsOperOper(d []interface{}) edpt.DdosMultiPuTrafficStatsOperOper {

	count1 := len(d)
	var ret edpt.DdosMultiPuTrafficStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Pu1CpuRate = in["pu1_cpu_rate"].(int)
		ret.Pu1ThroughputRate = in["pu1_throughput_rate"].(int)
		ret.Pu2CpuRate = in["pu2_cpu_rate"].(int)
		ret.Pu2ThroughputRate = in["pu2_throughput_rate"].(int)
		ret.PuList = getSliceDdosMultiPuTrafficStatsOperOperPuList(in["pu_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosMultiPuTrafficStatsOperOperPuList(d []interface{}) []edpt.DdosMultiPuTrafficStatsOperOperPuList {

	count1 := len(d)
	ret := make([]edpt.DdosMultiPuTrafficStatsOperOperPuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosMultiPuTrafficStatsOperOperPuList
		oi.PuIndex = in["pu_index"].(int)
		oi.KbitTop = getSliceDdosMultiPuTrafficStatsOperOperPuListKbitTop(in["kbit_top"].([]interface{}))
		oi.PktTop = getSliceDdosMultiPuTrafficStatsOperOperPuListPktTop(in["pkt_top"].([]interface{}))
		oi.CpuTop = getSliceDdosMultiPuTrafficStatsOperOperPuListCpuTop(in["cpu_top"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosMultiPuTrafficStatsOperOperPuListKbitTop(d []interface{}) []edpt.DdosMultiPuTrafficStatsOperOperPuListKbitTop {

	count1 := len(d)
	ret := make([]edpt.DdosMultiPuTrafficStatsOperOperPuListKbitTop, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosMultiPuTrafficStatsOperOperPuListKbitTop
		oi.ZoneName = in["zone_name"].(string)
		oi.Pu1Kbit = in["pu1_kbit"].(int)
		oi.Pu2Kbit = in["pu2_kbit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosMultiPuTrafficStatsOperOperPuListPktTop(d []interface{}) []edpt.DdosMultiPuTrafficStatsOperOperPuListPktTop {

	count1 := len(d)
	ret := make([]edpt.DdosMultiPuTrafficStatsOperOperPuListPktTop, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosMultiPuTrafficStatsOperOperPuListPktTop
		oi.ZoneName = in["zone_name"].(string)
		oi.Pu1Pkt = in["pu1_pkt"].(int)
		oi.Pu2Pkt = in["pu2_pkt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosMultiPuTrafficStatsOperOperPuListCpuTop(d []interface{}) []edpt.DdosMultiPuTrafficStatsOperOperPuListCpuTop {

	count1 := len(d)
	ret := make([]edpt.DdosMultiPuTrafficStatsOperOperPuListCpuTop, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosMultiPuTrafficStatsOperOperPuListCpuTop
		oi.ZoneName = in["zone_name"].(string)
		oi.Pu1Cpu = in["pu1_cpu"].(int)
		oi.Pu2Cpu = in["pu2_cpu"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosMultiPuTrafficStatsOper(d *schema.ResourceData) edpt.DdosMultiPuTrafficStatsOper {
	var ret edpt.DdosMultiPuTrafficStatsOper

	ret.Oper = getObjectDdosMultiPuTrafficStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
