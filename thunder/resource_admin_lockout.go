package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminLockout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_lockout`: Admin user lockout configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminLockoutCreate,
		UpdateContext: resourceAdminLockoutUpdate,
		ReadContext:   resourceAdminLockoutRead,
		DeleteContext: resourceAdminLockoutDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Admin user lockout duration, in minutes, by default 10 (Admin user lockout duration in minutes, 0 means forever)",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable admin user lockout",
			},
			"reset_time": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "After how long to reset the lockout counter, in minutes, by default 10 (Time in minutes after which to reset the lockout counter)",
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Admin user lockout threshold, by default 5",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAdminLockoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminLockoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminLockout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminLockoutRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminLockoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminLockoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminLockout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminLockoutRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminLockoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminLockoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminLockout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminLockoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminLockoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminLockout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminLockout(d *schema.ResourceData) edpt.AdminLockout {
	var ret edpt.AdminLockout
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.ResetTime = d.Get("reset_time").(int)
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	return ret
}
