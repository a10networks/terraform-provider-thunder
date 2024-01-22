package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_server_oper`: Operational Status for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceFwServerOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_group_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_consumed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_consumed_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_freed_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_failed": {
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

func resourceFwServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwServerOperOper := setObjectFwServerOperOper(res)
		d.Set("oper", FwServerOperOper)
		FwServerOperPortList := setSliceFwServerOperPortList(res)
		d.Set("port_list", FwServerOperPortList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwServerOperOper(ret edpt.DataFwServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state": ret.DtFwServerOper.Oper.State,
		},
	}
}

func setSliceFwServerOperPortList(d edpt.DataFwServerOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtFwServerOper.PortList {
		in := make(map[string]interface{})
		in["port_number"] = item.PortNumber
		in["protocol"] = item.Protocol
		in["oper"] = setObjectFwServerOperPortListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectFwServerOperPortListOper(d edpt.FwServerOperPortListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["ip"] = d.Ip

	in["ipv6"] = d.Ipv6

	in["vrid"] = d.Vrid

	in["ha_group_id"] = d.Ha_group_id

	in["ports_consumed"] = d.Ports_consumed

	in["ports_consumed_total"] = d.Ports_consumed_total

	in["ports_freed_total"] = d.Ports_freed_total

	in["alloc_failed"] = d.Alloc_failed
	result = append(result, in)
	return result
}

func getObjectFwServerOperOper(d []interface{}) edpt.FwServerOperOper {

	count1 := len(d)
	var ret edpt.FwServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getSliceFwServerOperPortList(d []interface{}) []edpt.FwServerOperPortList {

	count1 := len(d)
	ret := make([]edpt.FwServerOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwServerOperPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectFwServerOperPortListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwServerOperPortListOper(d []interface{}) edpt.FwServerOperPortListOper {

	count1 := len(d)
	var ret edpt.FwServerOperPortListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Vrid = in["vrid"].(int)
		ret.Ha_group_id = in["ha_group_id"].(int)
		ret.Ports_consumed = in["ports_consumed"].(int)
		ret.Ports_consumed_total = in["ports_consumed_total"].(int)
		ret.Ports_freed_total = in["ports_freed_total"].(int)
		ret.Alloc_failed = in["alloc_failed"].(int)
	}
	return ret
}

func dataToEndpointFwServerOper(d *schema.ResourceData) edpt.FwServerOper {
	var ret edpt.FwServerOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectFwServerOperOper(d.Get("oper").([]interface{}))

	ret.PortList = getSliceFwServerOperPortList(d.Get("port_list").([]interface{}))
	return ret
}
