package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOnePoolGroupMember() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_one_to_one_pool_group_member`: Add a CGNv6 one-to-one pool to this pool-group\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6OneToOnePoolGroupMemberCreate,
		UpdateContext: resourceCgnv6OneToOnePoolGroupMemberUpdate,
		ReadContext:   resourceCgnv6OneToOnePoolGroupMemberRead,
		DeleteContext: resourceCgnv6OneToOnePoolGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify CGNv6 one-to-one pool name",
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
func resourceCgnv6OneToOnePoolGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupMemberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroupMember(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOnePoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6OneToOnePoolGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupMemberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroupMember(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOnePoolGroupMemberRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6OneToOnePoolGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupMemberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroupMember(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6OneToOnePoolGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolGroupMemberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolGroupMember(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6OneToOnePoolGroupMember(d *schema.ResourceData) edpt.Cgnv6OneToOnePoolGroupMember {
	var ret edpt.Cgnv6OneToOnePoolGroupMember
	ret.Inst.PoolName = d.Get("pool_name").(string)
	//omit uuid
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	return ret
}
