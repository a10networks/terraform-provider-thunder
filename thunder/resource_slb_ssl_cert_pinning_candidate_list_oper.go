package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCertPinningCandidateListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cert_pinning_candidate_list_oper`: Operational Status for the object ssl-cert-pinning-candidate-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCertPinningCandidateListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"candidate_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"servername": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"conn_failure_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bypass_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"list_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"mid_request": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"central_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"local_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"bypassed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"alphabet_order": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stats_only": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslCertPinningCandidateListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertPinningCandidateListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertPinningCandidateListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCertPinningCandidateListOperOper := setObjectSlbSslCertPinningCandidateListOperOper(res)
		d.Set("oper", SlbSslCertPinningCandidateListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCertPinningCandidateListOperOper(ret edpt.DataSlbSslCertPinningCandidateListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"candidate_list": setSliceSlbSslCertPinningCandidateListOperOperCandidateList(ret.DtSlbSslCertPinningCandidateListOper.Oper.CandidateList),
			"mid_request":    ret.DtSlbSslCertPinningCandidateListOper.Oper.MidRequest,
			"server_name":    ret.DtSlbSslCertPinningCandidateListOper.Oper.ServerName,
			"central_list":   ret.DtSlbSslCertPinningCandidateListOper.Oper.CentralList,
			"local_list":     ret.DtSlbSslCertPinningCandidateListOper.Oper.LocalList,
			"bypassed":       ret.DtSlbSslCertPinningCandidateListOper.Oper.Bypassed,
			"alphabet_order": ret.DtSlbSslCertPinningCandidateListOper.Oper.AlphabetOrder,
			"all":            ret.DtSlbSslCertPinningCandidateListOper.Oper.All,
			"stats_only":     ret.DtSlbSslCertPinningCandidateListOper.Oper.StatsOnly,
		},
	}
}

func setSliceSlbSslCertPinningCandidateListOperOperCandidateList(d []edpt.SlbSslCertPinningCandidateListOperOperCandidateList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["servername"] = item.Servername
		in["conn_failure_count"] = item.ConnFailureCount
		in["bypass_count"] = item.BypassCount
		in["list_type"] = item.ListType
		in["ttl"] = item.Ttl
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslCertPinningCandidateListOperOper(d []interface{}) edpt.SlbSslCertPinningCandidateListOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCertPinningCandidateListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CandidateList = getSliceSlbSslCertPinningCandidateListOperOperCandidateList(in["candidate_list"].([]interface{}))
		ret.MidRequest = in["mid_request"].(int)
		ret.ServerName = in["server_name"].(string)
		ret.CentralList = in["central_list"].(int)
		ret.LocalList = in["local_list"].(int)
		ret.Bypassed = in["bypassed"].(int)
		ret.AlphabetOrder = in["alphabet_order"].(int)
		ret.All = in["all"].(int)
		ret.StatsOnly = in["stats_only"].(int)
	}
	return ret
}

func getSliceSlbSslCertPinningCandidateListOperOperCandidateList(d []interface{}) []edpt.SlbSslCertPinningCandidateListOperOperCandidateList {

	count1 := len(d)
	ret := make([]edpt.SlbSslCertPinningCandidateListOperOperCandidateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslCertPinningCandidateListOperOperCandidateList
		oi.Servername = in["servername"].(string)
		oi.ConnFailureCount = in["conn_failure_count"].(int)
		oi.BypassCount = in["bypass_count"].(int)
		oi.ListType = in["list_type"].(string)
		oi.Ttl = in["ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslCertPinningCandidateListOper(d *schema.ResourceData) edpt.SlbSslCertPinningCandidateListOper {
	var ret edpt.SlbSslCertPinningCandidateListOper

	ret.Oper = getObjectSlbSslCertPinningCandidateListOperOper(d.Get("oper").([]interface{}))
	return ret
}
