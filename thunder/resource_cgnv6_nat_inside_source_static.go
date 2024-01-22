package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatInsideSourceStatic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_inside_source_static`: Static Address Translations\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatInsideSourceStaticCreate,
		UpdateContext: resourceCgnv6NatInsideSourceStaticUpdate,
		ReadContext:   resourceCgnv6NatInsideSourceStaticRead,
		DeleteContext: resourceCgnv6NatInsideSourceStaticDelete,

		Schema: map[string]*schema.Schema{
			"nat_address": {
				Type: schema.TypeString, Optional: true, Description: "NAT Address",
			},
			"partition": {
				Type: schema.TypeString, Required: true, Description: "Inside User Partition (Partition Name)",
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
func resourceCgnv6NatInsideSourceStaticCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatInsideSourceStaticCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatInsideSourceStatic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatInsideSourceStaticRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatInsideSourceStaticUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatInsideSourceStaticUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatInsideSourceStatic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatInsideSourceStaticRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatInsideSourceStaticDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatInsideSourceStaticDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatInsideSourceStatic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatInsideSourceStaticRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatInsideSourceStaticRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatInsideSourceStatic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6NatInsideSourceStatic(d *schema.ResourceData) edpt.Cgnv6NatInsideSourceStatic {
	var ret edpt.Cgnv6NatInsideSourceStatic
	ret.Inst.NatAddress = d.Get("nat_address").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.SrcAddress = d.Get("src_address").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
