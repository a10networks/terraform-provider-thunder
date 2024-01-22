package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesMetaData() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_meta_data`: user-data Services configuration only works in shared partition\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesMetaDataCreate,
		UpdateContext: resourceCloudServicesMetaDataUpdate,
		ReadContext:   resourceCloudServicesMetaDataRead,
		DeleteContext: resourceCloudServicesMetaDataDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
			},
			"prevent_admin_passwd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents admin password from being changed if in YAML config",
			},
			"prevent_admin_ssh_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents admin ssh-key from being added if in YAML config",
			},
			"prevent_autofill": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "prevents use of meta-data to complete user_data configuration",
			},
			"prevent_blob": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents a10_blob from loading in YAML config",
			},
			"prevent_cloud_service": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents cloud-service configuration in YAML config",
			},
			"prevent_license": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents a10_license configuration in YAML config",
			},
			"prevent_user_ops": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents user from being added command is in user-data",
			},
			"prevent_webservice": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prevents a10_license configuration in YAML config",
			},
			"provider1": {
				Type: schema.TypeString, Optional: true, Description: "'openstack': OpenStack user-data services;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCloudServicesMetaDataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesMetaDataCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesMetaData(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesMetaDataRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesMetaDataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesMetaDataUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesMetaData(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesMetaDataRead(ctx, d, meta)
	}
	return diags
}
func resourceCloudServicesMetaDataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesMetaDataDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesMetaData(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesMetaDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesMetaDataRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesMetaData(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesMetaData(d *schema.ResourceData) edpt.CloudServicesMetaData {
	var ret edpt.CloudServicesMetaData
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.PreventAdminPasswd = d.Get("prevent_admin_passwd").(int)
	ret.Inst.PreventAdminSshKey = d.Get("prevent_admin_ssh_key").(int)
	ret.Inst.PreventAutofill = d.Get("prevent_autofill").(int)
	ret.Inst.PreventBlob = d.Get("prevent_blob").(int)
	ret.Inst.PreventCloudService = d.Get("prevent_cloud_service").(int)
	ret.Inst.PreventLicense = d.Get("prevent_license").(int)
	ret.Inst.PreventUserOps = d.Get("prevent_user_ops").(int)
	ret.Inst.PreventWebservice = d.Get("prevent_webservice").(int)
	ret.Inst.Provider1 = d.Get("provider1").(string)
	//omit uuid
	return ret
}
