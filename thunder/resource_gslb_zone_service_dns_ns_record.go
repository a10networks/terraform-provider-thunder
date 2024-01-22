package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsNsRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_ns_record`: Specify DNS NS Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsNsRecordCreate,
		UpdateContext: resourceGslbZoneServiceDnsNsRecordUpdate,
		ReadContext:   resourceGslbZoneServiceDnsNsRecordRead,
		DeleteContext: resourceGslbZoneServiceDnsNsRecordDelete,

		Schema: map[string]*schema.Schema{
			"ns_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the record has been used;",
						},
					},
				},
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
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
func resourceGslbZoneServiceDnsNsRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNsRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNsRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsNsRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsNsRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNsRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNsRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsNsRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsNsRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNsRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNsRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsNsRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNsRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNsRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneServiceDnsNsRecordSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsNsRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsNsRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsNsRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsNsRecord(d *schema.ResourceData) edpt.GslbZoneServiceDnsNsRecord {
	var ret edpt.GslbZoneServiceDnsNsRecord
	ret.Inst.NsName = d.Get("ns_name").(string)
	ret.Inst.SamplingEnable = getSliceGslbZoneServiceDnsNsRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
