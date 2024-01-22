package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateProxyServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_automatic_update_proxy_server`: Connect through proxy server\n\n__PLACEHOLDER__",
		CreateContext: resourceAutomaticUpdateProxyServerCreate,
		UpdateContext: resourceAutomaticUpdateProxyServerUpdate,
		ReadContext:   resourceAutomaticUpdateProxyServerRead,
		DeleteContext: resourceAutomaticUpdateProxyServerDelete,

		Schema: map[string]*schema.Schema{
			"auth_type": {
				Type: schema.TypeString, Optional: true, Default: "ntlm", Description: "'ntlm': NTLM authentication(default); 'basic': Basic authentication;",
			},
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Realm for NTLM authentication",
			},
			"https_port": {
				Type: schema.TypeInt, Optional: true, Description: "Proxy server HTTPs port",
			},
			"password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Password for proxy authentication",
			},
			"proxy_host": {
				Type: schema.TypeString, Optional: true, Description: "Proxy server hostname or IP address",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "password value",
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Username for proxy authentication",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAutomaticUpdateProxyServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateProxyServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateProxyServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateProxyServerRead(ctx, d, meta)
	}
	return diags
}

func resourceAutomaticUpdateProxyServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateProxyServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateProxyServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateProxyServerRead(ctx, d, meta)
	}
	return diags
}
func resourceAutomaticUpdateProxyServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateProxyServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateProxyServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAutomaticUpdateProxyServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateProxyServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateProxyServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAutomaticUpdateProxyServer(d *schema.ResourceData) edpt.AutomaticUpdateProxyServer {
	var ret edpt.AutomaticUpdateProxyServer
	ret.Inst.AuthType = d.Get("auth_type").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	//omit encrypted
	ret.Inst.HttpsPort = d.Get("https_port").(int)
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.ProxyHost = d.Get("proxy_host").(string)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	return ret
}
