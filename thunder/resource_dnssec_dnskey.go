package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssecDnskey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_dnssec_dnskey`: DNSKEY Resource Records of child zones\n\n__PLACEHOLDER__",
		CreateContext: resourceDnssecDnskeyCreate,
		UpdateContext: resourceDnssecDnskeyUpdate,
		ReadContext:   resourceDnssecDnskeyRead,
		DeleteContext: resourceDnssecDnskeyDelete,

		Schema: map[string]*schema.Schema{
			"key_delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete the DNSKEY file",
			},
			"zone_name": {
				Type: schema.TypeString, Optional: true, Description: "DNS zone name of the child zone",
			},
		},
	}
}
func resourceDnssecDnskeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDnskeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDnskey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecDnskeyRead(ctx, d, meta)
	}
	return diags
}

func resourceDnssecDnskeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDnskeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDnskey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecDnskeyRead(ctx, d, meta)
	}
	return diags
}
func resourceDnssecDnskeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDnskeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDnskey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDnssecDnskeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDnskeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDnskey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDnssecDnskey(d *schema.ResourceData) edpt.DnssecDnskey {
	var ret edpt.DnssecDnskey
	ret.Inst.KeyDelete = d.Get("key_delete").(int)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
