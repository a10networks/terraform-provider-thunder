package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetMgmtSnmpStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_net_mgmt_snmp_stats_oper`: Operational Status for the object stats\n\n__PLACEHOLDER__",
		ReadContext: resourceNetMgmtSnmpStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bad_version": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"unknown_community": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"illegal_operation": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"encoding_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"unknown_security_models": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"invalid_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"input_packets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"number_of_req_var": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"get_req_pdu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"get_next_pdu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"packet_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"too_big_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"no_such_name_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"bad_values_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"general_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"output_packets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"get_resp_pdu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"output_traps": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceNetMgmtSnmpStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetMgmtSnmpStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetMgmtSnmpStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetMgmtSnmpStatsOperOper := setObjectNetMgmtSnmpStatsOperOper(res)
		d.Set("oper", NetMgmtSnmpStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetMgmtSnmpStatsOperOper(ret edpt.DataNetMgmtSnmpStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"bad_version":             ret.DtNetMgmtSnmpStatsOper.Oper.BadVersion,
			"unknown_community":       ret.DtNetMgmtSnmpStatsOper.Oper.UnknownCommunity,
			"illegal_operation":       ret.DtNetMgmtSnmpStatsOper.Oper.IllegalOperation,
			"encoding_error":          ret.DtNetMgmtSnmpStatsOper.Oper.EncodingError,
			"unknown_security_models": ret.DtNetMgmtSnmpStatsOper.Oper.UnknownSecurityModels,
			"invalid_id":              ret.DtNetMgmtSnmpStatsOper.Oper.InvalidId,
			"input_packets":           ret.DtNetMgmtSnmpStatsOper.Oper.InputPackets,
			"number_of_req_var":       ret.DtNetMgmtSnmpStatsOper.Oper.NumberOfReqVar,
			"get_req_pdu":             ret.DtNetMgmtSnmpStatsOper.Oper.GetReqPdu,
			"get_next_pdu":            ret.DtNetMgmtSnmpStatsOper.Oper.GetNextPdu,
			"packet_drop":             ret.DtNetMgmtSnmpStatsOper.Oper.PacketDrop,
			"too_big_errors":          ret.DtNetMgmtSnmpStatsOper.Oper.TooBigErrors,
			"no_such_name_errors":     ret.DtNetMgmtSnmpStatsOper.Oper.NoSuchNameErrors,
			"bad_values_errors":       ret.DtNetMgmtSnmpStatsOper.Oper.BadValuesErrors,
			"general_errors":          ret.DtNetMgmtSnmpStatsOper.Oper.GeneralErrors,
			"output_packets":          ret.DtNetMgmtSnmpStatsOper.Oper.OutputPackets,
			"get_resp_pdu":            ret.DtNetMgmtSnmpStatsOper.Oper.GetRespPdu,
			"output_traps":            ret.DtNetMgmtSnmpStatsOper.Oper.OutputTraps,
		},
	}
}

func getObjectNetMgmtSnmpStatsOperOper(d []interface{}) edpt.NetMgmtSnmpStatsOperOper {

	count1 := len(d)
	var ret edpt.NetMgmtSnmpStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BadVersion = in["bad_version"].(int)
		ret.UnknownCommunity = in["unknown_community"].(int)
		ret.IllegalOperation = in["illegal_operation"].(int)
		ret.EncodingError = in["encoding_error"].(int)
		ret.UnknownSecurityModels = in["unknown_security_models"].(int)
		ret.InvalidId = in["invalid_id"].(int)
		ret.InputPackets = in["input_packets"].(int)
		ret.NumberOfReqVar = in["number_of_req_var"].(int)
		ret.GetReqPdu = in["get_req_pdu"].(int)
		ret.GetNextPdu = in["get_next_pdu"].(int)
		ret.PacketDrop = in["packet_drop"].(int)
		ret.TooBigErrors = in["too_big_errors"].(int)
		ret.NoSuchNameErrors = in["no_such_name_errors"].(int)
		ret.BadValuesErrors = in["bad_values_errors"].(int)
		ret.GeneralErrors = in["general_errors"].(int)
		ret.OutputPackets = in["output_packets"].(int)
		ret.GetRespPdu = in["get_resp_pdu"].(int)
		ret.OutputTraps = in["output_traps"].(int)
	}
	return ret
}

func dataToEndpointNetMgmtSnmpStatsOper(d *schema.ResourceData) edpt.NetMgmtSnmpStatsOper {
	var ret edpt.NetMgmtSnmpStatsOper

	ret.Oper = getObjectNetMgmtSnmpStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
