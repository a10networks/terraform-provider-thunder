package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerUser() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_user`: Deprecated SNMPv3 user command\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerUserCreate,
		UpdateContext: resourceSnmpServerUserUpdate,
		ReadContext:   resourceSnmpServerUserRead,
		DeleteContext: resourceSnmpServerUserDelete,

		Schema: map[string]*schema.Schema{
			"auth_val": {
				Type: schema.TypeString, Optional: true, Description: "'md5': Use HMAC MD5 algorithm for authentication; 'sha': Use HMAC SHA algorithm for authentication;",
			},
			"encpasswd": {
				Type: schema.TypeString, Optional: true, Description: "Passphrase for encryption",
			},
			"group": {
				Type: schema.TypeString, Optional: true, Description: "Group to which the user belongs",
			},
			"passwd": {
				Type: schema.TypeString, Optional: true, Description: "Password of this user",
			},
			"priv": {
				Type: schema.TypeString, Optional: true, Description: "'des': DES encryption alogrithm; 'aes': AES encryption alogrithm;  (Encryption type)",
			},
			"username": {
				Type: schema.TypeString, Required: true, Description: "Name of the user",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v3": {
				Type: schema.TypeString, Optional: true, Description: "'auth': Using the authNoPriv Security Level; 'noauth': Using the noAuthNoPriv Security Level;",
			},
		},
	}
}
func resourceSnmpServerUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerUserCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerUser(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerUserRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerUserUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerUser(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerUserRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerUserDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerUser(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerUserRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerUser(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerUser(d *schema.ResourceData) edpt.SnmpServerUser {
	var ret edpt.SnmpServerUser
	ret.Inst.AuthVal = d.Get("auth_val").(string)
	ret.Inst.Encpasswd = d.Get("encpasswd").(string)
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.Passwd = d.Get("passwd").(string)
	ret.Inst.Priv = d.Get("priv").(string)
	//omit priv_pw_encrypted
	//omit pw_encrypted
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	ret.Inst.V3 = d.Get("v3").(string)
	return ret
}
