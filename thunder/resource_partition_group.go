package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_partition_group`: Modify a partition group\n\n__PLACEHOLDER__",
		CreateContext: resourcePartitionGroupCreate,
		UpdateContext: resourcePartitionGroupUpdate,
		ReadContext:   resourcePartitionGroupRead,
		DeleteContext: resourcePartitionGroupDelete,

		Schema: map[string]*schema.Schema{
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member": {
							Type: schema.TypeString, Optional: true, Description: "Partition Name",
						},
					},
				},
			},
			"partition_group_name": {
				Type: schema.TypeString, Required: true, Description: "Partition Group Name",
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
func resourcePartitionGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionGroupRead(ctx, d, meta)
	}
	return diags
}

func resourcePartitionGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionGroupRead(ctx, d, meta)
	}
	return diags
}
func resourcePartitionGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePartitionGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSlicePartitionGroupMemberList(d []interface{}) []edpt.PartitionGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.PartitionGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PartitionGroupMemberList
		oi.Member = in["member"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPartitionGroup(d *schema.ResourceData) edpt.PartitionGroup {
	var ret edpt.PartitionGroup
	ret.Inst.MemberList = getSlicePartitionGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.PartitionGroupName = d.Get("partition_group_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
