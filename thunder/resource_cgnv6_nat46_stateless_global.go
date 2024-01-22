package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat46_stateless_global`: Stateless NAT46 Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat46StatelessGlobalCreate,
		UpdateContext: resourceCgnv6Nat46StatelessGlobalUpdate,
		ReadContext:   resourceCgnv6Nat46StatelessGlobalRead,
		DeleteContext: resourceCgnv6Nat46StatelessGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'outbound_ipv4_received': Outbound IPv4 packets received; 'outbound_ipv4_drop': Outbound IPv4 packets dropped; 'outbound_ipv4_fragment_received': Outbound IPv4 fragment packets received; 'outbound_ipv6_unreachable': Outbound IPv6 destination unreachable; 'outbound_ipv6_fragmented': Outbound IPv6 packets fragmented; 'inbound_ipv6_received': Inbound IPv6 packets received; 'inbound_ipv6_drop': Inbound IPv6 packets dropped; 'inbound_ipv6_fragment_received': Inbound IPv6 fragment packets received; 'inbound_ipv4_unreachable': Inbound IPv4 destination unreachable; 'inbound_ipv4_fragmented': Inbound IPv4 packets fragmented; 'packet_too_big': Packet too big; 'fragment_error': Fragment processing errors; 'icmpv6_to_icmp': ICMPv6 to ICMP; 'icmpv6_to_icmp_error': ICMPv6 to ICMP errors; 'icmp_to_icmpv6': ICMP to ICMPv6; 'icmp_to_icmpv6_error': ICMP to ICMPv6 errors; 'ha_standby': HA is standby; 'other_error': Other errors; 'conn_count': conn count;",
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
func resourceCgnv6Nat46StatelessGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat46StatelessGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat46StatelessGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat46StatelessGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Nat46StatelessGlobalSamplingEnable(d []interface{}) []edpt.Cgnv6Nat46StatelessGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Nat46StatelessGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Nat46StatelessGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Nat46StatelessGlobal(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessGlobal {
	var ret edpt.Cgnv6Nat46StatelessGlobal
	ret.Inst.SamplingEnable = getSliceCgnv6Nat46StatelessGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
