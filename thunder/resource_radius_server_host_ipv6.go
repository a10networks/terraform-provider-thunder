package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRadiusServerHostIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_radius_server_host_ipv6`: Specify the hostname of RADIUS server\n\n__PLACEHOLDER__",
		CreateContext: resourceRadiusServerHostIpv6Create,
		UpdateContext: resourceRadiusServerHostIpv6Update,
		ReadContext:   resourceRadiusServerHostIpv6Read,
		DeleteContext: resourceRadiusServerHostIpv6Delete,

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
	}
}
func resourceRadiusServerHostIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerHostIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceRadiusServerHostIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerHostIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceRadiusServerHostIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRadiusServerHostIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRadiusServerHostIpv6Secret(d []interface{}) edpt.RadiusServerHostIpv6Secret {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv6Secret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.PortCfg = getObjectRadiusServerHostIpv6SecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRadiusServerHostIpv6SecretPortCfg(d []interface{}) edpt.RadiusServerHostIpv6SecretPortCfg {

	count1 := len(d)
	var ret edpt.RadiusServerHostIpv6SecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthPort = in["auth_port"].(int)
		ret.AcctPort = in["acct_port"].(int)
		ret.Retransmit = in["retransmit"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointRadiusServerHostIpv6(d *schema.ResourceData) edpt.RadiusServerHostIpv6 {
	var ret edpt.RadiusServerHostIpv6
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Secret = getObjectRadiusServerHostIpv6Secret(d.Get("secret").([]interface{}))
	//omit uuid
	return ret
}
