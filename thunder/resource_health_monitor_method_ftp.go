package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_ftp`: FTP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodFtpCreate,
		UpdateContext: resourceHealthMonitorMethodFtpUpdate,
		ReadContext:   resourceHealthMonitorMethodFtpRead,
		DeleteContext: resourceHealthMonitorMethodFtpDelete,

		Schema: map[string]*schema.Schema{
			"ftp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FTP type",
			},
			"ftp_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"ftp_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Specify the user password, '' means empty password",
			},
			"ftp_port": {
				Type: schema.TypeInt, Optional: true, Default: 21, Description: "Specify FTP port (Specify port number, default is 21)",
			},
			"ftp_username": {
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
func resourceHealthMonitorMethodFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodFtp(d *schema.ResourceData) edpt.HealthMonitorMethodFtp {
	var ret edpt.HealthMonitorMethodFtp
	ret.Inst.Ftp = d.Get("ftp").(int)
	//omit ftp_encrypted
	ret.Inst.FtpPassword = d.Get("ftp_password").(int)
	ret.Inst.FtpPasswordString = d.Get("ftp_password_string").(string)
	ret.Inst.FtpPort = d.Get("ftp_port").(int)
	ret.Inst.FtpUsername = d.Get("ftp_username").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
