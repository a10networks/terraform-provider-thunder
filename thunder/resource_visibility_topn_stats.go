package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_topn_stats`: Statistics for the object topn\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityTopnStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"heap_alloc_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total heap node allocated",
						},
						"heap_alloc_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total heap node alloc failed",
						},
						"heap_alloc_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Total heap node alloc failed Out of Memory",
						},
						"obj_reg_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total object node allocated",
						},
						"obj_reg_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total object node alloc failed",
						},
						"obj_reg_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Total object node alloc failed Out of Memory",
						},
						"heap_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Total Heap node deleted",
						},
						"obj_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Total Object node deleted",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityTopnStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityTopnStatsStats := setObjectVisibilityTopnStatsStats(res)
		d.Set("stats", VisibilityTopnStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityTopnStatsStats(ret edpt.DataVisibilityTopnStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"heap_alloc_success": ret.DtVisibilityTopnStats.Stats.HeapAllocSuccess,
			"heap_alloc_failed":  ret.DtVisibilityTopnStats.Stats.HeapAllocFailed,
			"heap_alloc_oom":     ret.DtVisibilityTopnStats.Stats.HeapAllocOom,
			"obj_reg_success":    ret.DtVisibilityTopnStats.Stats.ObjRegSuccess,
			"obj_reg_failed":     ret.DtVisibilityTopnStats.Stats.ObjRegFailed,
			"obj_reg_oom":        ret.DtVisibilityTopnStats.Stats.ObjRegOom,
			"heap_deleted":       ret.DtVisibilityTopnStats.Stats.HeapDeleted,
			"obj_deleted":        ret.DtVisibilityTopnStats.Stats.ObjDeleted,
		},
	}
}

func getObjectVisibilityTopnStatsStats(d []interface{}) edpt.VisibilityTopnStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityTopnStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HeapAllocSuccess = in["heap_alloc_success"].(int)
		ret.HeapAllocFailed = in["heap_alloc_failed"].(int)
		ret.HeapAllocOom = in["heap_alloc_oom"].(int)
		ret.ObjRegSuccess = in["obj_reg_success"].(int)
		ret.ObjRegFailed = in["obj_reg_failed"].(int)
		ret.ObjRegOom = in["obj_reg_oom"].(int)
		ret.HeapDeleted = in["heap_deleted"].(int)
		ret.ObjDeleted = in["obj_deleted"].(int)
	}
	return ret
}

func dataToEndpointVisibilityTopnStats(d *schema.ResourceData) edpt.VisibilityTopnStats {
	var ret edpt.VisibilityTopnStats

	ret.Stats = getObjectVisibilityTopnStatsStats(d.Get("stats").([]interface{}))
	return ret
}
