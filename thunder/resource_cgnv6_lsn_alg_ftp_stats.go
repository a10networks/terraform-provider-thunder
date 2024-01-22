package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgFtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_ftp_stats`: Statistics for the object ftp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgFtpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_requests": {
							Type: schema.TypeInt, Optional: true, Description: "PORT Requests From Client",
						},
						"eprt_requests": {
							Type: schema.TypeInt, Optional: true, Description: "EPRT Requests From Client",
						},
						"lprt_requests": {
							Type: schema.TypeInt, Optional: true, Description: "LPRT Requests From Client",
						},
						"pasv_replies": {
							Type: schema.TypeInt, Optional: true, Description: "PASV Replies From Server",
						},
						"epsv_replies": {
							Type: schema.TypeInt, Optional: true, Description: "EPSV Replies From Server",
						},
						"lpsv_replies": {
							Type: schema.TypeInt, Optional: true, Description: "LPSV Replies From Server",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnAlgFtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgFtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgFtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgFtpStatsStats := setObjectCgnv6LsnAlgFtpStatsStats(res)
		d.Set("stats", Cgnv6LsnAlgFtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgFtpStatsStats(ret edpt.DataCgnv6LsnAlgFtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_requests": ret.DtCgnv6LsnAlgFtpStats.Stats.PortRequests,
			"eprt_requests": ret.DtCgnv6LsnAlgFtpStats.Stats.EprtRequests,
			"lprt_requests": ret.DtCgnv6LsnAlgFtpStats.Stats.LprtRequests,
			"pasv_replies":  ret.DtCgnv6LsnAlgFtpStats.Stats.PasvReplies,
			"epsv_replies":  ret.DtCgnv6LsnAlgFtpStats.Stats.EpsvReplies,
			"lpsv_replies":  ret.DtCgnv6LsnAlgFtpStats.Stats.LpsvReplies,
		},
	}
}

func getObjectCgnv6LsnAlgFtpStatsStats(d []interface{}) edpt.Cgnv6LsnAlgFtpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgFtpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortRequests = in["port_requests"].(int)
		ret.EprtRequests = in["eprt_requests"].(int)
		ret.LprtRequests = in["lprt_requests"].(int)
		ret.PasvReplies = in["pasv_replies"].(int)
		ret.EpsvReplies = in["epsv_replies"].(int)
		ret.LpsvReplies = in["lpsv_replies"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgFtpStats(d *schema.ResourceData) edpt.Cgnv6LsnAlgFtpStats {
	var ret edpt.Cgnv6LsnAlgFtpStats

	ret.Stats = getObjectCgnv6LsnAlgFtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
