package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbaGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rba_group`: RBA configuration for a group\n\n__PLACEHOLDER__",
		CreateContext: resourceRbaGroupCreate,
		UpdateContext: resourceRbaGroupUpdate,
		ReadContext:   resourceRbaGroupRead,
		DeleteContext: resourceRbaGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of a RBA group",
			},
			"partition_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition_name": {
							Type: schema.TypeString, Required: true, Description: "partition name",
						},
						"role_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"role": {
										Type: schema.TypeString, Optional: true, Description: "Role in a given partition",
									},
								},
							},
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"user_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user": {
							Type: schema.TypeString, Optional: true, Description: "Users in the group",
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
func resourceRbaGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceRbaGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceRbaGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRbaGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRbaGroupPartitionList(d []interface{}) []edpt.RbaGroupPartitionList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.RoleList = getSliceRbaGroupPartitionListRoleList(in["role_list"].([]interface{}))
		oi.RuleList = getSliceRbaGroupPartitionListRuleList(in["rule_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupPartitionListRoleList(d []interface{}) []edpt.RbaGroupPartitionListRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupPartitionListRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupPartitionListRoleList
		oi.Role = in["role"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupPartitionListRuleList(d []interface{}) []edpt.RbaGroupPartitionListRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupPartitionListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupPartitionListRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupUserList(d []interface{}) []edpt.RbaGroupUserList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupUserList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupUserList
		oi.User = in["user"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRbaGroup(d *schema.ResourceData) edpt.RbaGroup {
	var ret edpt.RbaGroup
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PartitionList = getSliceRbaGroupPartitionList(d.Get("partition_list").([]interface{}))
	ret.Inst.UserList = getSliceRbaGroupUserList(d.Get("user_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
