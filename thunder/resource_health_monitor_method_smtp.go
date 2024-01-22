package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodSmtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_smtp`: SMTP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodSmtpCreate,
		UpdateContext: resourceHealthMonitorMethodSmtpUpdate,
		ReadContext:   resourceHealthMonitorMethodSmtpRead,
		DeleteContext: resourceHealthMonitorMethodSmtpDelete,

		Schema: map[string]*schema.Schema{
			"mail_from": {
				Type: schema.TypeString, Optional: true, Description: "Specify SMTP Sender",
			},
			"rcpt_to": {
				Type: schema.TypeString, Optional: true, Description: "Specify SMTP Receiver",
			},
			"smtp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SMTP type",
			},
			"smtp_domain": {
				Type: schema.TypeString, Optional: true, Description: "Specify domain name of 'helo' command",
			},
			"smtp_port": {
				Type: schema.TypeInt, Optional: true, Default: 25, Description: "Specify SMTP port, default is 25 (Port Number (default 25))",
			},
			"smtp_starttls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check the STARTTLS support at helo response",
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
func resourceHealthMonitorMethodSmtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSmtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSmtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodSmtpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodSmtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSmtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSmtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodSmtpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodSmtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSmtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSmtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodSmtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSmtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSmtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodSmtp(d *schema.ResourceData) edpt.HealthMonitorMethodSmtp {
	var ret edpt.HealthMonitorMethodSmtp
	ret.Inst.MailFrom = d.Get("mail_from").(string)
	ret.Inst.RcptTo = d.Get("rcpt_to").(string)
	ret.Inst.Smtp = d.Get("smtp").(int)
	ret.Inst.SmtpDomain = d.Get("smtp_domain").(string)
	ret.Inst.SmtpPort = d.Get("smtp_port").(int)
	ret.Inst.SmtpStarttls = d.Get("smtp_starttls").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
