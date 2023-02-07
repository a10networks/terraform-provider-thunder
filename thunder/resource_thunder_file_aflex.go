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
func resourceFileAflex() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileAflexCreate,
		ReadContext:   resourceFileAflexRead,
		UpdateContext: resourceFileAflexUpdate,
		DeleteContext: resourceFileAflexDelete,
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
		},
	}
}

func resourceFileAflexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileAflex(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAflexRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAflexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.Aflex{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAflexUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileAflex(d)
		//file.Aflex is immutable. Overwrite it.
		obj.Overwrite = 1
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAflexRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAflexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.Aflex{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToFileAflex(d *schema.ResourceData) file.Aflex {
	var ret file.Aflex
	ret.Name = d.Get("name").(string)
	ret.Protocol = d.Get("protocol").(string)
	ret.Host = d.Get("host").(string)
	ret.Path = d.Get("path").(string)
	ret.Username = d.Get("username").(string)
	ret.Password = d.Get("password").(string)
	ret.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Overwrite = d.Get("overwrite").(int)
	return ret
}
