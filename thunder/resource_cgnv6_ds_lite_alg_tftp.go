package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteAlgTftp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_alg_tftp`: DS-Lite TFTP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteAlgTftpCreate,
		UpdateContext: resourceCgnv6DsLiteAlgTftpUpdate,
		ReadContext:   resourceCgnv6DsLiteAlgTftpRead,
		DeleteContext: resourceCgnv6DsLiteAlgTftpDelete,

		Schema: map[string]*schema.Schema{
			"tftp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteAlgTftpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgTftpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgTftp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgTftpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteAlgTftpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgTftpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgTftp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgTftpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteAlgTftpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgTftpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgTftp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteAlgTftpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgTftpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgTftp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLiteAlgTftp(d *schema.ResourceData) edpt.Cgnv6DsLiteAlgTftp {
	var ret edpt.Cgnv6DsLiteAlgTftp
	ret.Inst.TftpEnable = d.Get("tftp_enable").(string)
	//omit uuid
	return ret
}
