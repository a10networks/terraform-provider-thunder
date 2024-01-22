package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsMxRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_dns_mx_record`: Specify DNS MX Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneDnsMxRecordCreate,
		UpdateContext: resourceGslbZoneDnsMxRecordUpdate,
		ReadContext:   resourceGslbZoneDnsMxRecordRead,
		DeleteContext: resourceGslbZoneDnsMxRecordDelete,

		Schema: map[string]*schema.Schema{
			"mx_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
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
func resourceGslbZoneDnsMxRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsMxRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsMxRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneDnsMxRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneDnsMxRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsMxRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsMxRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneDnsMxRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneDnsMxRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsMxRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsMxRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneDnsMxRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsMxRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsMxRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneDnsMxRecordSamplingEnable(d []interface{}) []edpt.GslbZoneDnsMxRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneDnsMxRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneDnsMxRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneDnsMxRecord(d *schema.ResourceData) edpt.GslbZoneDnsMxRecord {
	var ret edpt.GslbZoneDnsMxRecord
	ret.Inst.MxName = d.Get("mx_name").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.SamplingEnable = getSliceGslbZoneDnsMxRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
