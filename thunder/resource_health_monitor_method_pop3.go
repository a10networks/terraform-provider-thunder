package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodPop3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_pop3`: POP3 type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodPop3Create,
		UpdateContext: resourceHealthMonitorMethodPop3Update,
		ReadContext:   resourceHealthMonitorMethodPop3Read,
		DeleteContext: resourceHealthMonitorMethodPop3Delete,

		Schema: map[string]*schema.Schema{
			"pop3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "POP3 type",
			},
			"pop3_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"pop3_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Specify the user password, '' means empty password",
			},
			"pop3_port": {
				Type: schema.TypeInt, Optional: true, Default: 110, Description: "Specify the POP3 port, default is 110 (Port Number (default 110))",
			},
			"pop3_username": {
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
func resourceHealthMonitorMethodPop3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodPop3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodPop3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodPop3Read(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodPop3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodPop3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodPop3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodPop3Read(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodPop3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodPop3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodPop3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodPop3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodPop3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodPop3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodPop3(d *schema.ResourceData) edpt.HealthMonitorMethodPop3 {
	var ret edpt.HealthMonitorMethodPop3
	ret.Inst.Pop3 = d.Get("pop3").(int)
	//omit pop3_encrypted
	ret.Inst.Pop3Password = d.Get("pop3_password").(int)
	ret.Inst.Pop3PasswordString = d.Get("pop3_password_string").(string)
	ret.Inst.Pop3Port = d.Get("pop3_port").(int)
	ret.Inst.Pop3Username = d.Get("pop3_username").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
