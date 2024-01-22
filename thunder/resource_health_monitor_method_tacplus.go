package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodTacplus() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_tacplus`: TACACS+ type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodTacplusCreate,
		UpdateContext: resourceHealthMonitorMethodTacplusUpdate,
		ReadContext:   resourceHealthMonitorMethodTacplusRead,
		DeleteContext: resourceHealthMonitorMethodTacplusDelete,

		Schema: map[string]*schema.Schema{
			"tacplus": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "TACACS+ type",
			},
			"tacplus_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"tacplus_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
			},
			"tacplus_port": {
				Type: schema.TypeInt, Optional: true, Default: 49, Description: "Specify the TACACS+ port, default 49 (Port number (default 49))",
			},
			"tacplus_secret": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the shared secret of TACACS+ server",
			},
			"tacplus_secret_string": {
				Type: schema.TypeString, Optional: true, Description: "Shared Crypto Key",
			},
			"tacplus_type": {
				Type: schema.TypeString, Optional: true, Default: "inbound-ascii-login", Description: "'inbound-ascii-login': Specify Inbound ASCII Login type;",
			},
			"tacplus_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the username",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodTacplusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTacplusCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTacplus(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodTacplusRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodTacplusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTacplusUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTacplus(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodTacplusRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodTacplusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTacplusDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTacplus(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodTacplusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTacplusRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTacplus(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodTacplus(d *schema.ResourceData) edpt.HealthMonitorMethodTacplus {
	var ret edpt.HealthMonitorMethodTacplus
	//omit secret_encrypted
	ret.Inst.Tacplus = d.Get("tacplus").(int)
	//omit tacplus_encrypted
	ret.Inst.TacplusPassword = d.Get("tacplus_password").(int)
	ret.Inst.TacplusPasswordString = d.Get("tacplus_password_string").(string)
	ret.Inst.TacplusPort = d.Get("tacplus_port").(int)
	ret.Inst.TacplusSecret = d.Get("tacplus_secret").(int)
	ret.Inst.TacplusSecretString = d.Get("tacplus_secret_string").(string)
	ret.Inst.TacplusType = d.Get("tacplus_type").(string)
	ret.Inst.TacplusUsername = d.Get("tacplus_username").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
