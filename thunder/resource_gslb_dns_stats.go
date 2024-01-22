package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbDnsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_dns_stats`: Statistics for the object dns\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbDnsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_query": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of DNS queries received",
						},
						"total_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of DNS replies sent to clients",
						},
						"bad_packet_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries with incorrect data length",
						},
						"bad_packet_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with incorrect data length",
						},
						"bad_header_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries with incorrect header",
						},
						"bad_header_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with incorrect header",
						},
						"bad_format_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries with incorrect format",
						},
						"bad_format_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with incorrect format",
						},
						"bad_service_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries with unknown service",
						},
						"bad_service_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with unknown service",
						},
						"bad_class_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries with incorrect class",
						},
						"bad_class_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with incorrect class",
						},
						"bad_type_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of queries with incorrect type",
						},
						"bad_type_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with incorrect type",
						},
						"no_answer": {
							Type: schema.TypeInt, Optional: true, Description: "Number of replies with unknown server IP",
						},
						"metric_health_check": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Health Check Hit",
						},
						"metric_weighted_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Weighted IP Hit",
						},
						"metric_weighted_site": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Weighted Site Hit",
						},
						"metric_capacity": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Capacity Hit",
						},
						"metric_active_server": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Active Server Hit",
						},
						"metric_easy_rdt": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Easy RDT Hit",
						},
						"metric_active_rdt": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Active RDT Hit",
						},
						"metric_geographic": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Geographic Hit",
						},
						"metric_connection_load": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Connection Load Hit",
						},
						"metric_number_of_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Number of Sessions Hit",
						},
						"metric_active_weight": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Active Weight Hit",
						},
						"metric_admin_preference": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Admin Preference Hit",
						},
						"metric_bandwidth_quality": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Bandwidth Quality Hit",
						},
						"metric_bandwidth_cost": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Bandwidth Cost Hit",
						},
						"metric_user": {
							Type: schema.TypeInt, Optional: true, Description: "Metric User Hit",
						},
						"metric_least_reponse": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Least Reponse Hit",
						},
						"metric_admin_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Admin IP Hit",
						},
						"metric_round_robin": {
							Type: schema.TypeInt, Optional: true, Description: "Metric Round Robin Hit",
						},
					},
				},
			},
		},
	}
}

func resourceGslbDnsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbDnsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbDnsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbDnsStatsStats := setObjectGslbDnsStatsStats(res)
		d.Set("stats", GslbDnsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbDnsStatsStats(ret edpt.DataGslbDnsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_query":               ret.DtGslbDnsStats.Stats.TotalQuery,
			"total_response":            ret.DtGslbDnsStats.Stats.TotalResponse,
			"bad_packet_query":          ret.DtGslbDnsStats.Stats.BadPacketQuery,
			"bad_packet_response":       ret.DtGslbDnsStats.Stats.BadPacketResponse,
			"bad_header_query":          ret.DtGslbDnsStats.Stats.BadHeaderQuery,
			"bad_header_response":       ret.DtGslbDnsStats.Stats.BadHeaderResponse,
			"bad_format_query":          ret.DtGslbDnsStats.Stats.BadFormatQuery,
			"bad_format_response":       ret.DtGslbDnsStats.Stats.BadFormatResponse,
			"bad_service_query":         ret.DtGslbDnsStats.Stats.BadServiceQuery,
			"bad_service_response":      ret.DtGslbDnsStats.Stats.BadServiceResponse,
			"bad_class_query":           ret.DtGslbDnsStats.Stats.BadClassQuery,
			"bad_class_response":        ret.DtGslbDnsStats.Stats.BadClassResponse,
			"bad_type_query":            ret.DtGslbDnsStats.Stats.BadTypeQuery,
			"bad_type_response":         ret.DtGslbDnsStats.Stats.BadTypeResponse,
			"no_answer":                 ret.DtGslbDnsStats.Stats.No_answer,
			"metric_health_check":       ret.DtGslbDnsStats.Stats.Metric_health_check,
			"metric_weighted_ip":        ret.DtGslbDnsStats.Stats.Metric_weighted_ip,
			"metric_weighted_site":      ret.DtGslbDnsStats.Stats.Metric_weighted_site,
			"metric_capacity":           ret.DtGslbDnsStats.Stats.Metric_capacity,
			"metric_active_server":      ret.DtGslbDnsStats.Stats.Metric_active_server,
			"metric_easy_rdt":           ret.DtGslbDnsStats.Stats.Metric_easy_rdt,
			"metric_active_rdt":         ret.DtGslbDnsStats.Stats.Metric_active_rdt,
			"metric_geographic":         ret.DtGslbDnsStats.Stats.Metric_geographic,
			"metric_connection_load":    ret.DtGslbDnsStats.Stats.Metric_connection_load,
			"metric_number_of_sessions": ret.DtGslbDnsStats.Stats.Metric_number_of_sessions,
			"metric_active_weight":      ret.DtGslbDnsStats.Stats.Metric_active_weight,
			"metric_admin_preference":   ret.DtGslbDnsStats.Stats.Metric_admin_preference,
			"metric_bandwidth_quality":  ret.DtGslbDnsStats.Stats.Metric_bandwidth_quality,
			"metric_bandwidth_cost":     ret.DtGslbDnsStats.Stats.Metric_bandwidth_cost,
			"metric_user":               ret.DtGslbDnsStats.Stats.Metric_user,
			"metric_least_reponse":      ret.DtGslbDnsStats.Stats.Metric_least_reponse,
			"metric_admin_ip":           ret.DtGslbDnsStats.Stats.Metric_admin_ip,
			"metric_round_robin":        ret.DtGslbDnsStats.Stats.Metric_round_robin,
		},
	}
}

func getObjectGslbDnsStatsStats(d []interface{}) edpt.GslbDnsStatsStats {

	count1 := len(d)
	var ret edpt.GslbDnsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalQuery = in["total_query"].(int)
		ret.TotalResponse = in["total_response"].(int)
		ret.BadPacketQuery = in["bad_packet_query"].(int)
		ret.BadPacketResponse = in["bad_packet_response"].(int)
		ret.BadHeaderQuery = in["bad_header_query"].(int)
		ret.BadHeaderResponse = in["bad_header_response"].(int)
		ret.BadFormatQuery = in["bad_format_query"].(int)
		ret.BadFormatResponse = in["bad_format_response"].(int)
		ret.BadServiceQuery = in["bad_service_query"].(int)
		ret.BadServiceResponse = in["bad_service_response"].(int)
		ret.BadClassQuery = in["bad_class_query"].(int)
		ret.BadClassResponse = in["bad_class_response"].(int)
		ret.BadTypeQuery = in["bad_type_query"].(int)
		ret.BadTypeResponse = in["bad_type_response"].(int)
		ret.No_answer = in["no_answer"].(int)
		ret.Metric_health_check = in["metric_health_check"].(int)
		ret.Metric_weighted_ip = in["metric_weighted_ip"].(int)
		ret.Metric_weighted_site = in["metric_weighted_site"].(int)
		ret.Metric_capacity = in["metric_capacity"].(int)
		ret.Metric_active_server = in["metric_active_server"].(int)
		ret.Metric_easy_rdt = in["metric_easy_rdt"].(int)
		ret.Metric_active_rdt = in["metric_active_rdt"].(int)
		ret.Metric_geographic = in["metric_geographic"].(int)
		ret.Metric_connection_load = in["metric_connection_load"].(int)
		ret.Metric_number_of_sessions = in["metric_number_of_sessions"].(int)
		ret.Metric_active_weight = in["metric_active_weight"].(int)
		ret.Metric_admin_preference = in["metric_admin_preference"].(int)
		ret.Metric_bandwidth_quality = in["metric_bandwidth_quality"].(int)
		ret.Metric_bandwidth_cost = in["metric_bandwidth_cost"].(int)
		ret.Metric_user = in["metric_user"].(int)
		ret.Metric_least_reponse = in["metric_least_reponse"].(int)
		ret.Metric_admin_ip = in["metric_admin_ip"].(int)
		ret.Metric_round_robin = in["metric_round_robin"].(int)
	}
	return ret
}

func dataToEndpointGslbDnsStats(d *schema.ResourceData) edpt.GslbDnsStats {
	var ret edpt.GslbDnsStats

	ret.Stats = getObjectGslbDnsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
