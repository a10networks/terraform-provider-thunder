package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRadiusServerHostName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_radius_server_host_name`: Specify the hostname of RADIUS server\n\n__PLACEHOLDER__",
		CreateContext: resourceRadiusServerHostNameCreate,
		UpdateContext: resourceRadiusServerHostNameUpdate,
		ReadContext:   resourceRadiusServerHostNameRead,
		DeleteContext: resourceRadiusServerHostNameDelete,

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
	}
}
func resourceRadiusServerHostNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerHostNameRead(ctx, d, meta)
	}
	return diags
}

func resourceRadiusServerHostNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRadiusServerHostNameRead(ctx, d, meta)
	}
	return diags
}
func resourceRadiusServerHostNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRadiusServerHostNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRadiusServerHostNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRadiusServerHostName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRadiusServerHostNameSecret(d []interface{}) edpt.RadiusServerHostNameSecret {

	count1 := len(d)
	var ret edpt.RadiusServerHostNameSecret
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecretValue = in["secret_value"].(string)
		//omit encrypted
		ret.PortCfg = getObjectRadiusServerHostNameSecretPortCfg(in["port_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRadiusServerHostNameSecretPortCfg(d []interface{}) edpt.RadiusServerHostNameSecretPortCfg {

	count1 := len(d)
	var ret edpt.RadiusServerHostNameSecretPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthPort = in["auth_port"].(int)
		ret.AcctPort = in["acct_port"].(int)
		ret.Retransmit = in["retransmit"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointRadiusServerHostName(d *schema.ResourceData) edpt.RadiusServerHostName {
	var ret edpt.RadiusServerHostName
	ret.Inst.Hostname = d.Get("hostname").(string)
	ret.Inst.Secret = getObjectRadiusServerHostNameSecret(d.Get("secret").([]interface{}))
	//omit uuid
	return ret
}
