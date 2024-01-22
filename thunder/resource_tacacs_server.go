package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tacacs_server`: Configure TACACS+ servers\n\n__PLACEHOLDER__",
		CreateContext: resourceTacacsServerCreate,
		UpdateContext: resourceTacacsServerUpdate,
		ReadContext:   resourceTacacsServerRead,
		DeleteContext: resourceTacacsServerDelete,

		Schema: map[string]*schema.Schema{
			"host": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
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
							},
						},
						"ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
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
							},
						},
						"tacacs_hostname_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
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
							},
						},
					},
				},
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "The moniter interval in seconds (default 60)",
			},
			"monitor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure TACACS+ servers",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTacacsServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerRead(ctx, d, meta)
	}
	return diags
}

func resourceTacacsServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTacacsServerRead(ctx, d, meta)
	}
	return diags
}
func resourceTacacsServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTacacsServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTacacsServerHost1896(d []interface{}) edpt.TacacsServerHost1896 {

	count1 := len(d)
	var ret edpt.TacacsServerHost1896
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4List = getSliceTacacsServerHostIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceTacacsServerHostIpv6List(in["ipv6_list"].([]interface{}))
		ret.TacacsHostnameList = getSliceTacacsServerHostTacacsHostnameList(in["tacacs_hostname_list"].([]interface{}))
	}
	return ret
}

func getSliceTacacsServerHostIpv4List(d []interface{}) []edpt.TacacsServerHostIpv4List {

	count1 := len(d)
	ret := make([]edpt.TacacsServerHostIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TacacsServerHostIpv4List
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Secret = getObjectTacacsServerHostIpv4ListSecret(in["secret"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectTacacsServerHostIpv4ListSecret(d []interface{}) edpt.TacacsServerHostIpv4ListSecret {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv4ListSecret
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
		ret.PortCfg = getObjectTacacsServerHostIpv4ListSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTacacsServerHostIpv4ListSecretPortCfg(d []interface{}) edpt.TacacsServerHostIpv4ListSecretPortCfg {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv4ListSecretPortCfg
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

func getSliceTacacsServerHostIpv6List(d []interface{}) []edpt.TacacsServerHostIpv6List {

	count1 := len(d)
	ret := make([]edpt.TacacsServerHostIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TacacsServerHostIpv6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Secret = getObjectTacacsServerHostIpv6ListSecret(in["secret"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectTacacsServerHostIpv6ListSecret(d []interface{}) edpt.TacacsServerHostIpv6ListSecret {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv6ListSecret
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
		ret.PortCfg = getObjectTacacsServerHostIpv6ListSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTacacsServerHostIpv6ListSecretPortCfg(d []interface{}) edpt.TacacsServerHostIpv6ListSecretPortCfg {

	count1 := len(d)
	var ret edpt.TacacsServerHostIpv6ListSecretPortCfg
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

func getSliceTacacsServerHostTacacsHostnameList(d []interface{}) []edpt.TacacsServerHostTacacsHostnameList {

	count1 := len(d)
	ret := make([]edpt.TacacsServerHostTacacsHostnameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TacacsServerHostTacacsHostnameList
		oi.Hostname = in["hostname"].(string)
		oi.Secret = getObjectTacacsServerHostTacacsHostnameListSecret(in["secret"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectTacacsServerHostTacacsHostnameListSecret(d []interface{}) edpt.TacacsServerHostTacacsHostnameListSecret {

	count1 := len(d)
	var ret edpt.TacacsServerHostTacacsHostnameListSecret
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
		ret.PortCfg = getObjectTacacsServerHostTacacsHostnameListSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTacacsServerHostTacacsHostnameListSecretPortCfg(d []interface{}) edpt.TacacsServerHostTacacsHostnameListSecretPortCfg {

	count1 := len(d)
	var ret edpt.TacacsServerHostTacacsHostnameListSecretPortCfg
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

func dataToEndpointTacacsServer(d *schema.ResourceData) edpt.TacacsServer {
	var ret edpt.TacacsServer
	ret.Inst.Host = getObjectTacacsServerHost1896(d.Get("host").([]interface{}))
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Monitor = d.Get("monitor").(int)
	//omit uuid
	return ret
}
