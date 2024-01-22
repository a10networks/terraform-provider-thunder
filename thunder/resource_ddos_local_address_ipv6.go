package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosLocalAddressIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_local_address_ipv6`: Configure DDoS ipv6 address\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosLocalAddressIpv6Create,
		UpdateContext: resourceDdosLocalAddressIpv6Update,
		ReadContext:   resourceDdosLocalAddressIpv6Read,
		DeleteContext: resourceDdosLocalAddressIpv6Delete,

		Schema: map[string]*schema.Schema{
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "DDoS IPv6 Address for syn cookie usage",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosLocalAddressIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosLocalAddressIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosLocalAddressIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosLocalAddressIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosLocalAddressIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosLocalAddressIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosLocalAddressIpv6(d *schema.ResourceData) edpt.DdosLocalAddressIpv6 {
	var ret edpt.DdosLocalAddressIpv6
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	//omit uuid
	return ret
}
