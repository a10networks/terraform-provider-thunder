package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportHealthExternal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_health_external`: Address the External Script Program\n\n__PLACEHOLDER__",
		CreateContext: resourceImportHealthExternalCreate,
		UpdateContext: resourceImportHealthExternalUpdate,
		ReadContext:   resourceImportHealthExternalRead,
		DeleteContext: resourceImportHealthExternalDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Describe the Program Function briefly",
			},
			"externalfilename": {
				Type: schema.TypeString, Optional: true, Description: "Specify the Program Name",
			},
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceImportHealthExternalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthExternalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthExternal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportHealthExternalRead(ctx, d, meta)
	}
	return diags
}

func resourceImportHealthExternalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthExternalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthExternal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportHealthExternalRead(ctx, d, meta)
	}
	return diags
}
func resourceImportHealthExternalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthExternalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthExternal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportHealthExternalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthExternalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthExternal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportHealthExternal(d *schema.ResourceData) edpt.ImportHealthExternal {
	var ret edpt.ImportHealthExternal
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.Externalfilename = d.Get("externalfilename").(string)
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
