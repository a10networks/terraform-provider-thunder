package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerV4Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_slb_dev_vip_server_vip_server_v4_oper`: Operational Status for the object vip-server-v4\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteSlbDevVipServerVipServerV4OperRead,

		Schema: map[string]*schema.Schema{
			"ipv4": {
				Type: schema.TypeString, Required: true, Description: "Specify IP address",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dev_vip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dev_vip_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"node_name": {
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
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"recent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dev_vip_port_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dev_vip_port_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dev_vip_port_state": {
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
			"device_name": {
				Type: schema.TypeString, Required: true, Description: "DeviceName",
			},
		},
	}
}

func resourceGslbSiteSlbDevVipServerVipServerV4OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV4OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV4Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteSlbDevVipServerVipServerV4OperOper := setObjectGslbSiteSlbDevVipServerVipServerV4OperOper(res)
		d.Set("oper", GslbSiteSlbDevVipServerVipServerV4OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteSlbDevVipServerVipServerV4OperOper(ret edpt.DataGslbSiteSlbDevVipServerVipServerV4Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dev_vip_addr":          ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Dev_vip_addr,
			"dev_vip_state":         ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Dev_vip_state,
			"node_name":             ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.NodeName,
			"service_ip":            ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.ServiceIp,
			"port_count":            ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.PortCount,
			"virtual_server":        ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.VirtualServer,
			"disabled":              ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Disabled,
			"gslb_protocol":         ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.GslbProtocol,
			"local_protocol":        ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.LocalProtocol,
			"manually_health_check": ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.ManuallyHealthCheck,
			"use_gslb_state":        ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Use_gslb_state,
			"dynamic":               ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Dynamic,
			"hits":                  ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Hits,
			"recent":                ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.Recent,
			"dev_vip_port_list":     setSliceGslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList(ret.DtGslbSiteSlbDevVipServerVipServerV4Oper.Oper.DevVipPortList),
		},
	}
}

func setSliceGslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList(d []edpt.GslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func getObjectGslbSiteSlbDevVipServerVipServerV4OperOper(d []interface{}) edpt.GslbSiteSlbDevVipServerVipServerV4OperOper {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevVipServerVipServerV4OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_addr = in["dev_vip_addr"].(string)
		ret.Dev_vip_state = in["dev_vip_state"].(string)
		ret.NodeName = in["node_name"].(string)
		ret.ServiceIp = in["service_ip"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
		ret.DevVipPortList = getSliceGslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerV4Oper(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerV4Oper {
	var ret edpt.GslbSiteSlbDevVipServerVipServerV4Oper

	ret.Ipv4 = d.Get("ipv4").(string)

	ret.Oper = getObjectGslbSiteSlbDevVipServerVipServerV4OperOper(d.Get("oper").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)

	ret.DeviceName = d.Get("device_name").(string)
	return ret
}
