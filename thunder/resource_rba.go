package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRba() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rba`: Role Based Access configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceRbaCreate,
		UpdateContext: resourceRbaUpdate,
		ReadContext:   resourceRbaRead,
		DeleteContext: resourceRbaDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable RBA; 'disable': Disable RBA;",
			},
			"group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Name of a RBA group",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"role_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Name for the RBA role",
						},
						"default_privilege": {
							Type: schema.TypeString, Optional: true, Default: "no-access", Description: "'no-access': no-access; 'read': read; 'write': write;",
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
						"name": {
							Type: schema.TypeString, Required: true, Description: "Name of a user account",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRbaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRba(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaRead(ctx, d, meta)
	}
	return diags
}

func resourceRbaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRba(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaRead(ctx, d, meta)
	}
	return diags
}
func resourceRbaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRba(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRbaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRba(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRbaGroupList(d []interface{}) []edpt.RbaGroupList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupList
		oi.Name = in["name"].(string)
		oi.UserList = getSliceRbaGroupListUserList(in["user_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PartitionList = getSliceRbaGroupListPartitionList(in["partition_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupListUserList(d []interface{}) []edpt.RbaGroupListUserList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupListUserList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupListUserList
		oi.User = in["user"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupListPartitionList(d []interface{}) []edpt.RbaGroupListPartitionList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupListPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupListPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.RoleList = getSliceRbaGroupListPartitionListRoleList(in["role_list"].([]interface{}))
		oi.RuleList = getSliceRbaGroupListPartitionListRuleList(in["rule_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupListPartitionListRoleList(d []interface{}) []edpt.RbaGroupListPartitionListRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupListPartitionListRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupListPartitionListRoleList
		oi.Role = in["role"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupListPartitionListRuleList(d []interface{}) []edpt.RbaGroupListPartitionListRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupListPartitionListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupListPartitionListRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaRoleList(d []interface{}) []edpt.RbaRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaRoleList
		oi.Name = in["name"].(string)
		oi.DefaultPrivilege = in["default_privilege"].(string)
		oi.PartitionOnly = in["partition_only"].(int)
		oi.RuleList = getSliceRbaRoleListRuleList(in["rule_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaRoleListRuleList(d []interface{}) []edpt.RbaRoleListRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaRoleListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaRoleListRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserList(d []interface{}) []edpt.RbaUserList {

	count1 := len(d)
	ret := make([]edpt.RbaUserList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PartitionList = getSliceRbaUserListPartitionList(in["partition_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserListPartitionList(d []interface{}) []edpt.RbaUserListPartitionList {

	count1 := len(d)
	ret := make([]edpt.RbaUserListPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserListPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.RoleList = getSliceRbaUserListPartitionListRoleList(in["role_list"].([]interface{}))
		oi.RuleList = getSliceRbaUserListPartitionListRuleList(in["rule_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserListPartitionListRoleList(d []interface{}) []edpt.RbaUserListPartitionListRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaUserListPartitionListRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserListPartitionListRoleList
		oi.Role = in["role"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserListPartitionListRuleList(d []interface{}) []edpt.RbaUserListPartitionListRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaUserListPartitionListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserListPartitionListRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRba(d *schema.ResourceData) edpt.Rba {
	var ret edpt.Rba
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.GroupList = getSliceRbaGroupList(d.Get("group_list").([]interface{}))
	ret.Inst.RoleList = getSliceRbaRoleList(d.Get("role_list").([]interface{}))
	ret.Inst.UserList = getSliceRbaUserList(d.Get("user_list").([]interface{}))
	//omit uuid
	return ret
}
