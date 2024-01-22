package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportHealthPostfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_health_postfile`: Address the HTTP Post data file\n\n__PLACEHOLDER__",
		CreateContext: resourceImportHealthPostfileCreate,
		UpdateContext: resourceImportHealthPostfileUpdate,
		ReadContext:   resourceImportHealthPostfileRead,
		DeleteContext: resourceImportHealthPostfileDelete,

		Schema: map[string]*schema.Schema{
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"postfilename": {
				Type: schema.TypeString, Optional: true, Description: "Specify the File Name",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceImportHealthPostfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthPostfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthPostfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportHealthPostfileRead(ctx, d, meta)
	}
	return diags
}

func resourceImportHealthPostfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthPostfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthPostfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportHealthPostfileRead(ctx, d, meta)
	}
	return diags
}
func resourceImportHealthPostfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthPostfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthPostfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportHealthPostfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportHealthPostfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportHealthPostfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportHealthPostfile(d *schema.ResourceData) edpt.ImportHealthPostfile {
	var ret edpt.ImportHealthPostfile
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.Postfilename = d.Get("postfilename").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
