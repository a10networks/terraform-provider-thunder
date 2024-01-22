package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminAwsAccesskey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_aws_accesskey`: Config AWS authorized access key\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminAwsAccesskeyCreate,
		UpdateContext: resourceAdminAwsAccesskeyUpdate,
		ReadContext:   resourceAdminAwsAccesskeyRead,
		DeleteContext: resourceAdminAwsAccesskeyDelete,

		Schema: map[string]*schema.Schema{
			"delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized aws accesskey",
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"import": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an aws-accesskey",
			},
			"show": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized aws accesskey",
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
func resourceAdminAwsAccesskeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAwsAccesskeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAwsAccesskey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminAwsAccesskeyRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminAwsAccesskeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAwsAccesskeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAwsAccesskey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminAwsAccesskeyRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminAwsAccesskeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAwsAccesskeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAwsAccesskey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminAwsAccesskeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminAwsAccesskeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminAwsAccesskey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminAwsAccesskey(d *schema.ResourceData) edpt.AdminAwsAccesskey {
	var ret edpt.AdminAwsAccesskey
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Import = d.Get("import").(int)
	ret.Inst.Show = d.Get("show").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.User = d.Get("user").(string)
	return ret
}
