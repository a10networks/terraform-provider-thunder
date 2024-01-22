package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapEncapsulationDomainStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_map_encapsulation_domain_stats`: Statistics for the object domain\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6MapEncapsulationDomainStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "MAP-E domain name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inbound_packet_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 Packets Received",
						},
						"inbound_frag_packet_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 Fragment Packets Received",
						},
						"inbound_addr_port_validation_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 Destination Address Port Validation Failed",
						},
						"inbound_rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 Reverse Route Lookup Failed",
						},
						"inbound_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv6 Destination Address Unreachable",
						},
						"outbound_packet_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 Packets Received",
						},
						"outbound_frag_packet_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 Fragment Packets Received",
						},
						"outbound_addr_validation_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 Source Address Validation Failed",
						},
						"outbound_rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 Reverse Route Lookup Failed",
						},
						"outbound_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv4 Destination Address Unreachable",
						},
						"packet_mtu_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Exceeded MTU",
						},
						"frag_icmp_sent": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Packet Too Big Sent",
						},
						"interface_not_configured": {
							Type: schema.TypeInt, Optional: true, Description: "Interfaces not Configured Dropped",
						},
						"bmr_prefixrules_configured": {
							Type: schema.TypeInt, Optional: true, Description: "BMR prefix rules configured",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6MapEncapsulationDomainStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6MapEncapsulationDomainStatsStats := setObjectCgnv6MapEncapsulationDomainStatsStats(res)
		d.Set("stats", Cgnv6MapEncapsulationDomainStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6MapEncapsulationDomainStatsStats(ret edpt.DataCgnv6MapEncapsulationDomainStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"inbound_packet_received":             ret.DtCgnv6MapEncapsulationDomainStats.Stats.Inbound_packet_received,
			"inbound_frag_packet_received":        ret.DtCgnv6MapEncapsulationDomainStats.Stats.Inbound_frag_packet_received,
			"inbound_addr_port_validation_failed": ret.DtCgnv6MapEncapsulationDomainStats.Stats.Inbound_addr_port_validation_failed,
			"inbound_rev_lookup_failed":           ret.DtCgnv6MapEncapsulationDomainStats.Stats.Inbound_rev_lookup_failed,
			"inbound_dest_unreachable":            ret.DtCgnv6MapEncapsulationDomainStats.Stats.Inbound_dest_unreachable,
			"outbound_packet_received":            ret.DtCgnv6MapEncapsulationDomainStats.Stats.Outbound_packet_received,
			"outbound_frag_packet_received":       ret.DtCgnv6MapEncapsulationDomainStats.Stats.Outbound_frag_packet_received,
			"outbound_addr_validation_failed":     ret.DtCgnv6MapEncapsulationDomainStats.Stats.Outbound_addr_validation_failed,
			"outbound_rev_lookup_failed":          ret.DtCgnv6MapEncapsulationDomainStats.Stats.Outbound_rev_lookup_failed,
			"outbound_dest_unreachable":           ret.DtCgnv6MapEncapsulationDomainStats.Stats.Outbound_dest_unreachable,
			"packet_mtu_exceeded":                 ret.DtCgnv6MapEncapsulationDomainStats.Stats.Packet_mtu_exceeded,
			"frag_icmp_sent":                      ret.DtCgnv6MapEncapsulationDomainStats.Stats.Frag_icmp_sent,
			"interface_not_configured":            ret.DtCgnv6MapEncapsulationDomainStats.Stats.Interface_not_configured,
			"bmr_prefixrules_configured":          ret.DtCgnv6MapEncapsulationDomainStats.Stats.Bmr_prefixrules_configured,
		},
	}
}

func getObjectCgnv6MapEncapsulationDomainStatsStats(d []interface{}) edpt.Cgnv6MapEncapsulationDomainStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6MapEncapsulationDomainStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inbound_packet_received = in["inbound_packet_received"].(int)
		ret.Inbound_frag_packet_received = in["inbound_frag_packet_received"].(int)
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_packet_received = in["outbound_packet_received"].(int)
		ret.Outbound_frag_packet_received = in["outbound_frag_packet_received"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Frag_icmp_sent = in["frag_icmp_sent"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		ret.Bmr_prefixrules_configured = in["bmr_prefixrules_configured"].(int)
	}
	return ret
}

func dataToEndpointCgnv6MapEncapsulationDomainStats(d *schema.ResourceData) edpt.Cgnv6MapEncapsulationDomainStats {
	var ret edpt.Cgnv6MapEncapsulationDomainStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectCgnv6MapEncapsulationDomainStatsStats(d.Get("stats").([]interface{}))
	return ret
}
