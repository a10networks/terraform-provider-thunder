package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsCnameRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_dns_cname_record`: Specify DNS CNAME Record\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceDnsCnameRecordCreate,
		UpdateContext: resourceGslbZoneServiceDnsCnameRecordUpdate,
		ReadContext:   resourceGslbZoneServiceDnsCnameRecordRead,
		DeleteContext: resourceGslbZoneServiceDnsCnameRecordDelete,

		Schema: map[string]*schema.Schema{
			"admin_preference": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Specify Administrative Preference, default is 100",
			},
			"alias_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the alias name",
			},
			"as_backup": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "As backup when fail",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'cname-hits': Number of times the CNAME has been used;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify Weight, default is 1",
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
func resourceGslbZoneServiceDnsCnameRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCnameRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCnameRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsCnameRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceDnsCnameRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCnameRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCnameRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceDnsCnameRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceDnsCnameRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCnameRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCnameRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceDnsCnameRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCnameRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCnameRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneServiceDnsCnameRecordSamplingEnable(d []interface{}) []edpt.GslbZoneServiceDnsCnameRecordSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceDnsCnameRecordSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceDnsCnameRecordSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsCnameRecord(d *schema.ResourceData) edpt.GslbZoneServiceDnsCnameRecord {
	var ret edpt.GslbZoneServiceDnsCnameRecord
	ret.Inst.AdminPreference = d.Get("admin_preference").(int)
	ret.Inst.AliasName = d.Get("alias_name").(string)
	ret.Inst.AsBackup = d.Get("as_backup").(int)
	ret.Inst.SamplingEnable = getSliceGslbZoneServiceDnsCnameRecordSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
