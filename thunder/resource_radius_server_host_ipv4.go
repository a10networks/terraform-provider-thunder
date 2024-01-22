package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRadiusServerHostIpv4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_radius_server_host_ipv4`: Specify the hostname of RADIUS server\n\n__PLACEHOLDER__",
		CreateContext: resourceRadiusServerHostIpv4Create,
		UpdateContext: resourceRadiusServerHostIpv4Update,
		ReadContext:   resourceRadiusServerHostIpv4Read,
		DeleteContext: resourceRadiusServerHostIpv4Delete,

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
	}
}
func resourceRadiusServerHostIpv4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerHostIpv4Read(ctx, d, meta)
	}
	return diags
}

func resourceRadiusServerHostIpv4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerHostIpv4Read(ctx, d, meta)
	}
	return diags
}
func resourceRadiusServerHostIpv4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRadiusServerHostIpv4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRadiusServerHostIpv4Secret(d []interface{}) edpt.RadiusServerHostIpv4Secret {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv4Secret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.PortCfg = getObjectRadiusServerHostIpv4SecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRadiusServerHostIpv4SecretPortCfg(d []interface{}) edpt.RadiusServerHostIpv4SecretPortCfg {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv4SecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthPort = in["auth_port"].(int)
		ret.AcctPort = in["acct_port"].(int)
		ret.Retransmit = in["retransmit"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointRadiusServerHostIpv4(d *schema.ResourceData) edpt.RadiusServerHostIpv4 {
	var ret edpt.RadiusServerHostIpv4
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Secret = getObjectRadiusServerHostIpv4Secret(d.Get("secret").([]interface{}))
	//omit uuid
	return ret
}
