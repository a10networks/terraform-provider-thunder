package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_alg_sip`: NAT64 SIP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64AlgSipCreate,
		UpdateContext: resourceCgnv6Nat64AlgSipUpdate,
		ReadContext:   resourceCgnv6Nat64AlgSipRead,
		DeleteContext: resourceCgnv6Nat64AlgSipDelete,

		Schema: map[string]*schema.Schema{
			"sip_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable NAT64 SIP ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat64AlgSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgSipRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64AlgSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgSipRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64AlgSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64AlgSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64AlgSip(d *schema.ResourceData) edpt.Cgnv6Nat64AlgSip {
	var ret edpt.Cgnv6Nat64AlgSip
	ret.Inst.SipEnable = d.Get("sip_enable").(string)
	//omit uuid
	return ret
}
