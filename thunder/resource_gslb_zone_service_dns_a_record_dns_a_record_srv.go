package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsARecordDnsARecordSrv() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_a_record_dns_a_record_srv`: Specify DNS Address Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsARecordDnsARecordSrvCreate,
		UpdateContext: resourceGslbZoneServiceDnsARecordDnsARecordSrvUpdate,
		ReadContext:   resourceGslbZoneServiceDnsARecordDnsARecordSrvRead,
		DeleteContext: resourceGslbZoneServiceDnsARecordDnsARecordSrvDelete,

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
			"no_resp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't use this Service-IP as DNS response",
			},
			"static": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return this Service-IP in DNS server mode",
			},
			"svrname": {
				Type: schema.TypeString, Required: true, Description: "Specify name",
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
func resourceGslbZoneServiceDnsARecordDnsARecordSrvCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordSrvCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrv(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsARecordDnsARecordSrvRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsARecordDnsARecordSrvUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordSrvUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrv(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsARecordDnsARecordSrvRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsARecordDnsARecordSrvDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordSrvDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrv(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsARecordDnsARecordSrvRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordSrvRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrv(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrv(d *schema.ResourceData) edpt.GslbZoneServiceDnsARecordDnsARecordSrv {
	var ret edpt.GslbZoneServiceDnsARecordDnsARecordSrv
	ret.Inst.AdminIp = d.Get("admin_ip").(int)
	ret.Inst.AsBackup = d.Get("as_backup").(int)
	ret.Inst.AsReplace = d.Get("as_replace").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.NoResp = d.Get("no_resp").(int)
	ret.Inst.Static = d.Get("static").(int)
	ret.Inst.Svrname = d.Get("svrname").(string)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
