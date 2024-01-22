package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryProxyServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_category_proxy_server`: Commands to connect web-category through proxy server\n\n__PLACEHOLDER__",
		CreateContext: resourceWebCategoryProxyServerCreate,
		UpdateContext: resourceWebCategoryProxyServerUpdate,
		ReadContext:   resourceWebCategoryProxyServerRead,
		DeleteContext: resourceWebCategoryProxyServerDelete,

		Schema: map[string]*schema.Schema{
			"auth_type": {
				Type: schema.TypeString, Optional: true, Default: "ntlm", Description: "'ntlm': NTLM authentication(default); 'basic': Basic authentication;",
			},
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Realm for NTLM authentication",
			},
			"http_port": {
				Type: schema.TypeInt, Optional: true, Description: "Proxy server HTTP port",
			},
			"https_port": {
				Type: schema.TypeInt, Optional: true, Description: "Proxy server HTTPS port(HTTP port will be used if not configured)",
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
func resourceWebCategoryProxyServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryProxyServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryProxyServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryProxyServerRead(ctx, d, meta)
	}
	return diags
}

func resourceWebCategoryProxyServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryProxyServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryProxyServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryProxyServerRead(ctx, d, meta)
	}
	return diags
}
func resourceWebCategoryProxyServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryProxyServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryProxyServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebCategoryProxyServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryProxyServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryProxyServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointWebCategoryProxyServer(d *schema.ResourceData) edpt.WebCategoryProxyServer {
	var ret edpt.WebCategoryProxyServer
	ret.Inst.AuthType = d.Get("auth_type").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	//omit encrypted
	ret.Inst.HttpPort = d.Get("http_port").(int)
	ret.Inst.HttpsPort = d.Get("https_port").(int)
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.ProxyHost = d.Get("proxy_host").(string)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	return ret
}
