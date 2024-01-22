package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMysqlStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_mysql_stats`: Statistics for the object mysql\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbMysqlStatsRead,

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
					},
				},
			},
		},
	}
}

func resourceSlbMysqlStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMysqlStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMysqlStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbMysqlStatsStats := setObjectSlbMysqlStatsStats(res)
		d.Set("stats", SlbMysqlStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbMysqlStatsStats(ret edpt.DataSlbMysqlStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":   ret.DtSlbMysqlStats.Stats.Curr_proxy,
			"total_proxy":  ret.DtSlbMysqlStats.Stats.Total_proxy,
			"curr_be_enc":  ret.DtSlbMysqlStats.Stats.Curr_be_enc,
			"total_be_enc": ret.DtSlbMysqlStats.Stats.Total_be_enc,
			"curr_fe_enc":  ret.DtSlbMysqlStats.Stats.Curr_fe_enc,
			"total_fe_enc": ret.DtSlbMysqlStats.Stats.Total_fe_enc,
			"client_fin":   ret.DtSlbMysqlStats.Stats.Client_fin,
			"server_fin":   ret.DtSlbMysqlStats.Stats.Server_fin,
			"session_err":  ret.DtSlbMysqlStats.Stats.Session_err,
			"queries":      ret.DtSlbMysqlStats.Stats.Queries,
			"commands":     ret.DtSlbMysqlStats.Stats.Commands,
		},
	}
}

func getObjectSlbMysqlStatsStats(d []interface{}) edpt.SlbMysqlStatsStats {

	count1 := len(d)
	var ret edpt.SlbMysqlStatsStats
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
	}
	return ret
}

func dataToEndpointSlbMysqlStats(d *schema.ResourceData) edpt.SlbMysqlStats {
	var ret edpt.SlbMysqlStats

	ret.Stats = getObjectSlbMysqlStatsStats(d.Get("stats").([]interface{}))
	return ret
}
