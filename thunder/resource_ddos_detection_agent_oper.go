package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgentOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_detection_agent_oper`: Operational Status for the object agent\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDetectionAgentOperRead,

		Schema: map[string]*schema.Schema{
			"agent_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name for the agent",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"brand": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sampler_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sampler_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sample_mode": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sampling_algorithm": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sampling_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inactive_timeout": {
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

func resourceDdosDetectionAgentOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDetectionAgentOperOper := setObjectDdosDetectionAgentOperOper(res)
		d.Set("oper", DdosDetectionAgentOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDetectionAgentOperOper(ret edpt.DataDdosDetectionAgentOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"brand":        ret.DtDdosDetectionAgentOper.Oper.Brand,
			"sampler_list": setSliceDdosDetectionAgentOperOperSamplerList(ret.DtDdosDetectionAgentOper.Oper.SamplerList),
		},
	}
}

func setSliceDdosDetectionAgentOperOperSamplerList(d []edpt.DdosDetectionAgentOperOperSamplerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sampler_id"] = item.SamplerId
		in["sample_mode"] = item.SampleMode
		in["sampling_algorithm"] = item.SamplingAlgorithm
		in["sampling_rate"] = item.SamplingRate
		in["active_timeout"] = item.ActiveTimeout
		in["inactive_timeout"] = item.InactiveTimeout
		result = append(result, in)
	}
	return result
}

func getObjectDdosDetectionAgentOperOper(d []interface{}) edpt.DdosDetectionAgentOperOper {

	count1 := len(d)
	var ret edpt.DdosDetectionAgentOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Brand = in["brand"].(string)
		ret.SamplerList = getSliceDdosDetectionAgentOperOperSamplerList(in["sampler_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDetectionAgentOperOperSamplerList(d []interface{}) []edpt.DdosDetectionAgentOperOperSamplerList {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentOperOperSamplerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentOperOperSamplerList
		oi.SamplerId = in["sampler_id"].(int)
		oi.SampleMode = in["sample_mode"].(int)
		oi.SamplingAlgorithm = in["sampling_algorithm"].(int)
		oi.SamplingRate = in["sampling_rate"].(int)
		oi.ActiveTimeout = in["active_timeout"].(int)
		oi.InactiveTimeout = in["inactive_timeout"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDetectionAgentOper(d *schema.ResourceData) edpt.DdosDetectionAgentOper {
	var ret edpt.DdosDetectionAgentOper

	ret.AgentName = d.Get("agent_name").(string)

	ret.Oper = getObjectDdosDetectionAgentOperOper(d.Get("oper").([]interface{}))
	return ret
}
