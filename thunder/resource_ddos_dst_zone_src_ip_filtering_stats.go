package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcIpFilteringStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_src_ip_filtering_stats`: Statistics for the object src-ip-filtering\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneSrcIpFilteringStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_1_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 1",
						},
						"class_list_2_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 2",
						},
						"class_list_3_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 3",
						},
						"class_list_4_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 4",
						},
						"class_list_5_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 5",
						},
						"class_list_6_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 6",
						},
						"class_list_7_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 7",
						},
						"class_list_8_match": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Match Class-List 8",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneSrcIpFilteringStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcIpFilteringStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcIpFilteringStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneSrcIpFilteringStatsStats := setObjectDdosDstZoneSrcIpFilteringStatsStats(res)
		d.Set("stats", DdosDstZoneSrcIpFilteringStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneSrcIpFilteringStatsStats(ret edpt.DataDdosDstZoneSrcIpFilteringStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"class_list_1_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList1Match,
			"class_list_2_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList2Match,
			"class_list_3_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList3Match,
			"class_list_4_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList4Match,
			"class_list_5_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList5Match,
			"class_list_6_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList6Match,
			"class_list_7_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList7Match,
			"class_list_8_match": ret.DtDdosDstZoneSrcIpFilteringStats.Stats.ClassList8Match,
		},
	}
}

func getObjectDdosDstZoneSrcIpFilteringStatsStats(d []interface{}) edpt.DdosDstZoneSrcIpFilteringStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcIpFilteringStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassList1Match = in["class_list_1_match"].(int)
		ret.ClassList2Match = in["class_list_2_match"].(int)
		ret.ClassList3Match = in["class_list_3_match"].(int)
		ret.ClassList4Match = in["class_list_4_match"].(int)
		ret.ClassList5Match = in["class_list_5_match"].(int)
		ret.ClassList6Match = in["class_list_6_match"].(int)
		ret.ClassList7Match = in["class_list_7_match"].(int)
		ret.ClassList8Match = in["class_list_8_match"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcIpFilteringStats(d *schema.ResourceData) edpt.DdosDstZoneSrcIpFilteringStats {
	var ret edpt.DdosDstZoneSrcIpFilteringStats

	ret.Stats = getObjectDdosDstZoneSrcIpFilteringStatsStats(d.Get("stats").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
