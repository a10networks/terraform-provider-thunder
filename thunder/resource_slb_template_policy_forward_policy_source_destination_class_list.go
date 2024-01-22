package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_source_destination_class_list`: Configure class-list for destination matching\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "Action to be performed",
			},
			"dest_class_list": {
				Type: schema.TypeString, Required: true, Description: "Destination Class List Name",
			},
			"dual_stack_action": {
				Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': Match URL; 'ip': Match destination IP address;",
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
func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassList(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassList {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DestClassList = d.Get("dest_class_list").(string)
	ret.Inst.DualStackAction = d.Get("dual_stack_action").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
