package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionAdmin() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_partition_admin`: Partition admin user configuration\n\n__PLACEHOLDER__",
		CreateContext: resourcePartitionAdminCreate,
		UpdateContext: resourcePartitionAdminUpdate,
		ReadContext:   resourcePartitionAdminRead,
		DeleteContext: resourcePartitionAdminDelete,

		Schema: map[string]*schema.Schema{
			"access_type": {
				Type: schema.TypeString, Optional: true, Default: "axapi,cli,web", Description: "",
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable user; 'disable': Disable user;",
			},
			"passwd_string": {
				Type: schema.TypeString, Optional: true, Description: "Config user password",
			},
			"privilege_partition": {
				Type: schema.TypeString, Optional: true, Default: "partition-read", Description: "'partition-enable-disable': Set per-partition enable-disable privilege; 'partition-read': Set per-partition read privilege; 'partition-write': Set per-partition write privilege;",
			},
			"trusted_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set trusted network administrator can login in",
			},
			"trusted_host_cidr": {
				Type: schema.TypeString, Optional: true, Description: "Trusted IP Address with network mask",
			},
			"user": {
				Type: schema.TypeString, Required: true, Description: "Partition admin user name",
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
func resourcePartitionAdminCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionAdminCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionAdmin(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionAdminRead(ctx, d, meta)
	}
	return diags
}

func resourcePartitionAdminUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionAdminUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionAdmin(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionAdminRead(ctx, d, meta)
	}
	return diags
}
func resourcePartitionAdminDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionAdminDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionAdmin(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePartitionAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionAdminRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionAdmin(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPartitionAdmin(d *schema.ResourceData) edpt.PartitionAdmin {
	var ret edpt.PartitionAdmin
	ret.Inst.AccessType = d.Get("access_type").(string)
	ret.Inst.Action = d.Get("action").(string)
	//omit encrypted
	ret.Inst.PasswdString = d.Get("passwd_string").(string)
	ret.Inst.PrivilegePartition = d.Get("privilege_partition").(string)
	ret.Inst.TrustedHost = d.Get("trusted_host").(int)
	ret.Inst.TrustedHostCidr = d.Get("trusted_host_cidr").(string)
	ret.Inst.User = d.Get("user").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
