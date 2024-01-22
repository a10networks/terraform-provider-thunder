package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsServerHostTacacsHostname() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tacacs_server_host_tacacs_hostname`: Specify the hostname of TACACS+ server\n\n__PLACEHOLDER__",
		CreateContext: resourceTacacsServerHostTacacsHostnameCreate,
		UpdateContext: resourceTacacsServerHostTacacsHostnameUpdate,
		ReadContext:   resourceTacacsServerHostTacacsHostnameRead,
		DeleteContext: resourceTacacsServerHostTacacsHostnameDelete,

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type: schema.TypeString, Required: true, Description: "Hostname of TACACS+ server",
			},
			"secret": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_value": {
							Type: schema.TypeString, Optional: true, Description: "The TACACS+ server's secret",
						},
						"source_ip": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"source_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 address",
						},
						"source_loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"source_eth": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"source_ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
						},
						"source_trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"source_lif": {
							Type: schema.TypeInt, Optional: true, Description: "Logical interface (Lif interface number)",
						},
						"port_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type: schema.TypeInt, Optional: true, Default: 49, Description: "Specify the port number used by TACACS+ server.( default port is 49) (Port number (default 49))",
									},
									"timeout": {
										Type: schema.TypeInt, Optional: true, Default: 12, Description: "Specify the maximum time allowed for setting up a connection with the TACACS+ server. (default timeout is 12 seconds) (Maximum time allowed for setting up a connection with the TACACS+ server, in seconds (default 12))",
									},
									"monitor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify monitor TACACS+ server",
									},
									"username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
									},
									"password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
									},
									"password_value": {
										Type: schema.TypeString, Optional: true, Description: "The user password",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTacacsServerHostTacacsHostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostTacacsHostnameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostTacacsHostname(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerHostTacacsHostnameRead(ctx, d, meta)
	}
	return diags
}

func resourceTacacsServerHostTacacsHostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostTacacsHostnameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostTacacsHostname(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerHostTacacsHostnameRead(ctx, d, meta)
	}
	return diags
}
func resourceTacacsServerHostTacacsHostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostTacacsHostnameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostTacacsHostname(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTacacsServerHostTacacsHostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostTacacsHostnameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostTacacsHostname(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTacacsServerHostTacacsHostnameSecret(d []interface{}) edpt.TacacsServerHostTacacsHostnameSecret {

	count1 := len(d)
	var ret edpt.TacacsServerHostTacacsHostnameSecret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.SourceIp = in["source_ip"].(string)
		ret.SourceIpv6 = in["source_ipv6"].(string)
		ret.SourceLoopback = in["source_loopback"].(int)
		ret.SourceEth = in["source_eth"].(int)
		ret.SourceVe = in["source_ve"].(int)
		ret.SourceTrunk = in["source_trunk"].(int)
		ret.SourceLif = in["source_lif"].(int)
		ret.PortCfg = getObjectTacacsServerHostTacacsHostnameSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTacacsServerHostTacacsHostnameSecretPortCfg(d []interface{}) edpt.TacacsServerHostTacacsHostnameSecretPortCfg {

	count1 := len(d)
	var ret edpt.TacacsServerHostTacacsHostnameSecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Monitor = in["monitor"].(int)
		ret.Username = in["username"].(string)
		ret.Password = in["password"].(int)
		ret.PasswordValue = in["password_value"].(string)
		//omit encrypted
	}
	return ret
}

func dataToEndpointTacacsServerHostTacacsHostname(d *schema.ResourceData) edpt.TacacsServerHostTacacsHostname {
	var ret edpt.TacacsServerHostTacacsHostname
	ret.Inst.Hostname = d.Get("hostname").(string)
	ret.Inst.Secret = getObjectTacacsServerHostTacacsHostnameSecret(d.Get("secret").([]interface{}))
	//omit uuid
	return ret
}
