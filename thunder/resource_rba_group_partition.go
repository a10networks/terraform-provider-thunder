package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbaGroupPartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rba_group_partition`: RBA configuration for the access privilege of a group within one partition\n\n__PLACEHOLDER__",
		CreateContext: resourceRbaGroupPartitionCreate,
		UpdateContext: resourceRbaGroupPartitionUpdate,
		ReadContext:   resourceRbaGroupPartitionRead,
		DeleteContext: resourceRbaGroupPartitionDelete,

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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceRbaGroupPartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupPartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroupPartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaGroupPartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceRbaGroupPartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupPartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroupPartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaGroupPartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceRbaGroupPartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupPartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroupPartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRbaGroupPartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaGroupPartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaGroupPartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRbaGroupPartitionRoleList(d []interface{}) []edpt.RbaGroupPartitionRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupPartitionRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupPartitionRoleList
		oi.Role = in["role"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaGroupPartitionRuleList(d []interface{}) []edpt.RbaGroupPartitionRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaGroupPartitionRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaGroupPartitionRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRbaGroupPartition(d *schema.ResourceData) edpt.RbaGroupPartition {
	var ret edpt.RbaGroupPartition
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.RoleList = getSliceRbaGroupPartitionRoleList(d.Get("role_list").([]interface{}))
	ret.Inst.RuleList = getSliceRbaGroupPartitionRuleList(d.Get("rule_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
