package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSystemDefaultOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_system_default_oper`: Operational Status for the object system-default\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSystemDefaultOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_bit_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_bit_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_bit_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_bit_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_bit_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_pkt_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_pkt_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_pkt_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_pkt_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_pkt_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_r_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_r_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_r_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_r_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_conn_r_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_frag_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_frag_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_frag_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_frag_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_frag_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_bit_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_bit_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_bit_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_bit_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_bit_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_pkt_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_pkt_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_pkt_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_pkt_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_pkt_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_r_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_r_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_r_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_r_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_conn_r_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_frag_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_frag_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_frag_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_frag_icmp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_frag_other": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosSystemDefaultOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefaultOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSystemDefaultOperOper := setObjectDdosSystemDefaultOperOper(res)
		d.Set("oper", DdosSystemDefaultOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSystemDefaultOperOper(ret edpt.DataDdosSystemDefaultOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dst_bit_entry":    ret.DtDdosSystemDefaultOper.Oper.DstBitEntry,
			"dst_bit_udp":      ret.DtDdosSystemDefaultOper.Oper.DstBitUdp,
			"dst_bit_tcp":      ret.DtDdosSystemDefaultOper.Oper.DstBitTcp,
			"dst_bit_icmp":     ret.DtDdosSystemDefaultOper.Oper.DstBitIcmp,
			"dst_bit_other":    ret.DtDdosSystemDefaultOper.Oper.DstBitOther,
			"dst_pkt_entry":    ret.DtDdosSystemDefaultOper.Oper.DstPktEntry,
			"dst_pkt_udp":      ret.DtDdosSystemDefaultOper.Oper.DstPktUdp,
			"dst_pkt_tcp":      ret.DtDdosSystemDefaultOper.Oper.DstPktTcp,
			"dst_pkt_icmp":     ret.DtDdosSystemDefaultOper.Oper.DstPktIcmp,
			"dst_pkt_other":    ret.DtDdosSystemDefaultOper.Oper.DstPktOther,
			"dst_conn_entry":   ret.DtDdosSystemDefaultOper.Oper.DstConnEntry,
			"dst_conn_udp":     ret.DtDdosSystemDefaultOper.Oper.DstConnUdp,
			"dst_conn_tcp":     ret.DtDdosSystemDefaultOper.Oper.DstConnTcp,
			"dst_conn_icmp":    ret.DtDdosSystemDefaultOper.Oper.DstConnIcmp,
			"dst_conn_other":   ret.DtDdosSystemDefaultOper.Oper.DstConnOther,
			"dst_conn_r_entry": ret.DtDdosSystemDefaultOper.Oper.DstConnREntry,
			"dst_conn_r_udp":   ret.DtDdosSystemDefaultOper.Oper.DstConnRUdp,
			"dst_conn_r_tcp":   ret.DtDdosSystemDefaultOper.Oper.DstConnRTcp,
			"dst_conn_r_icmp":  ret.DtDdosSystemDefaultOper.Oper.DstConnRIcmp,
			"dst_conn_r_other": ret.DtDdosSystemDefaultOper.Oper.DstConnROther,
			"dst_frag_entry":   ret.DtDdosSystemDefaultOper.Oper.DstFragEntry,
			"dst_frag_udp":     ret.DtDdosSystemDefaultOper.Oper.DstFragUdp,
			"dst_frag_tcp":     ret.DtDdosSystemDefaultOper.Oper.DstFragTcp,
			"dst_frag_icmp":    ret.DtDdosSystemDefaultOper.Oper.DstFragIcmp,
			"dst_frag_other":   ret.DtDdosSystemDefaultOper.Oper.DstFragOther,
			"src_bit_entry":    ret.DtDdosSystemDefaultOper.Oper.SrcBitEntry,
			"src_bit_udp":      ret.DtDdosSystemDefaultOper.Oper.SrcBitUdp,
			"src_bit_tcp":      ret.DtDdosSystemDefaultOper.Oper.SrcBitTcp,
			"src_bit_icmp":     ret.DtDdosSystemDefaultOper.Oper.SrcBitIcmp,
			"src_bit_other":    ret.DtDdosSystemDefaultOper.Oper.SrcBitOther,
			"src_pkt_entry":    ret.DtDdosSystemDefaultOper.Oper.SrcPktEntry,
			"src_pkt_udp":      ret.DtDdosSystemDefaultOper.Oper.SrcPktUdp,
			"src_pkt_tcp":      ret.DtDdosSystemDefaultOper.Oper.SrcPktTcp,
			"src_pkt_icmp":     ret.DtDdosSystemDefaultOper.Oper.SrcPktIcmp,
			"src_pkt_other":    ret.DtDdosSystemDefaultOper.Oper.SrcPktOther,
			"src_conn_entry":   ret.DtDdosSystemDefaultOper.Oper.SrcConnEntry,
			"src_conn_udp":     ret.DtDdosSystemDefaultOper.Oper.SrcConnUdp,
			"src_conn_tcp":     ret.DtDdosSystemDefaultOper.Oper.SrcConnTcp,
			"src_conn_icmp":    ret.DtDdosSystemDefaultOper.Oper.SrcConnIcmp,
			"src_conn_other":   ret.DtDdosSystemDefaultOper.Oper.SrcConnOther,
			"src_conn_r_entry": ret.DtDdosSystemDefaultOper.Oper.SrcConnREntry,
			"src_conn_r_udp":   ret.DtDdosSystemDefaultOper.Oper.SrcConnRUdp,
			"src_conn_r_tcp":   ret.DtDdosSystemDefaultOper.Oper.SrcConnRTcp,
			"src_conn_r_icmp":  ret.DtDdosSystemDefaultOper.Oper.SrcConnRIcmp,
			"src_conn_r_other": ret.DtDdosSystemDefaultOper.Oper.SrcConnROther,
			"src_frag_entry":   ret.DtDdosSystemDefaultOper.Oper.SrcFragEntry,
			"src_frag_tcp":     ret.DtDdosSystemDefaultOper.Oper.SrcFragTcp,
			"src_frag_udp":     ret.DtDdosSystemDefaultOper.Oper.SrcFragUdp,
			"src_frag_icmp":    ret.DtDdosSystemDefaultOper.Oper.SrcFragIcmp,
			"src_frag_other":   ret.DtDdosSystemDefaultOper.Oper.SrcFragOther,
		},
	}
}

func getObjectDdosSystemDefaultOperOper(d []interface{}) edpt.DdosSystemDefaultOperOper {

	count1 := len(d)
	var ret edpt.DdosSystemDefaultOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstBitEntry = in["dst_bit_entry"].(string)
		ret.DstBitUdp = in["dst_bit_udp"].(string)
		ret.DstBitTcp = in["dst_bit_tcp"].(string)
		ret.DstBitIcmp = in["dst_bit_icmp"].(string)
		ret.DstBitOther = in["dst_bit_other"].(string)
		ret.DstPktEntry = in["dst_pkt_entry"].(string)
		ret.DstPktUdp = in["dst_pkt_udp"].(string)
		ret.DstPktTcp = in["dst_pkt_tcp"].(string)
		ret.DstPktIcmp = in["dst_pkt_icmp"].(string)
		ret.DstPktOther = in["dst_pkt_other"].(string)
		ret.DstConnEntry = in["dst_conn_entry"].(string)
		ret.DstConnUdp = in["dst_conn_udp"].(string)
		ret.DstConnTcp = in["dst_conn_tcp"].(string)
		ret.DstConnIcmp = in["dst_conn_icmp"].(string)
		ret.DstConnOther = in["dst_conn_other"].(string)
		ret.DstConnREntry = in["dst_conn_r_entry"].(string)
		ret.DstConnRUdp = in["dst_conn_r_udp"].(string)
		ret.DstConnRTcp = in["dst_conn_r_tcp"].(string)
		ret.DstConnRIcmp = in["dst_conn_r_icmp"].(string)
		ret.DstConnROther = in["dst_conn_r_other"].(string)
		ret.DstFragEntry = in["dst_frag_entry"].(string)
		ret.DstFragUdp = in["dst_frag_udp"].(string)
		ret.DstFragTcp = in["dst_frag_tcp"].(string)
		ret.DstFragIcmp = in["dst_frag_icmp"].(string)
		ret.DstFragOther = in["dst_frag_other"].(string)
		ret.SrcBitEntry = in["src_bit_entry"].(string)
		ret.SrcBitUdp = in["src_bit_udp"].(string)
		ret.SrcBitTcp = in["src_bit_tcp"].(string)
		ret.SrcBitIcmp = in["src_bit_icmp"].(string)
		ret.SrcBitOther = in["src_bit_other"].(string)
		ret.SrcPktEntry = in["src_pkt_entry"].(string)
		ret.SrcPktUdp = in["src_pkt_udp"].(string)
		ret.SrcPktTcp = in["src_pkt_tcp"].(string)
		ret.SrcPktIcmp = in["src_pkt_icmp"].(string)
		ret.SrcPktOther = in["src_pkt_other"].(string)
		ret.SrcConnEntry = in["src_conn_entry"].(string)
		ret.SrcConnUdp = in["src_conn_udp"].(string)
		ret.SrcConnTcp = in["src_conn_tcp"].(string)
		ret.SrcConnIcmp = in["src_conn_icmp"].(string)
		ret.SrcConnOther = in["src_conn_other"].(string)
		ret.SrcConnREntry = in["src_conn_r_entry"].(string)
		ret.SrcConnRUdp = in["src_conn_r_udp"].(string)
		ret.SrcConnRTcp = in["src_conn_r_tcp"].(string)
		ret.SrcConnRIcmp = in["src_conn_r_icmp"].(string)
		ret.SrcConnROther = in["src_conn_r_other"].(string)
		ret.SrcFragEntry = in["src_frag_entry"].(string)
		ret.SrcFragTcp = in["src_frag_tcp"].(string)
		ret.SrcFragUdp = in["src_frag_udp"].(string)
		ret.SrcFragIcmp = in["src_frag_icmp"].(string)
		ret.SrcFragOther = in["src_frag_other"].(string)
	}
	return ret
}

func dataToEndpointDdosSystemDefaultOper(d *schema.ResourceData) edpt.DdosSystemDefaultOper {
	var ret edpt.DdosSystemDefaultOper

	ret.Oper = getObjectDdosSystemDefaultOperOper(d.Get("oper").([]interface{}))
	return ret
}
