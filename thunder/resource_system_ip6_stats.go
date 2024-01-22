package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIp6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ip6_stats`: IPv6 related statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIp6StatsRead,

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

func resourceSystemIp6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIp6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIp6StatsSamplingEnable := setSliceSystemIp6StatsSamplingEnable(res)
		d.Set("sampling_enable", SystemIp6StatsSamplingEnable)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceSystemIp6StatsSamplingEnable(d edpt.DataSystemIp6Stats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSystemIp6Stats.SamplingEnable {
		in := make(map[string]interface{})
		in["counters1"] = item.Counters1
		result = append(result, in)
	}
	return result
}

func getSliceSystemIp6StatsSamplingEnable(d []interface{}) []edpt.SystemIp6StatsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemIp6StatsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIp6StatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIp6Stats(d *schema.ResourceData) edpt.SystemIp6Stats {
	var ret edpt.SystemIp6Stats

	ret.SamplingEnable = getSliceSystemIp6StatsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
