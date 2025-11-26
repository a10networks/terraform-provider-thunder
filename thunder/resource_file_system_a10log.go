package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSystemA10log() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_system_a10log`: Manage logs of ACOS components\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSystemA10logCreate,
		UpdateContext: resourceFileSystemA10logUpdate,
		ReadContext:   resourceFileSystemA10logRead,
		DeleteContext: resourceFileSystemA10logDelete,

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
func resourceFileSystemA10logCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10log(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemA10logRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSystemA10logUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10log(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemA10logRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSystemA10logDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10log(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSystemA10logRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10log(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSystemA10log(d *schema.ResourceData) edpt.FileSystemA10log {
	var ret edpt.FileSystemA10log
	ret.Inst.Filename = d.Get("filename").(string)
	ret.Inst.Rm = d.Get("rm").(int)
	//omit uuid
	return ret
}
