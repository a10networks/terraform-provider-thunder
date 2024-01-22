package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatSharedPoolGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_shared_pool_group`: \n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatSharedPoolGroupCreate,
		UpdateContext: resourceCgnv6NatSharedPoolGroupUpdate,
		ReadContext:   resourceCgnv6NatSharedPoolGroupRead,
		DeleteContext: resourceCgnv6NatSharedPoolGroupDelete,

		Schema: map[string]*schema.Schema{
			"members": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceCgnv6NatSharedPoolGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatSharedPoolGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatSharedPoolGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatSharedPoolGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatSharedPoolGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatSharedPoolGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6NatSharedPoolGroupMembers108(d []interface{}) edpt.Cgnv6NatSharedPoolGroupMembers108 {

	var ret edpt.Cgnv6NatSharedPoolGroupMembers108
	return ret
}

func dataToEndpointCgnv6NatSharedPoolGroup(d *schema.ResourceData) edpt.Cgnv6NatSharedPoolGroup {
	var ret edpt.Cgnv6NatSharedPoolGroup
	ret.Inst.Members = getObjectCgnv6NatSharedPoolGroupMembers108(d.Get("members").([]interface{}))
	//omit uuid
	return ret
}
