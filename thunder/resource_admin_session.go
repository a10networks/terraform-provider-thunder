package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminSession() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_session`: Admin session management\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminSessionCreate,
		UpdateContext: resourceAdminSessionUpdate,
		ReadContext:   resourceAdminSessionRead,
		DeleteContext: resourceAdminSessionDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear all admin sessions",
			},
			"clear": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "clear admin session",
			},
			"sid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Session ID",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAdminSessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSessionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSession(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminSessionRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminSessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSessionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSession(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminSessionRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminSessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSessionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSession(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminSessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSessionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSession(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminSession(d *schema.ResourceData) edpt.AdminSession {
	var ret edpt.AdminSession
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Clear = d.Get("clear").(int)
	ret.Inst.Sid = d.Get("sid").(int)
	//omit uuid
	return ret
}
