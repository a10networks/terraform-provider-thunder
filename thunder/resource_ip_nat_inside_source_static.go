package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatInsideSourceStatic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_inside_source_static`: Static Address Translations\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatInsideSourceStaticCreate,
		UpdateContext: resourceIpNatInsideSourceStaticUpdate,
		ReadContext:   resourceIpNatInsideSourceStaticRead,
		DeleteContext: resourceIpNatInsideSourceStaticDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable static mapping (default); 'disable': Disable static mapping;",
			},
			"nat_address": {
				Type: schema.TypeString, Required: true, Description: "NAT Address",
			},
			"src_address": {
				Type: schema.TypeString, Required: true, Description: "Original Source Address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceIpNatInsideSourceStaticCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceStaticCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceStatic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceStaticRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatInsideSourceStaticUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceStaticUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceStatic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceStaticRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatInsideSourceStaticDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceStaticDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceStatic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatInsideSourceStaticRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceStaticRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceStatic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatInsideSourceStatic(d *schema.ResourceData) edpt.IpNatInsideSourceStatic {
	var ret edpt.IpNatInsideSourceStatic
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.NatAddress = d.Get("nat_address").(string)
	ret.Inst.SrcAddress = d.Get("src_address").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
