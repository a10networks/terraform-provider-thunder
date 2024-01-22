package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnCgnv6NatPoolTopnTmplMetrics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_cgnv6_nat_pool_topn_tmpl_metrics`: Configure topn metrics for cgnv6.nat.pool\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsCreate,
		UpdateContext: resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsUpdate,
		ReadContext:   resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsRead,
		DeleteContext: resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsDelete,

		Schema: map[string]*schema.Schema{
			"tcp_total": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for TCP total",
			},
			"udp_total": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for UDP Total",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmplMetrics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmplMetrics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmplMetrics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplMetricsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmplMetrics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmplMetrics(d *schema.ResourceData) edpt.VisibilityTopnCgnv6NatPoolTopnTmplMetrics {
	var ret edpt.VisibilityTopnCgnv6NatPoolTopnTmplMetrics
	ret.Inst.TcpTotal = d.Get("tcp_total").(int)
	ret.Inst.UdpTotal = d.Get("udp_total").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
