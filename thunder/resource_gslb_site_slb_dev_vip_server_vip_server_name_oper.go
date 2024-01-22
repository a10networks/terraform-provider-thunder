package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerNameOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_slb_dev_vip_server_vip_server_name_oper`: Operational Status for the object vip-server-name\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteSlbDevVipServerVipServerNameOperRead,

		Schema: map[string]*schema.Schema{
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
			"vip_name": {
				Type: schema.TypeString, Required: true, Description: "Specify a VIP name for the SLB device",
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

func resourceGslbSiteSlbDevVipServerVipServerNameOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerNameOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerNameOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteSlbDevVipServerVipServerNameOperOper := setObjectGslbSiteSlbDevVipServerVipServerNameOperOper(res)
		d.Set("oper", GslbSiteSlbDevVipServerVipServerNameOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteSlbDevVipServerVipServerNameOperOper(ret edpt.DataGslbSiteSlbDevVipServerVipServerNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dev_vip_addr":          ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Dev_vip_addr,
			"dev_vip_state":         ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Dev_vip_state,
			"node_name":             ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.NodeName,
			"service_ip":            ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.ServiceIp,
			"port_count":            ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.PortCount,
			"virtual_server":        ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.VirtualServer,
			"disabled":              ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Disabled,
			"gslb_protocol":         ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.GslbProtocol,
			"local_protocol":        ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.LocalProtocol,
			"manually_health_check": ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.ManuallyHealthCheck,
			"use_gslb_state":        ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Use_gslb_state,
			"dynamic":               ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Dynamic,
			"hits":                  ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Hits,
			"recent":                ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.Recent,
			"dev_vip_port_list":     setSliceGslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList(ret.DtGslbSiteSlbDevVipServerVipServerNameOper.Oper.DevVipPortList),
		},
	}
}

func setSliceGslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList(d []edpt.GslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func getObjectGslbSiteSlbDevVipServerVipServerNameOperOper(d []interface{}) edpt.GslbSiteSlbDevVipServerVipServerNameOperOper {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevVipServerVipServerNameOperOper
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
		ret.DevVipPortList = getSliceGslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerNameOperOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerNameOper(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerNameOper {
	var ret edpt.GslbSiteSlbDevVipServerVipServerNameOper

	ret.Oper = getObjectGslbSiteSlbDevVipServerVipServerNameOperOper(d.Get("oper").([]interface{}))

	ret.VipName = d.Get("vip_name").(string)

	ret.SiteName = d.Get("site_name").(string)

	ret.DeviceName = d.Get("device_name").(string)
	return ret
}
