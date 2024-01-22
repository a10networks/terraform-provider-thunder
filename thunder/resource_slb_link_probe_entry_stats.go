package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbLinkProbeEntryStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_link_probe_entry_stats`: Statistics for the object entry\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbLinkProbeEntryStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Count",
						},
						"total_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Created",
						},
						"total_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Inserted",
						},
						"total_ready_to_free": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Ready To Free",
						},
						"total_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Freed",
						},
						"err_entry_create_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Creation Failure",
						},
						"err_entry_create_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Creation Out of Memory",
						},
						"err_entry_insert_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Insert Failed",
						},
						"err_tmpl_probe_create_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Probe Template Creation Failure",
						},
						"err_tmpl_probe_create_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Probe Template Creation Out of Memory",
						},
						"total_http_probes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP Probes Sent out",
						},
						"total_http_response_received": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP responses received",
						},
						"total_http_response_good": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP responses matching probe template config",
						},
						"total_http_response_bad": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP responses not matching probe template config",
						},
						"total_tcp_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP errors in probes sent out",
						},
						"err_smart_nat_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Error creating Smart NAT Instance",
						},
						"err_smart_nat_port_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Error obtaining Smart NAT source port",
						},
						"err_l4_sess_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Error allocating L4 session for probe",
						},
						"err_probe_tcp_conn_send": {
							Type: schema.TypeInt, Optional: true, Description: "Error in initiating TCP connection for probe",
						},
						"probe_tcp_conn_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP connection sent for probe",
						},
					},
				},
			},
		},
	}
}

func resourceSlbLinkProbeEntryStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbLinkProbeEntryStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbLinkProbeEntryStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbLinkProbeEntryStatsStats := setObjectSlbLinkProbeEntryStatsStats(res)
		d.Set("stats", SlbLinkProbeEntryStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbLinkProbeEntryStatsStats(ret edpt.DataSlbLinkProbeEntryStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_entries":                 ret.DtSlbLinkProbeEntryStats.Stats.Curr_entries,
			"total_created":                ret.DtSlbLinkProbeEntryStats.Stats.Total_created,
			"total_inserted":               ret.DtSlbLinkProbeEntryStats.Stats.Total_inserted,
			"total_ready_to_free":          ret.DtSlbLinkProbeEntryStats.Stats.Total_ready_to_free,
			"total_freed":                  ret.DtSlbLinkProbeEntryStats.Stats.Total_freed,
			"err_entry_create_failed":      ret.DtSlbLinkProbeEntryStats.Stats.Err_entry_create_failed,
			"err_entry_create_oom":         ret.DtSlbLinkProbeEntryStats.Stats.Err_entry_create_oom,
			"err_entry_insert_failed":      ret.DtSlbLinkProbeEntryStats.Stats.Err_entry_insert_failed,
			"err_tmpl_probe_create_failed": ret.DtSlbLinkProbeEntryStats.Stats.Err_tmpl_probe_create_failed,
			"err_tmpl_probe_create_oom":    ret.DtSlbLinkProbeEntryStats.Stats.Err_tmpl_probe_create_oom,
			"total_http_probes_sent":       ret.DtSlbLinkProbeEntryStats.Stats.Total_http_probes_sent,
			"total_http_response_received": ret.DtSlbLinkProbeEntryStats.Stats.Total_http_response_received,
			"total_http_response_good":     ret.DtSlbLinkProbeEntryStats.Stats.Total_http_response_good,
			"total_http_response_bad":      ret.DtSlbLinkProbeEntryStats.Stats.Total_http_response_bad,
			"total_tcp_err":                ret.DtSlbLinkProbeEntryStats.Stats.Total_tcp_err,
			"err_smart_nat_alloc":          ret.DtSlbLinkProbeEntryStats.Stats.Err_smart_nat_alloc,
			"err_smart_nat_port_alloc":     ret.DtSlbLinkProbeEntryStats.Stats.Err_smart_nat_port_alloc,
			"err_l4_sess_alloc":            ret.DtSlbLinkProbeEntryStats.Stats.Err_l4_sess_alloc,
			"err_probe_tcp_conn_send":      ret.DtSlbLinkProbeEntryStats.Stats.Err_probe_tcp_conn_send,
			"probe_tcp_conn_sent":          ret.DtSlbLinkProbeEntryStats.Stats.Probe_tcp_conn_sent,
		},
	}
}

func getObjectSlbLinkProbeEntryStatsStats(d []interface{}) edpt.SlbLinkProbeEntryStatsStats {

	count1 := len(d)
	var ret edpt.SlbLinkProbeEntryStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_entries = in["curr_entries"].(int)
		ret.Total_created = in["total_created"].(int)
		ret.Total_inserted = in["total_inserted"].(int)
		ret.Total_ready_to_free = in["total_ready_to_free"].(int)
		ret.Total_freed = in["total_freed"].(int)
		ret.Err_entry_create_failed = in["err_entry_create_failed"].(int)
		ret.Err_entry_create_oom = in["err_entry_create_oom"].(int)
		ret.Err_entry_insert_failed = in["err_entry_insert_failed"].(int)
		ret.Err_tmpl_probe_create_failed = in["err_tmpl_probe_create_failed"].(int)
		ret.Err_tmpl_probe_create_oom = in["err_tmpl_probe_create_oom"].(int)
		ret.Total_http_probes_sent = in["total_http_probes_sent"].(int)
		ret.Total_http_response_received = in["total_http_response_received"].(int)
		ret.Total_http_response_good = in["total_http_response_good"].(int)
		ret.Total_http_response_bad = in["total_http_response_bad"].(int)
		ret.Total_tcp_err = in["total_tcp_err"].(int)
		ret.Err_smart_nat_alloc = in["err_smart_nat_alloc"].(int)
		ret.Err_smart_nat_port_alloc = in["err_smart_nat_port_alloc"].(int)
		ret.Err_l4_sess_alloc = in["err_l4_sess_alloc"].(int)
		ret.Err_probe_tcp_conn_send = in["err_probe_tcp_conn_send"].(int)
		ret.Probe_tcp_conn_sent = in["probe_tcp_conn_sent"].(int)
	}
	return ret
}

func dataToEndpointSlbLinkProbeEntryStats(d *schema.ResourceData) edpt.SlbLinkProbeEntryStats {
	var ret edpt.SlbLinkProbeEntryStats

	ret.Stats = getObjectSlbLinkProbeEntryStatsStats(d.Get("stats").([]interface{}))
	return ret
}
