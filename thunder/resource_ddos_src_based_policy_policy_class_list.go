package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcBasedPolicyPolicyClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_based_policy_policy_class_list`: Configure class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcBasedPolicyPolicyClassListCreate,
		UpdateContext: resourceDdosSrcBasedPolicyPolicyClassListUpdate,
		ReadContext:   resourceDdosSrcBasedPolicyPolicyClassListRead,
		DeleteContext: resourceDdosSrcBasedPolicyPolicyClassListDelete,

		Schema: map[string]*schema.Schema{
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
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
func resourceDdosSrcBasedPolicyPolicyClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyPolicyClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicyPolicyClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcBasedPolicyPolicyClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcBasedPolicyPolicyClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyPolicyClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicyPolicyClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcBasedPolicyPolicyClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcBasedPolicyPolicyClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyPolicyClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicyPolicyClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcBasedPolicyPolicyClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyPolicyClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicyPolicyClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosSrcBasedPolicyPolicyClassList(d *schema.ResourceData) edpt.DdosSrcBasedPolicyPolicyClassList {
	var ret edpt.DdosSrcBasedPolicyPolicyClassList
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
