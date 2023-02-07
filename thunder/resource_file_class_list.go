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
func resourceFileClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_class_list`: Class-List file\n\n",
		CreateContext: resourceFileClassListCreate,
		ReadContext:   resourceFileClassListRead,
		UpdateContext: resourceFileClassListUpdate,
		DeleteContext: resourceFileClassListDelete,
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

func resourceFileClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceFileClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.ClassList{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileClassList(d)
		//file.ClassList is immutable. Overwrite it.
		obj.Overwrite = 1
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceFileClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.ClassList{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToFileClassList(d *schema.ResourceData) file.ClassList {
	var ret file.ClassList
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
