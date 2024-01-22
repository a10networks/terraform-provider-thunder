package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatPoolGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_pool_group`: Configure CGNv6 NAT pool group\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatPoolGroupCreate,
		UpdateContext: resourceCgnv6NatPoolGroupUpdate,
		ReadContext:   resourceCgnv6NatPoolGroupRead,
		DeleteContext: resourceCgnv6NatPoolGroupDelete,

		Schema: map[string]*schema.Schema{
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool_name": {
							Type: schema.TypeString, Required: true, Description: "Specify CGNv6 NAT pool name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pool_group_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool group name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify VRRP-A vrid",
			},
		},
	}
}
func resourceCgnv6NatPoolGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatPoolGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatPoolGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatPoolGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatPoolGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatPoolGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6NatPoolGroupMemberList(d []interface{}) []edpt.Cgnv6NatPoolGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatPoolGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatPoolGroupMemberList
		oi.PoolName = in["pool_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatPoolGroup(d *schema.ResourceData) edpt.Cgnv6NatPoolGroup {
	var ret edpt.Cgnv6NatPoolGroup
	ret.Inst.MemberList = getSliceCgnv6NatPoolGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
