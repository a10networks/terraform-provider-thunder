package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgMgcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_alg_mgcp`: NAT64 MGCP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64AlgMgcpCreate,
		UpdateContext: resourceCgnv6Nat64AlgMgcpUpdate,
		ReadContext:   resourceCgnv6Nat64AlgMgcpRead,
		DeleteContext: resourceCgnv6Nat64AlgMgcpDelete,

		Schema: map[string]*schema.Schema{
			"mgcp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable NAT64 MGCP ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat64AlgMgcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgMgcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgMgcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgMgcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64AlgMgcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgMgcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgMgcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgMgcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64AlgMgcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgMgcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgMgcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64AlgMgcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgMgcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgMgcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64AlgMgcp(d *schema.ResourceData) edpt.Cgnv6Nat64AlgMgcp {
	var ret edpt.Cgnv6Nat64AlgMgcp
	ret.Inst.MgcpEnable = d.Get("mgcp_enable").(string)
	//omit uuid
	return ret
}
