package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nptv6DomainStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nptv6_domain_stats`: Statistics for the object domain\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nptv6DomainStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of NPTv6 domain",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outbound_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Packets",
						},
						"inbound_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets",
						},
						"hairpin_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Hairpin Packets",
						},
						"address_not_valid_for_translation": {
							Type: schema.TypeInt, Optional: true, Description: "Address Not Valid For Translation",
						},
						"inbound_packets_no_map": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets No Map",
						},
						"packets_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Destination Unreachable",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Nptv6DomainStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6DomainStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6DomainStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nptv6DomainStatsStats := setObjectCgnv6Nptv6DomainStatsStats(res)
		d.Set("stats", Cgnv6Nptv6DomainStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nptv6DomainStatsStats(ret edpt.DataCgnv6Nptv6DomainStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"outbound_packets":                  ret.DtCgnv6Nptv6DomainStats.Stats.OutboundPackets,
			"inbound_packets":                   ret.DtCgnv6Nptv6DomainStats.Stats.InboundPackets,
			"hairpin_packets":                   ret.DtCgnv6Nptv6DomainStats.Stats.HairpinPackets,
			"address_not_valid_for_translation": ret.DtCgnv6Nptv6DomainStats.Stats.AddressNotValidForTranslation,
			"inbound_packets_no_map":            ret.DtCgnv6Nptv6DomainStats.Stats.InboundPacketsNoMap,
			"packets_dest_unreachable":          ret.DtCgnv6Nptv6DomainStats.Stats.PacketsDestUnreachable,
		},
	}
}

func getObjectCgnv6Nptv6DomainStatsStats(d []interface{}) edpt.Cgnv6Nptv6DomainStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Nptv6DomainStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutboundPackets = in["outbound_packets"].(int)
		ret.InboundPackets = in["inbound_packets"].(int)
		ret.HairpinPackets = in["hairpin_packets"].(int)
		ret.AddressNotValidForTranslation = in["address_not_valid_for_translation"].(int)
		ret.InboundPacketsNoMap = in["inbound_packets_no_map"].(int)
		ret.PacketsDestUnreachable = in["packets_dest_unreachable"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Nptv6DomainStats(d *schema.ResourceData) edpt.Cgnv6Nptv6DomainStats {
	var ret edpt.Cgnv6Nptv6DomainStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectCgnv6Nptv6DomainStatsStats(d.Get("stats").([]interface{}))
	return ret
}
