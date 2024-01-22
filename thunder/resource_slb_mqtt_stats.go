package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMqttStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_mqtt_stats`: Statistics for the object mqtt\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbMqttStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recv_mqtt_connect": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Connect",
						},
						"recv_mqtt_connack": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Connack",
						},
						"recv_mqtt_publish": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Publish",
						},
						"recv_mqtt_puback": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Puback",
						},
						"recv_mqtt_pubrec": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Pubrec",
						},
						"recv_mqtt_pubrel": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Pubrel",
						},
						"recv_mqtt_pubcomp": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Pubcomp",
						},
						"recv_mqtt_subscribe": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Subscribe",
						},
						"recv_mqtt_suback": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Suback",
						},
						"recv_mqtt_unsubscribe": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Unsubscribe",
						},
						"recv_mqtt_unsuback": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Unsuback",
						},
						"recv_mqtt_pingreq": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Pingreq",
						},
						"recv_mqtt_pingresp": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Pingresp",
						},
						"recv_mqtt_disconnect": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Disconnect",
						},
						"recv_mqtt_auth": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Auth",
						},
						"recv_mqtt_other": {
							Type: schema.TypeInt, Optional: true, Description: "MQTT Unknown",
						},
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Current proxy conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total proxy conns",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Total MQTT Commands",
						},
						"parse_connect_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse connect failure",
						},
						"parse_publish_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse publish failure",
						},
						"parse_subscribe_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse subscribe failure",
						},
						"parse_unsubscribe_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse unsubscribe failure",
						},
						"tuple_not_linked": {
							Type: schema.TypeInt, Optional: true, Description: "tuple-not-linked failure",
						},
						"tuple_already_linked": {
							Type: schema.TypeInt, Optional: true, Description: "tuple-already-linked failure",
						},
						"conn_null": {
							Type: schema.TypeInt, Optional: true, Description: "Null conn",
						},
						"client_id_null": {
							Type: schema.TypeInt, Optional: true, Description: "Null client id",
						},
						"session_exist": {
							Type: schema.TypeInt, Optional: true, Description: "Session already exist",
						},
						"insertion_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Insertion failure",
						},
						"insertion_successful": {
							Type: schema.TypeInt, Optional: true, Description: "Insertion successful",
						},
					},
				},
			},
		},
	}
}

func resourceSlbMqttStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMqttStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMqttStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbMqttStatsStats := setObjectSlbMqttStatsStats(res)
		d.Set("stats", SlbMqttStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbMqttStatsStats(ret edpt.DataSlbMqttStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"recv_mqtt_connect":      ret.DtSlbMqttStats.Stats.Recv_mqtt_connect,
			"recv_mqtt_connack":      ret.DtSlbMqttStats.Stats.Recv_mqtt_connack,
			"recv_mqtt_publish":      ret.DtSlbMqttStats.Stats.Recv_mqtt_publish,
			"recv_mqtt_puback":       ret.DtSlbMqttStats.Stats.Recv_mqtt_puback,
			"recv_mqtt_pubrec":       ret.DtSlbMqttStats.Stats.Recv_mqtt_pubrec,
			"recv_mqtt_pubrel":       ret.DtSlbMqttStats.Stats.Recv_mqtt_pubrel,
			"recv_mqtt_pubcomp":      ret.DtSlbMqttStats.Stats.Recv_mqtt_pubcomp,
			"recv_mqtt_subscribe":    ret.DtSlbMqttStats.Stats.Recv_mqtt_subscribe,
			"recv_mqtt_suback":       ret.DtSlbMqttStats.Stats.Recv_mqtt_suback,
			"recv_mqtt_unsubscribe":  ret.DtSlbMqttStats.Stats.Recv_mqtt_unsubscribe,
			"recv_mqtt_unsuback":     ret.DtSlbMqttStats.Stats.Recv_mqtt_unsuback,
			"recv_mqtt_pingreq":      ret.DtSlbMqttStats.Stats.Recv_mqtt_pingreq,
			"recv_mqtt_pingresp":     ret.DtSlbMqttStats.Stats.Recv_mqtt_pingresp,
			"recv_mqtt_disconnect":   ret.DtSlbMqttStats.Stats.Recv_mqtt_disconnect,
			"recv_mqtt_auth":         ret.DtSlbMqttStats.Stats.Recv_mqtt_auth,
			"recv_mqtt_other":        ret.DtSlbMqttStats.Stats.Recv_mqtt_other,
			"curr_proxy":             ret.DtSlbMqttStats.Stats.Curr_proxy,
			"total_proxy":            ret.DtSlbMqttStats.Stats.Total_proxy,
			"request":                ret.DtSlbMqttStats.Stats.Request,
			"parse_connect_fail":     ret.DtSlbMqttStats.Stats.Parse_connect_fail,
			"parse_publish_fail":     ret.DtSlbMqttStats.Stats.Parse_publish_fail,
			"parse_subscribe_fail":   ret.DtSlbMqttStats.Stats.Parse_subscribe_fail,
			"parse_unsubscribe_fail": ret.DtSlbMqttStats.Stats.Parse_unsubscribe_fail,
			"tuple_not_linked":       ret.DtSlbMqttStats.Stats.Tuple_not_linked,
			"tuple_already_linked":   ret.DtSlbMqttStats.Stats.Tuple_already_linked,
			"conn_null":              ret.DtSlbMqttStats.Stats.Conn_null,
			"client_id_null":         ret.DtSlbMqttStats.Stats.Client_id_null,
			"session_exist":          ret.DtSlbMqttStats.Stats.Session_exist,
			"insertion_failed":       ret.DtSlbMqttStats.Stats.Insertion_failed,
			"insertion_successful":   ret.DtSlbMqttStats.Stats.Insertion_successful,
		},
	}
}

func getObjectSlbMqttStatsStats(d []interface{}) edpt.SlbMqttStatsStats {

	count1 := len(d)
	var ret edpt.SlbMqttStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Recv_mqtt_connect = in["recv_mqtt_connect"].(int)
		ret.Recv_mqtt_connack = in["recv_mqtt_connack"].(int)
		ret.Recv_mqtt_publish = in["recv_mqtt_publish"].(int)
		ret.Recv_mqtt_puback = in["recv_mqtt_puback"].(int)
		ret.Recv_mqtt_pubrec = in["recv_mqtt_pubrec"].(int)
		ret.Recv_mqtt_pubrel = in["recv_mqtt_pubrel"].(int)
		ret.Recv_mqtt_pubcomp = in["recv_mqtt_pubcomp"].(int)
		ret.Recv_mqtt_subscribe = in["recv_mqtt_subscribe"].(int)
		ret.Recv_mqtt_suback = in["recv_mqtt_suback"].(int)
		ret.Recv_mqtt_unsubscribe = in["recv_mqtt_unsubscribe"].(int)
		ret.Recv_mqtt_unsuback = in["recv_mqtt_unsuback"].(int)
		ret.Recv_mqtt_pingreq = in["recv_mqtt_pingreq"].(int)
		ret.Recv_mqtt_pingresp = in["recv_mqtt_pingresp"].(int)
		ret.Recv_mqtt_disconnect = in["recv_mqtt_disconnect"].(int)
		ret.Recv_mqtt_auth = in["recv_mqtt_auth"].(int)
		ret.Recv_mqtt_other = in["recv_mqtt_other"].(int)
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Request = in["request"].(int)
		ret.Parse_connect_fail = in["parse_connect_fail"].(int)
		ret.Parse_publish_fail = in["parse_publish_fail"].(int)
		ret.Parse_subscribe_fail = in["parse_subscribe_fail"].(int)
		ret.Parse_unsubscribe_fail = in["parse_unsubscribe_fail"].(int)
		ret.Tuple_not_linked = in["tuple_not_linked"].(int)
		ret.Tuple_already_linked = in["tuple_already_linked"].(int)
		ret.Conn_null = in["conn_null"].(int)
		ret.Client_id_null = in["client_id_null"].(int)
		ret.Session_exist = in["session_exist"].(int)
		ret.Insertion_failed = in["insertion_failed"].(int)
		ret.Insertion_successful = in["insertion_successful"].(int)
	}
	return ret
}

func dataToEndpointSlbMqttStats(d *schema.ResourceData) edpt.SlbMqttStats {
	var ret edpt.SlbMqttStats

	ret.Stats = getObjectSlbMqttStatsStats(d.Get("stats").([]interface{}))
	return ret
}
