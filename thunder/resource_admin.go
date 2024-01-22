package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdmin() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin`: System admin user configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminCreate,
		UpdateContext: resourceAdminUpdate,
		ReadContext:   resourceAdminRead,
		DeleteContext: resourceAdminDelete,

		Schema: map[string]*schema.Schema{
			"access": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_type": {
							Type: schema.TypeString, Optional: true, Default: "axapi,cli,web", Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"access_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify an ACL to classify a trusted host",
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable user; 'disable': Disable user;",
			},
			"aws_accesskey": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"import": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an aws-accesskey",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized aws accesskey",
						},
						"show": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized aws accesskey",
						},
					},
				},
			},
			"azure_cred": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"import": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an azure-credentials",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized Azure credentials",
						},
						"show": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized azure credentials",
						},
					},
				},
			},
			"cloud_cred": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'aws-cred': aws-cred; 'aws-config': aws-config; 'azure-cred': azure-cred; 'vmware-cred': vmware-cred;",
						},
						"import": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an Cloud-Cred and Cloud-Config",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized cloud credentials and cloud configuration",
						},
						"show": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized cloud credentials and cloud configuration",
						},
					},
				},
			},
			"gcp_cred": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"import": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an GCP-credentials",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an authorized GCP credentials",
						},
						"show": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show authorized GCP credentials",
						},
					},
				},
			},
			"passwd_string": {
				Type: schema.TypeString, Optional: true, Description: "Config admin user password",
			},
			"password": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password_in_module": {
							Type: schema.TypeString, Optional: true, Description: "Config admin user password",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"password_key": {
				Type: schema.TypeInt, Optional: true, Description: "Config admin user password",
			},
			"privilege_global": {
				Type: schema.TypeString, Optional: true, Description: "'read': Set read privilege; 'write': Set write privilege; 'hm': Set external health monitor script content operations privilege;",
			},
			"privilege_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"privilege_partition": {
							Type: schema.TypeString, Optional: true, Description: "'partition-enable-disable': Set per-partition enable/disable privilege; 'partition-read': Set per-partition read privilege; 'partition-write': Set per-partition write privilege;",
						},
						"partition_name": {
							Type: schema.TypeString, Optional: true, Description: "Partition Name",
						},
					},
				},
			},
			"privilege_shell_root": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set shell root privilege",
			},
			"ssh_pubkey": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"import": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an authorized public key",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
						"delete": {
							Type: schema.TypeInt, Optional: true, Description: "Delete an authorized public key (SSH key index)",
						},
						"list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "List all authorized public keys",
						},
					},
				},
			},
			"trusted_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set trusted network administrator can login in",
			},
			"trusted_host_acl_id": {
				Type: schema.TypeInt, Optional: true, Description: "ACL ID",
			},
			"trusted_host_cidr": {
				Type: schema.TypeString, Optional: true, Description: "Trusted IP Address with network mask",
			},
			"user": {
				Type: schema.TypeString, Required: true, Description: "System admin user name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAdminCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdmin(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdmin(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdmin(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdmin(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAdminAccess60(d []interface{}) edpt.AdminAccess60 {

	count1 := len(d)
	var ret edpt.AdminAccess60
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AccessType = in["access_type"].(string)
		//omit uuid
	}
	return ret
}

func getObjectAdminAwsAccesskey61(d []interface{}) edpt.AdminAwsAccesskey61 {

	count1 := len(d)
	var ret edpt.AdminAwsAccesskey61
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Import = in["import"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
		ret.Delete = in["delete"].(int)
		ret.Show = in["show"].(int)
	}
	return ret
}

func getObjectAdminAzureCred62(d []interface{}) edpt.AdminAzureCred62 {

	count1 := len(d)
	var ret edpt.AdminAzureCred62
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Import = in["import"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
		ret.Delete = in["delete"].(int)
		ret.Show = in["show"].(int)
	}
	return ret
}

func getObjectAdminCloudCred63(d []interface{}) edpt.AdminCloudCred63 {

	count1 := len(d)
	var ret edpt.AdminCloudCred63
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(string)
		ret.Import = in["import"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
		ret.Delete = in["delete"].(int)
		ret.Show = in["show"].(int)
	}
	return ret
}

func getObjectAdminGcpCred64(d []interface{}) edpt.AdminGcpCred64 {

	count1 := len(d)
	var ret edpt.AdminGcpCred64
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Import = in["import"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
		ret.Delete = in["delete"].(int)
		ret.Show = in["show"].(int)
	}
	return ret
}

func getObjectAdminPassword65(d []interface{}) edpt.AdminPassword65 {

	count1 := len(d)
	var ret edpt.AdminPassword65
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PasswordInModule = in["password_in_module"].(string)
		//omit encrypted_in_module
		//omit uuid
	}
	return ret
}

func getSliceAdminPrivilegeList(d []interface{}) []edpt.AdminPrivilegeList {

	count1 := len(d)
	ret := make([]edpt.AdminPrivilegeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AdminPrivilegeList
		oi.PrivilegePartition = in["privilege_partition"].(string)
		oi.PartitionName = in["partition_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAdminSshPubkey66(d []interface{}) edpt.AdminSshPubkey66 {

	count1 := len(d)
	var ret edpt.AdminSshPubkey66
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Import = in["import"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
		ret.Delete = in["delete"].(int)
		ret.List = in["list"].(int)
	}
	return ret
}

func dataToEndpointAdmin(d *schema.ResourceData) edpt.Admin {
	var ret edpt.Admin
	ret.Inst.Access = getObjectAdminAccess60(d.Get("access").([]interface{}))
	ret.Inst.AccessList = d.Get("access_list").(int)
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AwsAccesskey = getObjectAdminAwsAccesskey61(d.Get("aws_accesskey").([]interface{}))
	ret.Inst.AzureCred = getObjectAdminAzureCred62(d.Get("azure_cred").([]interface{}))
	ret.Inst.CloudCred = getObjectAdminCloudCred63(d.Get("cloud_cred").([]interface{}))
	//omit encrypted
	ret.Inst.GcpCred = getObjectAdminGcpCred64(d.Get("gcp_cred").([]interface{}))
	ret.Inst.PasswdString = d.Get("passwd_string").(string)
	ret.Inst.Password = getObjectAdminPassword65(d.Get("password").([]interface{}))
	ret.Inst.PasswordKey = d.Get("password_key").(int)
	ret.Inst.PrivilegeGlobal = d.Get("privilege_global").(string)
	ret.Inst.PrivilegeList = getSliceAdminPrivilegeList(d.Get("privilege_list").([]interface{}))
	ret.Inst.PrivilegeShellRoot = d.Get("privilege_shell_root").(int)
	ret.Inst.SshPubkey = getObjectAdminSshPubkey66(d.Get("ssh_pubkey").([]interface{}))
	ret.Inst.TrustedHost = d.Get("trusted_host").(int)
	ret.Inst.TrustedHostAclId = d.Get("trusted_host_acl_id").(int)
	ret.Inst.TrustedHostCidr = d.Get("trusted_host_cidr").(string)
	ret.Inst.User = d.Get("user").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
