package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbaUser() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rba_user`: RBA configuration for a user\n\n__PLACEHOLDER__",
		CreateContext: resourceRbaUserCreate,
		UpdateContext: resourceRbaUserUpdate,
		ReadContext:   resourceRbaUserRead,
		DeleteContext: resourceRbaUserDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of a user account",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRbaUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUser(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaUserRead(ctx, d, meta)
	}
	return diags
}

func resourceRbaUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUser(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaUserRead(ctx, d, meta)
	}
	return diags
}
func resourceRbaUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUser(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRbaUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUser(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRbaUserPartitionList(d []interface{}) []edpt.RbaUserPartitionList {

	count1 := len(d)
	ret := make([]edpt.RbaUserPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.RoleList = getSliceRbaUserPartitionListRoleList(in["role_list"].([]interface{}))
		oi.RuleList = getSliceRbaUserPartitionListRuleList(in["rule_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserPartitionListRoleList(d []interface{}) []edpt.RbaUserPartitionListRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaUserPartitionListRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserPartitionListRoleList
		oi.Role = in["role"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserPartitionListRuleList(d []interface{}) []edpt.RbaUserPartitionListRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaUserPartitionListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserPartitionListRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRbaUser(d *schema.ResourceData) edpt.RbaUser {
	var ret edpt.RbaUser
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PartitionList = getSliceRbaUserPartitionList(d.Get("partition_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
