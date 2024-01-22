package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlMetadataMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_saml_metadata_monitor`: Configure SAML metadata out-of-sync detection and recovery options\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationSamlMetadataMonitorCreate,
		UpdateContext: resourceAamAuthenticationSamlMetadataMonitorUpdate,
		ReadContext:   resourceAamAuthenticationSamlMetadataMonitorRead,
		DeleteContext: resourceAamAuthenticationSamlMetadataMonitorDelete,

		Schema: map[string]*schema.Schema{
			"acs_continuous_fail_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Specify how many ACS continuous fails will trigger metadata reload (ACS continuous fail threshold (default: 10))",
			},
			"acs_missing_period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify how long no acs request will trigger metadata reload (in seconds (default: 60))",
			},
			"acs_missing_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Specify how many ACS request missing in the period will trigger metadata reload (ACS request missing threshold (default: 100))",
			},
			"status": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable SAML metadata out-of-sync detection; 'disable': Disable SAML metadata out-of-sync detection;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationSamlMetadataMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlMetadataMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlMetadataMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlMetadataMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationSamlMetadataMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlMetadataMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlMetadataMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlMetadataMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationSamlMetadataMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlMetadataMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlMetadataMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationSamlMetadataMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlMetadataMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlMetadataMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthenticationSamlMetadataMonitor(d *schema.ResourceData) edpt.AamAuthenticationSamlMetadataMonitor {
	var ret edpt.AamAuthenticationSamlMetadataMonitor
	ret.Inst.AcsContinuousFailThreshold = d.Get("acs_continuous_fail_threshold").(int)
	ret.Inst.AcsMissingPeriod = d.Get("acs_missing_period").(int)
	ret.Inst.AcsMissingThreshold = d.Get("acs_missing_threshold").(int)
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
