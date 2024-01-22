package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePersistCookie() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_persist_cookie`: Cookie persistence\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePersistCookieCreate,
		UpdateContext: resourceSlbTemplatePersistCookieUpdate,
		ReadContext:   resourceSlbTemplatePersistCookieRead,
		DeleteContext: resourceSlbTemplatePersistCookieDelete,

		Schema: map[string]*schema.Schema{
			"cookie_name": {
				Type: schema.TypeString, Optional: true, Default: "sto-id", Description: "Set cookie name",
			},
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Set cookie domain",
			},
			"dont_honor_conn_rules": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not observe connection rate rules",
			},
			"encrypt_level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Encryption level for cookie name / value",
			},
			"expire": {
				Type: schema.TypeInt, Optional: true, Default: 31536000, Description: "Set cookie expiration time (Expiration in seconds)",
			},
			"httponly": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable HttpOnly attribute",
			},
			"insert_always": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert persist cookie to every reponse",
			},
			"match_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist for server, default is port",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Cookie persistence (Cookie persistence template name)",
			},
			"pass_phrase": {
				Type: schema.TypeString, Optional: true, Default: "ACOS4KEY", Description: "Set passphrase for encryption",
			},
			"pass_thru": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Pass thru mode - Server sends the persist cookie",
			},
			"path": {
				Type: schema.TypeString, Optional: true, Default: "/", Description: "Set cookie path (Cookie path, default is \"/\")",
			},
			"prefix": {
				Type: schema.TypeString, Optional: true, Description: "'host': the cookie will have been set with a Secure attribute, a Path attribute with a value of /, and no Domain attribute; 'secure': the cookie will have been set with a Secure attribute;",
			},
			"samesite": {
				Type: schema.TypeString, Optional: true, Description: "'none': none; 'lax': lax; 'strict': strict;",
			},
			"scan_all_members": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist within the same server SCAN",
			},
			"secure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable secure attribute",
			},
			"server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist to the same server, default is port",
			},
			"server_service_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist to the same server and within the same service group",
			},
			"service_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist within the same service group",
			},
			"use_attribute": {
				Type: schema.TypeString, Optional: true, Default: "expires", Description: "'max-age': Use the Max-Age attribute; 'expires': Use the Expires attribute; 'all': Use all attributes;",
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
func resourceSlbTemplatePersistCookieCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistCookieCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistCookie(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePersistCookieRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePersistCookieUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistCookieUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistCookie(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePersistCookieRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePersistCookieDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistCookieDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistCookie(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePersistCookieRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistCookieRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistCookie(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePersistCookie(d *schema.ResourceData) edpt.SlbTemplatePersistCookie {
	var ret edpt.SlbTemplatePersistCookie
	ret.Inst.CookieName = d.Get("cookie_name").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.DontHonorConnRules = d.Get("dont_honor_conn_rules").(int)
	ret.Inst.EncryptLevel = d.Get("encrypt_level").(int)
	//omit encrypted
	ret.Inst.Expire = d.Get("expire").(int)
	ret.Inst.Httponly = d.Get("httponly").(int)
	ret.Inst.InsertAlways = d.Get("insert_always").(int)
	ret.Inst.MatchType = d.Get("match_type").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PassPhrase = d.Get("pass_phrase").(string)
	ret.Inst.PassThru = d.Get("pass_thru").(int)
	ret.Inst.Path = d.Get("path").(string)
	ret.Inst.Prefix = d.Get("prefix").(string)
	ret.Inst.Samesite = d.Get("samesite").(string)
	ret.Inst.ScanAllMembers = d.Get("scan_all_members").(int)
	ret.Inst.Secure = d.Get("secure").(int)
	ret.Inst.Server = d.Get("server").(int)
	ret.Inst.ServerServiceGroup = d.Get("server_service_group").(int)
	ret.Inst.ServiceGroup = d.Get("service_group").(int)
	ret.Inst.UseAttribute = d.Get("use_attribute").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
