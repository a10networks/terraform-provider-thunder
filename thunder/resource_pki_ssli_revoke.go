package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiSsliRevoke() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_ssli_revoke`: Revoke an SSLi certificate\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiSsliRevokeCreate,
		UpdateContext: resourcePkiSsliRevokeUpdate,
		ReadContext:   resourcePkiSsliRevokeRead,
		DeleteContext: resourcePkiSsliRevokeDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "port number",
			},
			"serial": {
				Type: schema.TypeString, Optional: true, Description: "Serial number in hex",
			},
			"vip_name": {
				Type: schema.TypeString, Optional: true, Description: "VIP name",
			},
		},
	}
}
func resourcePkiSsliRevokeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliRevokeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliRevoke(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiSsliRevokeRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiSsliRevokeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliRevokeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliRevoke(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiSsliRevokeRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiSsliRevokeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliRevokeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliRevoke(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiSsliRevokeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliRevokeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliRevoke(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiSsliRevoke(d *schema.ResourceData) edpt.PkiSsliRevoke {
	var ret edpt.PkiSsliRevoke
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Serial = d.Get("serial").(string)
	ret.Inst.VipName = d.Get("vip_name").(string)
	return ret
}
