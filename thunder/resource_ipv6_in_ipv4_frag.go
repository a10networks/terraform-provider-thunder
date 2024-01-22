package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6InIpv4Frag() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_in_ipv4_frag`: IPv6-in-IPv4 fragmentation parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6InIpv4FragCreate,
		UpdateContext: resourceIpv6InIpv4FragUpdate,
		ReadContext:   resourceIpv6InIpv4FragRead,
		DeleteContext: resourceIpv6InIpv4FragDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-inserted': Session Inserted; 'session-expired': Session Expired; 'icmp-rcv': ICMP Received; 'icmpv6-rcv': ICMPv6 Received; 'udp-rcv': UDP Received; 'tcp-rcv': TCP Received; 'ipip-rcv': IP-in-IP Received; 'ipv6ip-rcv': IPv6-in-IP Received; 'other-rcv': Other Received; 'icmp-dropped': ICMP Dropped; 'icmpv6-dropped': ICMPv6 Dropped; 'udp-dropped': UDP Dropped; 'tcp-dropped': TCP Dropped; 'ipip-dropped': IP-in-IP Dropped; 'ipv6ip-dropped': IPv6-in-IP Dropped; 'other-dropped': Other Dropped; 'overlap-error': Overlapping Fragment Dropped; 'bad-ip-len': Bad IP Length; 'too-small': Fragment Too Small Drop; 'first-tcp-too-small': First TCP Fragment Too Small Drop; 'first-l4-too-small': First L4 Fragment Too Small Drop; 'total-sessions-exceeded': Total Sessions Exceeded Drop; 'no-session-memory': Out of Session Memory; 'fast-aging-set': Fragmentation Fast Aging Set; 'fast-aging-unset': Fragmentation Fast Aging Unset; 'fragment-queue-success': Fragment Queue Success; 'unaligned-len': Payload Length Unaligned; 'exceeded-len': Payload Length Out of Bounds; 'duplicate-first-frag': Duplicate First Fragment; 'duplicate-last-frag': Duplicate Last Fragment; 'total-fragments-exceeded': Total Queued Fragments Exceeded; 'fragment-queue-failure': Fragment Queue Failure; 'reassembly-success': Fragment Reassembly Success; 'max-len-exceeded': Fragment Max Data Length Exceeded; 'reassembly-failure': Fragment Reassembly Failure; 'policy-drop': MTU Exceeded Policy Drop; 'error-drop': Fragment Processing Drop; 'high-cpu-threshold': High CPU Threshold Reached; 'low-cpu-threshold': Low CPU Threshold Reached; 'cpu-threshold-drop': High CPU Drop; 'ipd-entry-drop': DDoS Protection Drop; 'max-packets-exceeded': Too Many Packets Per Reassembly Drop; 'session-packets-exceeded': Session Max Packets Exceeded; 'frag-session-count': Fragmentation Session Count; 'sctp-rcv': SCTP Received; 'sctp-dropped': SCTP Dropped; 'first-gtp-packet-too-small': First GTP Fragment Too Small Drop;",
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
func resourceIpv6InIpv4FragCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6InIpv4FragCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6InIpv4Frag(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6InIpv4FragRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6InIpv4FragUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6InIpv4FragUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6InIpv4Frag(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6InIpv4FragRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6InIpv4FragDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6InIpv4FragDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6InIpv4Frag(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6InIpv4FragRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6InIpv4FragRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6InIpv4Frag(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpv6InIpv4FragSamplingEnable(d []interface{}) []edpt.Ipv6InIpv4FragSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Ipv6InIpv4FragSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6InIpv4FragSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6InIpv4Frag(d *schema.ResourceData) edpt.Ipv6InIpv4Frag {
	var ret edpt.Ipv6InIpv4Frag
	ret.Inst.SamplingEnable = getSliceIpv6InIpv4FragSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
