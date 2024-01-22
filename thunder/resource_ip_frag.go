package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpFrag() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_frag`: IP fragmentation parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceIpFragCreate,
		UpdateContext: resourceIpFragUpdate,
		ReadContext:   resourceIpFragRead,
		DeleteContext: resourceIpFragDelete,

		Schema: map[string]*schema.Schema{
			"buff": {
				Type: schema.TypeInt, Optional: true, Description: "Max buff used for fragmentation (Buffer Value(10000-3000000))",
			},
			"cpu_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": {
							Type: schema.TypeInt, Optional: true, Default: 75, Description: "When CPU usage reaches this value, it will stop processing fragments (default: 75%)",
						},
						"low": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "When CPU usage remains under this value, it will resume processing fragments (default: 60%)",
						},
					},
				},
			},
			"max_packets_per_reassembly": {
				Type: schema.TypeInt, Optional: true, Description: "Max number of fragmented packets allowed per reassembly(0 is unlimited) (default 0)",
			},
			"max_reassembly_sessions": {
				Type: schema.TypeInt, Optional: true, Description: "Max number of pending reassembly sessions allowed (default 100000)",
			},
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
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 60000, Description: "Fragmentation timeout (in milliseconds 4 - 65535 (default is 60000))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpFragCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpFragCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpFrag(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpFragRead(ctx, d, meta)
	}
	return diags
}

func resourceIpFragUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpFragUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpFrag(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpFragRead(ctx, d, meta)
	}
	return diags
}
func resourceIpFragDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpFragDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpFrag(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpFragRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpFragRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpFrag(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpFragCpuThreshold(d []interface{}) edpt.IpFragCpuThreshold {

	count1 := len(d)
	var ret edpt.IpFragCpuThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.High = in["high"].(int)
		ret.Low = in["low"].(int)
	}
	return ret
}

func getSliceIpFragSamplingEnable(d []interface{}) []edpt.IpFragSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.IpFragSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpFragSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpFrag(d *schema.ResourceData) edpt.IpFrag {
	var ret edpt.IpFrag
	ret.Inst.Buff = d.Get("buff").(int)
	ret.Inst.CpuThreshold = getObjectIpFragCpuThreshold(d.Get("cpu_threshold").([]interface{}))
	ret.Inst.MaxPacketsPerReassembly = d.Get("max_packets_per_reassembly").(int)
	ret.Inst.MaxReassemblySessions = d.Get("max_reassembly_sessions").(int)
	ret.Inst.SamplingEnable = getSliceIpFragSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
