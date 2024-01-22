package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nptv6Domain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nptv6_domain`: Configure NPTv6 translation domain\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nptv6DomainCreate,
		UpdateContext: resourceCgnv6Nptv6DomainUpdate,
		ReadContext:   resourceCgnv6Nptv6DomainRead,
		DeleteContext: resourceCgnv6Nptv6DomainDelete,

		Schema: map[string]*schema.Schema{
			"inside_prefix": {
				Type: schema.TypeString, Optional: true, Description: "Configure inside network prefix (Inside IPv6 network prefix)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of NPTv6 domain",
			},
			"outside_prefix": {
				Type: schema.TypeString, Optional: true, Description: "Configure outside network prefix (Outside IPv6 network prefix)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'outbound-packets': Outbound Packets; 'inbound-packets': Inbound Packets; 'hairpin-packets': Hairpin Packets; 'address-not-valid-for-translation': Address Not Valid For Translation; 'inbound-packets-no-map': Inbound Packets No Map; 'packets-dest-unreachable': Packets Destination Unreachable;",
						},
					},
				},
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
func resourceCgnv6Nptv6DomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6DomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Domain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nptv6DomainRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nptv6DomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6DomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Domain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nptv6DomainRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nptv6DomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6DomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Domain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nptv6DomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6DomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Domain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Nptv6DomainSamplingEnable(d []interface{}) []edpt.Cgnv6Nptv6DomainSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Nptv6DomainSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Nptv6DomainSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Nptv6Domain(d *schema.ResourceData) edpt.Cgnv6Nptv6Domain {
	var ret edpt.Cgnv6Nptv6Domain
	ret.Inst.InsidePrefix = d.Get("inside_prefix").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OutsidePrefix = d.Get("outside_prefix").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6Nptv6DomainSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
