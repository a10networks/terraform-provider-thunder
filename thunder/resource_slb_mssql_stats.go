package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMssqlStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_mssql_stats`: Statistics for the object mssql\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbMssqlStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Curr Proxy Conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total Proxy Conns",
						},
						"curr_be_enc": {
							Type: schema.TypeInt, Optional: true, Description: "Curr BE Encryption Conns",
						},
						"total_be_enc": {
							Type: schema.TypeInt, Optional: true, Description: "Total BE Encryption Conns",
						},
						"curr_fe_enc": {
							Type: schema.TypeInt, Optional: true, Description: "Curr FE Encryption Conns",
						},
						"total_fe_enc": {
							Type: schema.TypeInt, Optional: true, Description: "Total FE Encryption Conns",
						},
						"client_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Client FIN",
						},
						"server_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Server FIN",
						},
						"session_err": {
							Type: schema.TypeInt, Optional: true, Description: "Session err",
						},
						"queries": {
							Type: schema.TypeInt, Optional: true, Description: "DB Queries",
						},
						"commands": {
							Type: schema.TypeInt, Optional: true, Description: "DB commands reply",
						},
						"auth_success": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication Success",
						},
						"auth_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication Failure",
						},
					},
				},
			},
		},
	}
}

func resourceSlbMssqlStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMssqlStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMssqlStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbMssqlStatsStats := setObjectSlbMssqlStatsStats(res)
		d.Set("stats", SlbMssqlStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbMssqlStatsStats(ret edpt.DataSlbMssqlStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":   ret.DtSlbMssqlStats.Stats.Curr_proxy,
			"total_proxy":  ret.DtSlbMssqlStats.Stats.Total_proxy,
			"curr_be_enc":  ret.DtSlbMssqlStats.Stats.Curr_be_enc,
			"total_be_enc": ret.DtSlbMssqlStats.Stats.Total_be_enc,
			"curr_fe_enc":  ret.DtSlbMssqlStats.Stats.Curr_fe_enc,
			"total_fe_enc": ret.DtSlbMssqlStats.Stats.Total_fe_enc,
			"client_fin":   ret.DtSlbMssqlStats.Stats.Client_fin,
			"server_fin":   ret.DtSlbMssqlStats.Stats.Server_fin,
			"session_err":  ret.DtSlbMssqlStats.Stats.Session_err,
			"queries":      ret.DtSlbMssqlStats.Stats.Queries,
			"commands":     ret.DtSlbMssqlStats.Stats.Commands,
			"auth_success": ret.DtSlbMssqlStats.Stats.Auth_success,
			"auth_failure": ret.DtSlbMssqlStats.Stats.Auth_failure,
		},
	}
}

func getObjectSlbMssqlStatsStats(d []interface{}) edpt.SlbMssqlStatsStats {

	count1 := len(d)
	var ret edpt.SlbMssqlStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Curr_be_enc = in["curr_be_enc"].(int)
		ret.Total_be_enc = in["total_be_enc"].(int)
		ret.Curr_fe_enc = in["curr_fe_enc"].(int)
		ret.Total_fe_enc = in["total_fe_enc"].(int)
		ret.Client_fin = in["client_fin"].(int)
		ret.Server_fin = in["server_fin"].(int)
		ret.Session_err = in["session_err"].(int)
		ret.Queries = in["queries"].(int)
		ret.Commands = in["commands"].(int)
		ret.Auth_success = in["auth_success"].(int)
		ret.Auth_failure = in["auth_failure"].(int)
	}
	return ret
}

func dataToEndpointSlbMssqlStats(d *schema.ResourceData) edpt.SlbMssqlStats {
	var ret edpt.SlbMssqlStats

	ret.Stats = getObjectSlbMssqlStatsStats(d.Get("stats").([]interface{}))
	return ret
}
