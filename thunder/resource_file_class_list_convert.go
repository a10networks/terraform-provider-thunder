package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/file"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Not direct mapping to any axapi endpoint
func resourceFileClassListConvert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_class_list_convert`: Upload and convert a class list file to A10 format\n\n",
		CreateContext: resourceFileClassListConvertCreate,
		ReadContext:   resourceFileClassListConvertRead,
		UpdateContext: resourceFileClassListConvertUpdate,
		DeleteContext: resourceFileClassListConvertDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Local file name",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"class_list_type": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Type of this class-list file",
				ValidateFunc: validation.StringInSlice([]string{"ac", "ipv4", "ipv6", "string", "string-case-insensitive"}, false),
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

func resourceFileClassListConvertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListConvertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileClassListConvert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileClassListConvertRead(ctx, d, meta)
	}
	return diags
}

func resourceFileClassListConvertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListConvertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.ClassListConvert{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileClassListConvertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListConvertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileClassListConvert(d)
		//file.ClassList is immutable. Overwrite it.
		obj.Overwrite = 1
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileClassListConvertRead(ctx, d, meta)
	}
	return diags
}

func resourceFileClassListConvertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListConvertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.ClassListConvert{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToFileClassListConvert(d *schema.ResourceData) file.ClassListConvert {
	var ret file.ClassListConvert
	ret.Name = d.Get("name").(string)
	ret.ClassListType = d.Get("class_list_type").(string)
	ret.Protocol = d.Get("protocol").(string)
	ret.Host = d.Get("host").(string)
	ret.Path = d.Get("path").(string)
	ret.Username = d.Get("username").(string)
	ret.Password = d.Get("password").(string)
	ret.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Overwrite = d.Get("overwrite").(int)
	return ret
}
