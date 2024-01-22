package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_alg_ftp`: NAT64 FTP ALG (default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64AlgFtpCreate,
		UpdateContext: resourceCgnv6Nat64AlgFtpUpdate,
		ReadContext:   resourceCgnv6Nat64AlgFtpRead,
		DeleteContext: resourceCgnv6Nat64AlgFtpDelete,

		Schema: map[string]*schema.Schema{
			"ftp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable NAT64 FTP ALG;",
			},
			"trans_eprt_to_port": {
				Type: schema.TypeString, Optional: true, Description: "'disable': disable;",
			},
			"trans_epsv_to_pasv": {
				Type: schema.TypeString, Optional: true, Description: "'disable': disable;",
			},
			"trans_lprt_to_port": {
				Type: schema.TypeString, Optional: true, Description: "'disable': disable;",
			},
			"trans_lpsv_to_pasv": {
				Type: schema.TypeString, Optional: true, Description: "'disable': disable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"xlat_no_trans_pasv": {
				Type: schema.TypeString, Optional: true, Description: "'enable': enable;",
			},
		},
	}
}
func resourceCgnv6Nat64AlgFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64AlgFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64AlgFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64AlgFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64AlgFtp(d *schema.ResourceData) edpt.Cgnv6Nat64AlgFtp {
	var ret edpt.Cgnv6Nat64AlgFtp
	ret.Inst.FtpEnable = d.Get("ftp_enable").(string)
	ret.Inst.TransEprtToPort = d.Get("trans_eprt_to_port").(string)
	ret.Inst.TransEpsvToPasv = d.Get("trans_epsv_to_pasv").(string)
	ret.Inst.TransLprtToPort = d.Get("trans_lprt_to_port").(string)
	ret.Inst.TransLpsvToPasv = d.Get("trans_lpsv_to_pasv").(string)
	//omit uuid
	ret.Inst.XlatNoTransPasv = d.Get("xlat_no_trans_pasv").(string)
	return ret
}
