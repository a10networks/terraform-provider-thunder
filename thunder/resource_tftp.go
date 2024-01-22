package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTftp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tftp`: TFTP client configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceTftpCreate,
		UpdateContext: resourceTftpUpdate,
		ReadContext:   resourceTftpRead,
		DeleteContext: resourceTftpDelete,

		Schema: map[string]*schema.Schema{
			"blksize": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "TFTP client block size value (Block size (Blksize/Max file size. Example: 1K/64M, 8K/512M, 32K/2048M),default is 32768)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTftpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTftpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTftp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTftpRead(ctx, d, meta)
	}
	return diags
}

func resourceTftpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTftpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTftp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTftpRead(ctx, d, meta)
	}
	return diags
}
func resourceTftpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTftpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTftp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTftpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTftpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTftp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTftp(d *schema.ResourceData) edpt.Tftp {
	var ret edpt.Tftp
	ret.Inst.Blksize = d.Get("blksize").(int)
	//omit uuid
	return ret
}
