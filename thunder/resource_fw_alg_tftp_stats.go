package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgTftpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_alg_tftp_stats`: Statistics for the object tftp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAlgTftpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TFTP Client Sessions Created",
						},
					},
				},
			},
		},
	}
}

func resourceFwAlgTftpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgTftpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgTftpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAlgTftpStatsStats := setObjectFwAlgTftpStatsStats(res)
		d.Set("stats", FwAlgTftpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAlgTftpStatsStats(ret edpt.DataFwAlgTftpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtFwAlgTftpStats.Stats.SessionCreated,
		},
	}
}

func getObjectFwAlgTftpStatsStats(d []interface{}) edpt.FwAlgTftpStatsStats {

	count1 := len(d)
	var ret edpt.FwAlgTftpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
	}
	return ret
}

func dataToEndpointFwAlgTftpStats(d *schema.ResourceData) edpt.FwAlgTftpStats {
	var ret edpt.FwAlgTftpStats

	ret.Stats = getObjectFwAlgTftpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
