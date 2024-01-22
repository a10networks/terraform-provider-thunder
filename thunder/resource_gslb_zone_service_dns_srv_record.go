package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsSrvRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_srv_record`: Specify DNS SRV Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsSrvRecordCreate,
		UpdateContext: resourceGslbZoneServiceDnsSrvRecordUpdate,
		ReadContext:   resourceGslbZoneServiceDnsSrvRecordRead,
		DeleteContext: resourceGslbZoneServiceDnsSrvRecordDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Specify Port (Port Number)",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Priority",
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
			"srv_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify TTL",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify Weight, default is 10",
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
func resourceGslbZoneServiceDnsSrvRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsSrvRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsSrvRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsSrvRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsSrvRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsSrvRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsSrvRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsSrvRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsSrvRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsSrvRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsSrvRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsSrvRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsSrvRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsSrvRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneServiceDnsSrvRecordSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsSrvRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsSrvRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsSrvRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsSrvRecord(d *schema.ResourceData) edpt.GslbZoneServiceDnsSrvRecord {
	var ret edpt.GslbZoneServiceDnsSrvRecord
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.SamplingEnable = getSliceGslbZoneServiceDnsSrvRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SrvName = d.Get("srv_name").(string)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
