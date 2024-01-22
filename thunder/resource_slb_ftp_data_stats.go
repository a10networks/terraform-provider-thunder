package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFtpDataStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ftp_data_stats`: Statistics for the object ftp-data\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFtpDataStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sessions_num": {
							Type: schema.TypeInt, Optional: true, Description: "Total Data Sessions",
						},
						"port_out_of_range": {
							Type: schema.TypeInt, Optional: true, Description: "Drop Data Port out of range",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFtpDataStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpDataStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpDataStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFtpDataStatsStats := setObjectSlbFtpDataStatsStats(res)
		d.Set("stats", SlbFtpDataStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFtpDataStatsStats(ret edpt.DataSlbFtpDataStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sessions_num":      ret.DtSlbFtpDataStats.Stats.Sessions_num,
			"port_out_of_range": ret.DtSlbFtpDataStats.Stats.Port_out_of_range,
		},
	}
}

func getObjectSlbFtpDataStatsStats(d []interface{}) edpt.SlbFtpDataStatsStats {

	count1 := len(d)
	var ret edpt.SlbFtpDataStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sessions_num = in["sessions_num"].(int)
		ret.Port_out_of_range = in["port_out_of_range"].(int)
	}
	return ret
}

func dataToEndpointSlbFtpDataStats(d *schema.ResourceData) edpt.SlbFtpDataStats {
	var ret edpt.SlbFtpDataStats

	ret.Stats = getObjectSlbFtpDataStatsStats(d.Get("stats").([]interface{}))
	return ret
}
