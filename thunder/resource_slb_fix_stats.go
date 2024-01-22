package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFixStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_fix_stats`: Statistics for the object fix\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFixStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Current proxy conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total proxy conns",
						},
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection failure",
						},
						"noroute": {
							Type: schema.TypeInt, Optional: true, Description: "No route failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT failure",
						},
						"client_err": {
							Type: schema.TypeInt, Optional: true, Description: "Client fail",
						},
						"server_err": {
							Type: schema.TypeInt, Optional: true, Description: "Server fail",
						},
						"insert_clientip": {
							Type: schema.TypeInt, Optional: true, Description: "Insert client IP",
						},
						"default_switching": {
							Type: schema.TypeInt, Optional: true, Description: "Default switching",
						},
						"sender_switching": {
							Type: schema.TypeInt, Optional: true, Description: "Sender ID switching",
						},
						"target_switching": {
							Type: schema.TypeInt, Optional: true, Description: "Target ID switching",
						},
						"client_tls_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Client TLS conn",
						},
						"server_tls_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Server TLS conn",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFixStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFixStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFixStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFixStatsStats := setObjectSlbFixStatsStats(res)
		d.Set("stats", SlbFixStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFixStatsStats(ret edpt.DataSlbFixStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":        ret.DtSlbFixStats.Stats.Curr_proxy,
			"total_proxy":       ret.DtSlbFixStats.Stats.Total_proxy,
			"svrsel_fail":       ret.DtSlbFixStats.Stats.Svrsel_fail,
			"noroute":           ret.DtSlbFixStats.Stats.Noroute,
			"snat_fail":         ret.DtSlbFixStats.Stats.Snat_fail,
			"client_err":        ret.DtSlbFixStats.Stats.Client_err,
			"server_err":        ret.DtSlbFixStats.Stats.Server_err,
			"insert_clientip":   ret.DtSlbFixStats.Stats.Insert_clientip,
			"default_switching": ret.DtSlbFixStats.Stats.Default_switching,
			"sender_switching":  ret.DtSlbFixStats.Stats.Sender_switching,
			"target_switching":  ret.DtSlbFixStats.Stats.Target_switching,
			"client_tls_conn":   ret.DtSlbFixStats.Stats.Client_tls_conn,
			"server_tls_conn":   ret.DtSlbFixStats.Stats.Server_tls_conn,
		},
	}
}

func getObjectSlbFixStatsStats(d []interface{}) edpt.SlbFixStatsStats {

	count1 := len(d)
	var ret edpt.SlbFixStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Noroute = in["noroute"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_err = in["client_err"].(int)
		ret.Server_err = in["server_err"].(int)
		ret.Insert_clientip = in["insert_clientip"].(int)
		ret.Default_switching = in["default_switching"].(int)
		ret.Sender_switching = in["sender_switching"].(int)
		ret.Target_switching = in["target_switching"].(int)
		ret.Client_tls_conn = in["client_tls_conn"].(int)
		ret.Server_tls_conn = in["server_tls_conn"].(int)
	}
	return ret
}

func dataToEndpointSlbFixStats(d *schema.ResourceData) edpt.SlbFixStats {
	var ret edpt.SlbFixStats

	ret.Stats = getObjectSlbFixStatsStats(d.Get("stats").([]interface{}))
	return ret
}
