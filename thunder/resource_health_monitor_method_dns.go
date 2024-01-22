package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_dns`: DNS type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodDnsCreate,
		UpdateContext: resourceHealthMonitorMethodDnsUpdate,
		ReadContext:   resourceHealthMonitorMethodDnsRead,
		DeleteContext: resourceHealthMonitorMethodDnsDelete,

		Schema: map[string]*schema.Schema{
			"dns": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "DNS type",
			},
			"dns_domain": {
				Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name of the host",
			},
			"dns_domain_expect": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_domain_response": {
							Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))",
						},
						"dns_domain_fqdn": {
							Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name expected in DNS response answer",
						},
						"dns_domain_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Specify expected resolved IPv4 address",
						},
						"dns_domain_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify expected resolved IPv6 address",
						},
					},
				},
			},
			"dns_domain_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Specify DNS port, default is 53 (DNS Port(default 53))",
			},
			"dns_domain_recurse": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;",
			},
			"dns_domain_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DNS transport over TCP, default is UDP",
			},
			"dns_domain_type": {
				Type: schema.TypeString, Optional: true, Default: "A", Description: "'A': Used for storing Ipv4 address (default); 'CNAME': Canonical name for a DNS alias; 'SOA': Start of authority; 'PTR': Domain name pointer; 'MX': Mail exchanger; 'TXT': Text string; 'AAAA': Used for storing Ipv6 128-bits address;",
			},
			"dns_ip_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse DNS lookup (Specify IPv4 or IPv6 address)",
			},
			"dns_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Specify IPv4 address",
			},
			"dns_ipv4_expect": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_ipv4_response": {
							Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))",
						},
						"dns_ipv4_fqdn": {
							Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name expected in DNS response answer",
						},
					},
				},
			},
			"dns_ipv4_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Specify DNS port, default is 53 (DNS Port(default 53))",
			},
			"dns_ipv4_recurse": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;",
			},
			"dns_ipv4_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DNS transport over TCP, default is UDP",
			},
			"dns_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
			},
			"dns_ipv6_expect": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_ipv6_response": {
							Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))",
						},
						"dns_ipv6_fqdn": {
							Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name expected in DNS response answer",
						},
					},
				},
			},
			"dns_ipv6_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Specify DNS port, default is 53 (DNS Port(default 53))",
			},
			"dns_ipv6_recurse": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;",
			},
			"dns_ipv6_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DNS transport over TCP, default is UDP",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethodDnsDnsDomainExpect(d []interface{}) edpt.HealthMonitorMethodDnsDnsDomainExpect {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodDnsDnsDomainExpect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsDomainResponse = in["dns_domain_response"].(string)
		ret.DnsDomainFqdn = in["dns_domain_fqdn"].(string)
		ret.DnsDomainIpv4 = in["dns_domain_ipv4"].(string)
		ret.DnsDomainIpv6 = in["dns_domain_ipv6"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodDnsDnsIpv4Expect(d []interface{}) edpt.HealthMonitorMethodDnsDnsIpv4Expect {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodDnsDnsIpv4Expect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsIpv4Response = in["dns_ipv4_response"].(string)
		ret.DnsIpv4Fqdn = in["dns_ipv4_fqdn"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodDnsDnsIpv6Expect(d []interface{}) edpt.HealthMonitorMethodDnsDnsIpv6Expect {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodDnsDnsIpv6Expect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsIpv6Response = in["dns_ipv6_response"].(string)
		ret.DnsIpv6Fqdn = in["dns_ipv6_fqdn"].(string)
	}
	return ret
}

func dataToEndpointHealthMonitorMethodDns(d *schema.ResourceData) edpt.HealthMonitorMethodDns {
	var ret edpt.HealthMonitorMethodDns
	ret.Inst.Dns = d.Get("dns").(int)
	ret.Inst.DnsDomain = d.Get("dns_domain").(string)
	ret.Inst.DnsDomainExpect = getObjectHealthMonitorMethodDnsDnsDomainExpect(d.Get("dns_domain_expect").([]interface{}))
	ret.Inst.DnsDomainPort = d.Get("dns_domain_port").(int)
	ret.Inst.DnsDomainRecurse = d.Get("dns_domain_recurse").(string)
	ret.Inst.DnsDomainTcp = d.Get("dns_domain_tcp").(int)
	ret.Inst.DnsDomainType = d.Get("dns_domain_type").(string)
	ret.Inst.DnsIpKey = d.Get("dns_ip_key").(int)
	ret.Inst.DnsIpv4Addr = d.Get("dns_ipv4_addr").(string)
	ret.Inst.DnsIpv4Expect = getObjectHealthMonitorMethodDnsDnsIpv4Expect(d.Get("dns_ipv4_expect").([]interface{}))
	ret.Inst.DnsIpv4Port = d.Get("dns_ipv4_port").(int)
	ret.Inst.DnsIpv4Recurse = d.Get("dns_ipv4_recurse").(string)
	ret.Inst.DnsIpv4Tcp = d.Get("dns_ipv4_tcp").(int)
	ret.Inst.DnsIpv6Addr = d.Get("dns_ipv6_addr").(string)
	ret.Inst.DnsIpv6Expect = getObjectHealthMonitorMethodDnsDnsIpv6Expect(d.Get("dns_ipv6_expect").([]interface{}))
	ret.Inst.DnsIpv6Port = d.Get("dns_ipv6_port").(int)
	ret.Inst.DnsIpv6Recurse = d.Get("dns_ipv6_recurse").(string)
	ret.Inst.DnsIpv6Tcp = d.Get("dns_ipv6_tcp").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
