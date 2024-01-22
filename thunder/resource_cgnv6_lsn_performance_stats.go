package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPerformanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_performance_stats`: Statistics for the object performance\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnPerformanceStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_sessions_current_epoch": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fullcone_created_current_epoch": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"user_quote_created_current_epoch": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"data_sessions_previous_epoch_first": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fullcone_created_previous_epoch_first": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"user_quote_created_previous_epoch_first": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"data_sessions_previous_epoch_last": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fullcone_created_previous_epoch_last": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"user_quote_created_previous_epoch_last": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnPerformanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerformanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerformanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnPerformanceStatsStats := setObjectCgnv6LsnPerformanceStatsStats(res)
		d.Set("stats", Cgnv6LsnPerformanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnPerformanceStatsStats(ret edpt.DataCgnv6LsnPerformanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"data_sessions_current_epoch":             ret.DtCgnv6LsnPerformanceStats.Stats.DataSessionsCurrentEpoch,
			"fullcone_created_current_epoch":          ret.DtCgnv6LsnPerformanceStats.Stats.FullconeCreatedCurrentEpoch,
			"user_quote_created_current_epoch":        ret.DtCgnv6LsnPerformanceStats.Stats.UserQuoteCreatedCurrentEpoch,
			"data_sessions_previous_epoch_first":      ret.DtCgnv6LsnPerformanceStats.Stats.DataSessionsPreviousEpochFirst,
			"fullcone_created_previous_epoch_first":   ret.DtCgnv6LsnPerformanceStats.Stats.FullconeCreatedPreviousEpochFirst,
			"user_quote_created_previous_epoch_first": ret.DtCgnv6LsnPerformanceStats.Stats.UserQuoteCreatedPreviousEpochFirst,
			"data_sessions_previous_epoch_last":       ret.DtCgnv6LsnPerformanceStats.Stats.DataSessionsPreviousEpochLast,
			"fullcone_created_previous_epoch_last":    ret.DtCgnv6LsnPerformanceStats.Stats.FullconeCreatedPreviousEpochLast,
			"user_quote_created_previous_epoch_last":  ret.DtCgnv6LsnPerformanceStats.Stats.UserQuoteCreatedPreviousEpochLast,
		},
	}
}

func getObjectCgnv6LsnPerformanceStatsStats(d []interface{}) edpt.Cgnv6LsnPerformanceStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnPerformanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DataSessionsCurrentEpoch = in["data_sessions_current_epoch"].(int)
		ret.FullconeCreatedCurrentEpoch = in["fullcone_created_current_epoch"].(int)
		ret.UserQuoteCreatedCurrentEpoch = in["user_quote_created_current_epoch"].(int)
		ret.DataSessionsPreviousEpochFirst = in["data_sessions_previous_epoch_first"].(int)
		ret.FullconeCreatedPreviousEpochFirst = in["fullcone_created_previous_epoch_first"].(int)
		ret.UserQuoteCreatedPreviousEpochFirst = in["user_quote_created_previous_epoch_first"].(int)
		ret.DataSessionsPreviousEpochLast = in["data_sessions_previous_epoch_last"].(int)
		ret.FullconeCreatedPreviousEpochLast = in["fullcone_created_previous_epoch_last"].(int)
		ret.UserQuoteCreatedPreviousEpochLast = in["user_quote_created_previous_epoch_last"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnPerformanceStats(d *schema.ResourceData) edpt.Cgnv6LsnPerformanceStats {
	var ret edpt.Cgnv6LsnPerformanceStats

	ret.Stats = getObjectCgnv6LsnPerformanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
