package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRadiusServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_radius_server`: Configure RADIUS server information\n\n__PLACEHOLDER__",
		CreateContext: resourceRadiusServerCreate,
		UpdateContext: resourceRadiusServerUpdate,
		ReadContext:   resourceRadiusServerRead,
		DeleteContext: resourceRadiusServerDelete,

		Schema: map[string]*schema.Schema{
			"default_privilege_read_write": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the RADIUS default privilege",
			},
			"host": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_addr": {
										Type: schema.TypeString, Required: true, Description: "IPV4 address of RADIUS server",
									},
									"secret": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"secret_value": {
													Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
												},
												"port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"auth_port": {
																Type: schema.TypeInt, Optional: true, Description: "Specify the RADIUS server's authentication port (default 1812)",
															},
															"acct_port": {
																Type: schema.TypeInt, Optional: true, Description: "Specify the RADIUS server's accounting port (default 1813)",
															},
															"retransmit": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the maximum times allowed for resending an request to the radius server (default 3)",
															},
															"timeout": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the maximum time allowed for waiting for a response from a radius server (default 3)",
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
										Type: schema.TypeString, Required: true, Description: "IPV6 address of RADIUS server",
									},
									"secret": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"secret_value": {
													Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
												},
												"port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"auth_port": {
																Type: schema.TypeInt, Optional: true, Description: "Specify the RADIUS server's authentication port (default 1812)",
															},
															"acct_port": {
																Type: schema.TypeInt, Optional: true, Description: "Specify the RADIUS server's accounting port (default 1813)",
															},
															"retransmit": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the maximum times allowed for resending an request to the radius server (default 3)",
															},
															"timeout": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the maximum time allowed for waiting for a response from a radius server (default 3)",
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
						"name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostname": {
										Type: schema.TypeString, Required: true, Description: "Hostname of RADIUS server",
									},
									"secret": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"secret_value": {
													Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
												},
												"port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"auth_port": {
																Type: schema.TypeInt, Optional: true, Description: "Specify the RADIUS server's authentication port (default 1812)",
															},
															"acct_port": {
																Type: schema.TypeInt, Optional: true, Description: "Specify the RADIUS server's accounting port (default 1813)",
															},
															"retransmit": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the maximum times allowed for resending an request to the radius server (default 3)",
															},
															"timeout": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the maximum time allowed for waiting for a response from a radius server (default 3)",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRadiusServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerRead(ctx, d, meta)
	}
	return diags
}

func resourceRadiusServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerRead(ctx, d, meta)
	}
	return diags
}
func resourceRadiusServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRadiusServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRadiusServerHost1093(d []interface{}) edpt.RadiusServerHost1093 {

	count1 := len(d)
	var ret edpt.RadiusServerHost1093
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4List = getSliceRadiusServerHostIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceRadiusServerHostIpv6List(in["ipv6_list"].([]interface{}))
		ret.NameList = getSliceRadiusServerHostNameList(in["name_list"].([]interface{}))
	}
	return ret
}

func getSliceRadiusServerHostIpv4List(d []interface{}) []edpt.RadiusServerHostIpv4List {

	count1 := len(d)
	ret := make([]edpt.RadiusServerHostIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RadiusServerHostIpv4List
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Secret = getObjectRadiusServerHostIpv4ListSecret(in["secret"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRadiusServerHostIpv4ListSecret(d []interface{}) edpt.RadiusServerHostIpv4ListSecret {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv4ListSecret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.PortCfg = getObjectRadiusServerHostIpv4ListSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRadiusServerHostIpv4ListSecretPortCfg(d []interface{}) edpt.RadiusServerHostIpv4ListSecretPortCfg {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv4ListSecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthPort = in["auth_port"].(int)
		ret.AcctPort = in["acct_port"].(int)
		ret.Retransmit = in["retransmit"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getSliceRadiusServerHostIpv6List(d []interface{}) []edpt.RadiusServerHostIpv6List {

	count1 := len(d)
	ret := make([]edpt.RadiusServerHostIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RadiusServerHostIpv6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Secret = getObjectRadiusServerHostIpv6ListSecret(in["secret"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRadiusServerHostIpv6ListSecret(d []interface{}) edpt.RadiusServerHostIpv6ListSecret {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv6ListSecret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.PortCfg = getObjectRadiusServerHostIpv6ListSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRadiusServerHostIpv6ListSecretPortCfg(d []interface{}) edpt.RadiusServerHostIpv6ListSecretPortCfg {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv6ListSecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthPort = in["auth_port"].(int)
		ret.AcctPort = in["acct_port"].(int)
		ret.Retransmit = in["retransmit"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getSliceRadiusServerHostNameList(d []interface{}) []edpt.RadiusServerHostNameList {

	count1 := len(d)
	ret := make([]edpt.RadiusServerHostNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RadiusServerHostNameList
		oi.Hostname = in["hostname"].(string)
		oi.Secret = getObjectRadiusServerHostNameListSecret(in["secret"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRadiusServerHostNameListSecret(d []interface{}) edpt.RadiusServerHostNameListSecret {

	count1 := len(d)
	var ret edpt.RadiusServerHostNameListSecret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.PortCfg = getObjectRadiusServerHostNameListSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRadiusServerHostNameListSecretPortCfg(d []interface{}) edpt.RadiusServerHostNameListSecretPortCfg {

	count1 := len(d)
	var ret edpt.RadiusServerHostNameListSecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthPort = in["auth_port"].(int)
		ret.AcctPort = in["acct_port"].(int)
		ret.Retransmit = in["retransmit"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointRadiusServer(d *schema.ResourceData) edpt.RadiusServer {
	var ret edpt.RadiusServer
	ret.Inst.DefaultPrivilegeReadWrite = d.Get("default_privilege_read_write").(int)
	ret.Inst.Host = getObjectRadiusServerHost1093(d.Get("host").([]interface{}))
	//omit uuid
	return ret
}
