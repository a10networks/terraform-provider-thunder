package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDynamicServiceDnsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dynamic_service_dns_stats`: Statistics for the object dynamic-service-dns\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDynamicServiceDnsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"n_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries",
						},
						"n_query_err": {
							Type: schema.TypeInt, Optional: true, Description: "Number of query errors",
						},
						"n_query_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries dropped due to full queue",
						},
						"n_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Number of responses",
						},
						"n_resp_err": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response errors",
						},
						"n_resp_f_formerr": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode FormErr",
						},
						"n_resp_f_servfail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode ServFail",
						},
						"n_resp_f_nxdomain": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode NXDomain",
						},
						"n_resp_f_notimp": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode NotImp",
						},
						"n_resp_f_refused": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode Refused",
						},
						"n_resp_f_yxdomain": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode YXDomain",
						},
						"n_resp_f_yxrrset": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode YXRRSet",
						},
						"n_resp_f_nxrrset": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode NXRRSet",
						},
						"n_resp_f_notauth": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode NotAuth",
						},
						"n_resp_f_notzone": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode NotZone",
						},
						"n_resp_f_dsotypeni": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode DSOTYPENI",
						},
						"n_resp_f_badvers": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADVERS",
						},
						"n_resp_f_badkey": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADKEY",
						},
						"n_resp_f_badtime": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADTIME",
						},
						"n_resp_f_badmode": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADMODE",
						},
						"n_resp_f_badname": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADNAME",
						},
						"n_resp_f_badalg": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADALG",
						},
						"n_resp_f_badtrunc": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADTRUNC",
						},
						"n_resp_f_badcookie": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with rcode BADCOOKIE",
						},
						"n_resp_f_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Number of response failure with invalid rcode",
						},
					},
				},
			},
		},
	}
}

func resourceSlbDynamicServiceDnsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDynamicServiceDnsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDynamicServiceDnsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDynamicServiceDnsStatsStats := setObjectSlbDynamicServiceDnsStatsStats(res)
		d.Set("stats", SlbDynamicServiceDnsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDynamicServiceDnsStatsStats(ret edpt.DataSlbDynamicServiceDnsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"n_query":            ret.DtSlbDynamicServiceDnsStats.Stats.N_query,
			"n_query_err":        ret.DtSlbDynamicServiceDnsStats.Stats.N_query_err,
			"n_query_drop":       ret.DtSlbDynamicServiceDnsStats.Stats.N_query_drop,
			"n_resp":             ret.DtSlbDynamicServiceDnsStats.Stats.N_resp,
			"n_resp_err":         ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_err,
			"n_resp_f_formerr":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_formerr,
			"n_resp_f_servfail":  ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_servfail,
			"n_resp_f_nxdomain":  ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_nxdomain,
			"n_resp_f_notimp":    ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_notimp,
			"n_resp_f_refused":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_refused,
			"n_resp_f_yxdomain":  ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_yxdomain,
			"n_resp_f_yxrrset":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_yxrrset,
			"n_resp_f_nxrrset":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_nxrrset,
			"n_resp_f_notauth":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_notauth,
			"n_resp_f_notzone":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_notzone,
			"n_resp_f_dsotypeni": ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_dsotypeni,
			"n_resp_f_badvers":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badvers,
			"n_resp_f_badkey":    ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badkey,
			"n_resp_f_badtime":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badtime,
			"n_resp_f_badmode":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badmode,
			"n_resp_f_badname":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badname,
			"n_resp_f_badalg":    ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badalg,
			"n_resp_f_badtrunc":  ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badtrunc,
			"n_resp_f_badcookie": ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_badcookie,
			"n_resp_f_invalid":   ret.DtSlbDynamicServiceDnsStats.Stats.N_resp_f_invalid,
		},
	}
}

func getObjectSlbDynamicServiceDnsStatsStats(d []interface{}) edpt.SlbDynamicServiceDnsStatsStats {

	count1 := len(d)
	var ret edpt.SlbDynamicServiceDnsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.N_query = in["n_query"].(int)
		ret.N_query_err = in["n_query_err"].(int)
		ret.N_query_drop = in["n_query_drop"].(int)
		ret.N_resp = in["n_resp"].(int)
		ret.N_resp_err = in["n_resp_err"].(int)
		ret.N_resp_f_formerr = in["n_resp_f_formerr"].(int)
		ret.N_resp_f_servfail = in["n_resp_f_servfail"].(int)
		ret.N_resp_f_nxdomain = in["n_resp_f_nxdomain"].(int)
		ret.N_resp_f_notimp = in["n_resp_f_notimp"].(int)
		ret.N_resp_f_refused = in["n_resp_f_refused"].(int)
		ret.N_resp_f_yxdomain = in["n_resp_f_yxdomain"].(int)
		ret.N_resp_f_yxrrset = in["n_resp_f_yxrrset"].(int)
		ret.N_resp_f_nxrrset = in["n_resp_f_nxrrset"].(int)
		ret.N_resp_f_notauth = in["n_resp_f_notauth"].(int)
		ret.N_resp_f_notzone = in["n_resp_f_notzone"].(int)
		ret.N_resp_f_dsotypeni = in["n_resp_f_dsotypeni"].(int)
		ret.N_resp_f_badvers = in["n_resp_f_badvers"].(int)
		ret.N_resp_f_badkey = in["n_resp_f_badkey"].(int)
		ret.N_resp_f_badtime = in["n_resp_f_badtime"].(int)
		ret.N_resp_f_badmode = in["n_resp_f_badmode"].(int)
		ret.N_resp_f_badname = in["n_resp_f_badname"].(int)
		ret.N_resp_f_badalg = in["n_resp_f_badalg"].(int)
		ret.N_resp_f_badtrunc = in["n_resp_f_badtrunc"].(int)
		ret.N_resp_f_badcookie = in["n_resp_f_badcookie"].(int)
		ret.N_resp_f_invalid = in["n_resp_f_invalid"].(int)
	}
	return ret
}

func dataToEndpointSlbDynamicServiceDnsStats(d *schema.ResourceData) edpt.SlbDynamicServiceDnsStats {
	var ret edpt.SlbDynamicServiceDnsStats

	ret.Stats = getObjectSlbDynamicServiceDnsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
