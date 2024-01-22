package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats`: Configure triggers for system.dpdk-stats object\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsRead,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packet drop",
						},
						"pkt_lnk_down_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packets link down drop",
						},
						"err_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total error packet drop",
						},
						"rx_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet error",
						},
						"tx_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet error",
						},
						"tx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet drop",
						},
						"rx_len_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet length error",
						},
						"rx_over_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet over error",
						},
						"rx_crc_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet CRC error",
						},
						"rx_frame_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet frame error",
						},
						"rx_no_buff_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet no buffer error",
						},
						"rx_miss_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet miss error",
						},
						"tx_abort_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet abort error",
						},
						"tx_carrier_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packert carrier error",
						},
						"tx_fifo_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet fifo error",
						},
						"tx_hbeat_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet HBEAT error",
						},
						"tx_windows_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX windows error",
						},
						"rx_long_len_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet long length error",
						},
						"rx_short_len_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet short length error",
						},
						"rx_align_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet align error",
						},
						"rx_csum_offload_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Rx packet check-sum offload error",
						},
						"io_rx_que_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core Rx queue drop",
						},
						"io_tx_que_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core TX queue drop",
						},
						"io_ring_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core ring drop",
						},
						"w_tx_que_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core queue drop",
						},
						"w_link_down_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core link down drop",
						},
						"w_ring_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core ring drop",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_rate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_exceeded_by": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
						},
						"duration": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
						},
						"pkt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packet drop",
						},
						"pkt_lnk_down_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packets link down drop",
						},
						"err_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total error packet drop",
						},
						"rx_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet error",
						},
						"tx_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet error",
						},
						"tx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet drop",
						},
						"rx_len_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet length error",
						},
						"rx_over_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet over error",
						},
						"rx_crc_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet CRC error",
						},
						"rx_frame_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet frame error",
						},
						"rx_no_buff_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet no buffer error",
						},
						"rx_miss_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet miss error",
						},
						"tx_abort_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet abort error",
						},
						"tx_carrier_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packert carrier error",
						},
						"tx_fifo_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet fifo error",
						},
						"tx_hbeat_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet HBEAT error",
						},
						"tx_windows_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX windows error",
						},
						"rx_long_len_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet long length error",
						},
						"rx_short_len_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet short length error",
						},
						"rx_align_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet align error",
						},
						"rx_csum_offload_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Rx packet check-sum offload error",
						},
						"io_rx_que_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core Rx queue drop",
						},
						"io_tx_que_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core TX queue drop",
						},
						"io_ring_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core ring drop",
						},
						"w_tx_que_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core queue drop",
						},
						"w_link_down_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core link down drop",
						},
						"w_ring_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core ring drop",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091 := setObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091(res)
		d.Set("trigger_stats_inc", VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091)
		VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092 := setObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092(res)
		d.Set("trigger_stats_rate", VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091(ret edpt.DataVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pkt_drop":            ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.PktDrop,
			"pkt_lnk_down_drop":   ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.PktLnkDownDrop,
			"err_pkt_drop":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.ErrPktDrop,
			"rx_err":              ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxErr,
			"tx_err":              ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxErr,
			"tx_drop":             ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxDrop,
			"rx_len_err":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxLenErr,
			"rx_over_err":         ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxOverErr,
			"rx_crc_err":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxCrcErr,
			"rx_frame_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxFrameErr,
			"rx_no_buff_err":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxNoBuffErr,
			"rx_miss_err":         ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxMissErr,
			"tx_abort_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxAbortErr,
			"tx_carrier_err":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxCarrierErr,
			"tx_fifo_err":         ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxFifoErr,
			"tx_hbeat_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxHbeatErr,
			"tx_windows_err":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.TxWindowsErr,
			"rx_long_len_err":     ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxLongLenErr,
			"rx_short_len_err":    ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxShortLenErr,
			"rx_align_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxAlignErr,
			"rx_csum_offload_err": ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.RxCsumOffloadErr,
			"io_rx_que_drop":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.IoRxQueDrop,
			"io_tx_que_drop":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.IoTxQueDrop,
			"io_ring_drop":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.IoRingDrop,
			"w_tx_que_drop":       ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.WTxQueDrop,
			"w_link_down_drop":    ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.WLinkDownDrop,
			"w_ring_drop":         ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsInc.WRingDrop,
			//omit uuid
		},
	}
}

func setObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092(ret edpt.DataVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"threshold_exceeded_by": ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.ThresholdExceededBy,
			"duration":              ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.Duration,
			"pkt_drop":              ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.PktDrop,
			"pkt_lnk_down_drop":     ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.PktLnkDownDrop,
			"err_pkt_drop":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.ErrPktDrop,
			"rx_err":                ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxErr,
			"tx_err":                ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxErr,
			"tx_drop":               ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxDrop,
			"rx_len_err":            ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxLenErr,
			"rx_over_err":           ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxOverErr,
			"rx_crc_err":            ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxCrcErr,
			"rx_frame_err":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxFrameErr,
			"rx_no_buff_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxNoBuffErr,
			"rx_miss_err":           ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxMissErr,
			"tx_abort_err":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxAbortErr,
			"tx_carrier_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxCarrierErr,
			"tx_fifo_err":           ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxFifoErr,
			"tx_hbeat_err":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxHbeatErr,
			"tx_windows_err":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.TxWindowsErr,
			"rx_long_len_err":       ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxLongLenErr,
			"rx_short_len_err":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxShortLenErr,
			"rx_align_err":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxAlignErr,
			"rx_csum_offload_err":   ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.RxCsumOffloadErr,
			"io_rx_que_drop":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.IoRxQueDrop,
			"io_tx_que_drop":        ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.IoTxQueDrop,
			"io_ring_drop":          ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.IoRingDrop,
			"w_tx_que_drop":         ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.WTxQueDrop,
			"w_link_down_drop":      ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.WLinkDownDrop,
			"w_ring_drop":           ret.DtVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats.TriggerStatsRate.WRingDrop,
			//omit uuid
		},
	}
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
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
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats

	ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2091(d.Get("trigger_stats_inc").([]interface{}))

	ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2092(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid

	ret.Name = d.Get("name").(string)
	return ret
}
