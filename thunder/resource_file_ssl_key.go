package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/file"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//Not direct mapping to any axapi endpoint
func resourceFileSslKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileSslKeyCreate,
		ReadContext:   resourceFileSslKeyRead,
		UpdateContext: resourceFileSslKeyUpdate,
		DeleteContext: resourceFileSslKeyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Local file name",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Transfer protocol",
				ValidateFunc: validation.StringInSlice([]string{"ftp", "http", "https", "scp", "sftp", "tftp"}, false),
			},
			"host": {
				Type: schema.TypeString, Required: true, Description: "Remote site (IP or domain name)",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"path": {
				Type: schema.TypeString, Required: true, Description: "Remote path",
				ValidateFunc: axapi.IsAbsolutePath,
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Username for the remote site",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "Password for the remote site",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark as non-exportable",
				ValidateFunc: validation.IntBetween(0, 1),
			},
		},
	}
}

func resourceFileSslKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileSslKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.SslKey{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSslKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileSslKey(d)
		//file.SslKey is immutable. Overwrite it.
		obj.Overwrite = 1
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.SslKey{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToFileSslKey(d *schema.ResourceData) file.SslKey {
	var ret file.SslKey
	ret.Name = d.Get("name").(string)
	ret.Protocol = d.Get("protocol").(string)
	ret.Host = d.Get("host").(string)
	ret.Path = d.Get("path").(string)
	ret.Username = d.Get("username").(string)
	ret.Password = d.Get("password").(string)
	ret.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Overwrite = d.Get("overwrite").(int)
	ret.Secured = d.Get("secured").(int)
	return ret
}
