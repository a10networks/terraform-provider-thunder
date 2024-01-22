package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeGatewayStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ike_gateway_stats`: Statistics for the object ike-gateway\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIkeGatewayStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "IKE-gateway name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v2_init_rekey": {
							Type: schema.TypeInt, Optional: true, Description: "Initiate Rekey",
						},
						"v2_rsp_rekey": {
							Type: schema.TypeInt, Optional: true, Description: "Respond Rekey",
						},
						"v2_child_sa_rekey": {
							Type: schema.TypeInt, Optional: true, Description: "Child SA Rekey",
						},
						"v2_in_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Invalid",
						},
						"v2_in_invalid_spi": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Invalid SPI",
						},
						"v2_in_init_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Init Request",
						},
						"v2_in_init_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Init Response",
						},
						"v2_out_init_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Init Request",
						},
						"v2_out_init_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Init Response",
						},
						"v2_in_auth_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Request",
						},
						"v2_in_auth_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Response",
						},
						"v2_out_auth_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Request",
						},
						"v2_out_auth_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Response",
						},
						"v2_in_create_child_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Create Child Request",
						},
						"v2_in_create_child_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Create Child Response",
						},
						"v2_out_create_child_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Create Child Request",
						},
						"v2_out_create_child_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Create Child Response",
						},
						"v2_in_info_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Info Request",
						},
						"v2_in_info_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Info Response",
						},
						"v2_out_info_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Request",
						},
						"v2_out_info_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Response",
						},
						"v1_in_id_prot_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming ID Protection Request",
						},
						"v1_in_id_prot_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming ID Protection Response",
						},
						"v1_out_id_prot_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing ID Protection Request",
						},
						"v1_out_id_prot_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing ID Protection Response",
						},
						"v1_in_auth_only_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Only Request",
						},
						"v1_in_auth_only_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Only Response",
						},
						"v1_out_auth_only_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Only Request",
						},
						"v1_out_auth_only_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Only Response",
						},
						"v1_in_aggressive_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Aggressive Request",
						},
						"v1_in_aggressive_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Aggressive Response",
						},
						"v1_out_aggressive_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Aggressive Request",
						},
						"v1_out_aggressive_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Aggressive Response",
						},
						"v1_in_info_v1_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Info Request",
						},
						"v1_in_info_v1_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Info Response",
						},
						"v1_out_info_v1_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Request",
						},
						"v1_out_info_v1_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Response",
						},
						"v1_in_transaction_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Transaction Request",
						},
						"v1_in_transaction_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Transaction Response",
						},
						"v1_out_transaction_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Transaction Request",
						},
						"v1_out_transaction_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Transaction Response",
						},
						"v1_in_quick_mode_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Quick Mode Request",
						},
						"v1_in_quick_mode_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming Quick Mode Response",
						},
						"v1_out_quick_mode_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Quick Mode Request",
						},
						"v1_out_quick_mode_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing Quick Mode Response",
						},
						"v1_in_new_group_mode_req": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming New Group Mode Request",
						},
						"v1_in_new_group_mode_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming New Group Mode Response",
						},
						"v1_out_new_group_mode_req": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing New Group Mode Request",
						},
						"v1_out_new_group_mode_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing New Group Mode Response",
						},
						"v1_child_sa_invalid_spi": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid SPI for Child SAs",
						},
						"v2_child_sa_invalid_spi": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid SPI for Child SAs",
						},
						"ike_current_version": {
							Type: schema.TypeInt, Optional: true, Description: "IKE version",
						},
					},
				},
			},
		},
	}
}

func resourceVpnIkeGatewayStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeGatewayStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeGatewayStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIkeGatewayStatsStats := setObjectVpnIkeGatewayStatsStats(res)
		d.Set("stats", VpnIkeGatewayStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIkeGatewayStatsStats(ret edpt.DataVpnIkeGatewayStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"v2_init_rekey":             ret.DtVpnIkeGatewayStats.Stats.V2InitRekey,
			"v2_rsp_rekey":              ret.DtVpnIkeGatewayStats.Stats.V2RspRekey,
			"v2_child_sa_rekey":         ret.DtVpnIkeGatewayStats.Stats.V2ChildSaRekey,
			"v2_in_invalid":             ret.DtVpnIkeGatewayStats.Stats.V2InInvalid,
			"v2_in_invalid_spi":         ret.DtVpnIkeGatewayStats.Stats.V2InInvalidSpi,
			"v2_in_init_req":            ret.DtVpnIkeGatewayStats.Stats.V2InInitReq,
			"v2_in_init_rsp":            ret.DtVpnIkeGatewayStats.Stats.V2InInitRsp,
			"v2_out_init_req":           ret.DtVpnIkeGatewayStats.Stats.V2OutInitReq,
			"v2_out_init_rsp":           ret.DtVpnIkeGatewayStats.Stats.V2OutInitRsp,
			"v2_in_auth_req":            ret.DtVpnIkeGatewayStats.Stats.V2InAuthReq,
			"v2_in_auth_rsp":            ret.DtVpnIkeGatewayStats.Stats.V2InAuthRsp,
			"v2_out_auth_req":           ret.DtVpnIkeGatewayStats.Stats.V2OutAuthReq,
			"v2_out_auth_rsp":           ret.DtVpnIkeGatewayStats.Stats.V2OutAuthRsp,
			"v2_in_create_child_req":    ret.DtVpnIkeGatewayStats.Stats.V2InCreateChildReq,
			"v2_in_create_child_rsp":    ret.DtVpnIkeGatewayStats.Stats.V2InCreateChildRsp,
			"v2_out_create_child_req":   ret.DtVpnIkeGatewayStats.Stats.V2OutCreateChildReq,
			"v2_out_create_child_rsp":   ret.DtVpnIkeGatewayStats.Stats.V2OutCreateChildRsp,
			"v2_in_info_req":            ret.DtVpnIkeGatewayStats.Stats.V2InInfoReq,
			"v2_in_info_rsp":            ret.DtVpnIkeGatewayStats.Stats.V2InInfoRsp,
			"v2_out_info_req":           ret.DtVpnIkeGatewayStats.Stats.V2OutInfoReq,
			"v2_out_info_rsp":           ret.DtVpnIkeGatewayStats.Stats.V2OutInfoRsp,
			"v1_in_id_prot_req":         ret.DtVpnIkeGatewayStats.Stats.V1InIdProtReq,
			"v1_in_id_prot_rsp":         ret.DtVpnIkeGatewayStats.Stats.V1InIdProtRsp,
			"v1_out_id_prot_req":        ret.DtVpnIkeGatewayStats.Stats.V1OutIdProtReq,
			"v1_out_id_prot_rsp":        ret.DtVpnIkeGatewayStats.Stats.V1OutIdProtRsp,
			"v1_in_auth_only_req":       ret.DtVpnIkeGatewayStats.Stats.V1InAuthOnlyReq,
			"v1_in_auth_only_rsp":       ret.DtVpnIkeGatewayStats.Stats.V1InAuthOnlyRsp,
			"v1_out_auth_only_req":      ret.DtVpnIkeGatewayStats.Stats.V1OutAuthOnlyReq,
			"v1_out_auth_only_rsp":      ret.DtVpnIkeGatewayStats.Stats.V1OutAuthOnlyRsp,
			"v1_in_aggressive_req":      ret.DtVpnIkeGatewayStats.Stats.V1InAggressiveReq,
			"v1_in_aggressive_rsp":      ret.DtVpnIkeGatewayStats.Stats.V1InAggressiveRsp,
			"v1_out_aggressive_req":     ret.DtVpnIkeGatewayStats.Stats.V1OutAggressiveReq,
			"v1_out_aggressive_rsp":     ret.DtVpnIkeGatewayStats.Stats.V1OutAggressiveRsp,
			"v1_in_info_v1_req":         ret.DtVpnIkeGatewayStats.Stats.V1InInfoV1Req,
			"v1_in_info_v1_rsp":         ret.DtVpnIkeGatewayStats.Stats.V1InInfoV1Rsp,
			"v1_out_info_v1_req":        ret.DtVpnIkeGatewayStats.Stats.V1OutInfoV1Req,
			"v1_out_info_v1_rsp":        ret.DtVpnIkeGatewayStats.Stats.V1OutInfoV1Rsp,
			"v1_in_transaction_req":     ret.DtVpnIkeGatewayStats.Stats.V1InTransactionReq,
			"v1_in_transaction_rsp":     ret.DtVpnIkeGatewayStats.Stats.V1InTransactionRsp,
			"v1_out_transaction_req":    ret.DtVpnIkeGatewayStats.Stats.V1OutTransactionReq,
			"v1_out_transaction_rsp":    ret.DtVpnIkeGatewayStats.Stats.V1OutTransactionRsp,
			"v1_in_quick_mode_req":      ret.DtVpnIkeGatewayStats.Stats.V1InQuickModeReq,
			"v1_in_quick_mode_rsp":      ret.DtVpnIkeGatewayStats.Stats.V1InQuickModeRsp,
			"v1_out_quick_mode_req":     ret.DtVpnIkeGatewayStats.Stats.V1OutQuickModeReq,
			"v1_out_quick_mode_rsp":     ret.DtVpnIkeGatewayStats.Stats.V1OutQuickModeRsp,
			"v1_in_new_group_mode_req":  ret.DtVpnIkeGatewayStats.Stats.V1InNewGroupModeReq,
			"v1_in_new_group_mode_rsp":  ret.DtVpnIkeGatewayStats.Stats.V1InNewGroupModeRsp,
			"v1_out_new_group_mode_req": ret.DtVpnIkeGatewayStats.Stats.V1OutNewGroupModeReq,
			"v1_out_new_group_mode_rsp": ret.DtVpnIkeGatewayStats.Stats.V1OutNewGroupModeRsp,
			"v1_child_sa_invalid_spi":   ret.DtVpnIkeGatewayStats.Stats.V1ChildSaInvalidSpi,
			"v2_child_sa_invalid_spi":   ret.DtVpnIkeGatewayStats.Stats.V2ChildSaInvalidSpi,
			"ike_current_version":       ret.DtVpnIkeGatewayStats.Stats.IkeCurrentVersion,
		},
	}
}

func getObjectVpnIkeGatewayStatsStats(d []interface{}) edpt.VpnIkeGatewayStatsStats {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V2InitRekey = in["v2_init_rekey"].(int)
		ret.V2RspRekey = in["v2_rsp_rekey"].(int)
		ret.V2ChildSaRekey = in["v2_child_sa_rekey"].(int)
		ret.V2InInvalid = in["v2_in_invalid"].(int)
		ret.V2InInvalidSpi = in["v2_in_invalid_spi"].(int)
		ret.V2InInitReq = in["v2_in_init_req"].(int)
		ret.V2InInitRsp = in["v2_in_init_rsp"].(int)
		ret.V2OutInitReq = in["v2_out_init_req"].(int)
		ret.V2OutInitRsp = in["v2_out_init_rsp"].(int)
		ret.V2InAuthReq = in["v2_in_auth_req"].(int)
		ret.V2InAuthRsp = in["v2_in_auth_rsp"].(int)
		ret.V2OutAuthReq = in["v2_out_auth_req"].(int)
		ret.V2OutAuthRsp = in["v2_out_auth_rsp"].(int)
		ret.V2InCreateChildReq = in["v2_in_create_child_req"].(int)
		ret.V2InCreateChildRsp = in["v2_in_create_child_rsp"].(int)
		ret.V2OutCreateChildReq = in["v2_out_create_child_req"].(int)
		ret.V2OutCreateChildRsp = in["v2_out_create_child_rsp"].(int)
		ret.V2InInfoReq = in["v2_in_info_req"].(int)
		ret.V2InInfoRsp = in["v2_in_info_rsp"].(int)
		ret.V2OutInfoReq = in["v2_out_info_req"].(int)
		ret.V2OutInfoRsp = in["v2_out_info_rsp"].(int)
		ret.V1InIdProtReq = in["v1_in_id_prot_req"].(int)
		ret.V1InIdProtRsp = in["v1_in_id_prot_rsp"].(int)
		ret.V1OutIdProtReq = in["v1_out_id_prot_req"].(int)
		ret.V1OutIdProtRsp = in["v1_out_id_prot_rsp"].(int)
		ret.V1InAuthOnlyReq = in["v1_in_auth_only_req"].(int)
		ret.V1InAuthOnlyRsp = in["v1_in_auth_only_rsp"].(int)
		ret.V1OutAuthOnlyReq = in["v1_out_auth_only_req"].(int)
		ret.V1OutAuthOnlyRsp = in["v1_out_auth_only_rsp"].(int)
		ret.V1InAggressiveReq = in["v1_in_aggressive_req"].(int)
		ret.V1InAggressiveRsp = in["v1_in_aggressive_rsp"].(int)
		ret.V1OutAggressiveReq = in["v1_out_aggressive_req"].(int)
		ret.V1OutAggressiveRsp = in["v1_out_aggressive_rsp"].(int)
		ret.V1InInfoV1Req = in["v1_in_info_v1_req"].(int)
		ret.V1InInfoV1Rsp = in["v1_in_info_v1_rsp"].(int)
		ret.V1OutInfoV1Req = in["v1_out_info_v1_req"].(int)
		ret.V1OutInfoV1Rsp = in["v1_out_info_v1_rsp"].(int)
		ret.V1InTransactionReq = in["v1_in_transaction_req"].(int)
		ret.V1InTransactionRsp = in["v1_in_transaction_rsp"].(int)
		ret.V1OutTransactionReq = in["v1_out_transaction_req"].(int)
		ret.V1OutTransactionRsp = in["v1_out_transaction_rsp"].(int)
		ret.V1InQuickModeReq = in["v1_in_quick_mode_req"].(int)
		ret.V1InQuickModeRsp = in["v1_in_quick_mode_rsp"].(int)
		ret.V1OutQuickModeReq = in["v1_out_quick_mode_req"].(int)
		ret.V1OutQuickModeRsp = in["v1_out_quick_mode_rsp"].(int)
		ret.V1InNewGroupModeReq = in["v1_in_new_group_mode_req"].(int)
		ret.V1InNewGroupModeRsp = in["v1_in_new_group_mode_rsp"].(int)
		ret.V1OutNewGroupModeReq = in["v1_out_new_group_mode_req"].(int)
		ret.V1OutNewGroupModeRsp = in["v1_out_new_group_mode_rsp"].(int)
		ret.V1ChildSaInvalidSpi = in["v1_child_sa_invalid_spi"].(int)
		ret.V2ChildSaInvalidSpi = in["v2_child_sa_invalid_spi"].(int)
		ret.IkeCurrentVersion = in["ike_current_version"].(int)
	}
	return ret
}

func dataToEndpointVpnIkeGatewayStats(d *schema.ResourceData) edpt.VpnIkeGatewayStats {
	var ret edpt.VpnIkeGatewayStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectVpnIkeGatewayStatsStats(d.Get("stats").([]interface{}))
	return ret
}
