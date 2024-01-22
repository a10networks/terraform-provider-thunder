package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCounterLibAccountingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_counter_lib_accounting_stats`: Statistics for the object counter-lib-accounting\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCounterLibAccountingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_ctr_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Total counters allocated",
						},
						"total_ctr_in_rml": {
							Type: schema.TypeInt, Optional: true, Description: "Total counters put in rml queue",
						},
						"total_ctr_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total counters actually freed",
						},
						"total_ctr_in_system": {
							Type: schema.TypeInt, Optional: true, Description: "Total counters currently allocated in system",
						},
						"total_oper_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Total oper blocks allocated",
						},
						"total_oper_free": {
							Type: schema.TypeInt, Optional: true, Description: "Total oper blocks freed",
						},
						"total_blocks_in_hash": {
							Type: schema.TypeInt, Optional: true, Description: "Total blocks in hash table",
						},
						"total_blocks_in_rml_hash": {
							Type: schema.TypeInt, Optional: true, Description: "Total blocks in rml in hash table",
						},
						"total_nodes_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Total nodes allocated(sflow/history/telemetry/cpu-compute)",
						},
						"total_nodes_in_rml": {
							Type: schema.TypeInt, Optional: true, Description: "Total nodes put in rml queue",
						},
						"total_nodes_free": {
							Type: schema.TypeInt, Optional: true, Description: "Total nodes freed(sflow/history/telemetry/cpu-compute)",
						},
						"total_nodes_free_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total nodes free failed",
						},
						"total_nodes_unlink_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total nodes unlink failed",
						},
						"total_nodes_in_system": {
							Type: schema.TypeInt, Optional: true, Description: "Total nodes currently allocated in system",
						},
						"total_memblocks_alloc_avro": {
							Type: schema.TypeInt, Optional: true, Description: "Total memory blocks allocated by avro lib",
						},
						"total_memblocks_free_avro": {
							Type: schema.TypeInt, Optional: true, Description: "Total memory blocks freed by avro lib",
						},
						"total_memblocks_realloc_avro": {
							Type: schema.TypeInt, Optional: true, Description: "Total memory blocks realloc by avro lib",
						},
					},
				},
			},
		},
	}
}

func resourceSystemCounterLibAccountingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCounterLibAccountingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCounterLibAccountingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCounterLibAccountingStatsStats := setObjectSystemCounterLibAccountingStatsStats(res)
		d.Set("stats", SystemCounterLibAccountingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCounterLibAccountingStatsStats(ret edpt.DataSystemCounterLibAccountingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_ctr_alloc":              ret.DtSystemCounterLibAccountingStats.Stats.TotalCtrAlloc,
			"total_ctr_in_rml":             ret.DtSystemCounterLibAccountingStats.Stats.TotalCtrInRml,
			"total_ctr_freed":              ret.DtSystemCounterLibAccountingStats.Stats.TotalCtrFreed,
			"total_ctr_in_system":          ret.DtSystemCounterLibAccountingStats.Stats.TotalCtrInSystem,
			"total_oper_alloc":             ret.DtSystemCounterLibAccountingStats.Stats.TotalOperAlloc,
			"total_oper_free":              ret.DtSystemCounterLibAccountingStats.Stats.TotalOperFree,
			"total_blocks_in_hash":         ret.DtSystemCounterLibAccountingStats.Stats.TotalBlocksInHash,
			"total_blocks_in_rml_hash":     ret.DtSystemCounterLibAccountingStats.Stats.TotalBlocksInRmlHash,
			"total_nodes_alloc":            ret.DtSystemCounterLibAccountingStats.Stats.TotalNodesAlloc,
			"total_nodes_in_rml":           ret.DtSystemCounterLibAccountingStats.Stats.TotalNodesInRml,
			"total_nodes_free":             ret.DtSystemCounterLibAccountingStats.Stats.TotalNodesFree,
			"total_nodes_free_failed":      ret.DtSystemCounterLibAccountingStats.Stats.TotalNodesFreeFailed,
			"total_nodes_unlink_failed":    ret.DtSystemCounterLibAccountingStats.Stats.TotalNodesUnlinkFailed,
			"total_nodes_in_system":        ret.DtSystemCounterLibAccountingStats.Stats.TotalNodesInSystem,
			"total_memblocks_alloc_avro":   ret.DtSystemCounterLibAccountingStats.Stats.TotalMemblocksAllocAvro,
			"total_memblocks_free_avro":    ret.DtSystemCounterLibAccountingStats.Stats.TotalMemblocksFreeAvro,
			"total_memblocks_realloc_avro": ret.DtSystemCounterLibAccountingStats.Stats.TotalMemblocksReallocAvro,
		},
	}
}

func getObjectSystemCounterLibAccountingStatsStats(d []interface{}) edpt.SystemCounterLibAccountingStatsStats {

	count1 := len(d)
	var ret edpt.SystemCounterLibAccountingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalCtrAlloc = in["total_ctr_alloc"].(int)
		ret.TotalCtrInRml = in["total_ctr_in_rml"].(int)
		ret.TotalCtrFreed = in["total_ctr_freed"].(int)
		ret.TotalCtrInSystem = in["total_ctr_in_system"].(int)
		ret.TotalOperAlloc = in["total_oper_alloc"].(int)
		ret.TotalOperFree = in["total_oper_free"].(int)
		ret.TotalBlocksInHash = in["total_blocks_in_hash"].(int)
		ret.TotalBlocksInRmlHash = in["total_blocks_in_rml_hash"].(int)
		ret.TotalNodesAlloc = in["total_nodes_alloc"].(int)
		ret.TotalNodesInRml = in["total_nodes_in_rml"].(int)
		ret.TotalNodesFree = in["total_nodes_free"].(int)
		ret.TotalNodesFreeFailed = in["total_nodes_free_failed"].(int)
		ret.TotalNodesUnlinkFailed = in["total_nodes_unlink_failed"].(int)
		ret.TotalNodesInSystem = in["total_nodes_in_system"].(int)
		ret.TotalMemblocksAllocAvro = in["total_memblocks_alloc_avro"].(int)
		ret.TotalMemblocksFreeAvro = in["total_memblocks_free_avro"].(int)
		ret.TotalMemblocksReallocAvro = in["total_memblocks_realloc_avro"].(int)
	}
	return ret
}

func dataToEndpointSystemCounterLibAccountingStats(d *schema.ResourceData) edpt.SystemCounterLibAccountingStats {
	var ret edpt.SystemCounterLibAccountingStats

	ret.Stats = getObjectSystemCounterLibAccountingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
