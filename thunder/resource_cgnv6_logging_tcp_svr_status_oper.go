package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LoggingTcpSvrStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_logging_tcp_svr_status_oper`: Operational Status for the object tcp-svr-status\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LoggingTcpSvrStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_svr_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"server_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"server_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"overall_status": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"num_cpus": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceCgnv6LoggingTcpSvrStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingTcpSvrStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingTcpSvrStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LoggingTcpSvrStatusOperOper := setObjectCgnv6LoggingTcpSvrStatusOperOper(res)
		d.Set("oper", Cgnv6LoggingTcpSvrStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LoggingTcpSvrStatusOperOper(ret edpt.DataCgnv6LoggingTcpSvrStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_svr_list": setObjectCgnv6LoggingTcpSvrStatusOperOperTcpSvrList(ret.DtCgnv6LoggingTcpSvrStatusOper.Oper.TcpSvrList),
		},
	}
}

func setObjectCgnv6LoggingTcpSvrStatusOperOperTcpSvrList(d edpt.Cgnv6LoggingTcpSvrStatusOperOperTcpSvrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["template"] = d.Template

	in["server_name"] = d.ServerName

	in["server_port"] = d.ServerPort

	in["overall_status"] = d.OverallStatus

	in["num_cpus"] = d.NumCpus

	in["status"] = d.Status
	result = append(result, in)
	return result
}

func getObjectCgnv6LoggingTcpSvrStatusOperOper(d []interface{}) edpt.Cgnv6LoggingTcpSvrStatusOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingTcpSvrStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpSvrList = getObjectCgnv6LoggingTcpSvrStatusOperOperTcpSvrList(in["tcp_svr_list"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6LoggingTcpSvrStatusOperOperTcpSvrList(d []interface{}) edpt.Cgnv6LoggingTcpSvrStatusOperOperTcpSvrList {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingTcpSvrStatusOperOperTcpSvrList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Template = in["template"].(string)
		ret.ServerName = in["server_name"].(string)
		ret.ServerPort = in["server_port"].(int)
		ret.OverallStatus = in["overall_status"].(int)
		ret.NumCpus = in["num_cpus"].(int)
		ret.Status = in["status"].(string)
	}
	return ret
}

func dataToEndpointCgnv6LoggingTcpSvrStatusOper(d *schema.ResourceData) edpt.Cgnv6LoggingTcpSvrStatusOper {
	var ret edpt.Cgnv6LoggingTcpSvrStatusOper

	ret.Oper = getObjectCgnv6LoggingTcpSvrStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
