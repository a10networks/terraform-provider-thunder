package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecSaStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ipsec_sa_stats`: IPsec sa statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIpsecSaStatsRead,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Required: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets-encrypted': Encrypted Packets; 'packets-decrypted': Decrypted Packets; 'anti-replay-num': Anti-Replay Failure; 'rekey-num': Rekey Times; 'packets-err-inactive': Inactive Error; 'packets-err-encryption': Encryption Error; 'packets-err-pad-check': Pad Check Error; 'packets-err-pkt-sanity': Packets Sanity Error; 'packets-err-icv-check': ICV Check Error; 'packets-err-lifetime-lifebytes': Lifetime Lifebytes Error; 'bytes-encrypted': Encrypted Bytes; 'bytes-decrypted': Decrypted Bytes; 'prefrag-success': Pre-frag Success; 'prefrag-error': Pre-frag Error; 'cavium-bytes-encrypted': CAVIUM Encrypted Bytes; 'cavium-bytes-decrypted': CAVIUM Decrypted Bytes; 'cavium-packets-encrypted': CAVIUM Encrypted Packets; 'cavium-packets-decrypted': CAVIUM Decrypted Packets; 'qat-bytes-encrypted': QAT Encrypted Bytes; 'qat-bytes-decrypted': QAT Decrypted Bytes; 'qat-packets-encrypted': QAT Encrypted Packets; 'qat-packets-decrypted': QAT Decrypted Packets; 'tunnel-intf-down': Packet dropped: Tunnel Interface Down; 'pkt-fail-prep-to-send': Packet dropped: Failed in prepare to send; 'no-next-hop': Packet dropped: No next hop; 'invalid-tunnel-id': Packet dropped: Invalid tunnel ID; 'no-tunnel-found': Packet dropped: No tunnel found; 'pkt-fail-to-send': Packet dropped: Failed to send; 'frag-after-encap-frag-packets': Frag-after-encap Fragment Generated; 'frag-received': Fragment Received; 'sequence-num': Sequence Number; 'sequence-num-rollover': Sequence Number Rollover; 'packets-err-nh-check': Next Header Check Error;",
						},
					},
				},
			},
		},
	}
}

func resourceVpnIpsecSaStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecSaStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecSaStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIpsecSaStatsSamplingEnable := setSliceVpnIpsecSaStatsSamplingEnable(res)
		d.Set("sampling_enable", VpnIpsecSaStatsSamplingEnable)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceVpnIpsecSaStatsSamplingEnable(d edpt.DataVpnIpsecSaStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVpnIpsecSaStats.SamplingEnable {
		in := make(map[string]interface{})
		in["counters1"] = item.Counters1
		result = append(result, in)
	}
	return result
}

func getSliceVpnIpsecSaStatsSamplingEnable(d []interface{}) []edpt.VpnIpsecSaStatsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaStatsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaStatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIpsecSaStats(d *schema.ResourceData) edpt.VpnIpsecSaStats {
	var ret edpt.VpnIpsecSaStats

	ret.SamplingEnable = getSliceVpnIpsecSaStatsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	return ret
}
