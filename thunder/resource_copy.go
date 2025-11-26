package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCopy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_copy`: Copy configuration to/from remote server\n\n__PLACEHOLDER__",
		CreateContext: resourceCopyCreate,
		UpdateContext: resourceCopyUpdate,
		ReadContext:   resourceCopyRead,
		DeleteContext: resourceCopyDelete,

		Schema: map[string]*schema.Schema{
			"dest_profile": {
				Type: schema.TypeString, Optional: true, Description: "Local Configuration Profile Name",
			},
			"dest_remote_file": {
				Type: schema.TypeString, Optional: true, Description: "Remote file path",
			},
			"dest_use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as destination port",
			},
			"profile": {
				Type: schema.TypeString, Optional: true, Description: "From Startup-config Profile",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "Remote file path",
			},
			"running_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "From Running Config",
			},
			"startup_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "From Startup Config",
			},
			"to_profile": {
				Type: schema.TypeString, Optional: true, Description: "To Local Configuration Profile",
			},
			"to_running_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "To Running Config",
			},
			"to_startup_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "To Startup Config",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceCopyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCopyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCopy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCopyRead(ctx, d, meta)
	}
	return diags
}

func resourceCopyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCopyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCopy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCopyRead(ctx, d, meta)
	}
	return diags
}
func resourceCopyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCopyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCopy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCopyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCopyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCopy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCopy(d *schema.ResourceData) edpt.Copy {
	var ret edpt.Copy
	ret.Inst.DestProfile = d.Get("dest_profile").(string)
	ret.Inst.DestRemoteFile = d.Get("dest_remote_file").(string)
	ret.Inst.DestUseMgmtPort = d.Get("dest_use_mgmt_port").(int)
	ret.Inst.Profile = d.Get("profile").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.RunningConfig = d.Get("running_config").(int)
	ret.Inst.StartupConfig = d.Get("startup_config").(int)
	ret.Inst.ToProfile = d.Get("to_profile").(string)
	ret.Inst.ToRunningConfig = d.Get("to_running_config").(int)
	ret.Inst.ToStartupConfig = d.Get("to_startup_config").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
