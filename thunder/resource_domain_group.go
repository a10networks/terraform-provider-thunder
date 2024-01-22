package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_domain_group`: Configure Domain group consisting of domain-lists\n\n__PLACEHOLDER__",
		CreateContext: resourceDomainGroupCreate,
		UpdateContext: resourceDomainGroupUpdate,
		ReadContext:   resourceDomainGroupRead,
		DeleteContext: resourceDomainGroupDelete,

		Schema: map[string]*schema.Schema{
			"domain_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify domain-list",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "permit", Description: "'permit': permit; 'deny': deny;",
						},
					},
				},
			},
			"domain_rate_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy_name": {
							Type: schema.TypeString, Required: true, Description: "'configuration': Configure domain group rate limit;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
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
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the domain group",
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
func resourceDomainGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceDomainGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceDomainGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDomainGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDomainGroupDomainGroupList(d []interface{}) []edpt.DomainGroupDomainGroupList {

	count1 := len(d)
	ret := make([]edpt.DomainGroupDomainGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainGroupDomainGroupList
		oi.DomainListName = in["domain_list_name"].(string)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainGroupDomainRateList(d []interface{}) []edpt.DomainGroupDomainRateList {

	count1 := len(d)
	ret := make([]edpt.DomainGroupDomainRateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainGroupDomainRateList
		oi.DummyName = in["dummy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.DomainListList = getSliceDomainGroupDomainRateListDomainListList(in["domain_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainGroupDomainRateListDomainListList(d []interface{}) []edpt.DomainGroupDomainRateListDomainListList {

	count1 := len(d)
	ret := make([]edpt.DomainGroupDomainRateListDomainListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainGroupDomainRateListDomainListList
		oi.Name = in["name"].(string)
		oi.PerSuffixRate = in["per_suffix_rate"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDomainGroup(d *schema.ResourceData) edpt.DomainGroup {
	var ret edpt.DomainGroup
	ret.Inst.DomainGroupList = getSliceDomainGroupDomainGroupList(d.Get("domain_group_list").([]interface{}))
	ret.Inst.DomainRateList = getSliceDomainGroupDomainRateList(d.Get("domain_rate_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
