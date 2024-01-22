package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_harmony_controller_profile_oper`: Operational Status for the object profile\n\n__PLACEHOLDER__",
		ReadContext: resourceHarmonyControllerProfileOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"overall_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"heartbeat_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"heartbeat_error_message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"service_registry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"service_registry_error_message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"registration_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"registration_status_code": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"registration_error_message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"deregistration_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"deregistration_status_code": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"deregistration_error_message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"schema_registry_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"broker_info": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"kafka_broker_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"number_of_tenant_mapped_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"number_of_tenant_unmapped_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tunnel_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tunnel_error_message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"peer_device_info": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceHarmonyControllerProfileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		HarmonyControllerProfileOperOper := setObjectHarmonyControllerProfileOperOper(res)
		d.Set("oper", HarmonyControllerProfileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectHarmonyControllerProfileOperOper(ret edpt.DataHarmonyControllerProfileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"overall_status":                       ret.DtHarmonyControllerProfileOper.Oper.OverallStatus,
			"heartbeat_status":                     ret.DtHarmonyControllerProfileOper.Oper.HeartbeatStatus,
			"heartbeat_error_message":              ret.DtHarmonyControllerProfileOper.Oper.HeartbeatErrorMessage,
			"service_registry":                     ret.DtHarmonyControllerProfileOper.Oper.ServiceRegistry,
			"service_registry_error_message":       ret.DtHarmonyControllerProfileOper.Oper.ServiceRegistryErrorMessage,
			"registration_status":                  ret.DtHarmonyControllerProfileOper.Oper.RegistrationStatus,
			"registration_status_code":             ret.DtHarmonyControllerProfileOper.Oper.RegistrationStatusCode,
			"registration_error_message":           ret.DtHarmonyControllerProfileOper.Oper.RegistrationErrorMessage,
			"deregistration_status":                ret.DtHarmonyControllerProfileOper.Oper.DeregistrationStatus,
			"deregistration_status_code":           ret.DtHarmonyControllerProfileOper.Oper.DeregistrationStatusCode,
			"deregistration_error_message":         ret.DtHarmonyControllerProfileOper.Oper.DeregistrationErrorMessage,
			"schema_registry_status":               ret.DtHarmonyControllerProfileOper.Oper.SchemaRegistryStatus,
			"broker_info":                          ret.DtHarmonyControllerProfileOper.Oper.Broker_info,
			"kafka_broker_state":                   ret.DtHarmonyControllerProfileOper.Oper.KafkaBrokerState,
			"number_of_tenant_mapped_partitions":   ret.DtHarmonyControllerProfileOper.Oper.NumberOfTenantMappedPartitions,
			"number_of_tenant_unmapped_partitions": ret.DtHarmonyControllerProfileOper.Oper.NumberOfTenantUnmappedPartitions,
			"tunnel_status":                        ret.DtHarmonyControllerProfileOper.Oper.TunnelStatus,
			"tunnel_error_message":                 ret.DtHarmonyControllerProfileOper.Oper.TunnelErrorMessage,
			"peer_device_info":                     ret.DtHarmonyControllerProfileOper.Oper.PeerDeviceInfo,
		},
	}
}

func getObjectHarmonyControllerProfileOperOper(d []interface{}) edpt.HarmonyControllerProfileOperOper {

	count1 := len(d)
	var ret edpt.HarmonyControllerProfileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OverallStatus = in["overall_status"].(string)
		ret.HeartbeatStatus = in["heartbeat_status"].(string)
		ret.HeartbeatErrorMessage = in["heartbeat_error_message"].(string)
		ret.ServiceRegistry = in["service_registry"].(string)
		ret.ServiceRegistryErrorMessage = in["service_registry_error_message"].(string)
		ret.RegistrationStatus = in["registration_status"].(string)
		ret.RegistrationStatusCode = in["registration_status_code"].(int)
		ret.RegistrationErrorMessage = in["registration_error_message"].(string)
		ret.DeregistrationStatus = in["deregistration_status"].(string)
		ret.DeregistrationStatusCode = in["deregistration_status_code"].(int)
		ret.DeregistrationErrorMessage = in["deregistration_error_message"].(string)
		ret.SchemaRegistryStatus = in["schema_registry_status"].(string)
		ret.Broker_info = in["broker_info"].(string)
		ret.KafkaBrokerState = in["kafka_broker_state"].(string)
		ret.NumberOfTenantMappedPartitions = in["number_of_tenant_mapped_partitions"].(int)
		ret.NumberOfTenantUnmappedPartitions = in["number_of_tenant_unmapped_partitions"].(int)
		ret.TunnelStatus = in["tunnel_status"].(string)
		ret.TunnelErrorMessage = in["tunnel_error_message"].(string)
		ret.PeerDeviceInfo = in["peer_device_info"].(string)
	}
	return ret
}

func dataToEndpointHarmonyControllerProfileOper(d *schema.ResourceData) edpt.HarmonyControllerProfileOper {
	var ret edpt.HarmonyControllerProfileOper

	ret.Oper = getObjectHarmonyControllerProfileOperOper(d.Get("oper").([]interface{}))
	return ret
}
