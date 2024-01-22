package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssecDs() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_dnssec_ds`: Delegation Signer(DS) Resource Records of child zones\n\n__PLACEHOLDER__",
		CreateContext: resourceDnssecDsCreate,
		UpdateContext: resourceDnssecDsUpdate,
		ReadContext:   resourceDnssecDsRead,
		DeleteContext: resourceDnssecDsDelete,

		Schema: map[string]*schema.Schema{
			"ds_delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete the DS file",
			},
			"zone_name": {
				Type: schema.TypeString, Optional: true, Description: "DNS zone name of the child zone",
			},
		},
	}
}
func resourceDnssecDsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDs(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecDsRead(ctx, d, meta)
	}
	return diags
}

func resourceDnssecDsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDs(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecDsRead(ctx, d, meta)
	}
	return diags
}
func resourceDnssecDsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDs(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDnssecDsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecDs(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDnssecDs(d *schema.ResourceData) edpt.DnssecDs {
	var ret edpt.DnssecDs
	ret.Inst.DsDelete = d.Get("ds_delete").(int)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
