package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOnePoolGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_one_to_one_pool_group`: Configure CGNv6 one-to-one pool group\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6OneToOnePoolGroupCreate,
		UpdateContext: resourceCgnv6OneToOnePoolGroupUpdate,
		ReadContext:   resourceCgnv6OneToOnePoolGroupRead,
		DeleteContext: resourceCgnv6OneToOnePoolGroupDelete,

		Schema: map[string]*schema.Schema{
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool_name": {
							Type: schema.TypeString, Required: true, Description: "Specify CGNv6 one-to-one pool name",
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
func resourceCgnv6OneToOnePoolGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOnePoolGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6OneToOnePoolGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOnePoolGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6OneToOnePoolGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6OneToOnePoolGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6OneToOnePoolGroupMemberList(d []interface{}) []edpt.Cgnv6OneToOnePoolGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOnePoolGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOnePoolGroupMemberList
		oi.PoolName = in["pool_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6OneToOnePoolGroup(d *schema.ResourceData) edpt.Cgnv6OneToOnePoolGroup {
	var ret edpt.Cgnv6OneToOnePoolGroup
	ret.Inst.MemberList = getSliceCgnv6OneToOnePoolGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
