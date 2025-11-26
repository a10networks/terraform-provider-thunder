package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipv6`: IPv6 Setting\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpv6Create,
		UpdateContext: resourceSystemIpv6Update,
		ReadContext:   resourceSystemIpv6Read,
		DeleteContext: resourceSystemIpv6Delete,

		Schema: map[string]*schema.Schema{
			"icmpv6_redirect_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmpv6 redirect messages",
			},
			"icmpv6_unreachable_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmpv6 unreachable messages",
			},
			"rpf_check_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable reverse path filter (strict mode)",
			},
			"source_route_pkt_drop_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 source routed packet drop",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIpv6(d *schema.ResourceData) edpt.SystemIpv6 {
	var ret edpt.SystemIpv6
	ret.Inst.Icmpv6RedirectDisable = d.Get("icmpv6_redirect_disable").(int)
	ret.Inst.Icmpv6UnreachableDisable = d.Get("icmpv6_unreachable_disable").(int)
	ret.Inst.RpfCheckEnable = d.Get("rpf_check_enable").(int)
	ret.Inst.SourceRoutePktDropEnable = d.Get("source_route_pkt_drop_enable").(int)
	//omit uuid
	return ret
}
