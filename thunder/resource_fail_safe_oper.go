package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFailSafeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fail_safe_oper`: Operational Status for the object fail-safe\n\n__PLACEHOLDER__",
		ReadContext: resourceFailSafeOperRead,

		Schema: map[string]*schema.Schema{
			"config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sw_error_mon": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hw_error_mon": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sw_recovery_timeout": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hw_recovery_timeout": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dataplane_recovery_timeout": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fpga_mon_enable": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fpga_mon_forced_reboot": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fpga_mon_interval": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fpga_mon_threshold": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mem_mon": {
										Type: schema.TypeString, Optional: true, Description: "",
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
						"free_session_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_session_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sess_mem_recovery_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_fpga_buffers": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"avail_fpga_buff_domain1": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"avail_fpga_buff_domain2": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_free_fpga_buff": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"free_fpga_buffers": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fpga_buff_recovery_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_system_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fpga_stats_num_cntrs": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fpga_stats_iochan": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fpga_stats_iochan_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fpga_stats_iochan_tx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fpga_stats_iochan_rx": {
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

func resourceFailSafeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FailSafeOperConfig := setObjectFailSafeOperConfig(res)
		d.Set("config", FailSafeOperConfig)
		FailSafeOperOper := setObjectFailSafeOperOper(res)
		d.Set("oper", FailSafeOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFailSafeOperConfig(ret edpt.DataFailSafeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectFailSafeOperConfigOper(ret.DtFailSafeOper.Config.Oper),
		},
	}
}

func setObjectFailSafeOperConfigOper(d edpt.FailSafeOperConfigOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["sw_error_mon"] = d.Sw_error_mon

	in["hw_error_mon"] = d.Hw_error_mon

	in["sw_recovery_timeout"] = d.Sw_recovery_timeout

	in["hw_recovery_timeout"] = d.Hw_recovery_timeout

	in["dataplane_recovery_timeout"] = d.Dataplane_recovery_timeout

	in["fpga_mon_enable"] = d.Fpga_mon_enable

	in["fpga_mon_forced_reboot"] = d.Fpga_mon_forced_reboot

	in["fpga_mon_interval"] = d.Fpga_mon_interval

	in["fpga_mon_threshold"] = d.Fpga_mon_threshold

	in["mem_mon"] = d.Mem_mon
	result = append(result, in)
	return result
}

func setObjectFailSafeOperOper(ret edpt.DataFailSafeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"free_session_memory":          ret.DtFailSafeOper.Oper.Free_session_memory,
			"total_session_memory":         ret.DtFailSafeOper.Oper.Total_session_memory,
			"sess_mem_recovery_threshold":  ret.DtFailSafeOper.Oper.Sess_mem_recovery_threshold,
			"total_fpga_buffers":           ret.DtFailSafeOper.Oper.Total_fpga_buffers,
			"avail_fpga_buff_domain1":      ret.DtFailSafeOper.Oper.Avail_fpga_buff_domain1,
			"avail_fpga_buff_domain2":      ret.DtFailSafeOper.Oper.Avail_fpga_buff_domain2,
			"total_free_fpga_buff":         ret.DtFailSafeOper.Oper.Total_free_fpga_buff,
			"free_fpga_buffers":            ret.DtFailSafeOper.Oper.Free_fpga_buffers,
			"fpga_buff_recovery_threshold": ret.DtFailSafeOper.Oper.Fpga_buff_recovery_threshold,
			"total_system_memory":          ret.DtFailSafeOper.Oper.Total_system_memory,
			"fpga_stats_num_cntrs":         ret.DtFailSafeOper.Oper.Fpga_stats_num_cntrs,
			"fpga_stats_iochan":            setSliceFailSafeOperOperFpga_stats_iochan(ret.DtFailSafeOper.Oper.Fpga_stats_iochan),
		},
	}
}

func setSliceFailSafeOperOperFpga_stats_iochan(d []edpt.FailSafeOperOperFpga_stats_iochan) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fpga_stats_iochan_id"] = item.Fpga_stats_iochan_id
		in["fpga_stats_iochan_tx"] = item.Fpga_stats_iochan_tx
		in["fpga_stats_iochan_rx"] = item.Fpga_stats_iochan_rx
		result = append(result, in)
	}
	return result
}

func getObjectFailSafeOperConfig(d []interface{}) edpt.FailSafeOperConfig {

	count1 := len(d)
	var ret edpt.FailSafeOperConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectFailSafeOperConfigOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectFailSafeOperConfigOper(d []interface{}) edpt.FailSafeOperConfigOper {

	count1 := len(d)
	var ret edpt.FailSafeOperConfigOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sw_error_mon = in["sw_error_mon"].(string)
		ret.Hw_error_mon = in["hw_error_mon"].(string)
		ret.Sw_recovery_timeout = in["sw_recovery_timeout"].(string)
		ret.Hw_recovery_timeout = in["hw_recovery_timeout"].(string)
		ret.Dataplane_recovery_timeout = in["dataplane_recovery_timeout"].(string)
		ret.Fpga_mon_enable = in["fpga_mon_enable"].(string)
		ret.Fpga_mon_forced_reboot = in["fpga_mon_forced_reboot"].(string)
		ret.Fpga_mon_interval = in["fpga_mon_interval"].(string)
		ret.Fpga_mon_threshold = in["fpga_mon_threshold"].(string)
		ret.Mem_mon = in["mem_mon"].(string)
	}
	return ret
}

func getObjectFailSafeOperOper(d []interface{}) edpt.FailSafeOperOper {

	count1 := len(d)
	var ret edpt.FailSafeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Free_session_memory = in["free_session_memory"].(int)
		ret.Total_session_memory = in["total_session_memory"].(int)
		ret.Sess_mem_recovery_threshold = in["sess_mem_recovery_threshold"].(int)
		ret.Total_fpga_buffers = in["total_fpga_buffers"].(int)
		ret.Avail_fpga_buff_domain1 = in["avail_fpga_buff_domain1"].(int)
		ret.Avail_fpga_buff_domain2 = in["avail_fpga_buff_domain2"].(int)
		ret.Total_free_fpga_buff = in["total_free_fpga_buff"].(int)
		ret.Free_fpga_buffers = in["free_fpga_buffers"].(int)
		ret.Fpga_buff_recovery_threshold = in["fpga_buff_recovery_threshold"].(int)
		ret.Total_system_memory = in["total_system_memory"].(int)
		ret.Fpga_stats_num_cntrs = in["fpga_stats_num_cntrs"].(int)
		ret.Fpga_stats_iochan = getSliceFailSafeOperOperFpga_stats_iochan(in["fpga_stats_iochan"].([]interface{}))
	}
	return ret
}

func getSliceFailSafeOperOperFpga_stats_iochan(d []interface{}) []edpt.FailSafeOperOperFpga_stats_iochan {

	count1 := len(d)
	ret := make([]edpt.FailSafeOperOperFpga_stats_iochan, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FailSafeOperOperFpga_stats_iochan
		oi.Fpga_stats_iochan_id = in["fpga_stats_iochan_id"].(int)
		oi.Fpga_stats_iochan_tx = in["fpga_stats_iochan_tx"].(int)
		oi.Fpga_stats_iochan_rx = in["fpga_stats_iochan_rx"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFailSafeOper(d *schema.ResourceData) edpt.FailSafeOper {
	var ret edpt.FailSafeOper

	ret.Config = getObjectFailSafeOperConfig(d.Get("config").([]interface{}))

	ret.Oper = getObjectFailSafeOperOper(d.Get("oper").([]interface{}))
	return ret
}
