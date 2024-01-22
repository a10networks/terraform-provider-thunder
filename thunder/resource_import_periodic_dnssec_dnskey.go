package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicDnssecDnskey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_dnssec_dnskey`: DNSSEC DNSKEY(KSK) file for child zone\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicDnssecDnskeyCreate,
		UpdateContext: resourceImportPeriodicDnssecDnskeyUpdate,
		ReadContext:   resourceImportPeriodicDnssecDnskeyRead,
		DeleteContext: resourceImportPeriodicDnssecDnskeyDelete,

		Schema: map[string]*schema.Schema{
			"dnssec_dnskey": {
				Type: schema.TypeString, Required: true, Description: "DNSSEC DNSKEY(KSK) file for child zone",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceImportPeriodicDnssecDnskeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicDnssecDnskeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicDnssecDnskey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicDnssecDnskeyRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicDnssecDnskeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicDnssecDnskeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicDnssecDnskey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicDnssecDnskeyRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicDnssecDnskeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicDnssecDnskeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicDnssecDnskey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicDnssecDnskeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicDnssecDnskeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicDnssecDnskey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicDnssecDnskey(d *schema.ResourceData) edpt.ImportPeriodicDnssecDnskey {
	var ret edpt.ImportPeriodicDnssecDnskey
	ret.Inst.DnssecDnskey = d.Get("dnssec_dnskey").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
