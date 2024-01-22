package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatAlgDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_alg_dns`: DNS ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatAlgDnsCreate,
		UpdateContext: resourceIpNatAlgDnsUpdate,
		ReadContext:   resourceIpNatAlgDnsRead,
		DeleteContext: resourceIpNatAlgDnsDelete,

		Schema: map[string]*schema.Schema{
			"dns_alg": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'disable': Disable DNS NAT ALG; 'enable': Enable DNS NAT ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpNatAlgDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatAlgDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatAlgDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatAlgDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatAlgDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatAlgDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatAlgDns(d *schema.ResourceData) edpt.IpNatAlgDns {
	var ret edpt.IpNatAlgDns
	ret.Inst.DnsAlg = d.Get("dns_alg").(string)
	//omit uuid
	return ret
}
