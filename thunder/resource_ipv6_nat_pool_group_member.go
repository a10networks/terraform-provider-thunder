package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatPoolGroupMember() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_nat_pool_group_member`: Add a NAT pool to this pool-group\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6NatPoolGroupMemberCreate,
		UpdateContext: resourceIpv6NatPoolGroupMemberUpdate,
		ReadContext:   resourceIpv6NatPoolGroupMemberRead,
		DeleteContext: resourceIpv6NatPoolGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify NAT pool name",
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
func resourceIpv6NatPoolGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupMemberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroupMember(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NatPoolGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupMemberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroupMember(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6NatPoolGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupMemberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroupMember(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6NatPoolGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupMemberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroupMember(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6NatPoolGroupMember(d *schema.ResourceData) edpt.Ipv6NatPoolGroupMember {
	var ret edpt.Ipv6NatPoolGroupMember
	ret.Inst.PoolName = d.Get("pool_name").(string)
	//omit uuid
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	return ret
}
