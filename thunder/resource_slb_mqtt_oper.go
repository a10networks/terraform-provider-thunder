package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMqttOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_mqtt_oper`: Operational Status for the object mqtt\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbMqttOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mqtt_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recv_mqtt_connect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_connack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_publish": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_puback": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_pubrec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_pubrel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_pubcomp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_subscribe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_suback": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_unsubscribe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_unsuback": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_pingreq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_pingresp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_disconnect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_auth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mqtt_other": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"curr_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_connect_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_publish_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_subscribe_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_unsubscribe_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tuple_not_linked": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tuple_already_linked": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_null": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_id_null": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_exist": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insertion_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insertion_successful": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbMqttOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMqttOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMqttOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbMqttOperOper := setObjectSlbMqttOperOper(res)
		d.Set("oper", SlbMqttOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbMqttOperOper(ret edpt.DataSlbMqttOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mqtt_cpu_list": setSliceSlbMqttOperOperMqttCpuList(ret.DtSlbMqttOper.Oper.MqttCpuList),
			"cpu_count":     ret.DtSlbMqttOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbMqttOperOperMqttCpuList(d []edpt.SlbMqttOperOperMqttCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["recv_mqtt_connect"] = item.Recv_mqtt_connect
		in["recv_mqtt_connack"] = item.Recv_mqtt_connack
		in["recv_mqtt_publish"] = item.Recv_mqtt_publish
		in["recv_mqtt_puback"] = item.Recv_mqtt_puback
		in["recv_mqtt_pubrec"] = item.Recv_mqtt_pubrec
		in["recv_mqtt_pubrel"] = item.Recv_mqtt_pubrel
		in["recv_mqtt_pubcomp"] = item.Recv_mqtt_pubcomp
		in["recv_mqtt_subscribe"] = item.Recv_mqtt_subscribe
		in["recv_mqtt_suback"] = item.Recv_mqtt_suback
		in["recv_mqtt_unsubscribe"] = item.Recv_mqtt_unsubscribe
		in["recv_mqtt_unsuback"] = item.Recv_mqtt_unsuback
		in["recv_mqtt_pingreq"] = item.Recv_mqtt_pingreq
		in["recv_mqtt_pingresp"] = item.Recv_mqtt_pingresp
		in["recv_mqtt_disconnect"] = item.Recv_mqtt_disconnect
		in["recv_mqtt_auth"] = item.Recv_mqtt_auth
		in["recv_mqtt_other"] = item.Recv_mqtt_other
		in["curr_proxy"] = item.Curr_proxy
		in["total_proxy"] = item.Total_proxy
		in["request"] = item.Request
		in["parse_connect_fail"] = item.Parse_connect_fail
		in["parse_publish_fail"] = item.Parse_publish_fail
		in["parse_subscribe_fail"] = item.Parse_subscribe_fail
		in["parse_unsubscribe_fail"] = item.Parse_unsubscribe_fail
		in["tuple_not_linked"] = item.Tuple_not_linked
		in["tuple_already_linked"] = item.Tuple_already_linked
		in["conn_null"] = item.Conn_null
		in["client_id_null"] = item.Client_id_null
		in["session_exist"] = item.Session_exist
		in["insertion_failed"] = item.Insertion_failed
		in["insertion_successful"] = item.Insertion_successful
		result = append(result, in)
	}
	return result
}

func getObjectSlbMqttOperOper(d []interface{}) edpt.SlbMqttOperOper {

	count1 := len(d)
	var ret edpt.SlbMqttOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MqttCpuList = getSliceSlbMqttOperOperMqttCpuList(in["mqtt_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbMqttOperOperMqttCpuList(d []interface{}) []edpt.SlbMqttOperOperMqttCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbMqttOperOperMqttCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbMqttOperOperMqttCpuList
		oi.Recv_mqtt_connect = in["recv_mqtt_connect"].(int)
		oi.Recv_mqtt_connack = in["recv_mqtt_connack"].(int)
		oi.Recv_mqtt_publish = in["recv_mqtt_publish"].(int)
		oi.Recv_mqtt_puback = in["recv_mqtt_puback"].(int)
		oi.Recv_mqtt_pubrec = in["recv_mqtt_pubrec"].(int)
		oi.Recv_mqtt_pubrel = in["recv_mqtt_pubrel"].(int)
		oi.Recv_mqtt_pubcomp = in["recv_mqtt_pubcomp"].(int)
		oi.Recv_mqtt_subscribe = in["recv_mqtt_subscribe"].(int)
		oi.Recv_mqtt_suback = in["recv_mqtt_suback"].(int)
		oi.Recv_mqtt_unsubscribe = in["recv_mqtt_unsubscribe"].(int)
		oi.Recv_mqtt_unsuback = in["recv_mqtt_unsuback"].(int)
		oi.Recv_mqtt_pingreq = in["recv_mqtt_pingreq"].(int)
		oi.Recv_mqtt_pingresp = in["recv_mqtt_pingresp"].(int)
		oi.Recv_mqtt_disconnect = in["recv_mqtt_disconnect"].(int)
		oi.Recv_mqtt_auth = in["recv_mqtt_auth"].(int)
		oi.Recv_mqtt_other = in["recv_mqtt_other"].(int)
		oi.Curr_proxy = in["curr_proxy"].(int)
		oi.Total_proxy = in["total_proxy"].(int)
		oi.Request = in["request"].(int)
		oi.Parse_connect_fail = in["parse_connect_fail"].(int)
		oi.Parse_publish_fail = in["parse_publish_fail"].(int)
		oi.Parse_subscribe_fail = in["parse_subscribe_fail"].(int)
		oi.Parse_unsubscribe_fail = in["parse_unsubscribe_fail"].(int)
		oi.Tuple_not_linked = in["tuple_not_linked"].(int)
		oi.Tuple_already_linked = in["tuple_already_linked"].(int)
		oi.Conn_null = in["conn_null"].(int)
		oi.Client_id_null = in["client_id_null"].(int)
		oi.Session_exist = in["session_exist"].(int)
		oi.Insertion_failed = in["insertion_failed"].(int)
		oi.Insertion_successful = in["insertion_successful"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbMqttOper(d *schema.ResourceData) edpt.SlbMqttOper {
	var ret edpt.SlbMqttOper

	ret.Oper = getObjectSlbMqttOperOper(d.Get("oper").([]interface{}))
	return ret
}
