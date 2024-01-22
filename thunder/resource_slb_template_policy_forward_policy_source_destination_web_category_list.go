package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_source_destination_web_category_list`: Configure web-category category-list for destination matching\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "Action to be performed",
			},
			"dual_stack_action": {
				Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': match URL;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"web_category_list": {
				Type: schema.TypeString, Required: true, Description: "Destination Web Category List Name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DualStackAction = d.Get("dual_stack_action").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.WebCategoryList = d.Get("web_category_list").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
