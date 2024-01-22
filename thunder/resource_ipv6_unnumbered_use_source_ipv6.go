package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6UnnumberedUseSourceIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_unnumbered_use_source_ipv6`: Source IPv6 address\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6UnnumberedUseSourceIpv6Create,
		UpdateContext: resourceIpv6UnnumberedUseSourceIpv6Update,
		ReadContext:   resourceIpv6UnnumberedUseSourceIpv6Read,
		DeleteContext: resourceIpv6UnnumberedUseSourceIpv6Delete,

		Schema: map[string]*schema.Schema{
			"update_source_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6UnnumberedUseSourceIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedUseSourceIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6UnnumberedUseSourceIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6UnnumberedUseSourceIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceIpv6UnnumberedUseSourceIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedUseSourceIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6UnnumberedUseSourceIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6UnnumberedUseSourceIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceIpv6UnnumberedUseSourceIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedUseSourceIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6UnnumberedUseSourceIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6UnnumberedUseSourceIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedUseSourceIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6UnnumberedUseSourceIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6UnnumberedUseSourceIpv6(d *schema.ResourceData) edpt.Ipv6UnnumberedUseSourceIpv6 {
	var ret edpt.Ipv6UnnumberedUseSourceIpv6
	ret.Inst.UpdateSourceIpv6 = d.Get("update_source_ipv6").(string)
	//omit uuid
	return ret
}
