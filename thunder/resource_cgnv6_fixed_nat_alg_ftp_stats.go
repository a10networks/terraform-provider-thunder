package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgFtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_ftp_stats`: Statistics for the object ftp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgFtpStatsRead,

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

func resourceCgnv6FixedNatAlgFtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgFtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgFtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgFtpStatsStats := setObjectCgnv6FixedNatAlgFtpStatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgFtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgFtpStatsStats(ret edpt.DataCgnv6FixedNatAlgFtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_requests": ret.DtCgnv6FixedNatAlgFtpStats.Stats.PortRequests,
			"eprt_requests": ret.DtCgnv6FixedNatAlgFtpStats.Stats.EprtRequests,
			"lprt_requests": ret.DtCgnv6FixedNatAlgFtpStats.Stats.LprtRequests,
			"pasv_replies":  ret.DtCgnv6FixedNatAlgFtpStats.Stats.PasvReplies,
			"epsv_replies":  ret.DtCgnv6FixedNatAlgFtpStats.Stats.EpsvReplies,
			"lpsv_replies":  ret.DtCgnv6FixedNatAlgFtpStats.Stats.LpsvReplies,
		},
	}
}

func getObjectCgnv6FixedNatAlgFtpStatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgFtpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgFtpStatsStats
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

func dataToEndpointCgnv6FixedNatAlgFtpStats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgFtpStats {
	var ret edpt.Cgnv6FixedNatAlgFtpStats

	ret.Stats = getObjectCgnv6FixedNatAlgFtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
