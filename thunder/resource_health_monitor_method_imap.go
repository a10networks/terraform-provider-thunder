package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodImap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_imap`: IMAP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodImapCreate,
		UpdateContext: resourceHealthMonitorMethodImapUpdate,
		ReadContext:   resourceHealthMonitorMethodImapRead,
		DeleteContext: resourceHealthMonitorMethodImapDelete,

		Schema: map[string]*schema.Schema{
			"imap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IMAP type",
			},
			"imap_cram_md5": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Challenge-response authentication mechanism",
			},
			"imap_login": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Simple login",
			},
			"imap_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"imap_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
			},
			"imap_plain": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Plain text",
			},
			"imap_port": {
				Type: schema.TypeInt, Optional: true, Default: 143, Description: "Specify the IMAP port, default is 143 (Port Number (default 143))",
			},
			"imap_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the username",
			},
			"pwd_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Authentication method",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodImapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodImapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodImap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodImapRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodImapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodImapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodImap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodImapRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodImapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodImapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodImap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodImapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodImapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodImap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodImap(d *schema.ResourceData) edpt.HealthMonitorMethodImap {
	var ret edpt.HealthMonitorMethodImap
	ret.Inst.Imap = d.Get("imap").(int)
	ret.Inst.ImapCramMd5 = d.Get("imap_cram_md5").(int)
	//omit imap_encrypted
	ret.Inst.ImapLogin = d.Get("imap_login").(int)
	ret.Inst.ImapPassword = d.Get("imap_password").(int)
	ret.Inst.ImapPasswordString = d.Get("imap_password_string").(string)
	ret.Inst.ImapPlain = d.Get("imap_plain").(int)
	ret.Inst.ImapPort = d.Get("imap_port").(int)
	ret.Inst.ImapUsername = d.Get("imap_username").(string)
	ret.Inst.PwdAuth = d.Get("pwd_auth").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
