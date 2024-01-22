package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseManagerReminder() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_license_manager_reminder`: Set the reminder for grace time\n\n__PLACEHOLDER__",
		CreateContext: resourceLicenseManagerReminderCreate,
		UpdateContext: resourceLicenseManagerReminderUpdate,
		ReadContext:   resourceLicenseManagerReminderRead,
		DeleteContext: resourceLicenseManagerReminderDelete,

		Schema: map[string]*schema.Schema{
			"reminder_value": {
				Type: schema.TypeInt, Required: true, Description: "Configure reminder for grace time (Hour)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLicenseManagerReminderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerReminderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerReminder(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerReminderRead(ctx, d, meta)
	}
	return diags
}

func resourceLicenseManagerReminderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerReminderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerReminder(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerReminderRead(ctx, d, meta)
	}
	return diags
}
func resourceLicenseManagerReminderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerReminderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerReminder(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLicenseManagerReminderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerReminderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerReminder(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLicenseManagerReminder(d *schema.ResourceData) edpt.LicenseManagerReminder {
	var ret edpt.LicenseManagerReminder
	ret.Inst.ReminderValue = d.Get("reminder_value").(int)
	//omit uuid
	return ret
}
