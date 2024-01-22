package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwFullConeSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_full_cone_session_oper`: Operational Status for the object full-cone-session\n\n__PLACEHOLDER__",
		ReadContext: resourceFwFullConeSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"outbound": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inbound": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cpu": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwFullConeSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwFullConeSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwFullConeSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwFullConeSessionOperOper := setObjectFwFullConeSessionOperOper(res)
		d.Set("oper", FwFullConeSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwFullConeSessionOperOper(ret edpt.DataFwFullConeSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list": setSliceFwFullConeSessionOperOperSessionList(ret.DtFwFullConeSessionOper.Oper.SessionList),
			"total":        ret.DtFwFullConeSessionOper.Oper.Total,
			"ipv4_addr":    ret.DtFwFullConeSessionOper.Oper.Ipv4Addr,
			"ipv6_addr":    ret.DtFwFullConeSessionOper.Oper.Ipv6Addr,
		},
	}
}

func setSliceFwFullConeSessionOperOperSessionList(d []edpt.FwFullConeSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["inside_address"] = item.InsideAddress
		in["inside_port"] = item.InsidePort
		in["outbound"] = item.Outbound
		in["inbound"] = item.Inbound
		in["cpu"] = item.Cpu
		in["age"] = item.Age
		result = append(result, in)
	}
	return result
}

func getObjectFwFullConeSessionOperOper(d []interface{}) edpt.FwFullConeSessionOperOper {

	count1 := len(d)
	var ret edpt.FwFullConeSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceFwFullConeSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.Total = in["total"].(int)
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
	}
	return ret
}

func getSliceFwFullConeSessionOperOperSessionList(d []interface{}) []edpt.FwFullConeSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.FwFullConeSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwFullConeSessionOperOperSessionList
		oi.Protocol = in["protocol"].(string)
		oi.InsideAddress = in["inside_address"].(string)
		oi.InsidePort = in["inside_port"].(int)
		oi.Outbound = in["outbound"].(int)
		oi.Inbound = in["inbound"].(int)
		oi.Cpu = in["cpu"].(int)
		oi.Age = in["age"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwFullConeSessionOper(d *schema.ResourceData) edpt.FwFullConeSessionOper {
	var ret edpt.FwFullConeSessionOper

	ret.Oper = getObjectFwFullConeSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
