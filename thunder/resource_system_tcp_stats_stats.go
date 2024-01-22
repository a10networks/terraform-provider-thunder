package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcpStatsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_tcp_stats_stats`: Statistics for the object tcp-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTcpStatsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connattempt": {
							Type: schema.TypeInt, Optional: true, Description: "Connect initiated",
						},
						"connects": {
							Type: schema.TypeInt, Optional: true, Description: "Connect established",
						},
						"drops": {
							Type: schema.TypeInt, Optional: true, Description: "Connect dropped",
						},
						"conndrops": {
							Type: schema.TypeInt, Optional: true, Description: "Embryonic connect dropped",
						},
						"closed": {
							Type: schema.TypeInt, Optional: true, Description: "Connect closed",
						},
						"segstimed": {
							Type: schema.TypeInt, Optional: true, Description: "Segs to get RTT",
						},
						"rttupdated": {
							Type: schema.TypeInt, Optional: true, Description: "Update RTT",
						},
						"delack": {
							Type: schema.TypeInt, Optional: true, Description: "Delayed acks sent",
						},
						"timeoutdrop": {
							Type: schema.TypeInt, Optional: true, Description: "Conn dropped in rxmt timeout",
						},
						"rexmttimeo": {
							Type: schema.TypeInt, Optional: true, Description: "Retransmit timeout",
						},
						"persisttimeo": {
							Type: schema.TypeInt, Optional: true, Description: "Persist timeout",
						},
						"keeptimeo": {
							Type: schema.TypeInt, Optional: true, Description: "Keepalive timeout",
						},
						"keepprobe": {
							Type: schema.TypeInt, Optional: true, Description: "Keepalive probe sent",
						},
						"keepdrops": {
							Type: schema.TypeInt, Optional: true, Description: "Connect dropped in keepalive",
						},
						"sndtotal": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet sent",
						},
						"sndpack": {
							Type: schema.TypeInt, Optional: true, Description: "Data packet sent",
						},
						"sndbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Data bytes sent",
						},
						"sndrexmitpack": {
							Type: schema.TypeInt, Optional: true, Description: "Data packet retransmit",
						},
						"sndrexmitbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Data byte retransmit",
						},
						"sndrexmitbad": {
							Type: schema.TypeInt, Optional: true, Description: "Unnecessary packet retransmit",
						},
						"sndacks": {
							Type: schema.TypeInt, Optional: true, Description: "Ack packet sent",
						},
						"sndprobe": {
							Type: schema.TypeInt, Optional: true, Description: "Window probe sent",
						},
						"sndurg": {
							Type: schema.TypeInt, Optional: true, Description: "URG packet sent",
						},
						"sndwinup": {
							Type: schema.TypeInt, Optional: true, Description: "Window update packet sent",
						},
						"sndctrl": {
							Type: schema.TypeInt, Optional: true, Description: "SYN|FIN|RST packet sent",
						},
						"sndrst": {
							Type: schema.TypeInt, Optional: true, Description: "RST packet sent",
						},
						"sndfin": {
							Type: schema.TypeInt, Optional: true, Description: "FIN packet sent",
						},
						"sndsyn": {
							Type: schema.TypeInt, Optional: true, Description: "SYN packet sent",
						},
						"rcvtotal": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet received",
						},
						"rcvpack": {
							Type: schema.TypeInt, Optional: true, Description: "Packet received",
						},
						"rcvbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received",
						},
						"rcvbadoff": {
							Type: schema.TypeInt, Optional: true, Description: "Packet received with bad offset",
						},
						"rcvmemdrop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped for lack of memory",
						},
						"rcvduppack": {
							Type: schema.TypeInt, Optional: true, Description: "Duplicate packet received",
						},
						"rcvdupbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Duplicate bytes received",
						},
						"rcvpartduppack": {
							Type: schema.TypeInt, Optional: true, Description: "Packet with some duplicate data",
						},
						"rcvpartdupbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Dup. bytes in part-dup. packets",
						},
						"rcvoopack": {
							Type: schema.TypeInt, Optional: true, Description: "Out-of-order packet received",
						},
						"rcvoobyte": {
							Type: schema.TypeInt, Optional: true, Description: "Out-of-order bytes received",
						},
						"rcvpackafterwin": {
							Type: schema.TypeInt, Optional: true, Description: "Packets with data after window",
						},
						"rcvbyteafterwin": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes rcvd after window",
						},
						"rcvwinprobe": {
							Type: schema.TypeInt, Optional: true, Description: "Rcvd window probe packet",
						},
						"rcvdupack": {
							Type: schema.TypeInt, Optional: true, Description: "Rcvd duplicate acks",
						},
						"rcvacktoomuch": {
							Type: schema.TypeInt, Optional: true, Description: "Rcvd acks for unsent data",
						},
						"rcvackpack": {
							Type: schema.TypeInt, Optional: true, Description: "Rcvd ack packets",
						},
						"rcvackbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes acked by rcvd acks",
						},
						"rcvwinupd": {
							Type: schema.TypeInt, Optional: true, Description: "Rcvd window update packets",
						},
						"pawsdrop": {
							Type: schema.TypeInt, Optional: true, Description: "Segments dropped due to PAWS",
						},
						"predack": {
							Type: schema.TypeInt, Optional: true, Description: "Hdr predict for acks",
						},
						"preddat": {
							Type: schema.TypeInt, Optional: true, Description: "Hdr predict for data pkts",
						},
						"persistdrop": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout in persist state",
						},
						"badrst": {
							Type: schema.TypeInt, Optional: true, Description: "Ignored RST",
						},
						"finwait2_drops": {
							Type: schema.TypeInt, Optional: true, Description: "Drop FIN_WAIT_2 connection after time limit",
						},
						"sack_recovery_episode": {
							Type: schema.TypeInt, Optional: true, Description: "SACK recovery episodes",
						},
						"sack_rexmits": {
							Type: schema.TypeInt, Optional: true, Description: "SACK rexmit segments",
						},
						"sack_rexmit_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "SACK rexmit bytes",
						},
						"sack_rcv_blocks": {
							Type: schema.TypeInt, Optional: true, Description: "SACK received",
						},
						"sack_send_blocks": {
							Type: schema.TypeInt, Optional: true, Description: "SACK sent",
						},
						"sndcack": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge ACK sent",
						},
						"cacklim": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge ACK limited",
						},
						"reassmemdrop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet dropped during reassembly",
						},
						"reasstimeout": {
							Type: schema.TypeInt, Optional: true, Description: "Reassembly Time Out",
						},
						"cc_idle": {
							Type: schema.TypeInt, Optional: true, Description: "Congestion control window set do to idle",
						},
						"cc_reduce": {
							Type: schema.TypeInt, Optional: true, Description: "Congestion control window reduced by event",
						},
						"rcvdsack": {
							Type: schema.TypeInt, Optional: true, Description: "Rcvd DSACK packets",
						},
						"a2brcvwnd": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP receive window",
						},
						"a2bsackpresent": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP SACK options present",
						},
						"a2bdupack": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP Dup/OO ACK",
						},
						"a2brxdata": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP Rxmitted data",
						},
						"a2btcpoptions": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP unsupported TCP options",
						},
						"a2boodata": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP oo data received",
						},
						"a2bpartialack": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP partial ack received",
						},
						"a2bfsmtransition": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP state machine transition",
						},
						"a2btransitionnum": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP total transitions",
						},
						"b2atransitionnum": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP total transitions",
						},
						"bad_iochan": {
							Type: schema.TypeInt, Optional: true, Description: "IO Channel Modified",
						},
						"atcpforward": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP forward",
						},
						"atcpsent": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP sent",
						},
						"atcprexmitsadrop": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP transmit SA drops",
						},
						"atcpsendbackack": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP sendback ACK",
						},
						"atcprexmit": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP retransmits",
						},
						"atcpbuffallocfail": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP buffer allocation fails",
						},
						"a2bappbuffering": {
							Type: schema.TypeInt, Optional: true, Description: "Transition to full stack on when application buffers too much data",
						},
						"atcpsendfail": {
							Type: schema.TypeInt, Optional: true, Description: "Adaptive TCP sent fails",
						},
						"earlyrexmit": {
							Type: schema.TypeInt, Optional: true, Description: "Early Retransmission sent",
						},
						"mburstlim": {
							Type: schema.TypeInt, Optional: true, Description: "Maxburst limited tx",
						},
						"a2bsndwnd": {
							Type: schema.TypeInt, Optional: true, Description: "ATCP to BTCP send window",
						},
						"proxyheaderv1": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy header v1",
						},
						"proxyheaderv2": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy header v2",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTcpStatsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpStatsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpStatsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTcpStatsStatsStats := setObjectSystemTcpStatsStatsStats(res)
		d.Set("stats", SystemTcpStatsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTcpStatsStatsStats(ret edpt.DataSystemTcpStatsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"connattempt":           ret.DtSystemTcpStatsStats.Stats.Connattempt,
			"connects":              ret.DtSystemTcpStatsStats.Stats.Connects,
			"drops":                 ret.DtSystemTcpStatsStats.Stats.Drops,
			"conndrops":             ret.DtSystemTcpStatsStats.Stats.Conndrops,
			"closed":                ret.DtSystemTcpStatsStats.Stats.Closed,
			"segstimed":             ret.DtSystemTcpStatsStats.Stats.Segstimed,
			"rttupdated":            ret.DtSystemTcpStatsStats.Stats.Rttupdated,
			"delack":                ret.DtSystemTcpStatsStats.Stats.Delack,
			"timeoutdrop":           ret.DtSystemTcpStatsStats.Stats.Timeoutdrop,
			"rexmttimeo":            ret.DtSystemTcpStatsStats.Stats.Rexmttimeo,
			"persisttimeo":          ret.DtSystemTcpStatsStats.Stats.Persisttimeo,
			"keeptimeo":             ret.DtSystemTcpStatsStats.Stats.Keeptimeo,
			"keepprobe":             ret.DtSystemTcpStatsStats.Stats.Keepprobe,
			"keepdrops":             ret.DtSystemTcpStatsStats.Stats.Keepdrops,
			"sndtotal":              ret.DtSystemTcpStatsStats.Stats.Sndtotal,
			"sndpack":               ret.DtSystemTcpStatsStats.Stats.Sndpack,
			"sndbyte":               ret.DtSystemTcpStatsStats.Stats.Sndbyte,
			"sndrexmitpack":         ret.DtSystemTcpStatsStats.Stats.Sndrexmitpack,
			"sndrexmitbyte":         ret.DtSystemTcpStatsStats.Stats.Sndrexmitbyte,
			"sndrexmitbad":          ret.DtSystemTcpStatsStats.Stats.Sndrexmitbad,
			"sndacks":               ret.DtSystemTcpStatsStats.Stats.Sndacks,
			"sndprobe":              ret.DtSystemTcpStatsStats.Stats.Sndprobe,
			"sndurg":                ret.DtSystemTcpStatsStats.Stats.Sndurg,
			"sndwinup":              ret.DtSystemTcpStatsStats.Stats.Sndwinup,
			"sndctrl":               ret.DtSystemTcpStatsStats.Stats.Sndctrl,
			"sndrst":                ret.DtSystemTcpStatsStats.Stats.Sndrst,
			"sndfin":                ret.DtSystemTcpStatsStats.Stats.Sndfin,
			"sndsyn":                ret.DtSystemTcpStatsStats.Stats.Sndsyn,
			"rcvtotal":              ret.DtSystemTcpStatsStats.Stats.Rcvtotal,
			"rcvpack":               ret.DtSystemTcpStatsStats.Stats.Rcvpack,
			"rcvbyte":               ret.DtSystemTcpStatsStats.Stats.Rcvbyte,
			"rcvbadoff":             ret.DtSystemTcpStatsStats.Stats.Rcvbadoff,
			"rcvmemdrop":            ret.DtSystemTcpStatsStats.Stats.Rcvmemdrop,
			"rcvduppack":            ret.DtSystemTcpStatsStats.Stats.Rcvduppack,
			"rcvdupbyte":            ret.DtSystemTcpStatsStats.Stats.Rcvdupbyte,
			"rcvpartduppack":        ret.DtSystemTcpStatsStats.Stats.Rcvpartduppack,
			"rcvpartdupbyte":        ret.DtSystemTcpStatsStats.Stats.Rcvpartdupbyte,
			"rcvoopack":             ret.DtSystemTcpStatsStats.Stats.Rcvoopack,
			"rcvoobyte":             ret.DtSystemTcpStatsStats.Stats.Rcvoobyte,
			"rcvpackafterwin":       ret.DtSystemTcpStatsStats.Stats.Rcvpackafterwin,
			"rcvbyteafterwin":       ret.DtSystemTcpStatsStats.Stats.Rcvbyteafterwin,
			"rcvwinprobe":           ret.DtSystemTcpStatsStats.Stats.Rcvwinprobe,
			"rcvdupack":             ret.DtSystemTcpStatsStats.Stats.Rcvdupack,
			"rcvacktoomuch":         ret.DtSystemTcpStatsStats.Stats.Rcvacktoomuch,
			"rcvackpack":            ret.DtSystemTcpStatsStats.Stats.Rcvackpack,
			"rcvackbyte":            ret.DtSystemTcpStatsStats.Stats.Rcvackbyte,
			"rcvwinupd":             ret.DtSystemTcpStatsStats.Stats.Rcvwinupd,
			"pawsdrop":              ret.DtSystemTcpStatsStats.Stats.Pawsdrop,
			"predack":               ret.DtSystemTcpStatsStats.Stats.Predack,
			"preddat":               ret.DtSystemTcpStatsStats.Stats.Preddat,
			"persistdrop":           ret.DtSystemTcpStatsStats.Stats.Persistdrop,
			"badrst":                ret.DtSystemTcpStatsStats.Stats.Badrst,
			"finwait2_drops":        ret.DtSystemTcpStatsStats.Stats.Finwait2_drops,
			"sack_recovery_episode": ret.DtSystemTcpStatsStats.Stats.Sack_recovery_episode,
			"sack_rexmits":          ret.DtSystemTcpStatsStats.Stats.Sack_rexmits,
			"sack_rexmit_bytes":     ret.DtSystemTcpStatsStats.Stats.Sack_rexmit_bytes,
			"sack_rcv_blocks":       ret.DtSystemTcpStatsStats.Stats.Sack_rcv_blocks,
			"sack_send_blocks":      ret.DtSystemTcpStatsStats.Stats.Sack_send_blocks,
			"sndcack":               ret.DtSystemTcpStatsStats.Stats.Sndcack,
			"cacklim":               ret.DtSystemTcpStatsStats.Stats.Cacklim,
			"reassmemdrop":          ret.DtSystemTcpStatsStats.Stats.Reassmemdrop,
			"reasstimeout":          ret.DtSystemTcpStatsStats.Stats.Reasstimeout,
			"cc_idle":               ret.DtSystemTcpStatsStats.Stats.Cc_idle,
			"cc_reduce":             ret.DtSystemTcpStatsStats.Stats.Cc_reduce,
			"rcvdsack":              ret.DtSystemTcpStatsStats.Stats.Rcvdsack,
			"a2brcvwnd":             ret.DtSystemTcpStatsStats.Stats.A2brcvwnd,
			"a2bsackpresent":        ret.DtSystemTcpStatsStats.Stats.A2bsackpresent,
			"a2bdupack":             ret.DtSystemTcpStatsStats.Stats.A2bdupack,
			"a2brxdata":             ret.DtSystemTcpStatsStats.Stats.A2brxdata,
			"a2btcpoptions":         ret.DtSystemTcpStatsStats.Stats.A2btcpoptions,
			"a2boodata":             ret.DtSystemTcpStatsStats.Stats.A2boodata,
			"a2bpartialack":         ret.DtSystemTcpStatsStats.Stats.A2bpartialack,
			"a2bfsmtransition":      ret.DtSystemTcpStatsStats.Stats.A2bfsmtransition,
			"a2btransitionnum":      ret.DtSystemTcpStatsStats.Stats.A2btransitionnum,
			"b2atransitionnum":      ret.DtSystemTcpStatsStats.Stats.B2atransitionnum,
			"bad_iochan":            ret.DtSystemTcpStatsStats.Stats.Bad_iochan,
			"atcpforward":           ret.DtSystemTcpStatsStats.Stats.Atcpforward,
			"atcpsent":              ret.DtSystemTcpStatsStats.Stats.Atcpsent,
			"atcprexmitsadrop":      ret.DtSystemTcpStatsStats.Stats.Atcprexmitsadrop,
			"atcpsendbackack":       ret.DtSystemTcpStatsStats.Stats.Atcpsendbackack,
			"atcprexmit":            ret.DtSystemTcpStatsStats.Stats.Atcprexmit,
			"atcpbuffallocfail":     ret.DtSystemTcpStatsStats.Stats.Atcpbuffallocfail,
			"a2bappbuffering":       ret.DtSystemTcpStatsStats.Stats.A2bappbuffering,
			"atcpsendfail":          ret.DtSystemTcpStatsStats.Stats.Atcpsendfail,
			"earlyrexmit":           ret.DtSystemTcpStatsStats.Stats.Earlyrexmit,
			"mburstlim":             ret.DtSystemTcpStatsStats.Stats.Mburstlim,
			"a2bsndwnd":             ret.DtSystemTcpStatsStats.Stats.A2bsndwnd,
			"proxyheaderv1":         ret.DtSystemTcpStatsStats.Stats.Proxyheaderv1,
			"proxyheaderv2":         ret.DtSystemTcpStatsStats.Stats.Proxyheaderv2,
		},
	}
}

func getObjectSystemTcpStatsStatsStats(d []interface{}) edpt.SystemTcpStatsStatsStats {

	count1 := len(d)
	var ret edpt.SystemTcpStatsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connattempt = in["connattempt"].(int)
		ret.Connects = in["connects"].(int)
		ret.Drops = in["drops"].(int)
		ret.Conndrops = in["conndrops"].(int)
		ret.Closed = in["closed"].(int)
		ret.Segstimed = in["segstimed"].(int)
		ret.Rttupdated = in["rttupdated"].(int)
		ret.Delack = in["delack"].(int)
		ret.Timeoutdrop = in["timeoutdrop"].(int)
		ret.Rexmttimeo = in["rexmttimeo"].(int)
		ret.Persisttimeo = in["persisttimeo"].(int)
		ret.Keeptimeo = in["keeptimeo"].(int)
		ret.Keepprobe = in["keepprobe"].(int)
		ret.Keepdrops = in["keepdrops"].(int)
		ret.Sndtotal = in["sndtotal"].(int)
		ret.Sndpack = in["sndpack"].(int)
		ret.Sndbyte = in["sndbyte"].(int)
		ret.Sndrexmitpack = in["sndrexmitpack"].(int)
		ret.Sndrexmitbyte = in["sndrexmitbyte"].(int)
		ret.Sndrexmitbad = in["sndrexmitbad"].(int)
		ret.Sndacks = in["sndacks"].(int)
		ret.Sndprobe = in["sndprobe"].(int)
		ret.Sndurg = in["sndurg"].(int)
		ret.Sndwinup = in["sndwinup"].(int)
		ret.Sndctrl = in["sndctrl"].(int)
		ret.Sndrst = in["sndrst"].(int)
		ret.Sndfin = in["sndfin"].(int)
		ret.Sndsyn = in["sndsyn"].(int)
		ret.Rcvtotal = in["rcvtotal"].(int)
		ret.Rcvpack = in["rcvpack"].(int)
		ret.Rcvbyte = in["rcvbyte"].(int)
		ret.Rcvbadoff = in["rcvbadoff"].(int)
		ret.Rcvmemdrop = in["rcvmemdrop"].(int)
		ret.Rcvduppack = in["rcvduppack"].(int)
		ret.Rcvdupbyte = in["rcvdupbyte"].(int)
		ret.Rcvpartduppack = in["rcvpartduppack"].(int)
		ret.Rcvpartdupbyte = in["rcvpartdupbyte"].(int)
		ret.Rcvoopack = in["rcvoopack"].(int)
		ret.Rcvoobyte = in["rcvoobyte"].(int)
		ret.Rcvpackafterwin = in["rcvpackafterwin"].(int)
		ret.Rcvbyteafterwin = in["rcvbyteafterwin"].(int)
		ret.Rcvwinprobe = in["rcvwinprobe"].(int)
		ret.Rcvdupack = in["rcvdupack"].(int)
		ret.Rcvacktoomuch = in["rcvacktoomuch"].(int)
		ret.Rcvackpack = in["rcvackpack"].(int)
		ret.Rcvackbyte = in["rcvackbyte"].(int)
		ret.Rcvwinupd = in["rcvwinupd"].(int)
		ret.Pawsdrop = in["pawsdrop"].(int)
		ret.Predack = in["predack"].(int)
		ret.Preddat = in["preddat"].(int)
		ret.Persistdrop = in["persistdrop"].(int)
		ret.Badrst = in["badrst"].(int)
		ret.Finwait2_drops = in["finwait2_drops"].(int)
		ret.Sack_recovery_episode = in["sack_recovery_episode"].(int)
		ret.Sack_rexmits = in["sack_rexmits"].(int)
		ret.Sack_rexmit_bytes = in["sack_rexmit_bytes"].(int)
		ret.Sack_rcv_blocks = in["sack_rcv_blocks"].(int)
		ret.Sack_send_blocks = in["sack_send_blocks"].(int)
		ret.Sndcack = in["sndcack"].(int)
		ret.Cacklim = in["cacklim"].(int)
		ret.Reassmemdrop = in["reassmemdrop"].(int)
		ret.Reasstimeout = in["reasstimeout"].(int)
		ret.Cc_idle = in["cc_idle"].(int)
		ret.Cc_reduce = in["cc_reduce"].(int)
		ret.Rcvdsack = in["rcvdsack"].(int)
		ret.A2brcvwnd = in["a2brcvwnd"].(int)
		ret.A2bsackpresent = in["a2bsackpresent"].(int)
		ret.A2bdupack = in["a2bdupack"].(int)
		ret.A2brxdata = in["a2brxdata"].(int)
		ret.A2btcpoptions = in["a2btcpoptions"].(int)
		ret.A2boodata = in["a2boodata"].(int)
		ret.A2bpartialack = in["a2bpartialack"].(int)
		ret.A2bfsmtransition = in["a2bfsmtransition"].(int)
		ret.A2btransitionnum = in["a2btransitionnum"].(int)
		ret.B2atransitionnum = in["b2atransitionnum"].(int)
		ret.Bad_iochan = in["bad_iochan"].(int)
		ret.Atcpforward = in["atcpforward"].(int)
		ret.Atcpsent = in["atcpsent"].(int)
		ret.Atcprexmitsadrop = in["atcprexmitsadrop"].(int)
		ret.Atcpsendbackack = in["atcpsendbackack"].(int)
		ret.Atcprexmit = in["atcprexmit"].(int)
		ret.Atcpbuffallocfail = in["atcpbuffallocfail"].(int)
		ret.A2bappbuffering = in["a2bappbuffering"].(int)
		ret.Atcpsendfail = in["atcpsendfail"].(int)
		ret.Earlyrexmit = in["earlyrexmit"].(int)
		ret.Mburstlim = in["mburstlim"].(int)
		ret.A2bsndwnd = in["a2bsndwnd"].(int)
		ret.Proxyheaderv1 = in["proxyheaderv1"].(int)
		ret.Proxyheaderv2 = in["proxyheaderv2"].(int)
	}
	return ret
}

func dataToEndpointSystemTcpStatsStats(d *schema.ResourceData) edpt.SystemTcpStatsStats {
	var ret edpt.SystemTcpStatsStats

	ret.Stats = getObjectSystemTcpStatsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
