package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_port_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosPortStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_port_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Undefined Dropped",
						},
						"dst_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Blacklist Packets Dropped",
						},
						"dst_port_exceed_drop_any": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Exceed Dropped",
						},
						"dst_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Pkt Rate Exceeded",
						},
						"dst_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port KiBit Rate Exceeded (KiBit)",
						},
						"dst_port_frag_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Frag Pkt Rate Exceeded",
						},
						"dst_port_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Conn Limit Exceeded",
						},
						"dst_port_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Conn Rate Exceeded",
						},
						"dst_sport_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Blacklist Packets Dropped",
						},
						"dst_sport_exceed_drop_any": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Exceed Dropped",
						},
						"dst_sport_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Pkt Rate Exceeded",
						},
						"dst_sport_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort KiBit Rate Exceeded (KiBit)",
						},
						"dst_sport_frag_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Frag Pkt Rate Exceeded",
						},
						"dst_sport_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Conn Limit Exceeded",
						},
						"dst_sport_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Conn Rate Exceeded",
						},
					},
				},
			},
		},
	}
}

func resourceDdosPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosPortStatsStats := setObjectDdosPortStatsStats(res)
		d.Set("stats", DdosPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosPortStatsStats(ret edpt.DataDdosPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dst_port_undef_drop":         ret.DtDdosPortStats.Stats.Dst_port_undef_drop,
			"dst_port_bl":                 ret.DtDdosPortStats.Stats.Dst_port_bl,
			"dst_port_exceed_drop_any":    ret.DtDdosPortStats.Stats.Dst_port_exceed_drop_any,
			"dst_port_pkt_rate_exceed":    ret.DtDdosPortStats.Stats.Dst_port_pkt_rate_exceed,
			"dst_port_kbit_rate_exceed":   ret.DtDdosPortStats.Stats.Dst_port_kbit_rate_exceed,
			"dst_port_frag_rate_exceed":   ret.DtDdosPortStats.Stats.Dst_port_frag_rate_exceed,
			"dst_port_conn_limit_exceed":  ret.DtDdosPortStats.Stats.Dst_port_conn_limit_exceed,
			"dst_port_conn_rate_exceed":   ret.DtDdosPortStats.Stats.Dst_port_conn_rate_exceed,
			"dst_sport_bl":                ret.DtDdosPortStats.Stats.Dst_sport_bl,
			"dst_sport_exceed_drop_any":   ret.DtDdosPortStats.Stats.Dst_sport_exceed_drop_any,
			"dst_sport_pkt_rate_exceed":   ret.DtDdosPortStats.Stats.Dst_sport_pkt_rate_exceed,
			"dst_sport_kbit_rate_exceed":  ret.DtDdosPortStats.Stats.Dst_sport_kbit_rate_exceed,
			"dst_sport_frag_rate_exceed":  ret.DtDdosPortStats.Stats.Dst_sport_frag_rate_exceed,
			"dst_sport_conn_limit_exceed": ret.DtDdosPortStats.Stats.Dst_sport_conn_limit_exceed,
			"dst_sport_conn_rate_exceed":  ret.DtDdosPortStats.Stats.Dst_sport_conn_rate_exceed,
		},
	}
}

func getObjectDdosPortStatsStats(d []interface{}) edpt.DdosPortStatsStats {

	count1 := len(d)
	var ret edpt.DdosPortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dst_port_undef_drop = in["dst_port_undef_drop"].(int)
		ret.Dst_port_bl = in["dst_port_bl"].(int)
		ret.Dst_port_exceed_drop_any = in["dst_port_exceed_drop_any"].(int)
		ret.Dst_port_pkt_rate_exceed = in["dst_port_pkt_rate_exceed"].(int)
		ret.Dst_port_kbit_rate_exceed = in["dst_port_kbit_rate_exceed"].(int)
		ret.Dst_port_frag_rate_exceed = in["dst_port_frag_rate_exceed"].(int)
		ret.Dst_port_conn_limit_exceed = in["dst_port_conn_limit_exceed"].(int)
		ret.Dst_port_conn_rate_exceed = in["dst_port_conn_rate_exceed"].(int)
		ret.Dst_sport_bl = in["dst_sport_bl"].(int)
		ret.Dst_sport_exceed_drop_any = in["dst_sport_exceed_drop_any"].(int)
		ret.Dst_sport_pkt_rate_exceed = in["dst_sport_pkt_rate_exceed"].(int)
		ret.Dst_sport_kbit_rate_exceed = in["dst_sport_kbit_rate_exceed"].(int)
		ret.Dst_sport_frag_rate_exceed = in["dst_sport_frag_rate_exceed"].(int)
		ret.Dst_sport_conn_limit_exceed = in["dst_sport_conn_limit_exceed"].(int)
		ret.Dst_sport_conn_rate_exceed = in["dst_sport_conn_rate_exceed"].(int)
	}
	return ret
}

func dataToEndpointDdosPortStats(d *schema.ResourceData) edpt.DdosPortStats {
	var ret edpt.DdosPortStats

	ret.Stats = getObjectDdosPortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
