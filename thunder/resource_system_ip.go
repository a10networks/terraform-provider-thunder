package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip`: IP Setting\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpCreate,
		UpdateContext: resourceSystemIpUpdate,
		ReadContext:   resourceSystemIpRead,
		DeleteContext: resourceSystemIpDelete,

		Schema: map[string]*schema.Schema{
			"icmp_redirect_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmp redirect messages",
			},
			"icmp_unreachable_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmp unreachable messages",
			},
			"rpf_check_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable reverse path filter (strict mode)",
			},
			"source_route_pkt_drop_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv4 source routed packet drop",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIp(d *schema.ResourceData) edpt.SystemIp {
	var ret edpt.SystemIp
	ret.Inst.IcmpRedirectDisable = d.Get("icmp_redirect_disable").(int)
	ret.Inst.IcmpUnreachableDisable = d.Get("icmp_unreachable_disable").(int)
	ret.Inst.RpfCheckEnable = d.Get("rpf_check_enable").(int)
	ret.Inst.SourceRoutePktDropEnable = d.Get("source_route_pkt_drop_enable").(int)
	//omit uuid
	return ret
}
