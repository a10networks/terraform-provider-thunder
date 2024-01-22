package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAwsAccesskeyLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_aws_accesskey_local`: The aws accesskey for admin user\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAwsAccesskeyLocalCreate,
		UpdateContext: resourceFileAwsAccesskeyLocalUpdate,
		ReadContext:   resourceFileAwsAccesskeyLocalRead,
		DeleteContext: resourceFileAwsAccesskeyLocalDelete,

		Schema: map[string]*schema.Schema{
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"user": {
				Type: schema.TypeString, Optional: true, Description: "user name of the pub key",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileAwsAccesskeyLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAwsAccesskeyLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAwsAccesskeyLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAwsAccesskeyLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAwsAccesskeyLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAwsAccesskeyLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileAwsAccesskeyLocal(d *schema.ResourceData) edpt.FileAwsAccesskeyLocal {
	var ret edpt.FileAwsAccesskeyLocal
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
	return ret
}
