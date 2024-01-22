package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcBasedPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_based_policy`: Configure src-based policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcBasedPolicyCreate,
		UpdateContext: resourceDdosSrcBasedPolicyUpdate,
		ReadContext:   resourceDdosSrcBasedPolicyRead,
		DeleteContext: resourceDdosSrcBasedPolicyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
			},
			"policy_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceDdosSrcBasedPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcBasedPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcBasedPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcBasedPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcBasedPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcBasedPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcBasedPolicyPolicyClassListList(d []interface{}) []edpt.DdosSrcBasedPolicyPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcBasedPolicyPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcBasedPolicyPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSrcBasedPolicy(d *schema.ResourceData) edpt.DdosSrcBasedPolicy {
	var ret edpt.DdosSrcBasedPolicy
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PolicyClassListList = getSliceDdosSrcBasedPolicyPolicyClassListList(d.Get("policy_class_list_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
