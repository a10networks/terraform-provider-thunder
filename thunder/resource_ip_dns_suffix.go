package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDnsSuffix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_dns_suffix`: DNS suffix\n\n__PLACEHOLDER__",
		CreateContext: resourceIpDnsSuffixCreate,
		UpdateContext: resourceIpDnsSuffixUpdate,
		ReadContext:   resourceIpDnsSuffixRead,
		DeleteContext: resourceIpDnsSuffixDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString, Optional: true, Description: "DNS suffix",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpDnsSuffixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSuffixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSuffix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDnsSuffixRead(ctx, d, meta)
	}
	return diags
}

func resourceIpDnsSuffixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSuffixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSuffix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDnsSuffixRead(ctx, d, meta)
	}
	return diags
}
func resourceIpDnsSuffixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSuffixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSuffix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpDnsSuffixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsSuffixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsSuffix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpDnsSuffix(d *schema.ResourceData) edpt.IpDnsSuffix {
	var ret edpt.IpDnsSuffix
	ret.Inst.DomainName = d.Get("domain_name").(string)
	//omit uuid
	return ret
}
