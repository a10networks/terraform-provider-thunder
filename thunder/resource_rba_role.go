package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbaRole() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rba_role`: Role configuration for RBA support\n\n__PLACEHOLDER__",
		CreateContext: resourceRbaRoleCreate,
		UpdateContext: resourceRbaRoleUpdate,
		ReadContext:   resourceRbaRoleRead,
		DeleteContext: resourceRbaRoleDelete,

		Schema: map[string]*schema.Schema{
			"default_privilege": {
				Type: schema.TypeString, Optional: true, Default: "no-access", Description: "'no-access': no-access; 'read': read; 'write': write;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name for the RBA role",
			},
			"partition_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Partition RBA Role",
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"object": {
							Type: schema.TypeString, Optional: true, Description: "Lineage of object class for permitted operation",
						},
						"operation": {
							Type: schema.TypeString, Optional: true, Description: "'no-access': no-access; 'read': read; 'oper': oper; 'write': write;",
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
func resourceRbaRoleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaRoleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaRole(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaRoleRead(ctx, d, meta)
	}
	return diags
}

func resourceRbaRoleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaRoleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaRole(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaRoleRead(ctx, d, meta)
	}
	return diags
}
func resourceRbaRoleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaRoleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaRole(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRbaRoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaRoleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaRole(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRbaRoleRuleList(d []interface{}) []edpt.RbaRoleRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaRoleRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaRoleRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRbaRole(d *schema.ResourceData) edpt.RbaRole {
	var ret edpt.RbaRole
	ret.Inst.DefaultPrivilege = d.Get("default_privilege").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PartitionOnly = d.Get("partition_only").(int)
	ret.Inst.RuleList = getSliceRbaRoleRuleList(d.Get("rule_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
