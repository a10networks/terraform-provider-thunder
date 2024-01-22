package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslReqQStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ssl_req_q_stats`: Statistics for the object ssl-req-q\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemSslReqQStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_ssl_queues": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ssl_req_q_depth_tot": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ssl_req_q_inuse_tot": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ssl_hw_q_depth_tot": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ssl_hw_q_inuse_tot": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemSslReqQStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslReqQStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslReqQStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemSslReqQStatsStats := setObjectSystemSslReqQStatsStats(res)
		d.Set("stats", SystemSslReqQStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemSslReqQStatsStats(ret edpt.DataSystemSslReqQStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_ssl_queues":      ret.DtSystemSslReqQStats.Stats.NumSslQueues,
			"ssl_req_q_depth_tot": ret.DtSystemSslReqQStats.Stats.SslReqQDepthTot,
			"ssl_req_q_inuse_tot": ret.DtSystemSslReqQStats.Stats.SslReqQInuseTot,
			"ssl_hw_q_depth_tot":  ret.DtSystemSslReqQStats.Stats.SslHwQDepthTot,
			"ssl_hw_q_inuse_tot":  ret.DtSystemSslReqQStats.Stats.SslHwQInuseTot,
		},
	}
}

func getObjectSystemSslReqQStatsStats(d []interface{}) edpt.SystemSslReqQStatsStats {

	count1 := len(d)
	var ret edpt.SystemSslReqQStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumSslQueues = in["num_ssl_queues"].(int)
		ret.SslReqQDepthTot = in["ssl_req_q_depth_tot"].(int)
		ret.SslReqQInuseTot = in["ssl_req_q_inuse_tot"].(int)
		ret.SslHwQDepthTot = in["ssl_hw_q_depth_tot"].(int)
		ret.SslHwQInuseTot = in["ssl_hw_q_inuse_tot"].(int)
	}
	return ret
}

func dataToEndpointSystemSslReqQStats(d *schema.ResourceData) edpt.SystemSslReqQStats {
	var ret edpt.SystemSslReqQStats

	ret.Stats = getObjectSystemSslReqQStatsStats(d.Get("stats").([]interface{}))
	return ret
}
