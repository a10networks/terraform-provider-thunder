package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainGroupDomainRateDomainList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_domain_group_domain_rate_domain_list`: Per domain list domain rate configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDomainGroupDomainRateDomainListCreate,
		UpdateContext: resourceDomainGroupDomainRateDomainListUpdate,
		ReadContext:   resourceDomainGroupDomainRateDomainListRead,
		DeleteContext: resourceDomainGroupDomainRateDomainListDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the domain list",
			},
			"per_suffix_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Per suffix domain rate limit",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dummy_name": {
				Type: schema.TypeString, Required: true, Description: "DummyName",
			},
		},
	}
}
func resourceDomainGroupDomainRateDomainListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateDomainListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRateDomainList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainGroupDomainRateDomainListRead(ctx, d, meta)
	}
	return diags
}

func resourceDomainGroupDomainRateDomainListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateDomainListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRateDomainList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainGroupDomainRateDomainListRead(ctx, d, meta)
	}
	return diags
}
func resourceDomainGroupDomainRateDomainListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateDomainListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRateDomainList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDomainGroupDomainRateDomainListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDomainRateDomainListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupDomainRateDomainList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDomainGroupDomainRateDomainList(d *schema.ResourceData) edpt.DomainGroupDomainRateDomainList {
	var ret edpt.DomainGroupDomainRateDomainList
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PerSuffixRate = d.Get("per_suffix_rate").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	return ret
}
