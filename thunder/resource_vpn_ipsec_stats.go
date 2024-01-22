package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ipsec_stats`: Statistics for the object ipsec\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIpsecStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "IPsec name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packets_encrypted": {
							Type: schema.TypeInt, Optional: true, Description: "Encrypted Packets",
						},
						"packets_decrypted": {
							Type: schema.TypeInt, Optional: true, Description: "Decrypted Packets",
						},
						"anti_replay_num": {
							Type: schema.TypeInt, Optional: true, Description: "Anti-Replay Failure",
						},
						"rekey_num": {
							Type: schema.TypeInt, Optional: true, Description: "Rekey Times",
						},
						"packets_err_inactive": {
							Type: schema.TypeInt, Optional: true, Description: "Inactive Error",
						},
						"packets_err_encryption": {
							Type: schema.TypeInt, Optional: true, Description: "Encryption Error",
						},
						"packets_err_pad_check": {
							Type: schema.TypeInt, Optional: true, Description: "Pad Check Error",
						},
						"packets_err_pkt_sanity": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Sanity Error",
						},
						"packets_err_icv_check": {
							Type: schema.TypeInt, Optional: true, Description: "ICV Check Error",
						},
						"packets_err_lifetime_lifebytes": {
							Type: schema.TypeInt, Optional: true, Description: "Lifetime Lifebytes Error",
						},
						"bytes_encrypted": {
							Type: schema.TypeInt, Optional: true, Description: "Encrypted Bytes",
						},
						"bytes_decrypted": {
							Type: schema.TypeInt, Optional: true, Description: "Decrypted Bytes",
						},
						"prefrag_success": {
							Type: schema.TypeInt, Optional: true, Description: "Pre-frag Success",
						},
						"prefrag_error": {
							Type: schema.TypeInt, Optional: true, Description: "Pre-frag Error",
						},
						"cavium_bytes_encrypted": {
							Type: schema.TypeInt, Optional: true, Description: "CAVIUM Encrypted Bytes",
						},
						"cavium_bytes_decrypted": {
							Type: schema.TypeInt, Optional: true, Description: "CAVIUM Decrypted Bytes",
						},
						"cavium_packets_encrypted": {
							Type: schema.TypeInt, Optional: true, Description: "CAVIUM Encrypted Packets",
						},
						"cavium_packets_decrypted": {
							Type: schema.TypeInt, Optional: true, Description: "CAVIUM Decrypted Packets",
						},
						"qat_bytes_encrypted": {
							Type: schema.TypeInt, Optional: true, Description: "QAT Encrypted Bytes",
						},
						"qat_bytes_decrypted": {
							Type: schema.TypeInt, Optional: true, Description: "QAT Decrypted Bytes",
						},
						"qat_packets_encrypted": {
							Type: schema.TypeInt, Optional: true, Description: "QAT Encrypted Packets",
						},
						"qat_packets_decrypted": {
							Type: schema.TypeInt, Optional: true, Description: "QAT Decrypted Packets",
						},
						"tunnel_intf_down": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Tunnel Interface Down",
						},
						"pkt_fail_prep_to_send": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Failed in prepare to send",
						},
						"no_next_hop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped: No next hop",
						},
						"invalid_tunnel_id": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Invalid tunnel ID",
						},
						"no_tunnel_found": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped: No tunnel found",
						},
						"pkt_fail_to_send": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Failed to send",
						},
						"frag_after_encap_frag_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Frag-after-encap Fragment Generated",
						},
						"frag_received": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Received",
						},
						"sequence_num": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Number",
						},
						"sequence_num_rollover": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Number Rollover",
						},
						"packets_err_nh_check": {
							Type: schema.TypeInt, Optional: true, Description: "Next Header Check Error",
						},
					},
				},
			},
		},
	}
}

func resourceVpnIpsecStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIpsecStatsStats := setObjectVpnIpsecStatsStats(res)
		d.Set("stats", VpnIpsecStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIpsecStatsStats(ret edpt.DataVpnIpsecStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packets_encrypted":              ret.DtVpnIpsecStats.Stats.PacketsEncrypted,
			"packets_decrypted":              ret.DtVpnIpsecStats.Stats.PacketsDecrypted,
			"anti_replay_num":                ret.DtVpnIpsecStats.Stats.AntiReplayNum,
			"rekey_num":                      ret.DtVpnIpsecStats.Stats.RekeyNum,
			"packets_err_inactive":           ret.DtVpnIpsecStats.Stats.PacketsErrInactive,
			"packets_err_encryption":         ret.DtVpnIpsecStats.Stats.PacketsErrEncryption,
			"packets_err_pad_check":          ret.DtVpnIpsecStats.Stats.PacketsErrPadCheck,
			"packets_err_pkt_sanity":         ret.DtVpnIpsecStats.Stats.PacketsErrPktSanity,
			"packets_err_icv_check":          ret.DtVpnIpsecStats.Stats.PacketsErrIcvCheck,
			"packets_err_lifetime_lifebytes": ret.DtVpnIpsecStats.Stats.PacketsErrLifetimeLifebytes,
			"bytes_encrypted":                ret.DtVpnIpsecStats.Stats.BytesEncrypted,
			"bytes_decrypted":                ret.DtVpnIpsecStats.Stats.BytesDecrypted,
			"prefrag_success":                ret.DtVpnIpsecStats.Stats.PrefragSuccess,
			"prefrag_error":                  ret.DtVpnIpsecStats.Stats.PrefragError,
			"cavium_bytes_encrypted":         ret.DtVpnIpsecStats.Stats.CaviumBytesEncrypted,
			"cavium_bytes_decrypted":         ret.DtVpnIpsecStats.Stats.CaviumBytesDecrypted,
			"cavium_packets_encrypted":       ret.DtVpnIpsecStats.Stats.CaviumPacketsEncrypted,
			"cavium_packets_decrypted":       ret.DtVpnIpsecStats.Stats.CaviumPacketsDecrypted,
			"qat_bytes_encrypted":            ret.DtVpnIpsecStats.Stats.QatBytesEncrypted,
			"qat_bytes_decrypted":            ret.DtVpnIpsecStats.Stats.QatBytesDecrypted,
			"qat_packets_encrypted":          ret.DtVpnIpsecStats.Stats.QatPacketsEncrypted,
			"qat_packets_decrypted":          ret.DtVpnIpsecStats.Stats.QatPacketsDecrypted,
			"tunnel_intf_down":               ret.DtVpnIpsecStats.Stats.TunnelIntfDown,
			"pkt_fail_prep_to_send":          ret.DtVpnIpsecStats.Stats.PktFailPrepToSend,
			"no_next_hop":                    ret.DtVpnIpsecStats.Stats.NoNextHop,
			"invalid_tunnel_id":              ret.DtVpnIpsecStats.Stats.InvalidTunnelId,
			"no_tunnel_found":                ret.DtVpnIpsecStats.Stats.NoTunnelFound,
			"pkt_fail_to_send":               ret.DtVpnIpsecStats.Stats.PktFailToSend,
			"frag_after_encap_frag_packets":  ret.DtVpnIpsecStats.Stats.FragAfterEncapFragPackets,
			"frag_received":                  ret.DtVpnIpsecStats.Stats.FragReceived,
			"sequence_num":                   ret.DtVpnIpsecStats.Stats.SequenceNum,
			"sequence_num_rollover":          ret.DtVpnIpsecStats.Stats.SequenceNumRollover,
			"packets_err_nh_check":           ret.DtVpnIpsecStats.Stats.PacketsErrNhCheck,
		},
	}
}

func getObjectVpnIpsecStatsStats(d []interface{}) edpt.VpnIpsecStatsStats {

	count1 := len(d)
	var ret edpt.VpnIpsecStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketsEncrypted = in["packets_encrypted"].(int)
		ret.PacketsDecrypted = in["packets_decrypted"].(int)
		ret.AntiReplayNum = in["anti_replay_num"].(int)
		ret.RekeyNum = in["rekey_num"].(int)
		ret.PacketsErrInactive = in["packets_err_inactive"].(int)
		ret.PacketsErrEncryption = in["packets_err_encryption"].(int)
		ret.PacketsErrPadCheck = in["packets_err_pad_check"].(int)
		ret.PacketsErrPktSanity = in["packets_err_pkt_sanity"].(int)
		ret.PacketsErrIcvCheck = in["packets_err_icv_check"].(int)
		ret.PacketsErrLifetimeLifebytes = in["packets_err_lifetime_lifebytes"].(int)
		ret.BytesEncrypted = in["bytes_encrypted"].(int)
		ret.BytesDecrypted = in["bytes_decrypted"].(int)
		ret.PrefragSuccess = in["prefrag_success"].(int)
		ret.PrefragError = in["prefrag_error"].(int)
		ret.CaviumBytesEncrypted = in["cavium_bytes_encrypted"].(int)
		ret.CaviumBytesDecrypted = in["cavium_bytes_decrypted"].(int)
		ret.CaviumPacketsEncrypted = in["cavium_packets_encrypted"].(int)
		ret.CaviumPacketsDecrypted = in["cavium_packets_decrypted"].(int)
		ret.QatBytesEncrypted = in["qat_bytes_encrypted"].(int)
		ret.QatBytesDecrypted = in["qat_bytes_decrypted"].(int)
		ret.QatPacketsEncrypted = in["qat_packets_encrypted"].(int)
		ret.QatPacketsDecrypted = in["qat_packets_decrypted"].(int)
		ret.TunnelIntfDown = in["tunnel_intf_down"].(int)
		ret.PktFailPrepToSend = in["pkt_fail_prep_to_send"].(int)
		ret.NoNextHop = in["no_next_hop"].(int)
		ret.InvalidTunnelId = in["invalid_tunnel_id"].(int)
		ret.NoTunnelFound = in["no_tunnel_found"].(int)
		ret.PktFailToSend = in["pkt_fail_to_send"].(int)
		ret.FragAfterEncapFragPackets = in["frag_after_encap_frag_packets"].(int)
		ret.FragReceived = in["frag_received"].(int)
		ret.SequenceNum = in["sequence_num"].(int)
		ret.SequenceNumRollover = in["sequence_num_rollover"].(int)
		ret.PacketsErrNhCheck = in["packets_err_nh_check"].(int)
	}
	return ret
}

func dataToEndpointVpnIpsecStats(d *schema.ResourceData) edpt.VpnIpsecStats {
	var ret edpt.VpnIpsecStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectVpnIpsecStatsStats(d.Get("stats").([]interface{}))
	return ret
}
