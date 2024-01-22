package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsServerHostIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tacacs_server_host_ipv6`: Specify the hostname of TACACS+ server\n\n__PLACEHOLDER__",
		CreateContext: resourceTacacsServerHostIpv6Create,
		UpdateContext: resourceTacacsServerHostIpv6Update,
		ReadContext:   resourceTacacsServerHostIpv6Read,
		DeleteContext: resourceTacacsServerHostIpv6Delete,

		Schema: map[string]*schema.Schema{
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address of TACACS+ server",
			},
			"secret": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_value": {
							Type: schema.TypeString, Optional: true, Description: "The TACACS+ server's secret",
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
func resourceTacacsServerHostIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerHostIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceTacacsServerHostIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerHostIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceTacacsServerHostIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTacacsServerHostIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerHostIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerHostIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTacacsServerHostIpv6Secret(d []interface{}) edpt.TacacsServerHostIpv6Secret {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv6Secret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.SourceIpv6 = in["source_ipv6"].(string)
		ret.SourceLoopback = in["source_loopback"].(int)
		ret.SourceEth = in["source_eth"].(int)
		ret.SourceVe = in["source_ve"].(int)
		ret.SourceTrunk = in["source_trunk"].(int)
		ret.SourceLif = in["source_lif"].(int)
		ret.PortCfg = getObjectTacacsServerHostIpv6SecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTacacsServerHostIpv6SecretPortCfg(d []interface{}) edpt.TacacsServerHostIpv6SecretPortCfg {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv6SecretPortCfg
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

func dataToEndpointTacacsServerHostIpv6(d *schema.ResourceData) edpt.TacacsServerHostIpv6 {
	var ret edpt.TacacsServerHostIpv6
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Secret = getObjectTacacsServerHostIpv6Secret(d.Get("secret").([]interface{}))
	//omit uuid
	return ret
}
