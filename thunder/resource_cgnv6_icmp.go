package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Icmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_icmp`: CGNV6 ICMP Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6IcmpCreate,
		UpdateContext: resourceCgnv6IcmpUpdate,
		ReadContext:   resourceCgnv6IcmpRead,
		DeleteContext: resourceCgnv6IcmpDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'icmp-unknown-type': ICMP Unknown Type; 'icmp-no-port-info': ICMP Port Info Not Included; 'icmp-no-session-drop': ICMP No Matching Session Drop; 'icmpv6-unknown-type': ICMPv6 Unknown Type; 'icmpv6-no-port-info': ICMPv6 Port Info Not Included; 'icmpv6-no-session-drop': ICMPv6 No Matching Session Drop; 'icmp-to-icmp': ICMP to ICMP Conversion; 'icmp-to-icmpv6': ICMP to ICMPv6 Conversion; 'icmpv6-to-icmp': ICMPv6 to ICMP Conversion; 'icmpv6-to-icmpv6': ICMPv6 to ICMPv6 Conversion; 'icmp-bad-type': Bad Embedded ICMP Type; 'icmpv6-bad-type': Bad Embedded ICMPv6 Type; '64-known-drop': NAT64 Forward Known ICMPv6 Drop; '64-unknown-drop': NAT64 Forward Unknown ICMPv6 Drop; '64-midpoint-hop': NAT64 Forward Unknown Source Drop; '46-known-drop': NAT64 Reverse Known ICMP Drop; '46-unknown-drop': NAT64 Reverse Known ICMPv6 Drop; '46-no-prefix-for-ipv4': NAT64 Reverse No Prefix Match for IPv4; '46-bad-encap-ip-header-len': 4to6 Bad Encapsulated IP Header Length; 'icmp-to-icmp-err': ICMP to ICMP Conversion Error; 'icmp-to-icmpv6-err': ICMP to ICMPv6 Conversion Error; 'icmpv6-to-icmp-err': ICMPv6 to ICMP Conversion Error; 'icmpv6-to-icmpv6-err': ICMPv6 to ICMPv6 Conversion Error; 'encap-cross-cpu-no-match': ICMP Embedded Cross CPU No Matching Session; 'encap-cross-cpu-preprocess-err': ICMP Embedded Cross CPU Preprocess Error; 'icmp-to-icmp-unknown-l4': ICMP Embedded Unknown L4 Protocol; 'icmp-to-icmpv6-unknown-l4': ICMP to ICMPv6 Embedded Unknown L4 Protocol; 'icmpv6-to-icmp-unknown-l4': ICMPv6 to ICMP Embedded Unknown L4 Protocol; 'icmpv6-to-icmpv6-unknown-l4': ICMPv6 to ICMPv6 Embedded Unknown L4 Protocol; 'static-nat': ICMP Static NAT; 'echo-to-pool-reply': Ping to Pool Reply; 'echo-to-pool-drop': Ping to Pool Drop; 'error-to-pool-drop': Error to Pool Drop; 'echo-to-pool-reply-v6': Ping6 to Pool Reply; 'echo-to-pool-drop-v6': Ping6 to Pool Drop; 'error-to-pool-drop-v6': Error to IPv6 Pool Drop; 'error-ip-mismatch': ICMP IP address mismatch;",
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
func resourceCgnv6IcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6IcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Icmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6IcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6IcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6IcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Icmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6IcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6IcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6IcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Icmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6IcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6IcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Icmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6IcmpSamplingEnable(d []interface{}) []edpt.Cgnv6IcmpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6IcmpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6IcmpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Icmp(d *schema.ResourceData) edpt.Cgnv6Icmp {
	var ret edpt.Cgnv6Icmp
	ret.Inst.SamplingEnable = getSliceCgnv6IcmpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
