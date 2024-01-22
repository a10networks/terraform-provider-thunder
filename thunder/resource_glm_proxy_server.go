package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlmProxyServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glm_proxy_server`: Connect to GLM through proxy server\n\n__PLACEHOLDER__",
		CreateContext: resourceGlmProxyServerCreate,
		UpdateContext: resourceGlmProxyServerUpdate,
		ReadContext:   resourceGlmProxyServerRead,
		DeleteContext: resourceGlmProxyServerDelete,

		Schema: map[string]*schema.Schema{
			"host": {
				Type: schema.TypeString, Optional: true, Description: "Proxy server hostname or IP address",
			},
			"password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Password for proxy authentication",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Proxy server port",
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
func resourceGlmProxyServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmProxyServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmProxyServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmProxyServerRead(ctx, d, meta)
	}
	return diags
}

func resourceGlmProxyServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmProxyServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmProxyServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmProxyServerRead(ctx, d, meta)
	}
	return diags
}
func resourceGlmProxyServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmProxyServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmProxyServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlmProxyServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmProxyServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmProxyServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGlmProxyServer(d *schema.ResourceData) edpt.GlmProxyServer {
	var ret edpt.GlmProxyServer
	//omit encrypted
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	return ret
}
