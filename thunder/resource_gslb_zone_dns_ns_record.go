package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsNsRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_dns_ns_record`: Specify DNS NS Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneDnsNsRecordCreate,
		UpdateContext: resourceGslbZoneDnsNsRecordUpdate,
		ReadContext:   resourceGslbZoneDnsNsRecordRead,
		DeleteContext: resourceGslbZoneDnsNsRecordDelete,

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
				Type: schema.TypeInt, Optional: true, Description: "Specify TTL",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbZoneDnsNsRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsNsRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsNsRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneDnsNsRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneDnsNsRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsNsRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsNsRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneDnsNsRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneDnsNsRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsNsRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsNsRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneDnsNsRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsNsRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsNsRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneDnsNsRecordSamplingEnable(d []interface{}) []edpt.GslbZoneDnsNsRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneDnsNsRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneDnsNsRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneDnsNsRecord(d *schema.ResourceData) edpt.GslbZoneDnsNsRecord {
	var ret edpt.GslbZoneDnsNsRecord
	ret.Inst.NsName = d.Get("ns_name").(string)
	ret.Inst.SamplingEnable = getSliceGslbZoneDnsNsRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
