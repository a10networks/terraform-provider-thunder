package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_ip_oper`: Operational Status for the object service-ip\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceIpOperRead,

		Schema: map[string]*schema.Schema{
			"node_name": {
				Type: schema.TypeString, Required: true, Description: "Service-IP Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"port_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_server": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"local_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"manually_health_check": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"use_gslb_state": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dynamic": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"port_proto": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"disabled": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"gslb_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"manually_health_check": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"use_gslb_state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dynamic": {
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

func resourceGslbServiceIpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceIpOperOper := setObjectGslbServiceIpOperOper(res)
		d.Set("oper", GslbServiceIpOperOper)
		GslbServiceIpOperPortList := setSliceGslbServiceIpOperPortList(res)
		d.Set("port_list", GslbServiceIpOperPortList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServiceIpOperOper(ret edpt.DataGslbServiceIpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"service_ip":            ret.DtGslbServiceIpOper.Oper.ServiceIp,
			"ip":                    ret.DtGslbServiceIpOper.Oper.Ip,
			"state":                 ret.DtGslbServiceIpOper.Oper.State,
			"port_count":            ret.DtGslbServiceIpOper.Oper.PortCount,
			"virtual_server":        ret.DtGslbServiceIpOper.Oper.VirtualServer,
			"disabled":              ret.DtGslbServiceIpOper.Oper.Disabled,
			"gslb_protocol":         ret.DtGslbServiceIpOper.Oper.GslbProtocol,
			"local_protocol":        ret.DtGslbServiceIpOper.Oper.LocalProtocol,
			"manually_health_check": ret.DtGslbServiceIpOper.Oper.ManuallyHealthCheck,
			"use_gslb_state":        ret.DtGslbServiceIpOper.Oper.Use_gslb_state,
			"dynamic":               ret.DtGslbServiceIpOper.Oper.Dynamic,
		},
	}
}

func setSliceGslbServiceIpOperPortList(d edpt.DataGslbServiceIpOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbServiceIpOper.PortList {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["port_proto"] = item.PortProto
		in["oper"] = setObjectGslbServiceIpOperPortListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbServiceIpOperPortListOper(d edpt.GslbServiceIpOperPortListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["service_port"] = d.ServicePort

	in["state"] = d.State

	in["disabled"] = d.Disabled

	in["gslb_protocol"] = d.GslbProtocol

	in["local_protocol"] = d.LocalProtocol

	in["tcp"] = d.Tcp

	in["manually_health_check"] = d.ManuallyHealthCheck

	in["use_gslb_state"] = d.Use_gslb_state

	in["dynamic"] = d.Dynamic
	result = append(result, in)
	return result
}

func getObjectGslbServiceIpOperOper(d []interface{}) edpt.GslbServiceIpOperOper {

	count1 := len(d)
	var ret edpt.GslbServiceIpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceIp = in["service_ip"].(string)
		ret.Ip = in["ip"].(string)
		ret.State = in["state"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
	}
	return ret
}

func getSliceGslbServiceIpOperPortList(d []interface{}) []edpt.GslbServiceIpOperPortList {

	count1 := len(d)
	ret := make([]edpt.GslbServiceIpOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceIpOperPortList
		oi.PortNum = in["port_num"].(int)
		oi.PortProto = in["port_proto"].(string)
		oi.Oper = getObjectGslbServiceIpOperPortListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbServiceIpOperPortListOper(d []interface{}) edpt.GslbServiceIpOperPortListOper {

	count1 := len(d)
	var ret edpt.GslbServiceIpOperPortListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServicePort = in["service_port"].(int)
		ret.State = in["state"].(string)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.Tcp = in["tcp"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
	}
	return ret
}

func dataToEndpointGslbServiceIpOper(d *schema.ResourceData) edpt.GslbServiceIpOper {
	var ret edpt.GslbServiceIpOper

	ret.NodeName = d.Get("node_name").(string)

	ret.Oper = getObjectGslbServiceIpOperOper(d.Get("oper").([]interface{}))

	ret.PortList = getSliceGslbServiceIpOperPortList(d.Get("port_list").([]interface{}))
	return ret
}
