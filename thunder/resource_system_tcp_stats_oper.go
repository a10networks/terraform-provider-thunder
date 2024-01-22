package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcpStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_tcp_stats_oper`: Operational Status for the object tcp-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTcpStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcpstats_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connattempt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connects": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conndrops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"segstimed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rttupdated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"delack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"timeoutdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rexmttimeo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"persisttimeo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"keeptimeo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"keepprobe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"keepdrops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndtotal": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndpack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndbyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndrexmitpack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndrexmitbyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndrexmitbad": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndacks": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndprobe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndurg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndwinup": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndctrl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvtotal": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvpack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvbyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvbadoff": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvmemdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reassmemdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reasstimeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvduppack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvdupbyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvpartduppack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvpartdupbyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvoopack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvoobyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvpackafterwin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvbyteafterwin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvwinprobe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvdupack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvacktoomuch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvackpack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvackbyte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvwinupd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pawsdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"predack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"preddat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"persistdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"badrst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"finwait2_drops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvdsack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sack_recovery_episode": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sack_rexmits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sack_rexmit_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sack_rcv_blocks": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sack_send_blocks": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndrst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndfin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndsyn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sndcack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cacklim": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cc_idle": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cc_reduce": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_iochan": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2btcpoptions": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2brcvwnd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2bsackpresent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2bdupack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2brxdata": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2boodata": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2bpartialack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2bfsmtransition": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2btransitionnum": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"b2atransitionnum": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"atcpforward": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"atcpsent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"atcpsendbackack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"atcprexmit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"atcpbuffallocfail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2bappbuffering": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"atcpsendfail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"earlyrexmit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mburstlim": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"a2bsndwnd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxyheaderv1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxyheaderv2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTcpStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTcpStatsOperOper := setObjectSystemTcpStatsOperOper(res)
		d.Set("oper", SystemTcpStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTcpStatsOperOper(ret edpt.DataSystemTcpStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcpstats_cpu_list": setSliceSystemTcpStatsOperOperTcpstatsCpuList(ret.DtSystemTcpStatsOper.Oper.TcpstatsCpuList),
			"cpu_count":         ret.DtSystemTcpStatsOper.Oper.CpuCount,
		},
	}
}

func setSliceSystemTcpStatsOperOperTcpstatsCpuList(d []edpt.SystemTcpStatsOperOperTcpstatsCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["connattempt"] = item.Connattempt
		in["connects"] = item.Connects
		in["drops"] = item.Drops
		in["conndrops"] = item.Conndrops
		in["closed"] = item.Closed
		in["segstimed"] = item.Segstimed
		in["rttupdated"] = item.Rttupdated
		in["delack"] = item.Delack
		in["timeoutdrop"] = item.Timeoutdrop
		in["rexmttimeo"] = item.Rexmttimeo
		in["persisttimeo"] = item.Persisttimeo
		in["keeptimeo"] = item.Keeptimeo
		in["keepprobe"] = item.Keepprobe
		in["keepdrops"] = item.Keepdrops
		in["sndtotal"] = item.Sndtotal
		in["sndpack"] = item.Sndpack
		in["sndbyte"] = item.Sndbyte
		in["sndrexmitpack"] = item.Sndrexmitpack
		in["sndrexmitbyte"] = item.Sndrexmitbyte
		in["sndrexmitbad"] = item.Sndrexmitbad
		in["sndacks"] = item.Sndacks
		in["sndprobe"] = item.Sndprobe
		in["sndurg"] = item.Sndurg
		in["sndwinup"] = item.Sndwinup
		in["sndctrl"] = item.Sndctrl
		in["rcvtotal"] = item.Rcvtotal
		in["rcvpack"] = item.Rcvpack
		in["rcvbyte"] = item.Rcvbyte
		in["rcvbadoff"] = item.Rcvbadoff
		in["rcvmemdrop"] = item.Rcvmemdrop
		in["reassmemdrop"] = item.Reassmemdrop
		in["reasstimeout"] = item.Reasstimeout
		in["rcvduppack"] = item.Rcvduppack
		in["rcvdupbyte"] = item.Rcvdupbyte
		in["rcvpartduppack"] = item.Rcvpartduppack
		in["rcvpartdupbyte"] = item.Rcvpartdupbyte
		in["rcvoopack"] = item.Rcvoopack
		in["rcvoobyte"] = item.Rcvoobyte
		in["rcvpackafterwin"] = item.Rcvpackafterwin
		in["rcvbyteafterwin"] = item.Rcvbyteafterwin
		in["rcvwinprobe"] = item.Rcvwinprobe
		in["rcvdupack"] = item.Rcvdupack
		in["rcvacktoomuch"] = item.Rcvacktoomuch
		in["rcvackpack"] = item.Rcvackpack
		in["rcvackbyte"] = item.Rcvackbyte
		in["rcvwinupd"] = item.Rcvwinupd
		in["pawsdrop"] = item.Pawsdrop
		in["predack"] = item.Predack
		in["preddat"] = item.Preddat
		in["persistdrop"] = item.Persistdrop
		in["badrst"] = item.Badrst
		in["finwait2_drops"] = item.Finwait2_drops
		in["rcvdsack"] = item.Rcvdsack
		in["sack_recovery_episode"] = item.Sack_recovery_episode
		in["sack_rexmits"] = item.Sack_rexmits
		in["sack_rexmit_bytes"] = item.Sack_rexmit_bytes
		in["sack_rcv_blocks"] = item.Sack_rcv_blocks
		in["sack_send_blocks"] = item.Sack_send_blocks
		in["sndrst"] = item.Sndrst
		in["sndfin"] = item.Sndfin
		in["sndsyn"] = item.Sndsyn
		in["sndcack"] = item.Sndcack
		in["cacklim"] = item.Cacklim
		in["cc_idle"] = item.Cc_idle
		in["cc_reduce"] = item.Cc_reduce
		in["bad_iochan"] = item.Bad_iochan
		in["a2btcpoptions"] = item.A2btcpoptions
		in["a2brcvwnd"] = item.A2brcvwnd
		in["a2bsackpresent"] = item.A2bsackpresent
		in["a2bdupack"] = item.A2bdupack
		in["a2brxdata"] = item.A2brxdata
		in["a2boodata"] = item.A2boodata
		in["a2bpartialack"] = item.A2bpartialack
		in["a2bfsmtransition"] = item.A2bfsmtransition
		in["a2btransitionnum"] = item.A2btransitionnum
		in["b2atransitionnum"] = item.B2atransitionnum
		in["atcpforward"] = item.Atcpforward
		in["atcpsent"] = item.Atcpsent
		in["atcpsendbackack"] = item.Atcpsendbackack
		in["atcprexmit"] = item.Atcprexmit
		in["atcpbuffallocfail"] = item.Atcpbuffallocfail
		in["a2bappbuffering"] = item.A2bappbuffering
		in["atcpsendfail"] = item.Atcpsendfail
		in["earlyrexmit"] = item.Earlyrexmit
		in["mburstlim"] = item.Mburstlim
		in["a2bsndwnd"] = item.A2bsndwnd
		in["proxyheaderv1"] = item.Proxyheaderv1
		in["proxyheaderv2"] = item.Proxyheaderv2
		result = append(result, in)
	}
	return result
}

func getObjectSystemTcpStatsOperOper(d []interface{}) edpt.SystemTcpStatsOperOper {

	count1 := len(d)
	var ret edpt.SystemTcpStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpstatsCpuList = getSliceSystemTcpStatsOperOperTcpstatsCpuList(in["tcpstats_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSystemTcpStatsOperOperTcpstatsCpuList(d []interface{}) []edpt.SystemTcpStatsOperOperTcpstatsCpuList {

	count1 := len(d)
	ret := make([]edpt.SystemTcpStatsOperOperTcpstatsCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpStatsOperOperTcpstatsCpuList
		oi.Connattempt = in["connattempt"].(int)
		oi.Connects = in["connects"].(int)
		oi.Drops = in["drops"].(int)
		oi.Conndrops = in["conndrops"].(int)
		oi.Closed = in["closed"].(int)
		oi.Segstimed = in["segstimed"].(int)
		oi.Rttupdated = in["rttupdated"].(int)
		oi.Delack = in["delack"].(int)
		oi.Timeoutdrop = in["timeoutdrop"].(int)
		oi.Rexmttimeo = in["rexmttimeo"].(int)
		oi.Persisttimeo = in["persisttimeo"].(int)
		oi.Keeptimeo = in["keeptimeo"].(int)
		oi.Keepprobe = in["keepprobe"].(int)
		oi.Keepdrops = in["keepdrops"].(int)
		oi.Sndtotal = in["sndtotal"].(int)
		oi.Sndpack = in["sndpack"].(int)
		oi.Sndbyte = in["sndbyte"].(int)
		oi.Sndrexmitpack = in["sndrexmitpack"].(int)
		oi.Sndrexmitbyte = in["sndrexmitbyte"].(int)
		oi.Sndrexmitbad = in["sndrexmitbad"].(int)
		oi.Sndacks = in["sndacks"].(int)
		oi.Sndprobe = in["sndprobe"].(int)
		oi.Sndurg = in["sndurg"].(int)
		oi.Sndwinup = in["sndwinup"].(int)
		oi.Sndctrl = in["sndctrl"].(int)
		oi.Rcvtotal = in["rcvtotal"].(int)
		oi.Rcvpack = in["rcvpack"].(int)
		oi.Rcvbyte = in["rcvbyte"].(int)
		oi.Rcvbadoff = in["rcvbadoff"].(int)
		oi.Rcvmemdrop = in["rcvmemdrop"].(int)
		oi.Reassmemdrop = in["reassmemdrop"].(int)
		oi.Reasstimeout = in["reasstimeout"].(int)
		oi.Rcvduppack = in["rcvduppack"].(int)
		oi.Rcvdupbyte = in["rcvdupbyte"].(int)
		oi.Rcvpartduppack = in["rcvpartduppack"].(int)
		oi.Rcvpartdupbyte = in["rcvpartdupbyte"].(int)
		oi.Rcvoopack = in["rcvoopack"].(int)
		oi.Rcvoobyte = in["rcvoobyte"].(int)
		oi.Rcvpackafterwin = in["rcvpackafterwin"].(int)
		oi.Rcvbyteafterwin = in["rcvbyteafterwin"].(int)
		oi.Rcvwinprobe = in["rcvwinprobe"].(int)
		oi.Rcvdupack = in["rcvdupack"].(int)
		oi.Rcvacktoomuch = in["rcvacktoomuch"].(int)
		oi.Rcvackpack = in["rcvackpack"].(int)
		oi.Rcvackbyte = in["rcvackbyte"].(int)
		oi.Rcvwinupd = in["rcvwinupd"].(int)
		oi.Pawsdrop = in["pawsdrop"].(int)
		oi.Predack = in["predack"].(int)
		oi.Preddat = in["preddat"].(int)
		oi.Persistdrop = in["persistdrop"].(int)
		oi.Badrst = in["badrst"].(int)
		oi.Finwait2_drops = in["finwait2_drops"].(int)
		oi.Rcvdsack = in["rcvdsack"].(int)
		oi.Sack_recovery_episode = in["sack_recovery_episode"].(int)
		oi.Sack_rexmits = in["sack_rexmits"].(int)
		oi.Sack_rexmit_bytes = in["sack_rexmit_bytes"].(int)
		oi.Sack_rcv_blocks = in["sack_rcv_blocks"].(int)
		oi.Sack_send_blocks = in["sack_send_blocks"].(int)
		oi.Sndrst = in["sndrst"].(int)
		oi.Sndfin = in["sndfin"].(int)
		oi.Sndsyn = in["sndsyn"].(int)
		oi.Sndcack = in["sndcack"].(int)
		oi.Cacklim = in["cacklim"].(int)
		oi.Cc_idle = in["cc_idle"].(int)
		oi.Cc_reduce = in["cc_reduce"].(int)
		oi.Bad_iochan = in["bad_iochan"].(int)
		oi.A2btcpoptions = in["a2btcpoptions"].(int)
		oi.A2brcvwnd = in["a2brcvwnd"].(int)
		oi.A2bsackpresent = in["a2bsackpresent"].(int)
		oi.A2bdupack = in["a2bdupack"].(int)
		oi.A2brxdata = in["a2brxdata"].(int)
		oi.A2boodata = in["a2boodata"].(int)
		oi.A2bpartialack = in["a2bpartialack"].(int)
		oi.A2bfsmtransition = in["a2bfsmtransition"].(int)
		oi.A2btransitionnum = in["a2btransitionnum"].(int)
		oi.B2atransitionnum = in["b2atransitionnum"].(int)
		oi.Atcpforward = in["atcpforward"].(int)
		oi.Atcpsent = in["atcpsent"].(int)
		oi.Atcpsendbackack = in["atcpsendbackack"].(int)
		oi.Atcprexmit = in["atcprexmit"].(int)
		oi.Atcpbuffallocfail = in["atcpbuffallocfail"].(int)
		oi.A2bappbuffering = in["a2bappbuffering"].(int)
		oi.Atcpsendfail = in["atcpsendfail"].(int)
		oi.Earlyrexmit = in["earlyrexmit"].(int)
		oi.Mburstlim = in["mburstlim"].(int)
		oi.A2bsndwnd = in["a2bsndwnd"].(int)
		oi.Proxyheaderv1 = in["proxyheaderv1"].(int)
		oi.Proxyheaderv2 = in["proxyheaderv2"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemTcpStatsOper(d *schema.ResourceData) edpt.SystemTcpStatsOper {
	var ret edpt.SystemTcpStatsOper

	ret.Oper = getObjectSystemTcpStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
