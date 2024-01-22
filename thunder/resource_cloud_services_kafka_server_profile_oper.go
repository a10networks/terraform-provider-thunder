package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesKafkaServerProfileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cloud_services_kafka_server_profile_oper`: Operational Status for the object kafka-server-profile\n\n__PLACEHOLDER__",
		ReadContext: resourceCloudServicesKafkaServerProfileOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_namespace": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"kafka_broker_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"registration_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCloudServicesKafkaServerProfileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesKafkaServerProfileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesKafkaServerProfileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		CloudServicesKafkaServerProfileOperOper := setObjectCloudServicesKafkaServerProfileOperOper(res)
		d.Set("oper", CloudServicesKafkaServerProfileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCloudServicesKafkaServerProfileOperOper(ret edpt.DataCloudServicesKafkaServerProfileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"active_namespace":    ret.DtCloudServicesKafkaServerProfileOper.Oper.ActiveNamespace,
			"kafka_broker_state":  ret.DtCloudServicesKafkaServerProfileOper.Oper.KafkaBrokerState,
			"registration_status": ret.DtCloudServicesKafkaServerProfileOper.Oper.RegistrationStatus,
		},
	}
}

func getObjectCloudServicesKafkaServerProfileOperOper(d []interface{}) edpt.CloudServicesKafkaServerProfileOperOper {

	count1 := len(d)
	var ret edpt.CloudServicesKafkaServerProfileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActiveNamespace = in["active_namespace"].(string)
		ret.KafkaBrokerState = in["kafka_broker_state"].(string)
		ret.RegistrationStatus = in["registration_status"].(string)
	}
	return ret
}

func dataToEndpointCloudServicesKafkaServerProfileOper(d *schema.ResourceData) edpt.CloudServicesKafkaServerProfileOper {
	var ret edpt.CloudServicesKafkaServerProfileOper

	ret.Oper = getObjectCloudServicesKafkaServerProfileOperOper(d.Get("oper").([]interface{}))
	return ret
}
