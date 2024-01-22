package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseManagerOverage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_license_manager_overage`: Set overage parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceLicenseManagerOverageCreate,
		UpdateContext: resourceLicenseManagerOverageUpdate,
		ReadContext:   resourceLicenseManagerOverageRead,
		DeleteContext: resourceLicenseManagerOverageDelete,

		Schema: map[string]*schema.Schema{
			"bytes": {
				Type: schema.TypeInt, Optional: true, Description: "Number of bytes",
			},
			"days": {
				Type: schema.TypeInt, Optional: true, Description: "Number of days",
			},
			"gb": {
				Type: schema.TypeInt, Optional: true, Description: "Number of GB",
			},
			"hours": {
				Type: schema.TypeInt, Optional: true, Description: "Number of hours",
			},
			"kb": {
				Type: schema.TypeInt, Optional: true, Description: "Number of KB",
			},
			"mb": {
				Type: schema.TypeInt, Optional: true, Description: "Number of MB",
			},
			"minutes": {
				Type: schema.TypeInt, Optional: true, Description: "Number of minutes",
			},
			"seconds": {
				Type: schema.TypeInt, Optional: true, Description: "Number of seconds",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLicenseManagerOverageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerOverageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerOverage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerOverageRead(ctx, d, meta)
	}
	return diags
}

func resourceLicenseManagerOverageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerOverageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerOverage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerOverageRead(ctx, d, meta)
	}
	return diags
}
func resourceLicenseManagerOverageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerOverageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerOverage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLicenseManagerOverageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerOverageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerOverage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLicenseManagerOverage(d *schema.ResourceData) edpt.LicenseManagerOverage {
	var ret edpt.LicenseManagerOverage
	ret.Inst.Bytes = d.Get("bytes").(int)
	ret.Inst.Days = d.Get("days").(int)
	ret.Inst.Gb = d.Get("gb").(int)
	ret.Inst.Hours = d.Get("hours").(int)
	ret.Inst.Kb = d.Get("kb").(int)
	ret.Inst.Mb = d.Get("mb").(int)
	ret.Inst.Minutes = d.Get("minutes").(int)
	ret.Inst.Seconds = d.Get("seconds").(int)
	//omit uuid
	return ret
}
