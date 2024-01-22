package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeStatsByGwOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ike_stats_by_gw_oper`: Operational Status for the object ike-stats-by-gw\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIkeStatsByGwOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway_name_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"remote_ip_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"remote_id_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"display_all_filter": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ike_stats_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"v1_in_id_prot_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_id_prot_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_id_prot_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_id_prot_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_auth_only_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_auth_only_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_auth_only_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_auth_only_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_aggressive_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_aggressive_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_aggressive_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_aggressive_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_info_v1_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_info_v1_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_info_v1_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_info_v1_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_transaction_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_transaction_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_transaction_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_transaction_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_quick_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_quick_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_quick_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_quick_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_new_group_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_in_new_group_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_new_group_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_out_new_group_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v1_child_sa_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_init_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_rsp_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_child_sa_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_invalid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_init_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_init_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_init_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_init_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_auth_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_auth_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_create_child_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_create_child_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_create_child_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_create_child_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_info_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_in_info_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_info_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_out_info_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"v2_child_sa_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVpnIkeStatsByGwOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeStatsByGwOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeStatsByGwOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIkeStatsByGwOperOper := setObjectVpnIkeStatsByGwOperOper(res)
		d.Set("oper", VpnIkeStatsByGwOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIkeStatsByGwOperOper(ret edpt.DataVpnIkeStatsByGwOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"gateway_name_filter": ret.DtVpnIkeStatsByGwOper.Oper.GatewayNameFilter,
			"remote_ip_filter":    ret.DtVpnIkeStatsByGwOper.Oper.RemoteIpFilter,
			"remote_id_filter":    ret.DtVpnIkeStatsByGwOper.Oper.RemoteIdFilter,
			"display_all_filter":  ret.DtVpnIkeStatsByGwOper.Oper.DisplayAllFilter,
			"ike_stats_list":      setSliceVpnIkeStatsByGwOperOperIkeStatsList(ret.DtVpnIkeStatsByGwOper.Oper.IkeStatsList),
		},
	}
}

func setSliceVpnIkeStatsByGwOperOperIkeStatsList(d []edpt.VpnIkeStatsByGwOperOperIkeStatsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["remote_id"] = item.RemoteId
		in["remote_ip"] = item.RemoteIp
		in["ike_version"] = item.IkeVersion
		in["v1_in_id_prot_req"] = item.V1InIdProtReq
		in["v1_in_id_prot_rsp"] = item.V1InIdProtRsp
		in["v1_out_id_prot_req"] = item.V1OutIdProtReq
		in["v1_out_id_prot_rsp"] = item.V1OutIdProtRsp
		in["v1_in_auth_only_req"] = item.V1InAuthOnlyReq
		in["v1_in_auth_only_rsp"] = item.V1InAuthOnlyRsp
		in["v1_out_auth_only_req"] = item.V1OutAuthOnlyReq
		in["v1_out_auth_only_rsp"] = item.V1OutAuthOnlyRsp
		in["v1_in_aggressive_req"] = item.V1InAggressiveReq
		in["v1_in_aggressive_rsp"] = item.V1InAggressiveRsp
		in["v1_out_aggressive_req"] = item.V1OutAggressiveReq
		in["v1_out_aggressive_rsp"] = item.V1OutAggressiveRsp
		in["v1_in_info_v1_req"] = item.V1InInfoV1Req
		in["v1_in_info_v1_rsp"] = item.V1InInfoV1Rsp
		in["v1_out_info_v1_req"] = item.V1OutInfoV1Req
		in["v1_out_info_v1_rsp"] = item.V1OutInfoV1Rsp
		in["v1_in_transaction_req"] = item.V1InTransactionReq
		in["v1_in_transaction_rsp"] = item.V1InTransactionRsp
		in["v1_out_transaction_req"] = item.V1OutTransactionReq
		in["v1_out_transaction_rsp"] = item.V1OutTransactionRsp
		in["v1_in_quick_mode_req"] = item.V1InQuickModeReq
		in["v1_in_quick_mode_rsp"] = item.V1InQuickModeRsp
		in["v1_out_quick_mode_req"] = item.V1OutQuickModeReq
		in["v1_out_quick_mode_rsp"] = item.V1OutQuickModeRsp
		in["v1_in_new_group_mode_req"] = item.V1InNewGroupModeReq
		in["v1_in_new_group_mode_rsp"] = item.V1InNewGroupModeRsp
		in["v1_out_new_group_mode_req"] = item.V1OutNewGroupModeReq
		in["v1_out_new_group_mode_rsp"] = item.V1OutNewGroupModeRsp
		in["v1_child_sa_invalid_spi"] = item.V1ChildSaInvalidSpi
		in["v2_init_rekey"] = item.V2InitRekey
		in["v2_rsp_rekey"] = item.V2RspRekey
		in["v2_child_sa_rekey"] = item.V2ChildSaRekey
		in["v2_in_invalid"] = item.V2InInvalid
		in["v2_in_invalid_spi"] = item.V2InInvalidSpi
		in["v2_in_init_req"] = item.V2InInitReq
		in["v2_in_init_rsp"] = item.V2InInitRsp
		in["v2_out_init_req"] = item.V2OutInitReq
		in["v2_out_init_rsp"] = item.V2OutInitRsp
		in["v2_in_auth_req"] = item.V2InAuthReq
		in["v2_in_auth_rsp"] = item.V2InAuthRsp
		in["v2_out_auth_req"] = item.V2OutAuthReq
		in["v2_out_auth_rsp"] = item.V2OutAuthRsp
		in["v2_in_create_child_req"] = item.V2InCreateChildReq
		in["v2_in_create_child_rsp"] = item.V2InCreateChildRsp
		in["v2_out_create_child_req"] = item.V2OutCreateChildReq
		in["v2_out_create_child_rsp"] = item.V2OutCreateChildRsp
		in["v2_in_info_req"] = item.V2InInfoReq
		in["v2_in_info_rsp"] = item.V2InInfoRsp
		in["v2_out_info_req"] = item.V2OutInfoReq
		in["v2_out_info_rsp"] = item.V2OutInfoRsp
		in["v2_child_sa_invalid_spi"] = item.V2ChildSaInvalidSpi
		result = append(result, in)
	}
	return result
}

func getObjectVpnIkeStatsByGwOperOper(d []interface{}) edpt.VpnIkeStatsByGwOperOper {

	count1 := len(d)
	var ret edpt.VpnIkeStatsByGwOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GatewayNameFilter = in["gateway_name_filter"].(string)
		ret.RemoteIpFilter = in["remote_ip_filter"].(string)
		ret.RemoteIdFilter = in["remote_id_filter"].(string)
		ret.DisplayAllFilter = in["display_all_filter"].(int)
		ret.IkeStatsList = getSliceVpnIkeStatsByGwOperOperIkeStatsList(in["ike_stats_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnIkeStatsByGwOperOperIkeStatsList(d []interface{}) []edpt.VpnIkeStatsByGwOperOperIkeStatsList {

	count1 := len(d)
	ret := make([]edpt.VpnIkeStatsByGwOperOperIkeStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeStatsByGwOperOperIkeStatsList
		oi.Name = in["name"].(string)
		oi.RemoteId = in["remote_id"].(string)
		oi.RemoteIp = in["remote_ip"].(string)
		oi.IkeVersion = in["ike_version"].(string)
		oi.V1InIdProtReq = in["v1_in_id_prot_req"].(int)
		oi.V1InIdProtRsp = in["v1_in_id_prot_rsp"].(int)
		oi.V1OutIdProtReq = in["v1_out_id_prot_req"].(int)
		oi.V1OutIdProtRsp = in["v1_out_id_prot_rsp"].(int)
		oi.V1InAuthOnlyReq = in["v1_in_auth_only_req"].(int)
		oi.V1InAuthOnlyRsp = in["v1_in_auth_only_rsp"].(int)
		oi.V1OutAuthOnlyReq = in["v1_out_auth_only_req"].(int)
		oi.V1OutAuthOnlyRsp = in["v1_out_auth_only_rsp"].(int)
		oi.V1InAggressiveReq = in["v1_in_aggressive_req"].(int)
		oi.V1InAggressiveRsp = in["v1_in_aggressive_rsp"].(int)
		oi.V1OutAggressiveReq = in["v1_out_aggressive_req"].(int)
		oi.V1OutAggressiveRsp = in["v1_out_aggressive_rsp"].(int)
		oi.V1InInfoV1Req = in["v1_in_info_v1_req"].(int)
		oi.V1InInfoV1Rsp = in["v1_in_info_v1_rsp"].(int)
		oi.V1OutInfoV1Req = in["v1_out_info_v1_req"].(int)
		oi.V1OutInfoV1Rsp = in["v1_out_info_v1_rsp"].(int)
		oi.V1InTransactionReq = in["v1_in_transaction_req"].(int)
		oi.V1InTransactionRsp = in["v1_in_transaction_rsp"].(int)
		oi.V1OutTransactionReq = in["v1_out_transaction_req"].(int)
		oi.V1OutTransactionRsp = in["v1_out_transaction_rsp"].(int)
		oi.V1InQuickModeReq = in["v1_in_quick_mode_req"].(int)
		oi.V1InQuickModeRsp = in["v1_in_quick_mode_rsp"].(int)
		oi.V1OutQuickModeReq = in["v1_out_quick_mode_req"].(int)
		oi.V1OutQuickModeRsp = in["v1_out_quick_mode_rsp"].(int)
		oi.V1InNewGroupModeReq = in["v1_in_new_group_mode_req"].(int)
		oi.V1InNewGroupModeRsp = in["v1_in_new_group_mode_rsp"].(int)
		oi.V1OutNewGroupModeReq = in["v1_out_new_group_mode_req"].(int)
		oi.V1OutNewGroupModeRsp = in["v1_out_new_group_mode_rsp"].(int)
		oi.V1ChildSaInvalidSpi = in["v1_child_sa_invalid_spi"].(int)
		oi.V2InitRekey = in["v2_init_rekey"].(int)
		oi.V2RspRekey = in["v2_rsp_rekey"].(int)
		oi.V2ChildSaRekey = in["v2_child_sa_rekey"].(int)
		oi.V2InInvalid = in["v2_in_invalid"].(int)
		oi.V2InInvalidSpi = in["v2_in_invalid_spi"].(int)
		oi.V2InInitReq = in["v2_in_init_req"].(int)
		oi.V2InInitRsp = in["v2_in_init_rsp"].(int)
		oi.V2OutInitReq = in["v2_out_init_req"].(int)
		oi.V2OutInitRsp = in["v2_out_init_rsp"].(int)
		oi.V2InAuthReq = in["v2_in_auth_req"].(int)
		oi.V2InAuthRsp = in["v2_in_auth_rsp"].(int)
		oi.V2OutAuthReq = in["v2_out_auth_req"].(int)
		oi.V2OutAuthRsp = in["v2_out_auth_rsp"].(int)
		oi.V2InCreateChildReq = in["v2_in_create_child_req"].(int)
		oi.V2InCreateChildRsp = in["v2_in_create_child_rsp"].(int)
		oi.V2OutCreateChildReq = in["v2_out_create_child_req"].(int)
		oi.V2OutCreateChildRsp = in["v2_out_create_child_rsp"].(int)
		oi.V2InInfoReq = in["v2_in_info_req"].(int)
		oi.V2InInfoRsp = in["v2_in_info_rsp"].(int)
		oi.V2OutInfoReq = in["v2_out_info_req"].(int)
		oi.V2OutInfoRsp = in["v2_out_info_rsp"].(int)
		oi.V2ChildSaInvalidSpi = in["v2_child_sa_invalid_spi"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIkeStatsByGwOper(d *schema.ResourceData) edpt.VpnIkeStatsByGwOper {
	var ret edpt.VpnIkeStatsByGwOper

	ret.Oper = getObjectVpnIkeStatsByGwOperOper(d.Get("oper").([]interface{}))
	return ret
}
