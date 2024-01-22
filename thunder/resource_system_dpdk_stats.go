package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDpdkStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_dpdk_stats`: List all DPDK error and drop counters\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDpdkStatsRead,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'pkt-drop': Total packet drop; 'pkt-lnk-down-drop': Total packets link down drop; 'err-pkt-drop': Total error packet drop; 'rx-err': Total RX packet error; 'tx-err': Total TX packet error; 'tx-drop': Total TX packet drop; 'rx-len-err': Total RX packet length error; 'rx-over-err': Total RX packet over error; 'rx-crc-err': Total RX packet CRC error; 'rx-frame-err': Total RX packet frame error; 'rx-no-buff-err': Total RX packet no buffer error; 'rx-miss-err': Total RX packet miss error; 'tx-abort-err': Total TX packet abort error; 'tx-carrier-err': Total TX packert carrier error; 'tx-fifo-err': Total TX packet fifo error; 'tx-hbeat-err': Total TX packet HBEAT error; 'tx-windows-err': Total TX windows error; 'rx-long-len-err': Total RX packet long length error; 'rx-short-len-err': Total RX packet short length error; 'rx-align-err': Total RX packet align error; 'rx-csum-offload-err': Total Rx packet check-sum offload error; 'io-rx-que-drop': Total IO core Rx queue drop; 'io-tx-que-drop': Total IO core TX queue drop; 'io-ring-drop': Total IO core ring drop; 'w-tx-que-drop': Total worker core queue drop; 'w-link-down-drop': Total worker core link down drop; 'w-ring-drop': Total worker core ring drop;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceSystemDpdkStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDpdkStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDpdkStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDpdkStatsSamplingEnable := setSliceSystemDpdkStatsSamplingEnable(res)
		d.Set("sampling_enable", SystemDpdkStatsSamplingEnable)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceSystemDpdkStatsSamplingEnable(d edpt.DataSystemDpdkStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSystemDpdkStats.SamplingEnable {
		in := make(map[string]interface{})
		in["counters1"] = item.Counters1
		result = append(result, in)
	}
	return result
}

func getSliceSystemDpdkStatsSamplingEnable(d []interface{}) []edpt.SystemDpdkStatsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemDpdkStatsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDpdkStatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemDpdkStats(d *schema.ResourceData) edpt.SystemDpdkStats {
	var ret edpt.SystemDpdkStats

	ret.SamplingEnable = getSliceSystemDpdkStatsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
