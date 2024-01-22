package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6PcpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_pcp_stats`: Statistics for the object pcp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6PcpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packets_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Received",
						},
						"lsn_map_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP MAP Request Processing Success (NAT44)",
						},
						"dslite_map_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP MAP Request Processing Success (DS-Lite)",
						},
						"nat64_map_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP MAP Request Processing Success (NAT64)",
						},
						"lsn_peer_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP PEER Request Processing Success (NAT44)",
						},
						"dslite_peer_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP PEER Request Processing Success (DS-Lite)",
						},
						"nat64_peer_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP PEER Request Processing Success (NAT64)",
						},
						"lsn_announce_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP ANNOUNCE Request Processing Success (NAT44)",
						},
						"dslite_announce_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP ANNOUNCE Request Processing Success (DS-Lite)",
						},
						"nat64_announce_process_success": {
							Type: schema.TypeInt, Optional: true, Description: "PCP ANNOUNCE Request Processing Success (NAT64)",
						},
						"pkt_not_request_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Not a PCP Request",
						},
						"pkt_too_short_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Too Short",
						},
						"noroute_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Response No Route",
						},
						"unsupported_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported PCP version",
						},
						"not_authorized": {
							Type: schema.TypeInt, Optional: true, Description: "PCP Request Not Authorized",
						},
						"malform_request": {
							Type: schema.TypeInt, Optional: true, Description: "PCP Request Malformed",
						},
						"unsupp_opcode": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported PCP Opcode",
						},
						"unsupp_option": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported PCP Option",
						},
						"malform_option": {
							Type: schema.TypeInt, Optional: true, Description: "PCP Option Malformed",
						},
						"no_resources": {
							Type: schema.TypeInt, Optional: true, Description: "No System or NAT Resources",
						},
						"unsupp_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported Mapping Protocol",
						},
						"user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "User Quota Exceeded",
						},
						"cannot_provide_suggest": {
							Type: schema.TypeInt, Optional: true, Description: "Cannot Provide Suggested Port When PREFER_FAILURE",
						},
						"address_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "PCP Client Address Mismatch",
						},
						"excessive_remote_peers": {
							Type: schema.TypeInt, Optional: true, Description: "Excessive Remote Peers",
						},
						"pkt_not_from_nat_inside": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Dropped For Not Coming From NAT Inside",
						},
						"l4_process_error": {
							Type: schema.TypeInt, Optional: true, Description: "L3/L4 Process Error",
						},
						"internal_error_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Internal Error",
						},
						"unsol_ance_sent_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Unsolicited Announce Sent",
						},
						"unsol_ance_sent_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Unsolicited Announce Send Failure",
						},
						"ha_sync_epoch_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HA Sync PCP Epoch Sent",
						},
						"ha_sync_epoch_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "HA Sync PCP Epoch Recv",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6PcpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PcpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PcpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6PcpStatsStats := setObjectCgnv6PcpStatsStats(res)
		d.Set("stats", Cgnv6PcpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6PcpStatsStats(ret edpt.DataCgnv6PcpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packets_rcv":                     ret.DtCgnv6PcpStats.Stats.PacketsRcv,
			"lsn_map_process_success":         ret.DtCgnv6PcpStats.Stats.LsnMapProcessSuccess,
			"dslite_map_process_success":      ret.DtCgnv6PcpStats.Stats.DsliteMapProcessSuccess,
			"nat64_map_process_success":       ret.DtCgnv6PcpStats.Stats.Nat64MapProcessSuccess,
			"lsn_peer_process_success":        ret.DtCgnv6PcpStats.Stats.LsnPeerProcessSuccess,
			"dslite_peer_process_success":     ret.DtCgnv6PcpStats.Stats.DslitePeerProcessSuccess,
			"nat64_peer_process_success":      ret.DtCgnv6PcpStats.Stats.Nat64PeerProcessSuccess,
			"lsn_announce_process_success":    ret.DtCgnv6PcpStats.Stats.LsnAnnounceProcessSuccess,
			"dslite_announce_process_success": ret.DtCgnv6PcpStats.Stats.DsliteAnnounceProcessSuccess,
			"nat64_announce_process_success":  ret.DtCgnv6PcpStats.Stats.Nat64AnnounceProcessSuccess,
			"pkt_not_request_drop":            ret.DtCgnv6PcpStats.Stats.PktNotRequestDrop,
			"pkt_too_short_drop":              ret.DtCgnv6PcpStats.Stats.PktTooShortDrop,
			"noroute_drop":                    ret.DtCgnv6PcpStats.Stats.NorouteDrop,
			"unsupported_version":             ret.DtCgnv6PcpStats.Stats.UnsupportedVersion,
			"not_authorized":                  ret.DtCgnv6PcpStats.Stats.NotAuthorized,
			"malform_request":                 ret.DtCgnv6PcpStats.Stats.MalformRequest,
			"unsupp_opcode":                   ret.DtCgnv6PcpStats.Stats.UnsuppOpcode,
			"unsupp_option":                   ret.DtCgnv6PcpStats.Stats.UnsuppOption,
			"malform_option":                  ret.DtCgnv6PcpStats.Stats.MalformOption,
			"no_resources":                    ret.DtCgnv6PcpStats.Stats.NoResources,
			"unsupp_protocol":                 ret.DtCgnv6PcpStats.Stats.UnsuppProtocol,
			"user_quota_exceeded":             ret.DtCgnv6PcpStats.Stats.UserQuotaExceeded,
			"cannot_provide_suggest":          ret.DtCgnv6PcpStats.Stats.CannotProvideSuggest,
			"address_mismatch":                ret.DtCgnv6PcpStats.Stats.AddressMismatch,
			"excessive_remote_peers":          ret.DtCgnv6PcpStats.Stats.ExcessiveRemotePeers,
			"pkt_not_from_nat_inside":         ret.DtCgnv6PcpStats.Stats.PktNotFromNatInside,
			"l4_process_error":                ret.DtCgnv6PcpStats.Stats.L4ProcessError,
			"internal_error_drop":             ret.DtCgnv6PcpStats.Stats.InternalErrorDrop,
			"unsol_ance_sent_succ":            ret.DtCgnv6PcpStats.Stats.Unsol_ance_sent_succ,
			"unsol_ance_sent_fail":            ret.DtCgnv6PcpStats.Stats.Unsol_ance_sent_fail,
			"ha_sync_epoch_sent":              ret.DtCgnv6PcpStats.Stats.Ha_sync_epoch_sent,
			"ha_sync_epoch_rcv":               ret.DtCgnv6PcpStats.Stats.Ha_sync_epoch_rcv,
		},
	}
}

func getObjectCgnv6PcpStatsStats(d []interface{}) edpt.Cgnv6PcpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6PcpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketsRcv = in["packets_rcv"].(int)
		ret.LsnMapProcessSuccess = in["lsn_map_process_success"].(int)
		ret.DsliteMapProcessSuccess = in["dslite_map_process_success"].(int)
		ret.Nat64MapProcessSuccess = in["nat64_map_process_success"].(int)
		ret.LsnPeerProcessSuccess = in["lsn_peer_process_success"].(int)
		ret.DslitePeerProcessSuccess = in["dslite_peer_process_success"].(int)
		ret.Nat64PeerProcessSuccess = in["nat64_peer_process_success"].(int)
		ret.LsnAnnounceProcessSuccess = in["lsn_announce_process_success"].(int)
		ret.DsliteAnnounceProcessSuccess = in["dslite_announce_process_success"].(int)
		ret.Nat64AnnounceProcessSuccess = in["nat64_announce_process_success"].(int)
		ret.PktNotRequestDrop = in["pkt_not_request_drop"].(int)
		ret.PktTooShortDrop = in["pkt_too_short_drop"].(int)
		ret.NorouteDrop = in["noroute_drop"].(int)
		ret.UnsupportedVersion = in["unsupported_version"].(int)
		ret.NotAuthorized = in["not_authorized"].(int)
		ret.MalformRequest = in["malform_request"].(int)
		ret.UnsuppOpcode = in["unsupp_opcode"].(int)
		ret.UnsuppOption = in["unsupp_option"].(int)
		ret.MalformOption = in["malform_option"].(int)
		ret.NoResources = in["no_resources"].(int)
		ret.UnsuppProtocol = in["unsupp_protocol"].(int)
		ret.UserQuotaExceeded = in["user_quota_exceeded"].(int)
		ret.CannotProvideSuggest = in["cannot_provide_suggest"].(int)
		ret.AddressMismatch = in["address_mismatch"].(int)
		ret.ExcessiveRemotePeers = in["excessive_remote_peers"].(int)
		ret.PktNotFromNatInside = in["pkt_not_from_nat_inside"].(int)
		ret.L4ProcessError = in["l4_process_error"].(int)
		ret.InternalErrorDrop = in["internal_error_drop"].(int)
		ret.Unsol_ance_sent_succ = in["unsol_ance_sent_succ"].(int)
		ret.Unsol_ance_sent_fail = in["unsol_ance_sent_fail"].(int)
		ret.Ha_sync_epoch_sent = in["ha_sync_epoch_sent"].(int)
		ret.Ha_sync_epoch_rcv = in["ha_sync_epoch_rcv"].(int)
	}
	return ret
}

func dataToEndpointCgnv6PcpStats(d *schema.ResourceData) edpt.Cgnv6PcpStats {
	var ret edpt.Cgnv6PcpStats

	ret.Stats = getObjectCgnv6PcpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
