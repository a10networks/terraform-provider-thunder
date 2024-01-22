package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminGcpCred() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_gcp_cred`: Config GCP credentials\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminGcpCredCreate,
		UpdateContext: resourceAdminGcpCredUpdate,
		ReadContext:   resourceAdminGcpCredRead,
		DeleteContext: resourceAdminGcpCredDelete,

		Schema: map[string]*schema.Schema{
			"delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized GCP credentials",
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"import": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an GCP-credentials",
			},
			"show": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized GCP credentials",
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
func resourceAdminGcpCredCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminGcpCredCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminGcpCred(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminGcpCredRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminGcpCredUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminGcpCredUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminGcpCred(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminGcpCredRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminGcpCredDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminGcpCredDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminGcpCred(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminGcpCredRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminGcpCredRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminGcpCred(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminGcpCred(d *schema.ResourceData) edpt.AdminGcpCred {
	var ret edpt.AdminGcpCred
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Import = d.Get("import").(int)
	ret.Inst.Show = d.Get("show").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.User = d.Get("user").(string)
	return ret
}
