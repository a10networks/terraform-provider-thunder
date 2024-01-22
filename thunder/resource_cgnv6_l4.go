package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6L4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_l4`: CGNV6 L4 Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6L4Create,
		UpdateContext: resourceCgnv6L4Update,
		ReadContext:   resourceCgnv6L4Read,
		DeleteContext: resourceCgnv6L4Delete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'no-fwd-route': No Forward Route for Session; 'no-rev-route': No Reverse Route for Session; 'out-of-session-memory': Out of Session Memory; 'tcp-rst-sent': TCP RST Sent; 'ipip-icmp-reply-sent': IPIP ICMP Echo Reply Sent; 'icmp-filtered-sent': ICMP Administratively Filtered Sent; 'icmp-host-unreachable-sent': ICMP Host Unreachable Sent; 'icmp-reply-no-session-drop': ICMP Reply No Session Drop; 'ipip-truncated': IPIP Truncated Packet; 'ip-src-invalid-unicast': IPv4 Source Not Valid Unicast; 'ip-dst-invalid-unicast': IPv4 Destination Not Valid Unicast; 'ipv6-src-invalid-unicast': IPv6 Source Not Valid Unicast; 'ipv6-dst-invalid-unicast': IPv6 Destination Not Valid Unicast; 'rate_drop_reset_unkn': Rate Drop reset; 'bad-l3-protocol': Bad Layer 3 Protocol; 'special-ipv4-no-route': Stateless IPv4 No Forward Route; 'special-ipv6-no-route': Stateless IPv6 No Forward Route; 'icmp-reply-sent': ICMP Echo Reply Sent; 'icmpv6-reply-sent': ICMPv6 Echo Reply Sent; 'out-of-state-dropped': L4 Out of State packets; 'ttl-exceeded-sent': ICMP TTL Exceeded Sent; 'cross-cpu-alg-gre-no-match': ALG GRE Cross CPU No Matching Session; 'cross-cpu-alg-gre-preprocess-err': ALG GRE Cross CPU Preprocess Error; 'lsn-fast-setup': LSN Fast Setup Attempt; 'lsn-fast-setup-err': LSN Fast Setup Error; 'nat64-fast-setup': NAT64 Fast Setup Attempt; 'nat64-fast-setup-err': NAT64 Fast Setup Error; 'dslite-fast-setup': DS-Lite Fast Setup Attempt; 'dslite-fast-setup-err': DS-Lite Fast Setup Error; 'fast-setup-delayed-err': Fast Setup Delayed Error; 'fast-setup-mtu-too-small': Fast Setup MTU Too Small; 'fixed-nat44-fast-setup': Fixed NAT Fast Setup Attempt; 'fixed-nat44-fast-setup-err': Fixed NAT Fast Setup Error; 'fixed-nat64-fast-setup': Fixed NAT Fast Setup Attempt; 'fixed-nat64-fast-setup-err': Fixed NAT Fast Setup Error; 'fixed-nat-dslite-fast-setup': Fixed NAT Fast Setup Attempt; 'fixed-nat-dslite-fast-setup-err': Fixed NAT Fast Setup Error; 'fixed-nat-fast-setup-delayed-err': Fixed NAT Fast Setup Delayed Error; 'fixed-nat-fast-setup-mtu-too-small': Fixed NAT Fast Setup MTU Too Small; 'static-nat-fast-setup': Static NAT Fast Setup Attempt; 'static-nat-fast-setup-err': Static NAT Fast Setup Error; 'dst-nat-needed-drop': Destination NAT Needed Drop; 'invalid-nat64-translated-addr': Invalid NAT64 Translated IPv4 Address; 'tcp-rst-loop-drop': RST Loop Drop; 'static-nat-alloc': Static NAT Alloc; 'static-nat-free': Static NAT Free; 'process-l4': Process L4; 'preprocess-error': Preprocess Error; 'process-special': Process Special; 'process-continue': Process Continue; 'process-error': Process Error; 'fw-match-no-rule-drop': Firewall Matched No CGNv6 Rule Drop; 'ip-unknown-process': Process IP Unknown; 'src-nat-pool-not-found': Src NAT Pool Not Found; 'dst-nat-pool-not-found': Dst NAT Pool Not Found; 'l3-ip-src-invalid-unicast': IPv4 L3 Source Invalid Unicast; 'l3-ip-dst-invalid-unicast': IPv4 L3 Destination Invalid Unicast; 'l3-ipv6-src-invalid-unicast': IPv6 L3 Source Invalid Unicast; 'l3-ipv6-dst-invalid-unicast': IPv6 L3 Destination Invalid Unicast; 'fw-zone-mismatch-rerouting-drop': Rerouting Zone Mismatch Drop; 'nat-range-list-acl-deny': Nat range-list ACL deny; 'nat-range-list-acl-permit': Nat range-list ACL permit; 'fw-next-action-incorrect-drop': FW Next Action Incorrect Drop;",
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
func resourceCgnv6L4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6L4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6L4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6L4Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6L4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6L4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6L4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6L4Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6L4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6L4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6L4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6L4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6L4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6L4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6L4SamplingEnable(d []interface{}) []edpt.Cgnv6L4SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6L4SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6L4SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6L4(d *schema.ResourceData) edpt.Cgnv6L4 {
	var ret edpt.Cgnv6L4
	ret.Inst.SamplingEnable = getSliceCgnv6L4SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
