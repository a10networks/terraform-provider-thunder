package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ipsec_oper`: Operational Status for the object ipsec\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIpsecOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "IPsec name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remote_ts_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"remote_ts_v6_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"in_spi_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"out_spi_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sa_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sa_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ts_proto": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peer_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peer_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"encryption_algorithm": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hash_algorithm": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lifetime": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lifebytes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dh_group": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_traversal": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"anti_replay": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"anti_replay_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rekey_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_inactive": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_encryption": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_pad_check": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_pkt_sanity": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_icv_check": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_lifetime_lifebytes": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},
									"bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"prefrag_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"prefrag_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cavium_bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cavium_bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cavium_packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cavium_packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"qat_bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"qat_bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"qat_packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"qat_packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tunnel_intf_down": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_fail_prep_to_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_next_hop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_tunnel_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_tunnel_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_fail_to_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_after_encap_frag_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sequence_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sequence_num_rollover": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packets_err_nh_check": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enforce_ts_encap_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enforce_ts_decap_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVpnIpsecOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIpsecOperOper := setObjectVpnIpsecOperOper(res)
		d.Set("oper", VpnIpsecOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIpsecOperOper(ret edpt.DataVpnIpsecOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"remote_ts_filter":    ret.DtVpnIpsecOper.Oper.RemoteTsFilter,
			"remote_ts_v6_filter": ret.DtVpnIpsecOper.Oper.RemoteTsV6Filter,
			"in_spi_filter":       ret.DtVpnIpsecOper.Oper.InSpiFilter,
			"out_spi_filter":      ret.DtVpnIpsecOper.Oper.OutSpiFilter,
			"sa_list":             setSliceVpnIpsecOperOperSaList(ret.DtVpnIpsecOper.Oper.SaList),
		},
	}
}

func setSliceVpnIpsecOperOperSaList(d []edpt.VpnIpsecOperOperSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["status"] = item.Status
		in["sa_index"] = item.SaIndex
		in["ts_proto"] = item.TsProto
		in["local_ip"] = item.LocalIp
		in["local_port"] = item.LocalPort
		in["peer_ip"] = item.PeerIp
		in["peer_port"] = item.PeerPort
		in["local_spi"] = item.LocalSpi
		in["remote_spi"] = item.RemoteSpi
		in["protocol"] = item.Protocol
		in["mode"] = item.Mode
		in["encryption_algorithm"] = item.EncryptionAlgorithm
		in["hash_algorithm"] = item.HashAlgorithm
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		in["dh_group"] = item.DhGroup
		in["nat_traversal"] = item.NatTraversal
		in["anti_replay"] = item.AntiReplay
		in["packets_encrypted"] = item.PacketsEncrypted
		in["packets_decrypted"] = item.PacketsDecrypted
		in["anti_replay_num"] = item.AntiReplayNum
		in["rekey_num"] = item.RekeyNum
		in["packets_err_inactive"] = item.PacketsErrInactive
		in["packets_err_encryption"] = item.PacketsErrEncryption
		in["packets_err_pad_check"] = item.PacketsErrPadCheck
		in["packets_err_pkt_sanity"] = item.PacketsErrPktSanity
		in["packets_err_icv_check"] = item.PacketsErrIcvCheck
		in["packets_err_lifetime_lifebytes"] = setObjectVpnIpsecOperOperSaListPacketsErrLifetimeLifebytes(item.PacketsErrLifetimeLifebytes)
		in["bytes_encrypted"] = item.BytesEncrypted
		in["bytes_decrypted"] = item.BytesDecrypted
		in["prefrag_success"] = item.PrefragSuccess
		in["prefrag_error"] = item.PrefragError
		in["cavium_bytes_encrypted"] = item.CaviumBytesEncrypted
		in["cavium_bytes_decrypted"] = item.CaviumBytesDecrypted
		in["cavium_packets_encrypted"] = item.CaviumPacketsEncrypted
		in["cavium_packets_decrypted"] = item.CaviumPacketsDecrypted
		in["qat_bytes_encrypted"] = item.QatBytesEncrypted
		in["qat_bytes_decrypted"] = item.QatBytesDecrypted
		in["qat_packets_encrypted"] = item.QatPacketsEncrypted
		in["qat_packets_decrypted"] = item.QatPacketsDecrypted
		in["tunnel_intf_down"] = item.TunnelIntfDown
		in["pkt_fail_prep_to_send"] = item.PktFailPrepToSend
		in["no_next_hop"] = item.NoNextHop
		in["invalid_tunnel_id"] = item.InvalidTunnelId
		in["no_tunnel_found"] = item.NoTunnelFound
		in["pkt_fail_to_send"] = item.PktFailToSend
		in["frag_after_encap_frag_packets"] = item.FragAfterEncapFragPackets
		in["frag_received"] = item.FragReceived
		in["sequence_num"] = item.SequenceNum
		in["sequence_num_rollover"] = item.SequenceNumRollover
		in["packets_err_nh_check"] = item.PacketsErrNhCheck
		in["enforce_ts_encap_drop"] = item.EnforceTsEncapDrop
		in["enforce_ts_decap_drop"] = item.EnforceTsDecapDrop
		result = append(result, in)
	}
	return result
}

func setObjectVpnIpsecOperOperSaListPacketsErrLifetimeLifebytes(d edpt.VpnIpsecOperOperSaListPacketsErrLifetimeLifebytes) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func getObjectVpnIpsecOperOper(d []interface{}) edpt.VpnIpsecOperOper {

	count1 := len(d)
	var ret edpt.VpnIpsecOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteTsFilter = in["remote_ts_filter"].(string)
		ret.RemoteTsV6Filter = in["remote_ts_v6_filter"].(string)
		ret.InSpiFilter = in["in_spi_filter"].(string)
		ret.OutSpiFilter = in["out_spi_filter"].(string)
		ret.SaList = getSliceVpnIpsecOperOperSaList(in["sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnIpsecOperOperSaList(d []interface{}) []edpt.VpnIpsecOperOperSaList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecOperOperSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecOperOperSaList
		oi.Status = in["status"].(string)
		oi.SaIndex = in["sa_index"].(int)
		oi.TsProto = in["ts_proto"].(int)
		oi.LocalIp = in["local_ip"].(string)
		oi.LocalPort = in["local_port"].(int)
		oi.PeerIp = in["peer_ip"].(string)
		oi.PeerPort = in["peer_port"].(int)
		oi.LocalSpi = in["local_spi"].(string)
		oi.RemoteSpi = in["remote_spi"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Mode = in["mode"].(string)
		oi.EncryptionAlgorithm = in["encryption_algorithm"].(string)
		oi.HashAlgorithm = in["hash_algorithm"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Lifebytes = in["lifebytes"].(string)
		oi.DhGroup = in["dh_group"].(int)
		oi.NatTraversal = in["nat_traversal"].(int)
		oi.AntiReplay = in["anti_replay"].(string)
		oi.PacketsEncrypted = in["packets_encrypted"].(int)
		oi.PacketsDecrypted = in["packets_decrypted"].(int)
		oi.AntiReplayNum = in["anti_replay_num"].(int)
		oi.RekeyNum = in["rekey_num"].(int)
		oi.PacketsErrInactive = in["packets_err_inactive"].(int)
		oi.PacketsErrEncryption = in["packets_err_encryption"].(int)
		oi.PacketsErrPadCheck = in["packets_err_pad_check"].(int)
		oi.PacketsErrPktSanity = in["packets_err_pkt_sanity"].(int)
		oi.PacketsErrIcvCheck = in["packets_err_icv_check"].(int)
		oi.PacketsErrLifetimeLifebytes = getObjectVpnIpsecOperOperSaListPacketsErrLifetimeLifebytes(in["packets_err_lifetime_lifebytes"].([]interface{}))
		oi.BytesEncrypted = in["bytes_encrypted"].(int)
		oi.BytesDecrypted = in["bytes_decrypted"].(int)
		oi.PrefragSuccess = in["prefrag_success"].(int)
		oi.PrefragError = in["prefrag_error"].(int)
		oi.CaviumBytesEncrypted = in["cavium_bytes_encrypted"].(int)
		oi.CaviumBytesDecrypted = in["cavium_bytes_decrypted"].(int)
		oi.CaviumPacketsEncrypted = in["cavium_packets_encrypted"].(int)
		oi.CaviumPacketsDecrypted = in["cavium_packets_decrypted"].(int)
		oi.QatBytesEncrypted = in["qat_bytes_encrypted"].(int)
		oi.QatBytesDecrypted = in["qat_bytes_decrypted"].(int)
		oi.QatPacketsEncrypted = in["qat_packets_encrypted"].(int)
		oi.QatPacketsDecrypted = in["qat_packets_decrypted"].(int)
		oi.TunnelIntfDown = in["tunnel_intf_down"].(int)
		oi.PktFailPrepToSend = in["pkt_fail_prep_to_send"].(int)
		oi.NoNextHop = in["no_next_hop"].(int)
		oi.InvalidTunnelId = in["invalid_tunnel_id"].(int)
		oi.NoTunnelFound = in["no_tunnel_found"].(int)
		oi.PktFailToSend = in["pkt_fail_to_send"].(int)
		oi.FragAfterEncapFragPackets = in["frag_after_encap_frag_packets"].(int)
		oi.FragReceived = in["frag_received"].(int)
		oi.SequenceNum = in["sequence_num"].(int)
		oi.SequenceNumRollover = in["sequence_num_rollover"].(int)
		oi.PacketsErrNhCheck = in["packets_err_nh_check"].(int)
		oi.EnforceTsEncapDrop = in["enforce_ts_encap_drop"].(int)
		oi.EnforceTsDecapDrop = in["enforce_ts_decap_drop"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIpsecOperOperSaListPacketsErrLifetimeLifebytes(d []interface{}) edpt.VpnIpsecOperOperSaListPacketsErrLifetimeLifebytes {

	var ret edpt.VpnIpsecOperOperSaListPacketsErrLifetimeLifebytes
	return ret
}

func dataToEndpointVpnIpsecOper(d *schema.ResourceData) edpt.VpnIpsecOper {
	var ret edpt.VpnIpsecOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectVpnIpsecOperOper(d.Get("oper").([]interface{}))
	return ret
}
