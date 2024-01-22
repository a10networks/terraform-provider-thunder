package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSyslogLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_syslog_local`: syslog file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSyslogLocalCreate,
		UpdateContext: resourceFileSyslogLocalUpdate,
		ReadContext:   resourceFileSyslogLocalRead,
		DeleteContext: resourceFileSyslogLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'export': export;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "syslog local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
		},
	}
}
func resourceFileSyslogLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSyslogLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSyslogLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSyslogLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSyslogLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSyslogLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSyslogLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSyslogLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSyslogLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSyslogLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSyslogLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSyslogLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSyslogLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSyslogLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSyslogLocal(d *schema.ResourceData) edpt.FileSyslogLocal {
	var ret edpt.FileSyslogLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	return ret
}
