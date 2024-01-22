package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryL4TypeStats() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_l4_type_stats_l4type_icmp`: Statistics for the object l4-type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryL4TypeStatsCreate,
		UpdateContext: resourceDdosDstEntryL4TypeStatsUpdate,
		ReadContext:   resourceDdosDstEntryL4TypeStatsRead,
		DeleteContext: resourceDdosDstEntryL4TypeStatsDelete,

		Schema: map[string]*schema.Schema{
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': L4-Type TCP; 'udp': L4-Type UDP; 'icmp': L4-Type ICMP; 'other': L4-Type OTHER;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4type_icmp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rate_type0_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Exceeded",
									},
									"rate_type1_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Exceeded",
									},
									"rate_type2_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Exceeded",
									},
									"type_deny_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dropped",
									},
									"icmpv4_rfc_undef_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMPv4 RFC Undef Type Dropped",
									},
									"icmpv6_rfc_undef_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMPv6 RFC Undef Type Dropped",
									},
									"wildcard_deny_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Dropped",
									},
									"port_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Received",
									},
									"port_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Dropped",
									},
									"port_pkt_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Forwarded",
									},
									"type": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type",
									},
									"type_bl": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Blacklisted",
									},
									"wildcard": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard",
									},
									"wildcard_bl": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Blacklisted",
									},
									"rate_type0_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Dropped",
									},
									"rate_type0_exceed_bl": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Blacklisted",
									},
									"rate_type1_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Dropped",
									},
									"rate_type1_exceed_bl": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Blacklisted",
									},
									"rate_type2_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Dropped",
									},
									"rate_type2_exceed_bl": {
										Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Blacklisted",
									},
									"port_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Received",
									},
									"outbound_port_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Received",
									},
									"outbound_port_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Received",
									},
									"outbound_port_pkt_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Forwarded",
									},
									"port_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Forwarded",
									},
									"port_bytes_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Dropped",
									},
									"port_src_bl": {
										Type: schema.TypeInt, Optional: true, Description: "Src Blacklisted",
									},
									"port_pkt_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Packet Rate Exceeded",
									},
									"port_kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded",
									},
									"exceed_drop_prate_src": {
										Type: schema.TypeInt, Optional: true, Description: "Src Pkt Rate Exceeded",
									},
									"exceed_drop_brate_src": {
										Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded",
									},
									"outbound_port_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Forwarded",
									},
									"outbound_port_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Dropped",
									},
									"outbound_port_bytes_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Dropped",
									},
									"exceed_drop_brate_src_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded Count",
									},
									"port_kbit_rate_exceed_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded Count",
									},
									"bl": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Blacklisted",
									},
									"src_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Packets Dropped",
									},
									"frag_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Received",
									},
									"frag_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Dropped",
									},
									"frag_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
									},
									"src_frag_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
									},
									"sflow_internal_samples_packed": {
										Type: schema.TypeInt, Optional: true, Description: "Sflow Internal Samples Packed",
									},
									"sflow_external_samples_packed": {
										Type: schema.TypeInt, Optional: true, Description: "Sflow External Samples Packed",
									},
									"sflow_internal_packets_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Sflow Internal Packets Sent",
									},
									"sflow_external_packets_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Sflow External Packets Sent",
									},
									"exceed_action_tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Tunnel",
									},
									"dst_hw_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Packets Dropped",
									},
									"exceed_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Dropped",
									},
								},
							},
						},
					},
				},
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntryL4TypeStatsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeStatsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypeStats(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryL4TypeStatsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryL4TypeStatsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeStatsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypeStats(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryL4TypeStatsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryL4TypeStatsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeStatsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypeStats(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryL4TypeStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypeStats(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntryL4TypeStatsStats(d []interface{}) edpt.DdosDstEntryL4TypeStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4typeIcmp = getObjectDdosDstEntryL4TypeStatsStatsL4typeIcmp(in["l4type_icmp"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeStatsStatsL4typeIcmp(d []interface{}) edpt.DdosDstEntryL4TypeStatsStatsL4typeIcmp {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeStatsStatsL4typeIcmp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rate_type0_exceed = in["rate_type0_exceed"].(int)
		ret.Rate_type1_exceed = in["rate_type1_exceed"].(int)
		ret.Rate_type2_exceed = in["rate_type2_exceed"].(int)
		ret.Type_deny_drop = in["type_deny_drop"].(int)
		ret.Icmpv4_rfc_undef_drop = in["icmpv4_rfc_undef_drop"].(int)
		ret.Icmpv6_rfc_undef_drop = in["icmpv6_rfc_undef_drop"].(int)
		ret.Wildcard_deny_drop = in["wildcard_deny_drop"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Type = in["type"].(int)
		ret.Type_bl = in["type_bl"].(int)
		ret.Wildcard = in["wildcard"].(int)
		ret.Wildcard_bl = in["wildcard_bl"].(int)
		ret.Rate_type0_exceed_drop = in["rate_type0_exceed_drop"].(int)
		ret.Rate_type0_exceed_bl = in["rate_type0_exceed_bl"].(int)
		ret.Rate_type1_exceed_drop = in["rate_type1_exceed_drop"].(int)
		ret.Rate_type1_exceed_bl = in["rate_type1_exceed_bl"].(int)
		ret.Rate_type2_exceed_drop = in["rate_type2_exceed_drop"].(int)
		ret.Rate_type2_exceed_bl = in["rate_type2_exceed_bl"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
		ret.Exceed_drop_prate_src = in["exceed_drop_prate_src"].(int)
		ret.Exceed_drop_brate_src = in["exceed_drop_brate_src"].(int)
		ret.Outbound_port_bytes_sent = in["outbound_port_bytes_sent"].(int)
		ret.Outbound_port_drop = in["outbound_port_drop"].(int)
		ret.Outbound_port_bytes_drop = in["outbound_port_bytes_drop"].(int)
		ret.Exceed_drop_brate_src_pkt = in["exceed_drop_brate_src_pkt"].(int)
		ret.Port_kbit_rate_exceed_pkt = in["port_kbit_rate_exceed_pkt"].(int)
		ret.Bl = in["bl"].(int)
		ret.Src_drop = in["src_drop"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.Frag_drop = in["frag_drop"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Exceed_action_drop = in["exceed_action_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosDstEntryL4TypeStats(d *schema.ResourceData) edpt.DdosDstEntryL4TypeStats {
	var ret edpt.DdosDstEntryL4TypeStats
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectDdosDstEntryL4TypeStatsStats(d.Get("stats").([]interface{}))
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
