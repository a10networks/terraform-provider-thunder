package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerSnmpv3User() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_SNMPv3_user`: Define a user who can access the SNMP engine\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerSnmpv3UserCreate,
		UpdateContext: resourceSnmpServerSnmpv3UserUpdate,
		ReadContext:   resourceSnmpServerSnmpv3UserRead,
		DeleteContext: resourceSnmpServerSnmpv3UserDelete,

		Schema: map[string]*schema.Schema{
			"auth_val": {
				Type: schema.TypeString, Optional: true, Description: "'md5': Use HMAC MD5 algorithm for authentication; 'sha': Use HMAC SHA algorithm for authentication; 'sha-512': Use HMAC SHA-512 algorithm for authentication; 'sha-384': Use HMAC SHA-384 algorithm for authentication; 'sha-256': Use HMAC SHA-256 algorithm for authentication; 'sha-224': Use HMAC SHA-224 algorithm for authentication;",
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
				Type: schema.TypeString, Optional: true, Description: "'des': DES encryption alogrithm; 'aes': AES encryption alogrithm; 'aes-192': AES-192 encryption alogrithm; 'aes-256': AES-256 encryption alogrithm;  (Encryption type)",
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
func resourceSnmpServerSnmpv3UserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv3UserCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv3User(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSnmpv3UserRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerSnmpv3UserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv3UserUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv3User(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSnmpv3UserRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerSnmpv3UserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv3UserDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv3User(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerSnmpv3UserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv3UserRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv3User(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerSnmpv3User(d *schema.ResourceData) edpt.SnmpServerSnmpv3User {
	var ret edpt.SnmpServerSnmpv3User
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
