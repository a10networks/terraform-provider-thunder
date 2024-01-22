package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosOutboundPolicyPolicyDefaultClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_outbound_policy_policy_default_class_list`: Configure default class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosOutboundPolicyPolicyDefaultClassListCreate,
		UpdateContext: resourceDdosOutboundPolicyPolicyDefaultClassListUpdate,
		ReadContext:   resourceDdosOutboundPolicyPolicyDefaultClassListRead,
		DeleteContext: resourceDdosOutboundPolicyPolicyDefaultClassListDelete,

		Schema: map[string]*schema.Schema{
			"class_list_glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"configuration": {
				Type: schema.TypeInt, Required: true, Description: "Default class-list configuration",
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
func resourceDdosOutboundPolicyPolicyDefaultClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyPolicyDefaultClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyPolicyDefaultClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyPolicyDefaultClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosOutboundPolicyPolicyDefaultClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyPolicyDefaultClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyPolicyDefaultClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyPolicyDefaultClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosOutboundPolicyPolicyDefaultClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyPolicyDefaultClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyPolicyDefaultClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosOutboundPolicyPolicyDefaultClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyPolicyDefaultClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyPolicyDefaultClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosOutboundPolicyPolicyDefaultClassList(d *schema.ResourceData) edpt.DdosOutboundPolicyPolicyDefaultClassList {
	var ret edpt.DdosOutboundPolicyPolicyDefaultClassList
	ret.Inst.ClassListGlid = d.Get("class_list_glid").(string)
	ret.Inst.Configuration = d.Get("configuration").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
