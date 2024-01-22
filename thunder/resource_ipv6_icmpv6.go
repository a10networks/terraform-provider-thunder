package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6Icmpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_icmpv6`: Global ICMPv6 commands\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6Icmpv6Create,
		UpdateContext: resourceIpv6Icmpv6Update,
		ReadContext:   resourceIpv6Icmpv6Read,
		DeleteContext: resourceIpv6Icmpv6Delete,

		Schema: map[string]*schema.Schema{
			"redirect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable outbound ICMPv6 redirect messages",
			},
			"unreachable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable outbound ICMPv6 unreachable messages",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6Icmpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6Icmpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Icmpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6Icmpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceIpv6Icmpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6Icmpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Icmpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6Icmpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceIpv6Icmpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6Icmpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Icmpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6Icmpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6Icmpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Icmpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6Icmpv6(d *schema.ResourceData) edpt.Ipv6Icmpv6 {
	var ret edpt.Ipv6Icmpv6
	ret.Inst.Redirect = d.Get("redirect").(int)
	ret.Inst.Unreachable = d.Get("unreachable").(int)
	//omit uuid
	return ret
}
