package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRemoveUpgradeLock() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_remove_upgrade_lock`: removing upgrade lock file mgmt_is_upgrade\n\n__PLACEHOLDER__",
		CreateContext: resourceRemoveUpgradeLockCreate,
		UpdateContext: resourceRemoveUpgradeLockUpdate,
		ReadContext:   resourceRemoveUpgradeLockRead,
		DeleteContext: resourceRemoveUpgradeLockDelete,

		Schema: map[string]*schema.Schema{
			"file": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lock is in the form of a file. This is not required for command.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRemoveUpgradeLockCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRemoveUpgradeLockCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRemoveUpgradeLock(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRemoveUpgradeLockRead(ctx, d, meta)
	}
	return diags
}

func resourceRemoveUpgradeLockUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRemoveUpgradeLockUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRemoveUpgradeLock(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRemoveUpgradeLockRead(ctx, d, meta)
	}
	return diags
}
func resourceRemoveUpgradeLockDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRemoveUpgradeLockDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRemoveUpgradeLock(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRemoveUpgradeLockRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRemoveUpgradeLockRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRemoveUpgradeLock(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRemoveUpgradeLock(d *schema.ResourceData) edpt.RemoveUpgradeLock {
	var ret edpt.RemoveUpgradeLock
	ret.Inst.File = d.Get("file").(int)
	//omit uuid
	return ret
}
