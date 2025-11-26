package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ip_stats`: IP related statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIpStatsRead,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'inreceives': Incoming packets received; 'inhdrerrors': Incoming packet header errors; 'intoobigerrors': Incoming packet too big errors; 'innoroutes': Incoming no route packet drops; 'inaddrerrors': Incoming packet address errors; 'inunknownprotos': Incoming unkown protocol packet drops; 'intruncatedpkts': Incoming truncated packets; 'indiscards': Incoming packets discarded; 'indelivers': Incoming packets delivered; 'outforwdatagrams': Outgoing forwarded datagrams; 'outrequests': Outgoing packets; 'outdiscards': Outgoing packets discarded; 'outnoroutes': Outgoing no route packet drops; 'reasmtimeout': Reassembly timed out packet drops; 'reasmreqds': Incoming reassembly requests; 'reasmoks': Incoming reassembled packets; 'reasmfails': Incoming reassembly requests failed; 'fragoks': Outgoing packets fragmented; 'fragfails': Outgoing packets fragmentation failed; 'fragcreates': Outgoing fragmented packets; 'inmcastpkts': Incoming multicast packets; 'outmcastpkts': Outgoing multicast packets;",
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

func resourceSystemIpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIpStatsSamplingEnable := setSliceSystemIpStatsSamplingEnable(res)
		d.Set("sampling_enable", SystemIpStatsSamplingEnable)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceSystemIpStatsSamplingEnable(d edpt.DataSystemIpStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSystemIpStats.SamplingEnable {
		in := make(map[string]interface{})
		in["counters1"] = item.Counters1
		result = append(result, in)
	}
	return result
}

func getSliceSystemIpStatsSamplingEnable(d []interface{}) []edpt.SystemIpStatsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemIpStatsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpStatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpStats(d *schema.ResourceData) edpt.SystemIpStats {
	var ret edpt.SystemIpStats

	ret.SamplingEnable = getSliceSystemIpStatsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
