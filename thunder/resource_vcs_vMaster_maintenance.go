package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsVmasterMaintenance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_vMaster_maintenance`: During this period, vMaster can leave and come back to be vMaster again\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsVmasterMaintenanceCreate,
		UpdateContext: resourceVcsVmasterMaintenanceUpdate,
		ReadContext:   resourceVcsVmasterMaintenanceRead,
		DeleteContext: resourceVcsVmasterMaintenanceDelete,

		Schema: map[string]*schema.Schema{
			"vmaster_maintenance": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "the seconds vMaster will be maintained, 60 by default",
			},
		},
	}
}
func resourceVcsVmasterMaintenanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterMaintenanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterMaintenance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVmasterMaintenanceRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsVmasterMaintenanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterMaintenanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterMaintenance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVmasterMaintenanceRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsVmasterMaintenanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterMaintenanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterMaintenance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsVmasterMaintenanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterMaintenanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterMaintenance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsVmasterMaintenance(d *schema.ResourceData) edpt.VcsVmasterMaintenance {
	var ret edpt.VcsVmasterMaintenance
	ret.Inst.VmasterMaintenance = d.Get("vmaster_maintenance").(int)
	return ret
}
