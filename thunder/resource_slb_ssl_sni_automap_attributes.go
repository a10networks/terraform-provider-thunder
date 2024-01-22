package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslSniAutomapAttributes() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ssl_sni_automap_attributes`: Server Name Automap global settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSslSniAutomapAttributesCreate,
		UpdateContext: resourceSlbSslSniAutomapAttributesUpdate,
		ReadContext:   resourceSlbSslSniAutomapAttributesRead,
		DeleteContext: resourceSlbSslSniAutomapAttributesDelete,

		Schema: map[string]*schema.Schema{
			"sni_delete_factor": {
				Type: schema.TypeInt, Optional: true, Default: 50, Description: "Contexts are deleted in groups of this value. Default is 50",
			},
			"sni_lower_limit": {
				Type: schema.TypeInt, Optional: true, Default: 500, Description: "Lower limit for free SNI contexts count. Default is 500",
			},
			"sni_upper_limit": {
				Type: schema.TypeInt, Optional: true, Default: 2000, Description: "Upper limit for free SNI contexts count. Default is 2000",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbSslSniAutomapAttributesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslSniAutomapAttributesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslSniAutomapAttributes(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslSniAutomapAttributesRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSslSniAutomapAttributesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslSniAutomapAttributesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslSniAutomapAttributes(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslSniAutomapAttributesRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSslSniAutomapAttributesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslSniAutomapAttributesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslSniAutomapAttributes(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSslSniAutomapAttributesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslSniAutomapAttributesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslSniAutomapAttributes(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbSslSniAutomapAttributes(d *schema.ResourceData) edpt.SlbSslSniAutomapAttributes {
	var ret edpt.SlbSslSniAutomapAttributes
	ret.Inst.SniDeleteFactor = d.Get("sni_delete_factor").(int)
	ret.Inst.SniLowerLimit = d.Get("sni_lower_limit").(int)
	ret.Inst.SniUpperLimit = d.Get("sni_upper_limit").(int)
	//omit uuid
	return ret
}
