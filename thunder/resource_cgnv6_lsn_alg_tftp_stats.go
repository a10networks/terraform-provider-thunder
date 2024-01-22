package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgTftpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_tftp_stats`: Statistics for the object tftp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgTftpStatsRead,

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

func resourceCgnv6LsnAlgTftpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgTftpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgTftpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgTftpStatsStats := setObjectCgnv6LsnAlgTftpStatsStats(res)
		d.Set("stats", Cgnv6LsnAlgTftpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgTftpStatsStats(ret edpt.DataCgnv6LsnAlgTftpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtCgnv6LsnAlgTftpStats.Stats.SessionCreated,
		},
	}
}

func getObjectCgnv6LsnAlgTftpStatsStats(d []interface{}) edpt.Cgnv6LsnAlgTftpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgTftpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgTftpStats(d *schema.ResourceData) edpt.Cgnv6LsnAlgTftpStats {
	var ret edpt.Cgnv6LsnAlgTftpStats

	ret.Stats = getObjectCgnv6LsnAlgTftpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
