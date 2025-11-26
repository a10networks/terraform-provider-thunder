package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSshPubkeyLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_ssh_pubkey_local`: The ssh pubkey for admin user\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSshPubkeyLocalCreate,
		UpdateContext: resourceFileSshPubkeyLocalUpdate,
		ReadContext:   resourceFileSshPubkeyLocalRead,
		DeleteContext: resourceFileSshPubkeyLocalDelete,

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
func resourceFileSshPubkeyLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSshPubkeyLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSshPubkeyLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSshPubkeyLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSshPubkeyLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSshPubkeyLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSshPubkeyLocal(d *schema.ResourceData) edpt.FileSshPubkeyLocal {
	var ret edpt.FileSshPubkeyLocal
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
	return ret
}
