package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpPortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_ip_port_oper`: Operational Status for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceIpPortOperRead,

		Schema: map[string]*schema.Schema{
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
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"port_proto": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"node_name": {
				Type: schema.TypeString, Required: true, Description: "NodeName",
			},
		},
	}
}

func resourceGslbServiceIpPortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpPortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpPortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceIpPortOperOper := setObjectGslbServiceIpPortOperOper(res)
		d.Set("oper", GslbServiceIpPortOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServiceIpPortOperOper(ret edpt.DataGslbServiceIpPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"service_port":          ret.DtGslbServiceIpPortOper.Oper.ServicePort,
			"state":                 ret.DtGslbServiceIpPortOper.Oper.State,
			"disabled":              ret.DtGslbServiceIpPortOper.Oper.Disabled,
			"gslb_protocol":         ret.DtGslbServiceIpPortOper.Oper.GslbProtocol,
			"local_protocol":        ret.DtGslbServiceIpPortOper.Oper.LocalProtocol,
			"tcp":                   ret.DtGslbServiceIpPortOper.Oper.Tcp,
			"manually_health_check": ret.DtGslbServiceIpPortOper.Oper.ManuallyHealthCheck,
			"use_gslb_state":        ret.DtGslbServiceIpPortOper.Oper.Use_gslb_state,
			"dynamic":               ret.DtGslbServiceIpPortOper.Oper.Dynamic,
		},
	}
}

func getObjectGslbServiceIpPortOperOper(d []interface{}) edpt.GslbServiceIpPortOperOper {

	count1 := len(d)
	var ret edpt.GslbServiceIpPortOperOper
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

func dataToEndpointGslbServiceIpPortOper(d *schema.ResourceData) edpt.GslbServiceIpPortOper {
	var ret edpt.GslbServiceIpPortOper

	ret.Oper = getObjectGslbServiceIpPortOperOper(d.Get("oper").([]interface{}))

	ret.PortNum = d.Get("port_num").(int)

	ret.PortProto = d.Get("port_proto").(string)

	ret.NodeName = d.Get("node_name").(string)
	return ret
}
