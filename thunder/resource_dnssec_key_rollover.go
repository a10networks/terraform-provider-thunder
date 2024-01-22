package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssecKeyRollover() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_dnssec_key_rollover`: DNSSEC Key rollover\n\n__PLACEHOLDER__",
		CreateContext: resourceDnssecKeyRolloverCreate,
		UpdateContext: resourceDnssecKeyRolloverUpdate,
		ReadContext:   resourceDnssecKeyRolloverRead,
		DeleteContext: resourceDnssecKeyRolloverDelete,

		Schema: map[string]*schema.Schema{
			"dnssec_key_type": {
				Type: schema.TypeString, Optional: true, Description: "'ZSK': Zone Signing Key; 'KSK': Key Signing Key;",
			},
			"ds_ready_in_parent_zone": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "DS RR is already ready in the parent zone",
			},
			"ksk_start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "start KSK rollover in emergency mode",
			},
			"zone_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the name for the DNS zone",
			},
			"zsk_start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "start ZSK rollover in emergency mode",
			},
		},
	}
}
func resourceDnssecKeyRolloverCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecKeyRolloverCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecKeyRollover(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecKeyRolloverRead(ctx, d, meta)
	}
	return diags
}

func resourceDnssecKeyRolloverUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecKeyRolloverUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecKeyRollover(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecKeyRolloverRead(ctx, d, meta)
	}
	return diags
}
func resourceDnssecKeyRolloverDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecKeyRolloverDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecKeyRollover(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDnssecKeyRolloverRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecKeyRolloverRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecKeyRollover(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDnssecKeyRollover(d *schema.ResourceData) edpt.DnssecKeyRollover {
	var ret edpt.DnssecKeyRollover
	ret.Inst.DnssecKeyType = d.Get("dnssec_key_type").(string)
	ret.Inst.DsReadyInParentZone = d.Get("ds_ready_in_parent_zone").(int)
	ret.Inst.KskStart = d.Get("ksk_start").(int)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ZskStart = d.Get("zsk_start").(int)
	return ret
}
