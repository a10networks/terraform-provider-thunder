package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDnsSecondary() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_dns_secondary`: Secondary DNS server\n\n__PLACEHOLDER__",
		CreateContext: resourceIpDnsSecondaryCreate,
		UpdateContext: resourceIpDnsSecondaryUpdate,
		ReadContext:   resourceIpDnsSecondaryRead,
		DeleteContext: resourceIpDnsSecondaryDelete,

		Schema: map[string]*schema.Schema{
			"ip_v4_addr": {
				Type: schema.TypeString, Optional: true, Description: "DNS server address",
			},
			"ip_v6_addr": {
				Type: schema.TypeString, Optional: true, Description: "DNS server address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpDnsSecondaryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSecondaryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSecondary(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDnsSecondaryRead(ctx, d, meta)
	}
	return diags
}

func resourceIpDnsSecondaryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSecondaryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSecondary(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDnsSecondaryRead(ctx, d, meta)
	}
	return diags
}
func resourceIpDnsSecondaryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSecondaryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSecondary(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpDnsSecondaryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSecondaryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSecondary(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpDnsSecondary(d *schema.ResourceData) edpt.IpDnsSecondary {
	var ret edpt.IpDnsSecondary
	ret.Inst.IpV4Addr = d.Get("ip_v4_addr").(string)
	ret.Inst.IpV6Addr = d.Get("ip_v6_addr").(string)
	//omit uuid
	return ret
}
