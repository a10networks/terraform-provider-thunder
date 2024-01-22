package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6Address() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_address`: Transparent mode IPv6 Address\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6AddressCreate,
		UpdateContext: resourceIpv6AddressUpdate,
		ReadContext:   resourceIpv6AddressRead,
		DeleteContext: resourceIpv6AddressDelete,

		Schema: map[string]*schema.Schema{
			"anycast": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 anycast address",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
			},
			"link_local": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 link local address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6AddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Address(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6AddressRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6AddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Address(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6AddressRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6AddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Address(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6AddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Address(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6Address(d *schema.ResourceData) edpt.Ipv6Address {
	var ret edpt.Ipv6Address
	ret.Inst.Anycast = d.Get("anycast").(int)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.LinkLocal = d.Get("link_local").(int)
	//omit uuid
	return ret
}
