package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventNotificationKafkaServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_event_notification_kafka_server_oper`: Operational Status for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceEventNotificationKafkaServerOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kafka_broker_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceEventNotificationKafkaServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventNotificationKafkaServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventNotificationKafkaServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		EventNotificationKafkaServerOperOper := setObjectEventNotificationKafkaServerOperOper(res)
		d.Set("oper", EventNotificationKafkaServerOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectEventNotificationKafkaServerOperOper(ret edpt.DataEventNotificationKafkaServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"kafka_broker_state": ret.DtEventNotificationKafkaServerOper.Oper.KafkaBrokerState,
		},
	}
}

func getObjectEventNotificationKafkaServerOperOper(d []interface{}) edpt.EventNotificationKafkaServerOperOper {

	count1 := len(d)
	var ret edpt.EventNotificationKafkaServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KafkaBrokerState = in["kafka_broker_state"].(string)
	}
	return ret
}

func dataToEndpointEventNotificationKafkaServerOper(d *schema.ResourceData) edpt.EventNotificationKafkaServerOper {
	var ret edpt.EventNotificationKafkaServerOper

	ret.Oper = getObjectEventNotificationKafkaServerOperOper(d.Get("oper").([]interface{}))
	return ret
}
