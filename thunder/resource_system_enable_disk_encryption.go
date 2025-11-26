package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemEnableDiskEncryption() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_enable_disk_encryption`: Disk encryption config\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemEnableDiskEncryptionCreate,
		UpdateContext: resourceSystemEnableDiskEncryptionUpdate,
		ReadContext:   resourceSystemEnableDiskEncryptionRead,
		DeleteContext: resourceSystemEnableDiskEncryptionDelete,

		Schema: map[string]*schema.Schema{
			"cipher": {
				Type: schema.TypeString, Optional: true, Default: "aes", Description: "'aes': cipher aes; 'serpent': cipher serpent; 'twofish': cipher twofish;",
			},
			"passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Enter phassphrase in plain text format",
			},
			"passphrase_base64": {
				Type: schema.TypeString, Optional: true, Description: "Enter phassphrase in base64 format",
			},
		},
	}
}
func resourceSystemEnableDiskEncryptionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemEnableDiskEncryptionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemEnableDiskEncryption(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemEnableDiskEncryptionRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemEnableDiskEncryptionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemEnableDiskEncryptionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemEnableDiskEncryption(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemEnableDiskEncryptionRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemEnableDiskEncryptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemEnableDiskEncryptionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemEnableDiskEncryption(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemEnableDiskEncryptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemEnableDiskEncryptionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemEnableDiskEncryption(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemEnableDiskEncryption(d *schema.ResourceData) edpt.SystemEnableDiskEncryption {
	var ret edpt.SystemEnableDiskEncryption
	ret.Inst.Cipher = d.Get("cipher").(string)
	ret.Inst.Passphrase = d.Get("passphrase").(string)
	ret.Inst.PassphraseBase64 = d.Get("passphrase_base64").(string)
	return ret
}
