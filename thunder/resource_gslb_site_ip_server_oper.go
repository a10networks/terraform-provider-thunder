package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteIpServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_ip_server_oper`: Operational Status for the object ip-server\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteIpServerOperRead,

		Schema: map[string]*schema.Schema{
			"ip_server_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the real server name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_server": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"service_ip": {
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
						"ip_server_port": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vport_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
		},
	}
}

func resourceGslbSiteIpServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteIpServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteIpServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteIpServerOperOper := setObjectGslbSiteIpServerOperOper(res)
		d.Set("oper", GslbSiteIpServerOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteIpServerOperOper(ret edpt.DataGslbSiteIpServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_server":             ret.DtGslbSiteIpServerOper.Oper.IpServer,
			"ip_address":            ret.DtGslbSiteIpServerOper.Oper.IpAddress,
			"state":                 ret.DtGslbSiteIpServerOper.Oper.State,
			"service_ip":            ret.DtGslbSiteIpServerOper.Oper.ServiceIp,
			"port_count":            ret.DtGslbSiteIpServerOper.Oper.PortCount,
			"virtual_server":        ret.DtGslbSiteIpServerOper.Oper.VirtualServer,
			"disabled":              ret.DtGslbSiteIpServerOper.Oper.Disabled,
			"gslb_protocol":         ret.DtGslbSiteIpServerOper.Oper.GslbProtocol,
			"local_protocol":        ret.DtGslbSiteIpServerOper.Oper.LocalProtocol,
			"manually_health_check": ret.DtGslbSiteIpServerOper.Oper.ManuallyHealthCheck,
			"use_gslb_state":        ret.DtGslbSiteIpServerOper.Oper.Use_gslb_state,
			"dynamic":               ret.DtGslbSiteIpServerOper.Oper.Dynamic,
			"ip_server_port":        setSliceGslbSiteIpServerOperOperIpServerPort(ret.DtGslbSiteIpServerOper.Oper.IpServerPort),
		},
	}
}

func setSliceGslbSiteIpServerOperOperIpServerPort(d []edpt.GslbSiteIpServerOperOperIpServerPort) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vport"] = item.Vport
		in["vport_state"] = item.VportState
		result = append(result, in)
	}
	return result
}

func getObjectGslbSiteIpServerOperOper(d []interface{}) edpt.GslbSiteIpServerOperOper {

	count1 := len(d)
	var ret edpt.GslbSiteIpServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpServer = in["ip_server"].(string)
		ret.IpAddress = in["ip_address"].(string)
		ret.State = in["state"].(string)
		ret.ServiceIp = in["service_ip"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
		ret.IpServerPort = getSliceGslbSiteIpServerOperOperIpServerPort(in["ip_server_port"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteIpServerOperOperIpServerPort(d []interface{}) []edpt.GslbSiteIpServerOperOperIpServerPort {

	count1 := len(d)
	ret := make([]edpt.GslbSiteIpServerOperOperIpServerPort, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteIpServerOperOperIpServerPort
		oi.Vport = in["vport"].(int)
		oi.VportState = in["vport_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteIpServerOper(d *schema.ResourceData) edpt.GslbSiteIpServerOper {
	var ret edpt.GslbSiteIpServerOper

	ret.IpServerName = d.Get("ip_server_name").(string)

	ret.Oper = getObjectGslbSiteIpServerOperOper(d.Get("oper").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)
	return ret
}
