package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminAccess() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_access`: Config access type\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminAccessCreate,
		UpdateContext: resourceAdminAccessUpdate,
		ReadContext:   resourceAdminAccessRead,
		DeleteContext: resourceAdminAccessDelete,

		Schema: map[string]*schema.Schema{
			"access_type": {
				Type: schema.TypeString, Optional: true, Default: "axapi,cli,web", Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"user": {
				Type: schema.TypeString, Required: true, Description: "User",
			},
		},
	}
}
func resourceAdminAccessCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAccessCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAccess(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminAccessRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminAccessUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAccessUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAccess(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminAccessRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminAccessDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAccessDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAccess(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminAccessRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAccessRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAccess(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminAccess(d *schema.ResourceData) edpt.AdminAccess {
	var ret edpt.AdminAccess
	ret.Inst.AccessType = d.Get("access_type").(string)
	//omit uuid
	ret.Inst.User = d.Get("user").(string)
	return ret
}
