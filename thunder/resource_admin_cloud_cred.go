package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminCloudCred() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_cloud_cred`: Config Cloud Credentials\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminCloudCredCreate,
		UpdateContext: resourceAdminCloudCredUpdate,
		ReadContext:   resourceAdminCloudCredRead,
		DeleteContext: resourceAdminCloudCredDelete,

		Schema: map[string]*schema.Schema{
			"delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized cloud credentials and cloud configuration",
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"import": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an Cloud-Cred and Cloud-Config",
			},
			"show": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized cloud credentials and cloud configuration",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'aws-cred': aws-cred; 'aws-config': aws-config; 'azure-cred': azure-cred; 'vmware-cred': vmware-cred;",
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
func resourceAdminCloudCredCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminCloudCredCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminCloudCred(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminCloudCredRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminCloudCredUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminCloudCredUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminCloudCred(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminCloudCredRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminCloudCredDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminCloudCredDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminCloudCred(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminCloudCredRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminCloudCredRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminCloudCred(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminCloudCred(d *schema.ResourceData) edpt.AdminCloudCred {
	var ret edpt.AdminCloudCred
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Import = d.Get("import").(int)
	ret.Inst.Show = d.Get("show").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.User = d.Get("user").(string)
	return ret
}
