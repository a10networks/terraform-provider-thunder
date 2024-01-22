package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemManagementInterfaceMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_management_interface_mode`: Configure management interface mode\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemManagementInterfaceModeCreate,
		UpdateContext: resourceSystemManagementInterfaceModeUpdate,
		ReadContext:   resourceSystemManagementInterfaceModeRead,
		DeleteContext: resourceSystemManagementInterfaceModeDelete,

		Schema: map[string]*schema.Schema{
			"dedicated": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set management interface in dedicated mode",
			},
			"non_dedicated": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set management interface in non-dedicated mode",
			},
		},
	}
}
func resourceSystemManagementInterfaceModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemManagementInterfaceModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemManagementInterfaceMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemManagementInterfaceModeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemManagementInterfaceModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemManagementInterfaceModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemManagementInterfaceMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemManagementInterfaceModeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemManagementInterfaceModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemManagementInterfaceModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemManagementInterfaceMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemManagementInterfaceModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemManagementInterfaceModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemManagementInterfaceMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemManagementInterfaceMode(d *schema.ResourceData) edpt.SystemManagementInterfaceMode {
	var ret edpt.SystemManagementInterfaceMode
	ret.Inst.Dedicated = d.Get("dedicated").(int)
	ret.Inst.NonDedicated = d.Get("non_dedicated").(int)
	return ret
}
