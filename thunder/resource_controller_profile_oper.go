package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProfileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_controller_profile_oper`: Operational Status for the object profile\n\n__PLACEHOLDER__",
		ReadContext: resourceControllerProfileOperRead,

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
						"number_of_orgunit_mapped_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"number_of_orgunit_unmapped_partitions": {
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

func resourceControllerProfileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ControllerProfileOperOper := setObjectControllerProfileOperOper(res)
		d.Set("oper", ControllerProfileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectControllerProfileOperOper(ret edpt.DataControllerProfileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"overall_status":                        ret.DtControllerProfileOper.Oper.OverallStatus,
			"heartbeat_status":                      ret.DtControllerProfileOper.Oper.HeartbeatStatus,
			"heartbeat_error_message":               ret.DtControllerProfileOper.Oper.HeartbeatErrorMessage,
			"service_registry":                      ret.DtControllerProfileOper.Oper.ServiceRegistry,
			"service_registry_error_message":        ret.DtControllerProfileOper.Oper.ServiceRegistryErrorMessage,
			"registration_status":                   ret.DtControllerProfileOper.Oper.RegistrationStatus,
			"registration_status_code":              ret.DtControllerProfileOper.Oper.RegistrationStatusCode,
			"registration_error_message":            ret.DtControllerProfileOper.Oper.RegistrationErrorMessage,
			"deregistration_status":                 ret.DtControllerProfileOper.Oper.DeregistrationStatus,
			"deregistration_status_code":            ret.DtControllerProfileOper.Oper.DeregistrationStatusCode,
			"deregistration_error_message":          ret.DtControllerProfileOper.Oper.DeregistrationErrorMessage,
			"schema_registry_status":                ret.DtControllerProfileOper.Oper.SchemaRegistryStatus,
			"broker_info":                           ret.DtControllerProfileOper.Oper.Broker_info,
			"kafka_broker_state":                    ret.DtControllerProfileOper.Oper.KafkaBrokerState,
			"number_of_orgunit_mapped_partitions":   ret.DtControllerProfileOper.Oper.NumberOfOrgunitMappedPartitions,
			"number_of_orgunit_unmapped_partitions": ret.DtControllerProfileOper.Oper.NumberOfOrgunitUnmappedPartitions,
			"tunnel_status":                         ret.DtControllerProfileOper.Oper.TunnelStatus,
			"tunnel_error_message":                  ret.DtControllerProfileOper.Oper.TunnelErrorMessage,
			"peer_device_info":                      ret.DtControllerProfileOper.Oper.PeerDeviceInfo,
		},
	}
}

func getObjectControllerProfileOperOper(d []interface{}) edpt.ControllerProfileOperOper {

	count1 := len(d)
	var ret edpt.ControllerProfileOperOper
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
		ret.NumberOfOrgunitMappedPartitions = in["number_of_orgunit_mapped_partitions"].(int)
		ret.NumberOfOrgunitUnmappedPartitions = in["number_of_orgunit_unmapped_partitions"].(int)
		ret.TunnelStatus = in["tunnel_status"].(string)
		ret.TunnelErrorMessage = in["tunnel_error_message"].(string)
		ret.PeerDeviceInfo = in["peer_device_info"].(string)
	}
	return ret
}

func dataToEndpointControllerProfileOper(d *schema.ResourceData) edpt.ControllerProfileOper {
	var ret edpt.ControllerProfileOper

	ret.Oper = getObjectControllerProfileOperOper(d.Get("oper").([]interface{}))
	return ret
}
