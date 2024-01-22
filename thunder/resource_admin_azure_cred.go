package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminAzureCred() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_azure_cred`: Config Azure credentials\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminAzureCredCreate,
		UpdateContext: resourceAdminAzureCredUpdate,
		ReadContext:   resourceAdminAzureCredRead,
		DeleteContext: resourceAdminAzureCredDelete,

		Schema: map[string]*schema.Schema{
			"delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized Azure credentials",
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"import": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an azure-credentials",
			},
			"show": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized azure credentials",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"user": {
				Type: schema.TypeString, Required: true, Description: "User",
			},
		},
	}
}
func resourceAdminAzureCredCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAzureCredCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAzureCred(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminAzureCredRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminAzureCredUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAzureCredUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAzureCred(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminAzureCredRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminAzureCredDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAzureCredDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAzureCred(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminAzureCredRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAzureCredRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAzureCred(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminAzureCred(d *schema.ResourceData) edpt.AdminAzureCred {
	var ret edpt.AdminAzureCred
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Import = d.Get("import").(int)
	ret.Inst.Show = d.Get("show").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.User = d.Get("user").(string)
	return ret
}
