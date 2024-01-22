package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosPatternRecognitionCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_pattern_recognition_cpu_oper`: Operational Status for the object cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosPatternRecognitionCpuOperRead,

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
	}
}

func resourceDdosPatternRecognitionCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPatternRecognitionCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPatternRecognitionCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosPatternRecognitionCpuOperOper := setObjectDdosPatternRecognitionCpuOperOper(res)
		d.Set("oper", DdosPatternRecognitionCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosPatternRecognitionCpuOperOper(ret edpt.DataDdosPatternRecognitionCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"number_of_cpus": ret.DtDdosPatternRecognitionCpuOper.Oper.NumberOfCpus,
			"cpu_info":       setSliceDdosPatternRecognitionCpuOperOperCpuInfo(ret.DtDdosPatternRecognitionCpuOper.Oper.CpuInfo),
		},
	}
}

func setSliceDdosPatternRecognitionCpuOperOperCpuInfo(d []edpt.DdosPatternRecognitionCpuOperOperCpuInfo) []map[string]interface{} {
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

func getObjectDdosPatternRecognitionCpuOperOper(d []interface{}) edpt.DdosPatternRecognitionCpuOperOper {

	count1 := len(d)
	var ret edpt.DdosPatternRecognitionCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumberOfCpus = in["number_of_cpus"].(int)
		ret.CpuInfo = getSliceDdosPatternRecognitionCpuOperOperCpuInfo(in["cpu_info"].([]interface{}))
	}
	return ret
}

func getSliceDdosPatternRecognitionCpuOperOperCpuInfo(d []interface{}) []edpt.DdosPatternRecognitionCpuOperOperCpuInfo {

	count1 := len(d)
	ret := make([]edpt.DdosPatternRecognitionCpuOperOperCpuInfo, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosPatternRecognitionCpuOperOperCpuInfo
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

func dataToEndpointDdosPatternRecognitionCpuOper(d *schema.ResourceData) edpt.DdosPatternRecognitionCpuOper {
	var ret edpt.DdosPatternRecognitionCpuOper

	ret.Oper = getObjectDdosPatternRecognitionCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
