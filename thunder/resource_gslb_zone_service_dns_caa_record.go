package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsCaaRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_caa_record`: Specify DNS CAA Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsCaaRecordCreate,
		UpdateContext: resourceGslbZoneServiceDnsCaaRecordUpdate,
		ReadContext:   resourceGslbZoneServiceDnsCaaRecordRead,
		DeleteContext: resourceGslbZoneServiceDnsCaaRecordDelete,

		Schema: map[string]*schema.Schema{
			"critical_flag": {
				Type: schema.TypeInt, Required: true, Description: "Issuer Critical Flag",
			},
			"property_tag": {
				Type: schema.TypeString, Required: true, Description: "Specify other property tags, only allowed lowercase alphanumeric",
			},
			"rdata": {
				Type: schema.TypeString, Required: true, Description: "Specify the Issuer Domain Name or a URL",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the CAA has been used;",
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
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbZoneServiceDnsCaaRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCaaRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCaaRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsCaaRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsCaaRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCaaRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCaaRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsCaaRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsCaaRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCaaRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCaaRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsCaaRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCaaRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCaaRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneServiceDnsCaaRecordSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsCaaRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsCaaRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsCaaRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsCaaRecord(d *schema.ResourceData) edpt.GslbZoneServiceDnsCaaRecord {
	var ret edpt.GslbZoneServiceDnsCaaRecord
	ret.Inst.CriticalFlag = d.Get("critical_flag").(int)
	ret.Inst.PropertyTag = d.Get("property_tag").(string)
	ret.Inst.Rdata = d.Get("rdata").(string)
	ret.Inst.SamplingEnable = getSliceGslbZoneServiceDnsCaaRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
