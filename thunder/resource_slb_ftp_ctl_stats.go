package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFtpCtlStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ftp_ctl_stats`: Statistics for the object ftp-ctl\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFtpCtlStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sessions_num": {
							Type: schema.TypeInt, Optional: true, Description: "Total Control Sessions",
						},
						"alg_pkts_num": {
							Type: schema.TypeInt, Optional: true, Description: "Total ALG packets",
						},
						"alg_pkts_xmitted_num": {
							Type: schema.TypeInt, Optional: true, Description: "ALG packets rexmitted",
						},
						"alg_port_helper_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total PORT helper sessions",
						},
						"alg_pasv_helper_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total PASV helper sessions",
						},
						"alg_port_helper_freed_unused": {
							Type: schema.TypeInt, Optional: true, Description: "PORT helper freed unused",
						},
						"alg_pasv_helper_freed_unused": {
							Type: schema.TypeInt, Optional: true, Description: "PASV helper freed unused",
						},
						"alg_port_helper_nat_free": {
							Type: schema.TypeInt, Optional: true, Description: "PORT helper NAT free",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFtpCtlStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpCtlStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpCtlStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFtpCtlStatsStats := setObjectSlbFtpCtlStatsStats(res)
		d.Set("stats", SlbFtpCtlStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFtpCtlStatsStats(ret edpt.DataSlbFtpCtlStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sessions_num":                 ret.DtSlbFtpCtlStats.Stats.Sessions_num,
			"alg_pkts_num":                 ret.DtSlbFtpCtlStats.Stats.Alg_pkts_num,
			"alg_pkts_xmitted_num":         ret.DtSlbFtpCtlStats.Stats.Alg_pkts_xmitted_num,
			"alg_port_helper_created":      ret.DtSlbFtpCtlStats.Stats.Alg_port_helper_created,
			"alg_pasv_helper_created":      ret.DtSlbFtpCtlStats.Stats.Alg_pasv_helper_created,
			"alg_port_helper_freed_unused": ret.DtSlbFtpCtlStats.Stats.Alg_port_helper_freed_unused,
			"alg_pasv_helper_freed_unused": ret.DtSlbFtpCtlStats.Stats.Alg_pasv_helper_freed_unused,
			"alg_port_helper_nat_free":     ret.DtSlbFtpCtlStats.Stats.Alg_port_helper_nat_free,
		},
	}
}

func getObjectSlbFtpCtlStatsStats(d []interface{}) edpt.SlbFtpCtlStatsStats {

	count1 := len(d)
	var ret edpt.SlbFtpCtlStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sessions_num = in["sessions_num"].(int)
		ret.Alg_pkts_num = in["alg_pkts_num"].(int)
		ret.Alg_pkts_xmitted_num = in["alg_pkts_xmitted_num"].(int)
		ret.Alg_port_helper_created = in["alg_port_helper_created"].(int)
		ret.Alg_pasv_helper_created = in["alg_pasv_helper_created"].(int)
		ret.Alg_port_helper_freed_unused = in["alg_port_helper_freed_unused"].(int)
		ret.Alg_pasv_helper_freed_unused = in["alg_pasv_helper_freed_unused"].(int)
		ret.Alg_port_helper_nat_free = in["alg_port_helper_nat_free"].(int)
	}
	return ret
}

func dataToEndpointSlbFtpCtlStats(d *schema.ResourceData) edpt.SlbFtpCtlStats {
	var ret edpt.SlbFtpCtlStats

	ret.Stats = getObjectSlbFtpCtlStatsStats(d.Get("stats").([]interface{}))
	return ret
}
