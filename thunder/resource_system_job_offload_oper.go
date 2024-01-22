package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemJobOffloadOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_job_offload_oper`: Operational Status for the object job-offload\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemJobOffloadOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_offload_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"jobs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"submit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"receive": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"execute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snt_home": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_home": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"complete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_submit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"q_no_space": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_execute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_complete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"offload_cpus": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemJobOffloadOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJobOffloadOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJobOffloadOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemJobOffloadOperOper := setObjectSystemJobOffloadOperOper(res)
		d.Set("oper", SystemJobOffloadOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemJobOffloadOperOper(ret edpt.DataSystemJobOffloadOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"job_offload_cpu_list": setSliceSystemJobOffloadOperOperJobOffloadCpuList(ret.DtSystemJobOffloadOper.Oper.JobOffloadCpuList),
			"cpu_count":            ret.DtSystemJobOffloadOper.Oper.CpuCount,
			"offload_cpus":         ret.DtSystemJobOffloadOper.Oper.OffloadCpus,
		},
	}
}

func setSliceSystemJobOffloadOperOperJobOffloadCpuList(d []edpt.SystemJobOffloadOperOperJobOffloadCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["jobs"] = item.Jobs
		in["submit"] = item.Submit
		in["receive"] = item.Receive
		in["execute"] = item.Execute
		in["snt_home"] = item.Snt_home
		in["rcv_home"] = item.Rcv_home
		in["complete"] = item.Complete
		in["fail_submit"] = item.Fail_submit
		in["q_no_space"] = item.Q_no_space
		in["fail_execute"] = item.Fail_execute
		in["fail_complete"] = item.Fail_complete
		result = append(result, in)
	}
	return result
}

func getObjectSystemJobOffloadOperOper(d []interface{}) edpt.SystemJobOffloadOperOper {

	count1 := len(d)
	var ret edpt.SystemJobOffloadOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.JobOffloadCpuList = getSliceSystemJobOffloadOperOperJobOffloadCpuList(in["job_offload_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
		ret.OffloadCpus = in["offload_cpus"].(int)
	}
	return ret
}

func getSliceSystemJobOffloadOperOperJobOffloadCpuList(d []interface{}) []edpt.SystemJobOffloadOperOperJobOffloadCpuList {

	count1 := len(d)
	ret := make([]edpt.SystemJobOffloadOperOperJobOffloadCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemJobOffloadOperOperJobOffloadCpuList
		oi.Jobs = in["jobs"].(int)
		oi.Submit = in["submit"].(int)
		oi.Receive = in["receive"].(int)
		oi.Execute = in["execute"].(int)
		oi.Snt_home = in["snt_home"].(int)
		oi.Rcv_home = in["rcv_home"].(int)
		oi.Complete = in["complete"].(int)
		oi.Fail_submit = in["fail_submit"].(int)
		oi.Q_no_space = in["q_no_space"].(int)
		oi.Fail_execute = in["fail_execute"].(int)
		oi.Fail_complete = in["fail_complete"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemJobOffloadOper(d *schema.ResourceData) edpt.SystemJobOffloadOper {
	var ret edpt.SystemJobOffloadOper

	ret.Oper = getObjectSystemJobOffloadOperOper(d.Get("oper").([]interface{}))
	return ret
}
