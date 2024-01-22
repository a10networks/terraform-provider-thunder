package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosPatternRecognitionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_pattern_recognition_stats`: Statistics for the object pattern-recognition\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosPatternRecognitionStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Engine Started",
						},
						"not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Pattern Not Found",
						},
						"generic_error": {
							Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Exceptions",
						},
						"filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Extracted Filter Match",
						},
						"filter_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Extracted Filter Drop",
						},
					},
				},
			},
		},
	}
}

func resourceDdosPatternRecognitionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPatternRecognitionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPatternRecognitionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosPatternRecognitionStatsStats := setObjectDdosPatternRecognitionStatsStats(res)
		d.Set("stats", DdosPatternRecognitionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosPatternRecognitionStatsStats(ret edpt.DataDdosPatternRecognitionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"proceeded":     ret.DtDdosPatternRecognitionStats.Stats.Proceeded,
			"not_found":     ret.DtDdosPatternRecognitionStats.Stats.Not_found,
			"generic_error": ret.DtDdosPatternRecognitionStats.Stats.Generic_error,
			"filter_match":  ret.DtDdosPatternRecognitionStats.Stats.Filter_match,
			"filter_drop":   ret.DtDdosPatternRecognitionStats.Stats.Filter_drop,
		},
	}
}

func getObjectDdosPatternRecognitionStatsStats(d []interface{}) edpt.DdosPatternRecognitionStatsStats {

	count1 := len(d)
	var ret edpt.DdosPatternRecognitionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Proceeded = in["proceeded"].(int)
		ret.Not_found = in["not_found"].(int)
		ret.Generic_error = in["generic_error"].(int)
		ret.Filter_match = in["filter_match"].(int)
		ret.Filter_drop = in["filter_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosPatternRecognitionStats(d *schema.ResourceData) edpt.DdosPatternRecognitionStats {
	var ret edpt.DdosPatternRecognitionStats

	ret.Stats = getObjectDdosPatternRecognitionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
