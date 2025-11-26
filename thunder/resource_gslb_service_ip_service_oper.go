package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpServiceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_ip_service_oper`: Operational Status for the object service\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceIpServiceOperRead,

		Schema: map[string]*schema.Schema{
			"label": {
				Type: schema.TypeString, Required: true, Description: "Service Label",
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

func resourceGslbServiceIpServiceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpServiceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpServiceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceIpServiceOperOper := setObjectGslbServiceIpServiceOperOper(res)
		d.Set("oper", GslbServiceIpServiceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServiceIpServiceOperOper(ret edpt.DataGslbServiceIpServiceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"service_port":          ret.DtGslbServiceIpServiceOper.Oper.ServicePort,
			"state":                 ret.DtGslbServiceIpServiceOper.Oper.State,
			"disabled":              ret.DtGslbServiceIpServiceOper.Oper.Disabled,
			"gslb_protocol":         ret.DtGslbServiceIpServiceOper.Oper.GslbProtocol,
			"local_protocol":        ret.DtGslbServiceIpServiceOper.Oper.LocalProtocol,
			"tcp":                   ret.DtGslbServiceIpServiceOper.Oper.Tcp,
			"manually_health_check": ret.DtGslbServiceIpServiceOper.Oper.ManuallyHealthCheck,
			"use_gslb_state":        ret.DtGslbServiceIpServiceOper.Oper.Use_gslb_state,
			"dynamic":               ret.DtGslbServiceIpServiceOper.Oper.Dynamic,
		},
	}
}

func getObjectGslbServiceIpServiceOperOper(d []interface{}) edpt.GslbServiceIpServiceOperOper {

	count1 := len(d)
	var ret edpt.GslbServiceIpServiceOperOper
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

func dataToEndpointGslbServiceIpServiceOper(d *schema.ResourceData) edpt.GslbServiceIpServiceOper {
	var ret edpt.GslbServiceIpServiceOper

	ret.Label = d.Get("label").(string)

	ret.Oper = getObjectGslbServiceIpServiceOperOper(d.Get("oper").([]interface{}))

	ret.PortNum = d.Get("port_num").(int)

	ret.PortProto = d.Get("port_proto").(string)

	ret.NodeName = d.Get("node_name").(string)
	return ret
}
