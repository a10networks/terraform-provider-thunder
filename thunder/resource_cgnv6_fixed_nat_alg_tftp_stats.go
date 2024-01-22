package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgTftpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_tftp_stats`: Statistics for the object tftp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgTftpStatsRead,

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

func resourceCgnv6FixedNatAlgTftpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgTftpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgTftpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgTftpStatsStats := setObjectCgnv6FixedNatAlgTftpStatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgTftpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgTftpStatsStats(ret edpt.DataCgnv6FixedNatAlgTftpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtCgnv6FixedNatAlgTftpStats.Stats.SessionCreated,
		},
	}
}

func getObjectCgnv6FixedNatAlgTftpStatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgTftpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgTftpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgTftpStats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgTftpStats {
	var ret edpt.Cgnv6FixedNatAlgTftpStats

	ret.Stats = getObjectCgnv6FixedNatAlgTftpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
