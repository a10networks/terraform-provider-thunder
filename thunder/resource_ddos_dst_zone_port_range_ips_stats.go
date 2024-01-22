package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeIpsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_ips_stats`: Statistics for the object ips\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeIpsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ips_matched_total": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Total",
						},
						"ips_matched_action_pass": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Action Pass",
						},
						"ips_matched_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Action Drop",
						},
						"ips_matched_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Action Blacklist",
						},
						"ips_matched_severity_high": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Severity High",
						},
						"ips_matched_severity_medium": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Severity Medium",
						},
						"ips_matched_severity_low": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Severity Low",
						},
						"src_ips_matched_action_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src IPS Matched Action Pass",
						},
						"src_ips_matched_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src IPS Matched Action Drop",
						},
						"src_ips_matched_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src IPS Matched Action Blacklist",
						},
						"src_ips_matched_severity_high": {
							Type: schema.TypeInt, Optional: true, Description: "Src IPS Matched Severity High",
						},
						"src_ips_matched_severity_medium": {
							Type: schema.TypeInt, Optional: true, Description: "Src IPS Matched Severity Medium",
						},
						"src_ips_matched_severity_low": {
							Type: schema.TypeInt, Optional: true, Description: "Src IPS Matched Severity Low",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}

func resourceDdosDstZonePortRangeIpsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIpsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeIpsStatsStats := setObjectDdosDstZonePortRangeIpsStatsStats(res)
		d.Set("stats", DdosDstZonePortRangeIpsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeIpsStatsStats(ret edpt.DataDdosDstZonePortRangeIpsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ips_matched_total":                ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_total,
			"ips_matched_action_pass":          ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_action_pass,
			"ips_matched_action_drop":          ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_action_drop,
			"ips_matched_action_blacklist":     ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_action_blacklist,
			"ips_matched_severity_high":        ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_severity_high,
			"ips_matched_severity_medium":      ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_severity_medium,
			"ips_matched_severity_low":         ret.DtDdosDstZonePortRangeIpsStats.Stats.Ips_matched_severity_low,
			"src_ips_matched_action_pass":      ret.DtDdosDstZonePortRangeIpsStats.Stats.Src_ips_matched_action_pass,
			"src_ips_matched_action_drop":      ret.DtDdosDstZonePortRangeIpsStats.Stats.Src_ips_matched_action_drop,
			"src_ips_matched_action_blacklist": ret.DtDdosDstZonePortRangeIpsStats.Stats.Src_ips_matched_action_blacklist,
			"src_ips_matched_severity_high":    ret.DtDdosDstZonePortRangeIpsStats.Stats.Src_ips_matched_severity_high,
			"src_ips_matched_severity_medium":  ret.DtDdosDstZonePortRangeIpsStats.Stats.Src_ips_matched_severity_medium,
			"src_ips_matched_severity_low":     ret.DtDdosDstZonePortRangeIpsStats.Stats.Src_ips_matched_severity_low,
		},
	}
}

func getObjectDdosDstZonePortRangeIpsStatsStats(d []interface{}) edpt.DdosDstZonePortRangeIpsStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeIpsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ips_matched_total = in["ips_matched_total"].(int)
		ret.Ips_matched_action_pass = in["ips_matched_action_pass"].(int)
		ret.Ips_matched_action_drop = in["ips_matched_action_drop"].(int)
		ret.Ips_matched_action_blacklist = in["ips_matched_action_blacklist"].(int)
		ret.Ips_matched_severity_high = in["ips_matched_severity_high"].(int)
		ret.Ips_matched_severity_medium = in["ips_matched_severity_medium"].(int)
		ret.Ips_matched_severity_low = in["ips_matched_severity_low"].(int)
		ret.Src_ips_matched_action_pass = in["src_ips_matched_action_pass"].(int)
		ret.Src_ips_matched_action_drop = in["src_ips_matched_action_drop"].(int)
		ret.Src_ips_matched_action_blacklist = in["src_ips_matched_action_blacklist"].(int)
		ret.Src_ips_matched_severity_high = in["src_ips_matched_severity_high"].(int)
		ret.Src_ips_matched_severity_medium = in["src_ips_matched_severity_medium"].(int)
		ret.Src_ips_matched_severity_low = in["src_ips_matched_severity_low"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeIpsStats(d *schema.ResourceData) edpt.DdosDstZonePortRangeIpsStats {
	var ret edpt.DdosDstZonePortRangeIpsStats

	ret.Stats = getObjectDdosDstZonePortRangeIpsStatsStats(d.Get("stats").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
