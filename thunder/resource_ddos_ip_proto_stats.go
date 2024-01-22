package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosIpProtoStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_ip_proto_stats`: Statistics for the object ip-proto\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosIpProtoStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_ipproto_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Total Packets Received",
						},
						"dst_ipproto_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Total Packets Dropped",
						},
						"dst_ipproto_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Blacklist Packets Dropped",
						},
						"dst_ipproto_exceed_drop_any": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Exceed Dropped",
						},
						"dst_ipproto_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Pkt Rate Exceeded",
						},
						"dst_ipproto_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto KiBit Rate Exceeded (KiBit)",
						},
						"dst_ipproto_frag_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Frag Pkt Rate Exceeded",
						},
					},
				},
			},
		},
	}
}

func resourceDdosIpProtoStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpProtoStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpProtoStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosIpProtoStatsStats := setObjectDdosIpProtoStatsStats(res)
		d.Set("stats", DdosIpProtoStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosIpProtoStatsStats(ret edpt.DataDdosIpProtoStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dst_ipproto_rcvd":             ret.DtDdosIpProtoStats.Stats.Dst_ipproto_rcvd,
			"dst_ipproto_drop":             ret.DtDdosIpProtoStats.Stats.Dst_ipproto_drop,
			"dst_ipproto_bl":               ret.DtDdosIpProtoStats.Stats.Dst_ipproto_bl,
			"dst_ipproto_exceed_drop_any":  ret.DtDdosIpProtoStats.Stats.Dst_ipproto_exceed_drop_any,
			"dst_ipproto_pkt_rate_exceed":  ret.DtDdosIpProtoStats.Stats.Dst_ipproto_pkt_rate_exceed,
			"dst_ipproto_kbit_rate_exceed": ret.DtDdosIpProtoStats.Stats.Dst_ipproto_kbit_rate_exceed,
			"dst_ipproto_frag_rate_exceed": ret.DtDdosIpProtoStats.Stats.Dst_ipproto_frag_rate_exceed,
		},
	}
}

func getObjectDdosIpProtoStatsStats(d []interface{}) edpt.DdosIpProtoStatsStats {

	count1 := len(d)
	var ret edpt.DdosIpProtoStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dst_ipproto_rcvd = in["dst_ipproto_rcvd"].(int)
		ret.Dst_ipproto_drop = in["dst_ipproto_drop"].(int)
		ret.Dst_ipproto_bl = in["dst_ipproto_bl"].(int)
		ret.Dst_ipproto_exceed_drop_any = in["dst_ipproto_exceed_drop_any"].(int)
		ret.Dst_ipproto_pkt_rate_exceed = in["dst_ipproto_pkt_rate_exceed"].(int)
		ret.Dst_ipproto_kbit_rate_exceed = in["dst_ipproto_kbit_rate_exceed"].(int)
		ret.Dst_ipproto_frag_rate_exceed = in["dst_ipproto_frag_rate_exceed"].(int)
	}
	return ret
}

func dataToEndpointDdosIpProtoStats(d *schema.ResourceData) edpt.DdosIpProtoStats {
	var ret edpt.DdosIpProtoStats

	ret.Stats = getObjectDdosIpProtoStatsStats(d.Get("stats").([]interface{}))
	return ret
}
