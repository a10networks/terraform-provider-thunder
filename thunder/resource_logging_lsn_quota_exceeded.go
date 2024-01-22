package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLsnQuotaExceeded() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_lsn_quota_exceeded`: Set LSN quota exceeded log parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingLsnQuotaExceededCreate,
		UpdateContext: resourceLoggingLsnQuotaExceededUpdate,
		ReadContext:   resourceLoggingLsnQuotaExceededRead,
		DeleteContext: resourceLoggingLsnQuotaExceededDelete,

		Schema: map[string]*schema.Schema{
			"custom1": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Customized attribute No.1",
			},
			"custom2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Customized attribute No.2",
			},
			"custom3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Customized attribute No.3",
			},
			"custom4": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Customized attribute No.4",
			},
			"custom5": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Customized attribute No.5",
			},
			"custom6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Customized attribute No.6",
			},
			"disable_pool_based": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable log LSN user quota exceeded based on LSN pool(Default: enabled)",
			},
			"imei": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "International Mobile Equipment Identity (IMEI)",
			},
			"imsi": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "International Mobile Subscriber Identity (IMSI)",
			},
			"ip_based": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log LSN user quota exceeded based on private IP(Default: disabled)",
			},
			"msisdn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mobile Subscriber Integrated Services Digital Netwrok-Number (MSISDN)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"with_radius_attribute": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log with radius attribute",
			},
		},
	}
}
func resourceLoggingLsnQuotaExceededCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnQuotaExceededCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnQuotaExceeded(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLsnQuotaExceededRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingLsnQuotaExceededUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnQuotaExceededUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnQuotaExceeded(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLsnQuotaExceededRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingLsnQuotaExceededDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnQuotaExceededDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnQuotaExceeded(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingLsnQuotaExceededRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnQuotaExceededRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnQuotaExceeded(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingLsnQuotaExceeded(d *schema.ResourceData) edpt.LoggingLsnQuotaExceeded {
	var ret edpt.LoggingLsnQuotaExceeded
	ret.Inst.Custom1 = d.Get("custom1").(int)
	ret.Inst.Custom2 = d.Get("custom2").(int)
	ret.Inst.Custom3 = d.Get("custom3").(int)
	ret.Inst.Custom4 = d.Get("custom4").(int)
	ret.Inst.Custom5 = d.Get("custom5").(int)
	ret.Inst.Custom6 = d.Get("custom6").(int)
	ret.Inst.DisablePoolBased = d.Get("disable_pool_based").(int)
	ret.Inst.Imei = d.Get("imei").(int)
	ret.Inst.Imsi = d.Get("imsi").(int)
	ret.Inst.IpBased = d.Get("ip_based").(int)
	ret.Inst.Msisdn = d.Get("msisdn").(int)
	//omit uuid
	ret.Inst.WithRadiusAttribute = d.Get("with_radius_attribute").(int)
	return ret
}
