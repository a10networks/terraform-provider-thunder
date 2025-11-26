package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpsecFpgaDecrypt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipsec_fpga_decrypt`: Enable or disable FPGA decryption offload\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpsecFpgaDecryptCreate,
		UpdateContext: resourceSystemIpsecFpgaDecryptUpdate,
		ReadContext:   resourceSystemIpsecFpgaDecryptRead,
		DeleteContext: resourceSystemIpsecFpgaDecryptDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable FPGA decryption offload; 'disable': Disable FPGA decryption offload;",
			},
		},
	}
}
func resourceSystemIpsecFpgaDecryptCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecFpgaDecryptCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsecFpgaDecrypt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpsecFpgaDecryptRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpsecFpgaDecryptUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecFpgaDecryptUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsecFpgaDecrypt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpsecFpgaDecryptRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpsecFpgaDecryptDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecFpgaDecryptDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsecFpgaDecrypt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpsecFpgaDecryptRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecFpgaDecryptRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsecFpgaDecrypt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIpsecFpgaDecrypt(d *schema.ResourceData) edpt.SystemIpsecFpgaDecrypt {
	var ret edpt.SystemIpsecFpgaDecrypt
	ret.Inst.Action = d.Get("action").(string)
	return ret
}
