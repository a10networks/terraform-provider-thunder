package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsARecordDnsARecordIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6`: Specify DNS Address Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsARecordDnsARecordIpv6Create,
		UpdateContext: resourceGslbZoneServiceDnsARecordDnsARecordIpv6Update,
		ReadContext:   resourceGslbZoneServiceDnsARecordDnsARecordIpv6Read,
		DeleteContext: resourceGslbZoneServiceDnsARecordDnsARecordIpv6Delete,

		Schema: map[string]*schema.Schema{
			"admin_ip": {
				Type: schema.TypeInt, Optional: true, Description: "Specify admin priority of Service-IP (Specify the priority)",
			},
			"as_backup": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "As backup when fail",
			},
			"as_replace": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP when enable ip-replace",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this Service-IP",
			},
			"dns_a_record_ipv6": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address",
			},
			"no_resp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't use this Service-IP as DNS response",
			},
			"static": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP in DNS server mode",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL for Service-IP",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Description: "Specify weight for Service-IP (Weight value)",
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbZoneServiceDnsARecordDnsARecordIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsARecordDnsARecordIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsARecordDnsARecordIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsARecordDnsARecordIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsARecordDnsARecordIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsARecordDnsARecordIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6(d *schema.ResourceData) edpt.GslbZoneServiceDnsARecordDnsARecordIpv6 {
	var ret edpt.GslbZoneServiceDnsARecordDnsARecordIpv6
	ret.Inst.AdminIp = d.Get("admin_ip").(int)
	ret.Inst.AsBackup = d.Get("as_backup").(int)
	ret.Inst.AsReplace = d.Get("as_replace").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DnsARecordIpv6 = d.Get("dns_a_record_ipv6").(string)
	ret.Inst.NoResp = d.Get("no_resp").(int)
	ret.Inst.Static = d.Get("static").(int)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
