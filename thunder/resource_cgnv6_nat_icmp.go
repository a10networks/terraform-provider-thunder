package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatIcmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_icmp`: Configure NAT ICMP settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatIcmpCreate,
		UpdateContext: resourceCgnv6NatIcmpUpdate,
		ReadContext:   resourceCgnv6NatIcmpRead,
		DeleteContext: resourceCgnv6NatIcmpDelete,

		Schema: map[string]*schema.Schema{
			"always_source_nat_errors": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source NAT intermediate routers' IPs for ICMP errors (default: disabled)",
			},
			"respond_to_ping": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond to ICMP echo requests to NAT pool IPs (default: disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6NatIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatIcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatIcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6NatIcmp(d *schema.ResourceData) edpt.Cgnv6NatIcmp {
	var ret edpt.Cgnv6NatIcmp
	ret.Inst.AlwaysSourceNatErrors = d.Get("always_source_nat_errors").(int)
	ret.Inst.RespondToPing = d.Get("respond_to_ping").(int)
	//omit uuid
	return ret
}
