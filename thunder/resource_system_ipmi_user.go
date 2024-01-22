package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpmiUser() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipmi_user`: Add, Change or Disable IPMI users\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpmiUserCreate,
		UpdateContext: resourceSystemIpmiUserUpdate,
		ReadContext:   resourceSystemIpmiUserRead,
		DeleteContext: resourceSystemIpmiUserDelete,

		Schema: map[string]*schema.Schema{
			"add": {
				Type: schema.TypeString, Optional: true, Description: "Add a new IPMI user (IPMI User Name)",
			},
			"administrator": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Full control",
			},
			"callback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lowest privilege level",
			},
			"disable": {
				Type: schema.TypeString, Optional: true, Description: "Disable an existing IPMI user (IPMI User Name)",
			},
			"newname": {
				Type: schema.TypeString, Optional: true, Description: "New IPMI User Name",
			},
			"newpass": {
				Type: schema.TypeString, Optional: true, Description: "New Password",
			},
			"operator": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most BMC commands are allowed",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "Password",
			},
			"privilege": {
				Type: schema.TypeString, Optional: true, Description: "Change an existing IPMI user privilege (IPMI User Name)",
			},
			"setname": {
				Type: schema.TypeString, Optional: true, Description: "Change User Name (Current IPMI User Name)",
			},
			"setpass": {
				Type: schema.TypeString, Optional: true, Description: "Change Password (IPMI User Name)",
			},
			"user": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only 'benign' commands are allowed",
			},
		},
	}
}
func resourceSystemIpmiUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiUserCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiUser(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiUserRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpmiUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiUserUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiUser(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiUserRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpmiUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiUserDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiUser(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpmiUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiUserRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiUser(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIpmiUser(d *schema.ResourceData) edpt.SystemIpmiUser {
	var ret edpt.SystemIpmiUser
	ret.Inst.Add = d.Get("add").(string)
	ret.Inst.Administrator = d.Get("administrator").(int)
	ret.Inst.Callback = d.Get("callback").(int)
	ret.Inst.Disable = d.Get("disable").(string)
	ret.Inst.Newname = d.Get("newname").(string)
	ret.Inst.Newpass = d.Get("newpass").(string)
	ret.Inst.Operator = d.Get("operator").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.Privilege = d.Get("privilege").(string)
	ret.Inst.Setname = d.Get("setname").(string)
	ret.Inst.Setpass = d.Get("setpass").(string)
	ret.Inst.User = d.Get("user").(int)
	return ret
}
