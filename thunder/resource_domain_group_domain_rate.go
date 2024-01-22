package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainGroupDomainRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_domain_group_domain_rate`: Configure Domain group per domain rate\n\n__PLACEHOLDER__",
		CreateContext: resourceDomainGroupDomainRateCreate,
		UpdateContext: resourceDomainGroupDomainRateUpdate,
		ReadContext:   resourceDomainGroupDomainRateRead,
		DeleteContext: resourceDomainGroupDomainRateDelete,

		Schema: map[string]*schema.Schema{
			"domain_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify name of the domain list",
						},
						"per_suffix_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Per suffix domain rate limit",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"dummy_name": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure domain group rate limit;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDomainGroupDomainRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainGroupDomainRateRead(ctx, d, meta)
	}
	return diags
}

func resourceDomainGroupDomainRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainGroupDomainRateRead(ctx, d, meta)
	}
	return diags
}
func resourceDomainGroupDomainRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDomainGroupDomainRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDomainGroupDomainRateDomainListList(d []interface{}) []edpt.DomainGroupDomainRateDomainListList {

	count1 := len(d)
	ret := make([]edpt.DomainGroupDomainRateDomainListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainGroupDomainRateDomainListList
		oi.Name = in["name"].(string)
		oi.PerSuffixRate = in["per_suffix_rate"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDomainGroupDomainRate(d *schema.ResourceData) edpt.DomainGroupDomainRate {
	var ret edpt.DomainGroupDomainRate
	ret.Inst.DomainListList = getSliceDomainGroupDomainRateDomainListList(d.Get("domain_list_list").([]interface{}))
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
