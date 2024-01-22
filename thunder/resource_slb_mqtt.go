package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMqtt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_mqtt`: Show MQTT Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbMqttCreate,
		UpdateContext: resourceSlbMqttUpdate,
		ReadContext:   resourceSlbMqttRead,
		DeleteContext: resourceSlbMqttDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'recv_mqtt_connect': MQTT Connect; 'recv_mqtt_connack': MQTT Connack; 'recv_mqtt_publish': MQTT Publish; 'recv_mqtt_puback': MQTT Puback; 'recv_mqtt_pubrec': MQTT Pubrec; 'recv_mqtt_pubrel': MQTT Pubrel; 'recv_mqtt_pubcomp': MQTT Pubcomp; 'recv_mqtt_subscribe': MQTT Subscribe; 'recv_mqtt_suback': MQTT Suback; 'recv_mqtt_unsubscribe': MQTT Unsubscribe; 'recv_mqtt_unsuback': MQTT Unsuback; 'recv_mqtt_pingreq': MQTT Pingreq; 'recv_mqtt_pingresp': MQTT Pingresp; 'recv_mqtt_disconnect': MQTT Disconnect; 'recv_mqtt_auth': MQTT Auth; 'recv_mqtt_other': MQTT Unknown; 'curr_proxy': Current proxy conns; 'total_proxy': Total proxy conns; 'request': Total MQTT Commands; 'parse_connect_fail': Parse connect failure; 'parse_publish_fail': Parse publish failure; 'parse_subscribe_fail': Parse subscribe failure; 'parse_unsubscribe_fail': Parse unsubscribe failure; 'tuple_not_linked': tuple-not-linked failure; 'tuple_already_linked': tuple-already-linked failure; 'conn_null': Null conn; 'client_id_null': Null client id; 'session_exist': Session already exist; 'insertion_failed': Insertion failure; 'insertion_successful': Insertion successful;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbMqttCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMqttCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMqtt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMqttRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbMqttUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMqttUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMqtt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMqttRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbMqttDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMqttDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMqtt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbMqttRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMqttRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMqtt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbMqttSamplingEnable(d []interface{}) []edpt.SlbMqttSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbMqttSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbMqttSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbMqtt(d *schema.ResourceData) edpt.SlbMqtt {
	var ret edpt.SlbMqtt
	ret.Inst.SamplingEnable = getSliceSlbMqttSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
