package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_alg_pptp`: NAT64 PPTP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64AlgPptpCreate,
		UpdateContext: resourceCgnv6Nat64AlgPptpUpdate,
		ReadContext:   resourceCgnv6Nat64AlgPptpRead,
		DeleteContext: resourceCgnv6Nat64AlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"pptp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable NAT64 PPTP ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat64AlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64AlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64AlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64AlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64AlgPptp(d *schema.ResourceData) edpt.Cgnv6Nat64AlgPptp {
	var ret edpt.Cgnv6Nat64AlgPptp
	ret.Inst.PptpEnable = d.Get("pptp_enable").(string)
	//omit uuid
	return ret
}
