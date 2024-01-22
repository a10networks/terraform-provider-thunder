package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePlatBuffStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_plat_buff_stats_oper`: Operational Status for the object plat-buff-stats\n\n__PLACEHOLDER__",
		ReadContext: resourcePlatBuffStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buff_thr_data": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"thr_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"num_buff_cache": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buff_app_state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buff_app_inq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buff_misc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_buf_a_state": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_inapp_cp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_b_incache_cp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_b_incache": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_b_inthrq": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_b_dcmsg_q": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_b_misc": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_b_dfree": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_inhw": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"g_num_thr_poll": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"g_num_domains": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"threads_domain0": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"thr_name_buf_stat": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"capsules_domain0": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buf_fpga": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"buf_fpga_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"buf_fpga_s": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"buf_stat_fpga": {
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
						"threads_domain1": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"thr_name_buf_stat": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"capsules_domain1": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buf_fpga": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"buf_fpga_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"buf_fpga_s": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"buf_stat_fpga": {
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
						"g_buff_pool": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"g_stat_fpga_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"g_stat_gets0": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"g_stat_puts0": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"g_stat_gets1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"g_stat_puts1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_num_buffers": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"debug_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_ncache": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_ncap": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_npart": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_nddom": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_nbuf": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_buf_total_cache_pfpga": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_buf_total_cache_pcpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dbg_buf_in_cache_per_fpga": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dbg_fpga_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_fpga_ncache": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_fpga_ncap": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_fpga_npart": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_fpga_nddom": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"dbg_buf_in_cache_per_cpu": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dbg_cpu_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_cpu_ncache": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_cpu_ncap": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_cpu_npart": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dbg_cpu_nddom": {
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

func resourcePlatBuffStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePlatBuffStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPlatBuffStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PlatBuffStatsOperOper := setObjectPlatBuffStatsOperOper(res)
		d.Set("oper", PlatBuffStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPlatBuffStatsOperOper(ret edpt.DataPlatBuffStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"buff_thr_data":             setSlicePlatBuffStatsOperOperBuffThrData(ret.DtPlatBuffStatsOper.Oper.BuffThrData),
			"total_buf_a_state":         ret.DtPlatBuffStatsOper.Oper.TotalBufAState,
			"total_inapp_cp":            ret.DtPlatBuffStatsOper.Oper.TotalInappCp,
			"total_b_incache_cp":        ret.DtPlatBuffStatsOper.Oper.TotalBIncacheCp,
			"total_b_incache":           ret.DtPlatBuffStatsOper.Oper.TotalBIncache,
			"total_b_inthrq":            ret.DtPlatBuffStatsOper.Oper.TotalBInthrq,
			"total_b_dcmsg_q":           ret.DtPlatBuffStatsOper.Oper.TotalBDcmsgQ,
			"total_b_misc":              ret.DtPlatBuffStatsOper.Oper.TotalBMisc,
			"total_b_dfree":             ret.DtPlatBuffStatsOper.Oper.TotalBDfree,
			"total_free":                ret.DtPlatBuffStatsOper.Oper.TotalFree,
			"total_inhw":                ret.DtPlatBuffStatsOper.Oper.TotalInhw,
			"g_num_thr_poll":            ret.DtPlatBuffStatsOper.Oper.G_num_thr_poll,
			"g_num_domains":             ret.DtPlatBuffStatsOper.Oper.G_num_domains,
			"threads_domain0":           setSlicePlatBuffStatsOperOperThreadsDomain0(ret.DtPlatBuffStatsOper.Oper.ThreadsDomain0),
			"capsules_domain0":          setSlicePlatBuffStatsOperOperCapsulesDomain0(ret.DtPlatBuffStatsOper.Oper.CapsulesDomain0),
			"threads_domain1":           setSlicePlatBuffStatsOperOperThreadsDomain1(ret.DtPlatBuffStatsOper.Oper.ThreadsDomain1),
			"capsules_domain1":          setSlicePlatBuffStatsOperOperCapsulesDomain1(ret.DtPlatBuffStatsOper.Oper.CapsulesDomain1),
			"g_buff_pool":               setSlicePlatBuffStatsOperOperGBuffPool(ret.DtPlatBuffStatsOper.Oper.GBuffPool),
			"total_num_buffers":         ret.DtPlatBuffStatsOper.Oper.TotalNumBuffers,
			"debug_count":               ret.DtPlatBuffStatsOper.Oper.DebugCount,
			"dbg_ncache":                ret.DtPlatBuffStatsOper.Oper.DbgNcache,
			"dbg_ncap":                  ret.DtPlatBuffStatsOper.Oper.DbgNcap,
			"dbg_npart":                 ret.DtPlatBuffStatsOper.Oper.DbgNpart,
			"dbg_nddom":                 ret.DtPlatBuffStatsOper.Oper.DbgNddom,
			"dbg_nbuf":                  ret.DtPlatBuffStatsOper.Oper.DbgNbuf,
			"dbg_buf_total_cache_pfpga": ret.DtPlatBuffStatsOper.Oper.DbgBufTotalCachePfpga,
			"dbg_buf_total_cache_pcpu":  ret.DtPlatBuffStatsOper.Oper.DbgBufTotalCachePcpu,
			"dbg_buf_in_cache_per_fpga": setSlicePlatBuffStatsOperOperDbgBufInCachePerFpga(ret.DtPlatBuffStatsOper.Oper.DbgBufInCachePerFpga),
			"dbg_buf_in_cache_per_cpu":  setSlicePlatBuffStatsOperOperDbgBufInCachePerCpu(ret.DtPlatBuffStatsOper.Oper.DbgBufInCachePerCpu),
		},
	}
}

func setSlicePlatBuffStatsOperOperBuffThrData(d []edpt.PlatBuffStatsOperOperBuffThrData) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["thr_num"] = item.ThrNum
		in["num_buff_cache"] = item.NumBuffCache
		in["buff_app_state"] = item.BuffAppState
		in["buff_app_inq"] = item.BuffAppInq
		in["buff_misc"] = item.BuffMisc
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperThreadsDomain0(d []edpt.PlatBuffStatsOperOperThreadsDomain0) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["thr_name_buf_stat"] = item.ThrNameBufStat
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperCapsulesDomain0(d []edpt.PlatBuffStatsOperOperCapsulesDomain0) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_num"] = item.DomainNum
		in["buf_fpga"] = setSlicePlatBuffStatsOperOperCapsulesDomain0BufFpga(item.BufFpga)
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperCapsulesDomain0BufFpga(d []edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpga) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["buf_fpga_num"] = item.BufFpgaNum
		in["buf_fpga_s"] = setSlicePlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS(item.BufFpgaS)
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS(d []edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["buf_stat_fpga"] = item.BufStatFpga
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperThreadsDomain1(d []edpt.PlatBuffStatsOperOperThreadsDomain1) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["thr_name_buf_stat"] = item.ThrNameBufStat
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperCapsulesDomain1(d []edpt.PlatBuffStatsOperOperCapsulesDomain1) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_num"] = item.DomainNum
		in["buf_fpga"] = setSlicePlatBuffStatsOperOperCapsulesDomain1BufFpga(item.BufFpga)
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperCapsulesDomain1BufFpga(d []edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpga) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["buf_fpga_num"] = item.BufFpgaNum
		in["buf_fpga_s"] = setSlicePlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS(item.BufFpgaS)
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS(d []edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["buf_stat_fpga"] = item.BufStatFpga
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperGBuffPool(d []edpt.PlatBuffStatsOperOperGBuffPool) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["g_stat_fpga_name"] = item.GStatFpgaName
		in["g_stat_gets0"] = item.GStatGets0
		in["g_stat_puts0"] = item.GStatPuts0
		in["g_stat_gets1"] = item.GStatGets1
		in["g_stat_puts1"] = item.GStatPuts1
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperDbgBufInCachePerFpga(d []edpt.PlatBuffStatsOperOperDbgBufInCachePerFpga) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dbg_fpga_num"] = item.DbgFpgaNum
		in["dbg_fpga_ncache"] = item.DbgFpgaNcache
		in["dbg_fpga_ncap"] = item.DbgFpgaNcap
		in["dbg_fpga_npart"] = item.DbgFpgaNpart
		in["dbg_fpga_nddom"] = item.DbgFpgaNddom
		result = append(result, in)
	}
	return result
}

func setSlicePlatBuffStatsOperOperDbgBufInCachePerCpu(d []edpt.PlatBuffStatsOperOperDbgBufInCachePerCpu) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dbg_cpu_num"] = item.DbgCpuNum
		in["dbg_cpu_ncache"] = item.DbgCpuNcache
		in["dbg_cpu_ncap"] = item.DbgCpuNcap
		in["dbg_cpu_npart"] = item.DbgCpuNpart
		in["dbg_cpu_nddom"] = item.DbgCpuNddom
		result = append(result, in)
	}
	return result
}

func getObjectPlatBuffStatsOperOper(d []interface{}) edpt.PlatBuffStatsOperOper {

	count1 := len(d)
	var ret edpt.PlatBuffStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BuffThrData = getSlicePlatBuffStatsOperOperBuffThrData(in["buff_thr_data"].([]interface{}))
		ret.TotalBufAState = in["total_buf_a_state"].(int)
		ret.TotalInappCp = in["total_inapp_cp"].(int)
		ret.TotalBIncacheCp = in["total_b_incache_cp"].(int)
		ret.TotalBIncache = in["total_b_incache"].(int)
		ret.TotalBInthrq = in["total_b_inthrq"].(int)
		ret.TotalBDcmsgQ = in["total_b_dcmsg_q"].(int)
		ret.TotalBMisc = in["total_b_misc"].(int)
		ret.TotalBDfree = in["total_b_dfree"].(int)
		ret.TotalFree = in["total_free"].(int)
		ret.TotalInhw = in["total_inhw"].(int)
		ret.G_num_thr_poll = in["g_num_thr_poll"].(int)
		ret.G_num_domains = in["g_num_domains"].(int)
		ret.ThreadsDomain0 = getSlicePlatBuffStatsOperOperThreadsDomain0(in["threads_domain0"].([]interface{}))
		ret.CapsulesDomain0 = getSlicePlatBuffStatsOperOperCapsulesDomain0(in["capsules_domain0"].([]interface{}))
		ret.ThreadsDomain1 = getSlicePlatBuffStatsOperOperThreadsDomain1(in["threads_domain1"].([]interface{}))
		ret.CapsulesDomain1 = getSlicePlatBuffStatsOperOperCapsulesDomain1(in["capsules_domain1"].([]interface{}))
		ret.GBuffPool = getSlicePlatBuffStatsOperOperGBuffPool(in["g_buff_pool"].([]interface{}))
		ret.TotalNumBuffers = in["total_num_buffers"].(int)
		ret.DebugCount = in["debug_count"].(int)
		ret.DbgNcache = in["dbg_ncache"].(int)
		ret.DbgNcap = in["dbg_ncap"].(int)
		ret.DbgNpart = in["dbg_npart"].(int)
		ret.DbgNddom = in["dbg_nddom"].(int)
		ret.DbgNbuf = in["dbg_nbuf"].(int)
		ret.DbgBufTotalCachePfpga = in["dbg_buf_total_cache_pfpga"].(int)
		ret.DbgBufTotalCachePcpu = in["dbg_buf_total_cache_pcpu"].(int)
		ret.DbgBufInCachePerFpga = getSlicePlatBuffStatsOperOperDbgBufInCachePerFpga(in["dbg_buf_in_cache_per_fpga"].([]interface{}))
		ret.DbgBufInCachePerCpu = getSlicePlatBuffStatsOperOperDbgBufInCachePerCpu(in["dbg_buf_in_cache_per_cpu"].([]interface{}))
	}
	return ret
}

func getSlicePlatBuffStatsOperOperBuffThrData(d []interface{}) []edpt.PlatBuffStatsOperOperBuffThrData {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperBuffThrData, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperBuffThrData
		oi.ThrNum = in["thr_num"].(int)
		oi.NumBuffCache = in["num_buff_cache"].(int)
		oi.BuffAppState = in["buff_app_state"].(int)
		oi.BuffAppInq = in["buff_app_inq"].(int)
		oi.BuffMisc = in["buff_misc"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperThreadsDomain0(d []interface{}) []edpt.PlatBuffStatsOperOperThreadsDomain0 {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperThreadsDomain0, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperThreadsDomain0
		oi.ThrNameBufStat = in["thr_name_buf_stat"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperCapsulesDomain0(d []interface{}) []edpt.PlatBuffStatsOperOperCapsulesDomain0 {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperCapsulesDomain0, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperCapsulesDomain0
		oi.DomainNum = in["domain_num"].(int)
		oi.BufFpga = getSlicePlatBuffStatsOperOperCapsulesDomain0BufFpga(in["buf_fpga"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperCapsulesDomain0BufFpga(d []interface{}) []edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpga {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpga, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpga
		oi.BufFpgaNum = in["buf_fpga_num"].(int)
		oi.BufFpgaS = getSlicePlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS(in["buf_fpga_s"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS(d []interface{}) []edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS
		oi.BufStatFpga = in["buf_stat_fpga"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperThreadsDomain1(d []interface{}) []edpt.PlatBuffStatsOperOperThreadsDomain1 {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperThreadsDomain1, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperThreadsDomain1
		oi.ThrNameBufStat = in["thr_name_buf_stat"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperCapsulesDomain1(d []interface{}) []edpt.PlatBuffStatsOperOperCapsulesDomain1 {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperCapsulesDomain1, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperCapsulesDomain1
		oi.DomainNum = in["domain_num"].(int)
		oi.BufFpga = getSlicePlatBuffStatsOperOperCapsulesDomain1BufFpga(in["buf_fpga"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperCapsulesDomain1BufFpga(d []interface{}) []edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpga {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpga, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpga
		oi.BufFpgaNum = in["buf_fpga_num"].(int)
		oi.BufFpgaS = getSlicePlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS(in["buf_fpga_s"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS(d []interface{}) []edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS
		oi.BufStatFpga = in["buf_stat_fpga"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperGBuffPool(d []interface{}) []edpt.PlatBuffStatsOperOperGBuffPool {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperGBuffPool, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperGBuffPool
		oi.GStatFpgaName = in["g_stat_fpga_name"].(string)
		oi.GStatGets0 = in["g_stat_gets0"].(string)
		oi.GStatPuts0 = in["g_stat_puts0"].(string)
		oi.GStatGets1 = in["g_stat_gets1"].(string)
		oi.GStatPuts1 = in["g_stat_puts1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperDbgBufInCachePerFpga(d []interface{}) []edpt.PlatBuffStatsOperOperDbgBufInCachePerFpga {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperDbgBufInCachePerFpga, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperDbgBufInCachePerFpga
		oi.DbgFpgaNum = in["dbg_fpga_num"].(int)
		oi.DbgFpgaNcache = in["dbg_fpga_ncache"].(int)
		oi.DbgFpgaNcap = in["dbg_fpga_ncap"].(int)
		oi.DbgFpgaNpart = in["dbg_fpga_npart"].(int)
		oi.DbgFpgaNddom = in["dbg_fpga_nddom"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatBuffStatsOperOperDbgBufInCachePerCpu(d []interface{}) []edpt.PlatBuffStatsOperOperDbgBufInCachePerCpu {

	count1 := len(d)
	ret := make([]edpt.PlatBuffStatsOperOperDbgBufInCachePerCpu, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatBuffStatsOperOperDbgBufInCachePerCpu
		oi.DbgCpuNum = in["dbg_cpu_num"].(int)
		oi.DbgCpuNcache = in["dbg_cpu_ncache"].(int)
		oi.DbgCpuNcap = in["dbg_cpu_ncap"].(int)
		oi.DbgCpuNpart = in["dbg_cpu_npart"].(int)
		oi.DbgCpuNddom = in["dbg_cpu_nddom"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPlatBuffStatsOper(d *schema.ResourceData) edpt.PlatBuffStatsOper {
	var ret edpt.PlatBuffStatsOper

	ret.Oper = getObjectPlatBuffStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
