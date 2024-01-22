package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerMigrateVip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_migrate_vip`: Migrate this virtual server\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerMigrateVipCreate,
		UpdateContext: resourceSlbVirtualServerMigrateVipUpdate,
		ReadContext:   resourceSlbVirtualServerMigrateVipRead,
		DeleteContext: resourceSlbVirtualServerMigrateVipDelete,

		Schema: map[string]*schema.Schema{
			"cancel_migration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cancel migration",
			},
			"finish_migration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Complete the migration",
			},
			"target_data_cpu": {
				Type: schema.TypeInt, Optional: true, Description: "Number of CPUs on the target platform",
			},
			"target_floating_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Specify IP address",
			},
			"target_floating_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbVirtualServerMigrateVipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerMigrateVipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerMigrateVip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerMigrateVipRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerMigrateVipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerMigrateVipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerMigrateVip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerMigrateVipRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerMigrateVipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerMigrateVipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerMigrateVip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerMigrateVipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerMigrateVipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerMigrateVip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbVirtualServerMigrateVip(d *schema.ResourceData) edpt.SlbVirtualServerMigrateVip {
	var ret edpt.SlbVirtualServerMigrateVip
	ret.Inst.CancelMigration = d.Get("cancel_migration").(int)
	ret.Inst.FinishMigration = d.Get("finish_migration").(int)
	ret.Inst.TargetDataCpu = d.Get("target_data_cpu").(int)
	ret.Inst.TargetFloatingIpv4 = d.Get("target_floating_ipv4").(string)
	ret.Inst.TargetFloatingIpv6 = d.Get("target_floating_ipv6").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
