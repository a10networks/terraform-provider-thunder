package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsNaptrRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_naptr_record`: Specify DNS NAPTR Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsNaptrRecordCreate,
		UpdateContext: resourceGslbZoneServiceDnsNaptrRecordUpdate,
		ReadContext:   resourceGslbZoneServiceDnsNaptrRecordRead,
		DeleteContext: resourceGslbZoneServiceDnsNaptrRecordDelete,

		Schema: map[string]*schema.Schema{
			"flag": {
				Type: schema.TypeString, Required: true, Description: "Specify the flag (e.g., a, s). Default is empty flag",
			},
			"naptr_target": {
				Type: schema.TypeString, Required: true, Description: "Specify the replacement or regular expression",
			},
			"order": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Order",
			},
			"preference": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Preference",
			},
			"regexp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Return the regular expression",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'naptr-hits': Number of times the NAPTR has been used;",
						},
					},
				},
			},
			"service_proto": {
				Type: schema.TypeString, Required: true, Description: "Specify Service and Protocol",
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
func resourceGslbZoneServiceDnsNaptrRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNaptrRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNaptrRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsNaptrRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsNaptrRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNaptrRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNaptrRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsNaptrRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsNaptrRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNaptrRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNaptrRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsNaptrRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNaptrRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNaptrRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneServiceDnsNaptrRecordSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsNaptrRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsNaptrRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsNaptrRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsNaptrRecord(d *schema.ResourceData) edpt.GslbZoneServiceDnsNaptrRecord {
	var ret edpt.GslbZoneServiceDnsNaptrRecord
	ret.Inst.Flag = d.Get("flag").(string)
	ret.Inst.NaptrTarget = d.Get("naptr_target").(string)
	ret.Inst.Order = d.Get("order").(int)
	ret.Inst.Preference = d.Get("preference").(int)
	ret.Inst.Regexp = d.Get("regexp").(int)
	ret.Inst.SamplingEnable = getSliceGslbZoneServiceDnsNaptrRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceProto = d.Get("service_proto").(string)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
