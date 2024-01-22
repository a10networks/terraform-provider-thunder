package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsReload() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_reload`: VCS reload\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsReloadCreate,
		UpdateContext: resourceVcsReloadUpdate,
		ReadContext:   resourceVcsReloadRead,
		DeleteContext: resourceVcsReloadDelete,

		Schema: map[string]*schema.Schema{
			"cluster_discovery": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "apply the configuration change on the aVCS cluster",
			},
			"complete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"db_safe": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "keep database safe",
			},
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "Reload a specific device when VCS is enabled (device id)",
			},
			"disable_merge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "don't merge this vBlade's configuration to aVCS chassis",
			},
			"force": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "force to complete",
			},
			"start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Description: "automatically complete the db-safe mode when timeout",
			},
		},
	}
}
func resourceVcsReloadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsReloadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsReload(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsReloadRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsReloadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsReloadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsReload(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsReloadRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsReloadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsReloadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsReload(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsReloadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsReloadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsReload(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsReload(d *schema.ResourceData) edpt.VcsReload {
	var ret edpt.VcsReload
	ret.Inst.ClusterDiscovery = d.Get("cluster_discovery").(int)
	ret.Inst.Complete = d.Get("complete").(int)
	ret.Inst.DbSafe = d.Get("db_safe").(int)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.DisableMerge = d.Get("disable_merge").(int)
	ret.Inst.Force = d.Get("force").(int)
	ret.Inst.Start = d.Get("start").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	return ret
}
