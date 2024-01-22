package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_record`: Specify DNS Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsRecordCreate,
		UpdateContext: resourceGslbZoneServiceDnsRecordUpdate,
		ReadContext:   resourceGslbZoneServiceDnsRecordRead,
		DeleteContext: resourceGslbZoneServiceDnsRecordDelete,

		Schema: map[string]*schema.Schema{
			"data": {
				Type: schema.TypeString, Optional: true, Description: "Specify DNS Data",
			},
			"type": {
				Type: schema.TypeInt, Required: true, Description: "Specify DNS Type",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceGslbZoneServiceDnsRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbZoneServiceDnsRecord(d *schema.ResourceData) edpt.GslbZoneServiceDnsRecord {
	var ret edpt.GslbZoneServiceDnsRecord
	ret.Inst.Data = d.Get("data").(string)
	ret.Inst.Type = d.Get("type").(int)
	//omit uuid
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
