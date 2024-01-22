package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminPassword() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_password`: Config admin user password\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminPasswordCreate,
		UpdateContext: resourceAdminPasswordUpdate,
		ReadContext:   resourceAdminPasswordRead,
		DeleteContext: resourceAdminPasswordDelete,

		Schema: map[string]*schema.Schema{
			"password_in_module": {
				Type: schema.TypeString, Optional: true, Description: "Config admin user password",
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
func resourceAdminPasswordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminPasswordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminPassword(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminPasswordRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminPasswordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminPasswordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminPassword(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminPasswordRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminPasswordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminPasswordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminPassword(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminPasswordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminPasswordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminPassword(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminPassword(d *schema.ResourceData) edpt.AdminPassword {
	var ret edpt.AdminPassword
	//omit encrypted_in_module
	ret.Inst.PasswordInModule = d.Get("password_in_module").(string)
	//omit uuid
	ret.Inst.User = d.Get("user").(string)
	return ret
}
