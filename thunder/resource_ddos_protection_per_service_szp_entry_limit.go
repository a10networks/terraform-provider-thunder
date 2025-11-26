package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtectionPerServiceSzpEntryLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_protection_per_service_szp_entry_limit`: Global per service type src-zone-port entry limit\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosProtectionPerServiceSzpEntryLimitCreate,
		UpdateContext: resourceDdosProtectionPerServiceSzpEntryLimitUpdate,
		ReadContext:   resourceDdosProtectionPerServiceSzpEntryLimitRead,
		DeleteContext: resourceDdosProtectionPerServiceSzpEntryLimitDelete,

		Schema: map[string]*schema.Schema{
			"dns_tcp_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range dns-tcp",
			},
			"dns_udp_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range dns-udp",
			},
			"http_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range http",
			},
			"ip_proto_custom_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for custom ip-proto",
			},
			"ip_proto_gre_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto gre",
			},
			"ip_proto_icmp_v4_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto icmp-v4",
			},
			"ip_proto_icmp_v6_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto icmp-v6",
			},
			"ip_proto_ipv4_encap_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto ipv4-encap",
			},
			"ip_proto_ipv6_encap_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto ipv6-encap",
			},
			"ip_proto_other_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto other",
			},
			"quic_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range quic",
			},
			"sip_tcp_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range sip-tcp",
			},
			"sip_udp_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range sip-udp",
			},
			"ssl_l4_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range ssl-l4",
			},
			"tcp_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range tcp",
			},
			"udp_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range udp",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosProtectionPerServiceSzpEntryLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionPerServiceSzpEntryLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionPerServiceSzpEntryLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionPerServiceSzpEntryLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosProtectionPerServiceSzpEntryLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionPerServiceSzpEntryLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionPerServiceSzpEntryLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionPerServiceSzpEntryLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosProtectionPerServiceSzpEntryLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionPerServiceSzpEntryLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionPerServiceSzpEntryLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosProtectionPerServiceSzpEntryLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionPerServiceSzpEntryLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionPerServiceSzpEntryLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosProtectionPerServiceSzpEntryLimit(d *schema.ResourceData) edpt.DdosProtectionPerServiceSzpEntryLimit {
	var ret edpt.DdosProtectionPerServiceSzpEntryLimit
	ret.Inst.DnsTcpLimit = d.Get("dns_tcp_limit").(int)
	ret.Inst.DnsUdpLimit = d.Get("dns_udp_limit").(int)
	ret.Inst.HttpLimit = d.Get("http_limit").(int)
	ret.Inst.IpProtoCustomLimit = d.Get("ip_proto_custom_limit").(int)
	ret.Inst.IpProtoGreLimit = d.Get("ip_proto_gre_limit").(int)
	ret.Inst.IpProtoIcmpV4Limit = d.Get("ip_proto_icmp_v4_limit").(int)
	ret.Inst.IpProtoIcmpV6Limit = d.Get("ip_proto_icmp_v6_limit").(int)
	ret.Inst.IpProtoIpv4EncapLimit = d.Get("ip_proto_ipv4_encap_limit").(int)
	ret.Inst.IpProtoIpv6EncapLimit = d.Get("ip_proto_ipv6_encap_limit").(int)
	ret.Inst.IpProtoOtherLimit = d.Get("ip_proto_other_limit").(int)
	ret.Inst.QuicLimit = d.Get("quic_limit").(int)
	ret.Inst.SipTcpLimit = d.Get("sip_tcp_limit").(int)
	ret.Inst.SipUdpLimit = d.Get("sip_udp_limit").(int)
	ret.Inst.SslL4Limit = d.Get("ssl_l4_limit").(int)
	ret.Inst.TcpLimit = d.Get("tcp_limit").(int)
	ret.Inst.UdpLimit = d.Get("udp_limit").(int)
	//omit uuid
	return ret
}
