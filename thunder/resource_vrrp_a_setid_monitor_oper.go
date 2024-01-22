package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpASetidMonitorOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_setid_monitor_oper`: Operational Status for the object setid-monitor\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpASetidMonitorOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"setid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lif": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"device_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ipaddress": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"mac": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"vlan": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"last_observed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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

func resourceVrrpASetidMonitorOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpASetidMonitorOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpASetidMonitorOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpASetidMonitorOperOper := setObjectVrrpASetidMonitorOperOper(res)
		d.Set("oper", VrrpASetidMonitorOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpASetidMonitorOperOper(ret edpt.DataVrrpASetidMonitorOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"set_id":     ret.DtVrrpASetidMonitorOper.Oper.SetId,
			"setid_list": setSliceVrrpASetidMonitorOperOperSetidList(ret.DtVrrpASetidMonitorOper.Oper.SetidList),
		},
	}
}

func setSliceVrrpASetidMonitorOperOperSetidList(d []edpt.VrrpASetidMonitorOperOperSetidList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["interface_type"] = item.Interface_type
		in["lif"] = item.Lif
		in["ip_list"] = setSliceVrrpASetidMonitorOperOperSetidListIpList(item.IpList)
		result = append(result, in)
	}
	return result
}

func setSliceVrrpASetidMonitorOperOperSetidListIpList(d []edpt.VrrpASetidMonitorOperOperSetidListIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.DeviceId
		in["ipaddress"] = item.Ipaddress
		in["ipv6"] = item.Ipv6
		in["mac"] = item.Mac
		in["vlan"] = item.Vlan
		in["last_observed"] = item.Last_observed
		result = append(result, in)
	}
	return result
}

func getObjectVrrpASetidMonitorOperOper(d []interface{}) edpt.VrrpASetidMonitorOperOper {

	count1 := len(d)
	var ret edpt.VrrpASetidMonitorOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SetId = in["set_id"].(int)
		ret.SetidList = getSliceVrrpASetidMonitorOperOperSetidList(in["setid_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpASetidMonitorOperOperSetidList(d []interface{}) []edpt.VrrpASetidMonitorOperOperSetidList {

	count1 := len(d)
	ret := make([]edpt.VrrpASetidMonitorOperOperSetidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpASetidMonitorOperOperSetidList
		oi.Interface_type = in["interface_type"].(string)
		oi.Lif = in["lif"].(int)
		oi.IpList = getSliceVrrpASetidMonitorOperOperSetidListIpList(in["ip_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpASetidMonitorOperOperSetidListIpList(d []interface{}) []edpt.VrrpASetidMonitorOperOperSetidListIpList {

	count1 := len(d)
	ret := make([]edpt.VrrpASetidMonitorOperOperSetidListIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpASetidMonitorOperOperSetidListIpList
		oi.DeviceId = in["device_id"].(int)
		oi.Ipaddress = in["ipaddress"].(string)
		oi.Ipv6 = in["ipv6"].(string)
		oi.Mac = in["mac"].(string)
		oi.Vlan = in["vlan"].(int)
		oi.Last_observed = in["last_observed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpASetidMonitorOper(d *schema.ResourceData) edpt.VrrpASetidMonitorOper {
	var ret edpt.VrrpASetidMonitorOper

	ret.Oper = getObjectVrrpASetidMonitorOperOper(d.Get("oper").([]interface{}))
	return ret
}
