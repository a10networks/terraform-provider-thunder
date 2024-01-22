package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPoolGroupMember() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_pool_group_member`: Add a NAT pool to this pool-group\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatPoolGroupMemberCreate,
		UpdateContext: resourceIpNatPoolGroupMemberUpdate,
		ReadContext:   resourceIpNatPoolGroupMemberRead,
		DeleteContext: resourceIpNatPoolGroupMemberDelete,

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
func resourceIpNatPoolGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupMemberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroupMember(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatPoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatPoolGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupMemberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroupMember(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatPoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatPoolGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupMemberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroupMember(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatPoolGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupMemberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroupMember(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatPoolGroupMember(d *schema.ResourceData) edpt.IpNatPoolGroupMember {
	var ret edpt.IpNatPoolGroupMember
	ret.Inst.PoolName = d.Get("pool_name").(string)
	//omit uuid
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	return ret
}
