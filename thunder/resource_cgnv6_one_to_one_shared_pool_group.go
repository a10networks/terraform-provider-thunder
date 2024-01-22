package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneSharedPoolGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_one_to_one_shared_pool_group`: \n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6OneToOneSharedPoolGroupCreate,
		UpdateContext: resourceCgnv6OneToOneSharedPoolGroupUpdate,
		ReadContext:   resourceCgnv6OneToOneSharedPoolGroupRead,
		DeleteContext: resourceCgnv6OneToOneSharedPoolGroupDelete,

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
func resourceCgnv6OneToOneSharedPoolGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOneSharedPoolGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6OneToOneSharedPoolGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOneSharedPoolGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6OneToOneSharedPoolGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6OneToOneSharedPoolGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6OneToOneSharedPoolGroupMembers109(d []interface{}) edpt.Cgnv6OneToOneSharedPoolGroupMembers109 {

	var ret edpt.Cgnv6OneToOneSharedPoolGroupMembers109
	return ret
}

func dataToEndpointCgnv6OneToOneSharedPoolGroup(d *schema.ResourceData) edpt.Cgnv6OneToOneSharedPoolGroup {
	var ret edpt.Cgnv6OneToOneSharedPoolGroup
	ret.Inst.Members = getObjectCgnv6OneToOneSharedPoolGroupMembers109(d.Get("members").([]interface{}))
	//omit uuid
	return ret
}
