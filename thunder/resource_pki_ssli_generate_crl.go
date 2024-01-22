package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiSsliGenerateCrl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_ssli_generate_crl`: Generate a CRL for for a vport\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiSsliGenerateCrlCreate,
		UpdateContext: resourcePkiSsliGenerateCrlUpdate,
		ReadContext:   resourcePkiSsliGenerateCrlRead,
		DeleteContext: resourcePkiSsliGenerateCrlDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "port number",
			},
			"vip_name": {
				Type: schema.TypeString, Optional: true, Description: "VIP name",
			},
		},
	}
}
func resourcePkiSsliGenerateCrlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliGenerateCrlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliGenerateCrl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiSsliGenerateCrlRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiSsliGenerateCrlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliGenerateCrlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliGenerateCrl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiSsliGenerateCrlRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiSsliGenerateCrlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliGenerateCrlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliGenerateCrl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiSsliGenerateCrlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiSsliGenerateCrlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiSsliGenerateCrl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiSsliGenerateCrl(d *schema.ResourceData) edpt.PkiSsliGenerateCrl {
	var ret edpt.PkiSsliGenerateCrl
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.VipName = d.Get("vip_name").(string)
	return ret
}
