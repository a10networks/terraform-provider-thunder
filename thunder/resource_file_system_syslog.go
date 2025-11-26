package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSystemSyslog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_system_syslog`: Manage var logs of Linux system\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSystemSyslogCreate,
		UpdateContext: resourceFileSystemSyslogUpdate,
		ReadContext:   resourceFileSystemSyslogRead,
		DeleteContext: resourceFileSystemSyslogDelete,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type: schema.TypeString, Optional: true, Description: "Filename to remove",
			},
			"rm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove file",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileSystemSyslogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemSyslogRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSystemSyslogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemSyslogRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSystemSyslogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSystemSyslogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSystemSyslog(d *schema.ResourceData) edpt.FileSystemSyslog {
	var ret edpt.FileSystemSyslog
	ret.Inst.Filename = d.Get("filename").(string)
	ret.Inst.Rm = d.Get("rm").(int)
	//omit uuid
	return ret
}
