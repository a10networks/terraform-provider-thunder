package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDpdkStatsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_dpdk_stats_stats`: Statistics for the object dpdk-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDpdkStatsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet drop",
						},
						"pkt_lnk_down_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total packets link down drop",
						},
						"err_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total error packet drop",
						},
						"rx_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet error",
						},
						"tx_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX packet error",
						},
						"tx_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX packet drop",
						},
						"rx_len_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet length error",
						},
						"rx_over_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet over error",
						},
						"rx_crc_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet CRC error",
						},
						"rx_frame_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet frame error",
						},
						"rx_no_buff_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet no buffer error",
						},
						"rx_miss_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet miss error",
						},
						"tx_abort_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX packet abort error",
						},
						"tx_carrier_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX packert carrier error",
						},
						"tx_fifo_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX packet fifo error",
						},
						"tx_hbeat_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX packet HBEAT error",
						},
						"tx_windows_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TX windows error",
						},
						"rx_long_len_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet long length error",
						},
						"rx_short_len_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet short length error",
						},
						"rx_align_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total RX packet align error",
						},
						"rx_csum_offload_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total Rx packet check-sum offload error",
						},
						"io_rx_que_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total IO core Rx queue drop",
						},
						"io_tx_que_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total IO core TX queue drop",
						},
						"io_ring_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total IO core ring drop",
						},
						"w_tx_que_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total worker core queue drop",
						},
						"w_link_down_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total worker core link down drop",
						},
						"w_ring_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total worker core ring drop",
						},
					},
				},
			},
		},
	}
}

func resourceSystemDpdkStatsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDpdkStatsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDpdkStatsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDpdkStatsStatsStats := setObjectSystemDpdkStatsStatsStats(res)
		d.Set("stats", SystemDpdkStatsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemDpdkStatsStatsStats(ret edpt.DataSystemDpdkStatsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pkt_drop":            ret.DtSystemDpdkStatsStats.Stats.PktDrop,
			"pkt_lnk_down_drop":   ret.DtSystemDpdkStatsStats.Stats.PktLnkDownDrop,
			"err_pkt_drop":        ret.DtSystemDpdkStatsStats.Stats.ErrPktDrop,
			"rx_err":              ret.DtSystemDpdkStatsStats.Stats.RxErr,
			"tx_err":              ret.DtSystemDpdkStatsStats.Stats.TxErr,
			"tx_drop":             ret.DtSystemDpdkStatsStats.Stats.TxDrop,
			"rx_len_err":          ret.DtSystemDpdkStatsStats.Stats.RxLenErr,
			"rx_over_err":         ret.DtSystemDpdkStatsStats.Stats.RxOverErr,
			"rx_crc_err":          ret.DtSystemDpdkStatsStats.Stats.RxCrcErr,
			"rx_frame_err":        ret.DtSystemDpdkStatsStats.Stats.RxFrameErr,
			"rx_no_buff_err":      ret.DtSystemDpdkStatsStats.Stats.RxNoBuffErr,
			"rx_miss_err":         ret.DtSystemDpdkStatsStats.Stats.RxMissErr,
			"tx_abort_err":        ret.DtSystemDpdkStatsStats.Stats.TxAbortErr,
			"tx_carrier_err":      ret.DtSystemDpdkStatsStats.Stats.TxCarrierErr,
			"tx_fifo_err":         ret.DtSystemDpdkStatsStats.Stats.TxFifoErr,
			"tx_hbeat_err":        ret.DtSystemDpdkStatsStats.Stats.TxHbeatErr,
			"tx_windows_err":      ret.DtSystemDpdkStatsStats.Stats.TxWindowsErr,
			"rx_long_len_err":     ret.DtSystemDpdkStatsStats.Stats.RxLongLenErr,
			"rx_short_len_err":    ret.DtSystemDpdkStatsStats.Stats.RxShortLenErr,
			"rx_align_err":        ret.DtSystemDpdkStatsStats.Stats.RxAlignErr,
			"rx_csum_offload_err": ret.DtSystemDpdkStatsStats.Stats.RxCsumOffloadErr,
			"io_rx_que_drop":      ret.DtSystemDpdkStatsStats.Stats.IoRxQueDrop,
			"io_tx_que_drop":      ret.DtSystemDpdkStatsStats.Stats.IoTxQueDrop,
			"io_ring_drop":        ret.DtSystemDpdkStatsStats.Stats.IoRingDrop,
			"w_tx_que_drop":       ret.DtSystemDpdkStatsStats.Stats.WTxQueDrop,
			"w_link_down_drop":    ret.DtSystemDpdkStatsStats.Stats.WLinkDownDrop,
			"w_ring_drop":         ret.DtSystemDpdkStatsStats.Stats.WRingDrop,
		},
	}
}

func getObjectSystemDpdkStatsStatsStats(d []interface{}) edpt.SystemDpdkStatsStatsStats {

	count1 := len(d)
	var ret edpt.SystemDpdkStatsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktDrop = in["pkt_drop"].(int)
		ret.PktLnkDownDrop = in["pkt_lnk_down_drop"].(int)
		ret.ErrPktDrop = in["err_pkt_drop"].(int)
		ret.RxErr = in["rx_err"].(int)
		ret.TxErr = in["tx_err"].(int)
		ret.TxDrop = in["tx_drop"].(int)
		ret.RxLenErr = in["rx_len_err"].(int)
		ret.RxOverErr = in["rx_over_err"].(int)
		ret.RxCrcErr = in["rx_crc_err"].(int)
		ret.RxFrameErr = in["rx_frame_err"].(int)
		ret.RxNoBuffErr = in["rx_no_buff_err"].(int)
		ret.RxMissErr = in["rx_miss_err"].(int)
		ret.TxAbortErr = in["tx_abort_err"].(int)
		ret.TxCarrierErr = in["tx_carrier_err"].(int)
		ret.TxFifoErr = in["tx_fifo_err"].(int)
		ret.TxHbeatErr = in["tx_hbeat_err"].(int)
		ret.TxWindowsErr = in["tx_windows_err"].(int)
		ret.RxLongLenErr = in["rx_long_len_err"].(int)
		ret.RxShortLenErr = in["rx_short_len_err"].(int)
		ret.RxAlignErr = in["rx_align_err"].(int)
		ret.RxCsumOffloadErr = in["rx_csum_offload_err"].(int)
		ret.IoRxQueDrop = in["io_rx_que_drop"].(int)
		ret.IoTxQueDrop = in["io_tx_que_drop"].(int)
		ret.IoRingDrop = in["io_ring_drop"].(int)
		ret.WTxQueDrop = in["w_tx_que_drop"].(int)
		ret.WLinkDownDrop = in["w_link_down_drop"].(int)
		ret.WRingDrop = in["w_ring_drop"].(int)
	}
	return ret
}

func dataToEndpointSystemDpdkStatsStats(d *schema.ResourceData) edpt.SystemDpdkStatsStats {
	var ret edpt.SystemDpdkStatsStats

	ret.Stats = getObjectSystemDpdkStatsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
