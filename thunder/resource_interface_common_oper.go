package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceCommonOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_common_oper`: Operational Status for the object common\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceCommonOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interfaces": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_num": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_part_default_vlan": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"span_tree": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rate_pkt_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rate_byte_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rate_pkt_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rate_byte_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"input_utilization": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"output_utilization": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type_vendor_part": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_lane": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"transceivers_info": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"transceiver_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lane": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"curr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hi_alarm": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hi_warn": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lo_warn": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lo_alarm": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vnp_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tot_num_phy_intf": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceCommonOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceCommonOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceCommonOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceCommonOperOper := setObjectInterfaceCommonOperOper(res)
		d.Set("oper", InterfaceCommonOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceCommonOperOper(ret edpt.DataInterfaceCommonOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interfaces":       setSliceInterfaceCommonOperOperInterfaces(ret.DtInterfaceCommonOper.Oper.Interfaces),
			"time":             ret.DtInterfaceCommonOper.Oper.Time,
			"interval":         ret.DtInterfaceCommonOper.Oper.Interval,
			"vnp_id":           ret.DtInterfaceCommonOper.Oper.Vnp_id,
			"tot_num_phy_intf": ret.DtInterfaceCommonOper.Oper.Tot_num_phy_intf,
		},
	}
}

func setSliceInterfaceCommonOperOperInterfaces(d []edpt.InterfaceCommonOperOperInterfaces) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["port_type"] = item.PortType
		in["is_part_default_vlan"] = item.Is_part_default_vlan
		in["span_tree"] = item.SpanTree
		in["rate_pkt_sent"] = item.RatePktSent
		in["rate_byte_sent"] = item.RateByteSent
		in["rate_pkt_rcvd"] = item.RatePktRcvd
		in["rate_byte_rcvd"] = item.RateByteRcvd
		in["input_utilization"] = item.InputUtilization
		in["output_utilization"] = item.OutputUtilization
		in["type_vendor_part"] = item.Type_vendor_part
		in["total_lane"] = item.Total_lane
		in["transceivers_info"] = setSliceInterfaceCommonOperOperInterfacesTransceivers_info(item.Transceivers_info)
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceCommonOperOperInterfacesTransceivers_info(d []edpt.InterfaceCommonOperOperInterfacesTransceivers_info) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["transceiver_type"] = item.TransceiverType
		in["lane"] = item.Lane
		in["curr"] = item.Curr
		in["hi_alarm"] = item.Hi_alarm
		in["hi_warn"] = item.Hi_warn
		in["lo_warn"] = item.Lo_warn
		in["lo_alarm"] = item.Lo_alarm
		result = append(result, in)
	}
	return result
}

func getObjectInterfaceCommonOperOper(d []interface{}) edpt.InterfaceCommonOperOper {

	count1 := len(d)
	var ret edpt.InterfaceCommonOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interfaces = getSliceInterfaceCommonOperOperInterfaces(in["interfaces"].([]interface{}))
		ret.Time = in["time"].(string)
		ret.Interval = in["interval"].(int)
		ret.Vnp_id = in["vnp_id"].(int)
		ret.Tot_num_phy_intf = in["tot_num_phy_intf"].(int)
	}
	return ret
}

func getSliceInterfaceCommonOperOperInterfaces(d []interface{}) []edpt.InterfaceCommonOperOperInterfaces {

	count1 := len(d)
	ret := make([]edpt.InterfaceCommonOperOperInterfaces, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceCommonOperOperInterfaces
		oi.PortNum = in["port_num"].(string)
		oi.PortType = in["port_type"].(string)
		oi.Is_part_default_vlan = in["is_part_default_vlan"].(string)
		oi.SpanTree = in["span_tree"].(string)
		oi.RatePktSent = in["rate_pkt_sent"].(int)
		oi.RateByteSent = in["rate_byte_sent"].(int)
		oi.RatePktRcvd = in["rate_pkt_rcvd"].(int)
		oi.RateByteRcvd = in["rate_byte_rcvd"].(int)
		oi.InputUtilization = in["input_utilization"].(int)
		oi.OutputUtilization = in["output_utilization"].(int)
		oi.Type_vendor_part = in["type_vendor_part"].(string)
		oi.Total_lane = in["total_lane"].(int)
		oi.Transceivers_info = getSliceInterfaceCommonOperOperInterfacesTransceivers_info(in["transceivers_info"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceCommonOperOperInterfacesTransceivers_info(d []interface{}) []edpt.InterfaceCommonOperOperInterfacesTransceivers_info {

	count1 := len(d)
	ret := make([]edpt.InterfaceCommonOperOperInterfacesTransceivers_info, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceCommonOperOperInterfacesTransceivers_info
		oi.TransceiverType = in["transceiver_type"].(string)
		oi.Lane = in["lane"].(int)
		oi.Curr = in["curr"].(string)
		oi.Hi_alarm = in["hi_alarm"].(string)
		oi.Hi_warn = in["hi_warn"].(string)
		oi.Lo_warn = in["lo_warn"].(string)
		oi.Lo_alarm = in["lo_alarm"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceCommonOper(d *schema.ResourceData) edpt.InterfaceCommonOper {
	var ret edpt.InterfaceCommonOper

	ret.Oper = getObjectInterfaceCommonOperOper(d.Get("oper").([]interface{}))
	return ret
}
