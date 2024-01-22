package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnCgnv6NatPoolTopnTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_cgnv6_nat_pool_topn_tmpl`: Configure template for cgnv6.nat.pool\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnCgnv6NatPoolTopnTmplCreate,
		UpdateContext: resourceVisibilityTopnCgnv6NatPoolTopnTmplUpdate,
		ReadContext:   resourceVisibilityTopnCgnv6NatPoolTopnTmplRead,
		DeleteContext: resourceVisibilityTopnCgnv6NatPoolTopnTmplDelete,

		Schema: map[string]*schema.Schema{
			"interval": {
				Type: schema.TypeString, Optional: true, Description: "'5': 5 minutes; '15': 15 minutes; '30': 30 minutes; '60': 60 minutes; 'all-time': Since template is activated;",
			},
			"metrics": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for UDP Total",
						},
						"tcp_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for TCP total",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Template Name",
			},
			"topn_size": {
				Type: schema.TypeInt, Optional: true, Description: "Congure value of N for topn",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityTopnCgnv6NatPoolTopnTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnCgnv6NatPoolTopnTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnCgnv6NatPoolTopnTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnCgnv6NatPoolTopnTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnCgnv6NatPoolTopnTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnCgnv6NatPoolTopnTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityTopnCgnv6NatPoolTopnTmplMetrics3131(d []interface{}) edpt.VisibilityTopnCgnv6NatPoolTopnTmplMetrics3131 {

	count1 := len(d)
	var ret edpt.VisibilityTopnCgnv6NatPoolTopnTmplMetrics3131
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpTotal = in["udp_total"].(int)
		ret.TcpTotal = in["tcp_total"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityTopnCgnv6NatPoolTopnTmpl(d *schema.ResourceData) edpt.VisibilityTopnCgnv6NatPoolTopnTmpl {
	var ret edpt.VisibilityTopnCgnv6NatPoolTopnTmpl
	ret.Inst.Interval = d.Get("interval").(string)
	ret.Inst.Metrics = getObjectVisibilityTopnCgnv6NatPoolTopnTmplMetrics3131(d.Get("metrics").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TopnSize = d.Get("topn_size").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
