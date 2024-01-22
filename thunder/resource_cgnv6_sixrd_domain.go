package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SixrdDomain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sixrd_domain`: sixrd Domain\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SixrdDomainCreate,
		UpdateContext: resourceCgnv6SixrdDomainUpdate,
		ReadContext:   resourceCgnv6SixrdDomainRead,
		DeleteContext: resourceCgnv6SixrdDomainDelete,

		Schema: map[string]*schema.Schema{
			"br_ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "6rd BR IPv4 address",
			},
			"ce_ipv4_netmask": {
				Type: schema.TypeString, Optional: true, Description: "Mask length",
			},
			"ce_ipv4_network": {
				Type: schema.TypeString, Optional: true, Description: "Customer Edge IPv4 network",
			},
			"ipv6_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
			},
			"mtu": {
				Type: schema.TypeInt, Optional: true, Description: "Tunnel MTU",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "6rd Domain name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'outbound-tcp-packets-received': Outbound TCP packets received; 'outbound-udp-packets-received': Outbound UDP packets received; 'outbound-icmp-packets-received': Outbound ICMP packets received; 'outbound-other-packets-received': Outbound other packets received; 'outbound-packets-drop': Outbound packets dropped; 'outbound-ipv6-dest-unreachable': Outbound IPv6 destination unreachable; 'outbound-fragment-ipv6': Outbound Fragmented IPv6; 'inbound-tcp-packets-received': Inbound TCP packets received; 'inbound-udp-packets-received': Inbound UDP packets received; 'inbound-icmp-packets-received': Inbound ICMP packets received; 'inbound-other-packets-received': Inbound other packets received; 'inbound-packets-drop': Inbound packets dropped; 'inbound-ipv4-dest-unreachable': Inbound IPv4 destination unreachable; 'inbound-fragment-ipv4': Inbound Fragmented IPv4; 'inbound-tunnel-fragment-ipv6': Inbound Fragmented IPv6 in tunnel; 'vport-matched': Traffic match SLB virtual port; 'unknown-delegated-prefix': Unknown 6rd delegated prefix; 'packet-too-big': Packet too big; 'not-local-ip': Not local IP; 'fragment-error': Fragment processing errors; 'other-error': Other errors;",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SixrdDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdDomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdDomain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SixrdDomainRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SixrdDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdDomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdDomain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SixrdDomainRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SixrdDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdDomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdDomain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SixrdDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdDomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdDomain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6SixrdDomainSamplingEnable(d []interface{}) []edpt.Cgnv6SixrdDomainSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6SixrdDomainSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6SixrdDomainSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6SixrdDomain(d *schema.ResourceData) edpt.Cgnv6SixrdDomain {
	var ret edpt.Cgnv6SixrdDomain
	ret.Inst.BrIpv4Address = d.Get("br_ipv4_address").(string)
	ret.Inst.CeIpv4Netmask = d.Get("ce_ipv4_netmask").(string)
	ret.Inst.CeIpv4Network = d.Get("ce_ipv4_network").(string)
	ret.Inst.Ipv6Prefix = d.Get("ipv6_prefix").(string)
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6SixrdDomainSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
