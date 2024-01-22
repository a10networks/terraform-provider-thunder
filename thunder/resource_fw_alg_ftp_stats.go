package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgFtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_alg_ftp_stats`: Statistics for the object ftp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAlgFtpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_port_request": {
							Type: schema.TypeInt, Optional: true, Description: "PORT Requests From Client",
						},
						"client_eprt_request": {
							Type: schema.TypeInt, Optional: true, Description: "EPRT Requests From Client",
						},
						"server_pasv_reply": {
							Type: schema.TypeInt, Optional: true, Description: "PASV Replies From Server",
						},
						"server_epsv_reply": {
							Type: schema.TypeInt, Optional: true, Description: "EPSV Replies From Server",
						},
					},
				},
			},
		},
	}
}

func resourceFwAlgFtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgFtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgFtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAlgFtpStatsStats := setObjectFwAlgFtpStatsStats(res)
		d.Set("stats", FwAlgFtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAlgFtpStatsStats(ret edpt.DataFwAlgFtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"client_port_request": ret.DtFwAlgFtpStats.Stats.ClientPortRequest,
			"client_eprt_request": ret.DtFwAlgFtpStats.Stats.ClientEprtRequest,
			"server_pasv_reply":   ret.DtFwAlgFtpStats.Stats.ServerPasvReply,
			"server_epsv_reply":   ret.DtFwAlgFtpStats.Stats.ServerEpsvReply,
		},
	}
}

func getObjectFwAlgFtpStatsStats(d []interface{}) edpt.FwAlgFtpStatsStats {

	count1 := len(d)
	var ret edpt.FwAlgFtpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientPortRequest = in["client_port_request"].(int)
		ret.ClientEprtRequest = in["client_eprt_request"].(int)
		ret.ServerPasvReply = in["server_pasv_reply"].(int)
		ret.ServerEpsvReply = in["server_epsv_reply"].(int)
	}
	return ret
}

func dataToEndpointFwAlgFtpStats(d *schema.ResourceData) edpt.FwAlgFtpStats {
	var ret edpt.FwAlgFtpStats

	ret.Stats = getObjectFwAlgFtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
