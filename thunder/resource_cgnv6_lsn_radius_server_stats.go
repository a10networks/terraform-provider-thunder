package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRadiusServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_radius_server_stats`: Statistics for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRadiusServerStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msisdn_received": {
							Type: schema.TypeInt, Optional: true, Description: "MSISDN Received",
						},
						"imei_received": {
							Type: schema.TypeInt, Optional: true, Description: "IMEI Received",
						},
						"imsi_received": {
							Type: schema.TypeInt, Optional: true, Description: "IMSI Received",
						},
						"custom_received": {
							Type: schema.TypeInt, Optional: true, Description: "Custom attribute Received",
						},
						"radius_request_received": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request Received",
						},
						"radius_request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request Dropped (Malformed Packet)",
						},
						"request_bad_secret_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request Bad Secret Dropped",
						},
						"request_no_key_vap_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request No Key Attribute Dropped",
						},
						"request_malformed_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request Malformed Dropped",
						},
						"request_ignored": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request Ignored",
						},
						"radius_table_full": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Request Dropped (Table Full)",
						},
						"secret_not_configured_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS Secret Not Configured Dropped",
						},
						"ha_standby_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "HA Standby Dropped",
						},
						"ipv6_prefix_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Framed IPV6 Prefix Length Mismatch",
						},
						"invalid_key": {
							Type: schema.TypeInt, Optional: true, Description: "Radius Request has Invalid Key Field",
						},
						"smp_created": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS SMP Created",
						},
						"smp_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS SMP Deleted",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnRadiusServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRadiusServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRadiusServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRadiusServerStatsStats := setObjectCgnv6LsnRadiusServerStatsStats(res)
		d.Set("stats", Cgnv6LsnRadiusServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRadiusServerStatsStats(ret edpt.DataCgnv6LsnRadiusServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"msisdn_received":               ret.DtCgnv6LsnRadiusServerStats.Stats.MsisdnReceived,
			"imei_received":                 ret.DtCgnv6LsnRadiusServerStats.Stats.ImeiReceived,
			"imsi_received":                 ret.DtCgnv6LsnRadiusServerStats.Stats.ImsiReceived,
			"custom_received":               ret.DtCgnv6LsnRadiusServerStats.Stats.CustomReceived,
			"radius_request_received":       ret.DtCgnv6LsnRadiusServerStats.Stats.RadiusRequestReceived,
			"radius_request_dropped":        ret.DtCgnv6LsnRadiusServerStats.Stats.RadiusRequestDropped,
			"request_bad_secret_dropped":    ret.DtCgnv6LsnRadiusServerStats.Stats.RequestBadSecretDropped,
			"request_no_key_vap_dropped":    ret.DtCgnv6LsnRadiusServerStats.Stats.RequestNoKeyVapDropped,
			"request_malformed_dropped":     ret.DtCgnv6LsnRadiusServerStats.Stats.RequestMalformedDropped,
			"request_ignored":               ret.DtCgnv6LsnRadiusServerStats.Stats.RequestIgnored,
			"radius_table_full":             ret.DtCgnv6LsnRadiusServerStats.Stats.RadiusTableFull,
			"secret_not_configured_dropped": ret.DtCgnv6LsnRadiusServerStats.Stats.SecretNotConfiguredDropped,
			"ha_standby_dropped":            ret.DtCgnv6LsnRadiusServerStats.Stats.HaStandbyDropped,
			"ipv6_prefix_length_mismatch":   ret.DtCgnv6LsnRadiusServerStats.Stats.Ipv6PrefixLengthMismatch,
			"invalid_key":                   ret.DtCgnv6LsnRadiusServerStats.Stats.InvalidKey,
			"smp_created":                   ret.DtCgnv6LsnRadiusServerStats.Stats.SmpCreated,
			"smp_deleted":                   ret.DtCgnv6LsnRadiusServerStats.Stats.SmpDeleted,
		},
	}
}

func getObjectCgnv6LsnRadiusServerStatsStats(d []interface{}) edpt.Cgnv6LsnRadiusServerStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRadiusServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MsisdnReceived = in["msisdn_received"].(int)
		ret.ImeiReceived = in["imei_received"].(int)
		ret.ImsiReceived = in["imsi_received"].(int)
		ret.CustomReceived = in["custom_received"].(int)
		ret.RadiusRequestReceived = in["radius_request_received"].(int)
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RequestIgnored = in["request_ignored"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.SecretNotConfiguredDropped = in["secret_not_configured_dropped"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.Ipv6PrefixLengthMismatch = in["ipv6_prefix_length_mismatch"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		ret.SmpCreated = in["smp_created"].(int)
		ret.SmpDeleted = in["smp_deleted"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnRadiusServerStats(d *schema.ResourceData) edpt.Cgnv6LsnRadiusServerStats {
	var ret edpt.Cgnv6LsnRadiusServerStats

	ret.Stats = getObjectCgnv6LsnRadiusServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
