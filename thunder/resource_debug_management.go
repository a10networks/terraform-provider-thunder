package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugManagement() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_management`: management module parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugManagementCreate,
		UpdateContext: resourceDebugManagementUpdate,
		ReadContext:   resourceDebugManagementRead,
		DeleteContext: resourceDebugManagementDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "management all debug switch",
			},
			"system": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "management system debug switch",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugManagementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugManagementCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugManagement(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugManagementRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugManagementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugManagementUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugManagement(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugManagementRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugManagementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugManagementDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugManagement(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugManagementRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugManagement(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugManagement(d *schema.ResourceData) edpt.DebugManagement {
	var ret edpt.DebugManagement
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.System = d.Get("system").(int)
	//omit uuid
	return ret
}
