package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatPoolGroupMember() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_pool_group_member`: Add a CGNv6 NAT pool to this pool-group\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatPoolGroupMemberCreate,
		UpdateContext: resourceCgnv6NatPoolGroupMemberUpdate,
		ReadContext:   resourceCgnv6NatPoolGroupMemberRead,
		DeleteContext: resourceCgnv6NatPoolGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify CGNv6 NAT pool name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"pool_group_name": {
				Type: schema.TypeString, Required: true, Description: "PoolGroupName",
			},
		},
	}
}
func resourceCgnv6NatPoolGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupMemberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroupMember(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatPoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatPoolGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupMemberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroupMember(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatPoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatPoolGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupMemberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroupMember(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatPoolGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolGroupMemberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolGroupMember(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6NatPoolGroupMember(d *schema.ResourceData) edpt.Cgnv6NatPoolGroupMember {
	var ret edpt.Cgnv6NatPoolGroupMember
	ret.Inst.PoolName = d.Get("pool_name").(string)
	//omit uuid
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	return ret
}
