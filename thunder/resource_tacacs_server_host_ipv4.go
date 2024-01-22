package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsServerHostIpv4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tacacs_server_host_ipv4`: Specify the hostname of TACACS+ server\n\n__PLACEHOLDER__",
		CreateContext: resourceTacacsServerHostIpv4Create,
		UpdateContext: resourceTacacsServerHostIpv4Update,
		ReadContext:   resourceTacacsServerHostIpv4Read,
		DeleteContext: resourceTacacsServerHostIpv4Delete,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV4 address of TACACS+ server",
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
func resourceTacacsServerHostIpv4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerHostIpv4Read(ctx, d, meta)
	}
	return diags
}

func resourceTacacsServerHostIpv4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerHostIpv4Read(ctx, d, meta)
	}
	return diags
}
func resourceTacacsServerHostIpv4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTacacsServerHostIpv4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTacacsServerHostIpv4Secret(d []interface{}) edpt.TacacsServerHostIpv4Secret {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv4Secret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.SourceIp = in["source_ip"].(string)
		ret.SourceLoopback = in["source_loopback"].(int)
		ret.SourceEth = in["source_eth"].(int)
		ret.SourceVe = in["source_ve"].(int)
		ret.SourceTrunk = in["source_trunk"].(int)
		ret.SourceLif = in["source_lif"].(int)
		ret.PortCfg = getObjectTacacsServerHostIpv4SecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTacacsServerHostIpv4SecretPortCfg(d []interface{}) edpt.TacacsServerHostIpv4SecretPortCfg {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv4SecretPortCfg
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

func dataToEndpointTacacsServerHostIpv4(d *schema.ResourceData) edpt.TacacsServerHostIpv4 {
	var ret edpt.TacacsServerHostIpv4
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Secret = getObjectTacacsServerHostIpv4Secret(d.Get("secret").([]interface{}))
	//omit uuid
	return ret
}
