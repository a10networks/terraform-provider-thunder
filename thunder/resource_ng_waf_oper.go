package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNgWafOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ng_waf_oper`: Operational Status for the object ng-waf\n\n__PLACEHOLDER__",
		ReadContext: resourceNgWafOperRead,

		Schema: map[string]*schema.Schema{
			"cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"number_of_cpus": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cpu_info": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cpu_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sec1": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sec5": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sec10": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sec30": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sec60": {
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
			"custom_page": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"file": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"size": {
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
			"custom_signals": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"signal_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"signal": {
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
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ngwaf_stats_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"count1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"vserver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vport": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cache_vserver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cache_vport": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"clear_all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ngwaf_version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"partition_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"agent_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"access_key_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"secret_access_key": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"cache_entries": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tracked_custom_signal": {
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

func resourceNgWafOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWafOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NgWafOperCpu := setObjectNgWafOperCpu(res)
		d.Set("cpu", NgWafOperCpu)
		NgWafOperCustomPage := setObjectNgWafOperCustomPage(res)
		d.Set("custom_page", NgWafOperCustomPage)
		NgWafOperCustomSignals := setObjectNgWafOperCustomSignals(res)
		d.Set("custom_signals", NgWafOperCustomSignals)
		NgWafOperOper := setObjectNgWafOperOper(res)
		d.Set("oper", NgWafOperOper)
		NgWafOperStatus := setObjectNgWafOperStatus(res)
		d.Set("status", NgWafOperStatus)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNgWafOperCpu(ret edpt.DataNgWafOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectNgWafOperCpuOper(ret.DtNgWafOper.Cpu.Oper),
		},
	}
}

func setObjectNgWafOperCpuOper(d edpt.NgWafOperCpuOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["number_of_cpus"] = d.NumberOfCpus
	in["cpu_info"] = setSliceNgWafOperCpuOperCpuInfo(d.CpuInfo)
	result = append(result, in)
	return result
}

func setSliceNgWafOperCpuOperCpuInfo(d []edpt.NgWafOperCpuOperCpuInfo) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_id"] = item.CpuId
		in["sec1"] = item.Sec1
		in["sec5"] = item.Sec5
		in["sec10"] = item.Sec10
		in["sec30"] = item.Sec30
		in["sec60"] = item.Sec60
		result = append(result, in)
	}
	return result
}

func setObjectNgWafOperCustomPage(ret edpt.DataNgWafOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectNgWafOperCustomPageOper(ret.DtNgWafOper.CustomPage.Oper),
		},
	}
}

func setObjectNgWafOperCustomPageOper(d edpt.NgWafOperCustomPageOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["file_list"] = setSliceNgWafOperCustomPageOperFileList(d.FileList)
	result = append(result, in)
	return result
}

func setSliceNgWafOperCustomPageOperFileList(d []edpt.NgWafOperCustomPageOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["size"] = item.Size
		result = append(result, in)
	}
	return result
}

func setObjectNgWafOperCustomSignals(ret edpt.DataNgWafOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectNgWafOperCustomSignalsOper(ret.DtNgWafOper.CustomSignals.Oper),
		},
	}
}

func setObjectNgWafOperCustomSignalsOper(d edpt.NgWafOperCustomSignalsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["signal_list"] = setSliceNgWafOperCustomSignalsOperSignalList(d.SignalList)
	result = append(result, in)
	return result
}

func setSliceNgWafOperCustomSignalsOperSignalList(d []edpt.NgWafOperCustomSignalsOperSignalList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["signal"] = item.Signal
		result = append(result, in)
	}
	return result
}

func setObjectNgWafOperOper(ret edpt.DataNgWafOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ngwaf_stats_list": setSliceNgWafOperOperNgwafStatsList(ret.DtNgWafOper.Oper.NgwafStatsList),
			"vserver":          ret.DtNgWafOper.Oper.Vserver,
			"vport":            ret.DtNgWafOper.Oper.Vport,
			"cache_vserver":    ret.DtNgWafOper.Oper.CacheVserver,
			"cache_vport":      ret.DtNgWafOper.Oper.CacheVport,
			"clear_all":        ret.DtNgWafOper.Oper.ClearAll,
		},
	}
}

func setSliceNgWafOperOperNgwafStatsList(d []edpt.NgWafOperOperNgwafStatsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["name"] = item.Name
		in["count1"] = item.Count1
		result = append(result, in)
	}
	return result
}

func setObjectNgWafOperStatus(ret edpt.DataNgWafOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectNgWafOperStatusOper(ret.DtNgWafOper.Status.Oper),
		},
	}
}

func setObjectNgWafOperStatusOper(d edpt.NgWafOperStatusOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ngwaf_version"] = d.NgwafVersion
	in["partition_list"] = setSliceNgWafOperStatusOperPartitionList(d.PartitionList)
	result = append(result, in)
	return result
}

func setSliceNgWafOperStatusOperPartitionList(d []edpt.NgWafOperStatusOperPartitionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["partition_name"] = item.PartitionName
		in["status"] = item.Status
		in["agent_name"] = item.AgentName
		in["access_key_id"] = item.AccessKeyId
		in["secret_access_key"] = item.SecretAccessKey
		in["cache_entries"] = item.CacheEntries
		in["tracked_custom_signal"] = item.TrackedCustomSignal
		result = append(result, in)
	}
	return result
}

func getObjectNgWafOperCpu(d []interface{}) edpt.NgWafOperCpu {

	count1 := len(d)
	var ret edpt.NgWafOperCpu
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectNgWafOperCpuOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectNgWafOperCpuOper(d []interface{}) edpt.NgWafOperCpuOper {

	count1 := len(d)
	var ret edpt.NgWafOperCpuOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumberOfCpus = in["number_of_cpus"].(int)
		ret.CpuInfo = getSliceNgWafOperCpuOperCpuInfo(in["cpu_info"].([]interface{}))
	}
	return ret
}

func getSliceNgWafOperCpuOperCpuInfo(d []interface{}) []edpt.NgWafOperCpuOperCpuInfo {

	count1 := len(d)
	ret := make([]edpt.NgWafOperCpuOperCpuInfo, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafOperCpuOperCpuInfo
		oi.CpuId = in["cpu_id"].(int)
		oi.Sec1 = in["sec1"].(int)
		oi.Sec5 = in["sec5"].(int)
		oi.Sec10 = in["sec10"].(int)
		oi.Sec30 = in["sec30"].(int)
		oi.Sec60 = in["sec60"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNgWafOperCustomPage(d []interface{}) edpt.NgWafOperCustomPage {

	count1 := len(d)
	var ret edpt.NgWafOperCustomPage
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectNgWafOperCustomPageOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectNgWafOperCustomPageOper(d []interface{}) edpt.NgWafOperCustomPageOper {

	count1 := len(d)
	var ret edpt.NgWafOperCustomPageOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceNgWafOperCustomPageOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceNgWafOperCustomPageOperFileList(d []interface{}) []edpt.NgWafOperCustomPageOperFileList {

	count1 := len(d)
	ret := make([]edpt.NgWafOperCustomPageOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafOperCustomPageOperFileList
		oi.File = in["file"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNgWafOperCustomSignals(d []interface{}) edpt.NgWafOperCustomSignals {

	count1 := len(d)
	var ret edpt.NgWafOperCustomSignals
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectNgWafOperCustomSignalsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectNgWafOperCustomSignalsOper(d []interface{}) edpt.NgWafOperCustomSignalsOper {

	count1 := len(d)
	var ret edpt.NgWafOperCustomSignalsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignalList = getSliceNgWafOperCustomSignalsOperSignalList(in["signal_list"].([]interface{}))
	}
	return ret
}

func getSliceNgWafOperCustomSignalsOperSignalList(d []interface{}) []edpt.NgWafOperCustomSignalsOperSignalList {

	count1 := len(d)
	ret := make([]edpt.NgWafOperCustomSignalsOperSignalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafOperCustomSignalsOperSignalList
		oi.Signal = in["signal"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNgWafOperOper(d []interface{}) edpt.NgWafOperOper {

	count1 := len(d)
	var ret edpt.NgWafOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NgwafStatsList = getSliceNgWafOperOperNgwafStatsList(in["ngwaf_stats_list"].([]interface{}))
		ret.Vserver = in["vserver"].(string)
		ret.Vport = in["vport"].(string)
		ret.CacheVserver = in["cache_vserver"].(string)
		ret.CacheVport = in["cache_vport"].(string)
		ret.ClearAll = in["clear_all"].(int)
	}
	return ret
}

func getSliceNgWafOperOperNgwafStatsList(d []interface{}) []edpt.NgWafOperOperNgwafStatsList {

	count1 := len(d)
	ret := make([]edpt.NgWafOperOperNgwafStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafOperOperNgwafStatsList
		oi.Type = in["type"].(string)
		oi.Name = in["name"].(string)
		oi.Count1 = in["count1"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNgWafOperStatus(d []interface{}) edpt.NgWafOperStatus {

	count1 := len(d)
	var ret edpt.NgWafOperStatus
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectNgWafOperStatusOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectNgWafOperStatusOper(d []interface{}) edpt.NgWafOperStatusOper {

	count1 := len(d)
	var ret edpt.NgWafOperStatusOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NgwafVersion = in["ngwaf_version"].(string)
		ret.PartitionList = getSliceNgWafOperStatusOperPartitionList(in["partition_list"].([]interface{}))
	}
	return ret
}

func getSliceNgWafOperStatusOperPartitionList(d []interface{}) []edpt.NgWafOperStatusOperPartitionList {

	count1 := len(d)
	ret := make([]edpt.NgWafOperStatusOperPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafOperStatusOperPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.Status = in["status"].(string)
		oi.AgentName = in["agent_name"].(string)
		oi.AccessKeyId = in["access_key_id"].(string)
		oi.SecretAccessKey = in["secret_access_key"].(string)
		oi.CacheEntries = in["cache_entries"].(int)
		oi.TrackedCustomSignal = in["tracked_custom_signal"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNgWafOper(d *schema.ResourceData) edpt.NgWafOper {
	var ret edpt.NgWafOper

	ret.Cpu = getObjectNgWafOperCpu(d.Get("cpu").([]interface{}))

	ret.CustomPage = getObjectNgWafOperCustomPage(d.Get("custom_page").([]interface{}))

	ret.CustomSignals = getObjectNgWafOperCustomSignals(d.Get("custom_signals").([]interface{}))

	ret.Oper = getObjectNgWafOperOper(d.Get("oper").([]interface{}))

	ret.Status = getObjectNgWafOperStatus(d.Get("status").([]interface{}))
	return ret
}
