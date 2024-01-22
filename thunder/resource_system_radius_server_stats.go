package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemRadiusServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_radius_server_stats`: Statistics for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemRadiusServerStatsRead,

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
					},
				},
			},
		},
	}
}

func resourceSystemRadiusServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemRadiusServerStatsStats := setObjectSystemRadiusServerStatsStats(res)
		d.Set("stats", SystemRadiusServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemRadiusServerStatsStats(ret edpt.DataSystemRadiusServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"msisdn_received":               ret.DtSystemRadiusServerStats.Stats.MsisdnReceived,
			"imei_received":                 ret.DtSystemRadiusServerStats.Stats.ImeiReceived,
			"imsi_received":                 ret.DtSystemRadiusServerStats.Stats.ImsiReceived,
			"custom_received":               ret.DtSystemRadiusServerStats.Stats.CustomReceived,
			"radius_request_received":       ret.DtSystemRadiusServerStats.Stats.RadiusRequestReceived,
			"radius_request_dropped":        ret.DtSystemRadiusServerStats.Stats.RadiusRequestDropped,
			"request_bad_secret_dropped":    ret.DtSystemRadiusServerStats.Stats.RequestBadSecretDropped,
			"request_no_key_vap_dropped":    ret.DtSystemRadiusServerStats.Stats.RequestNoKeyVapDropped,
			"request_malformed_dropped":     ret.DtSystemRadiusServerStats.Stats.RequestMalformedDropped,
			"request_ignored":               ret.DtSystemRadiusServerStats.Stats.RequestIgnored,
			"radius_table_full":             ret.DtSystemRadiusServerStats.Stats.RadiusTableFull,
			"secret_not_configured_dropped": ret.DtSystemRadiusServerStats.Stats.SecretNotConfiguredDropped,
			"ha_standby_dropped":            ret.DtSystemRadiusServerStats.Stats.HaStandbyDropped,
			"ipv6_prefix_length_mismatch":   ret.DtSystemRadiusServerStats.Stats.Ipv6PrefixLengthMismatch,
			"invalid_key":                   ret.DtSystemRadiusServerStats.Stats.InvalidKey,
		},
	}
}

func getObjectSystemRadiusServerStatsStats(d []interface{}) edpt.SystemRadiusServerStatsStats {

	count1 := len(d)
	var ret edpt.SystemRadiusServerStatsStats
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
	}
	return ret
}

func dataToEndpointSystemRadiusServerStats(d *schema.ResourceData) edpt.SystemRadiusServerStats {
	var ret edpt.SystemRadiusServerStats

	ret.Stats = getObjectSystemRadiusServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
