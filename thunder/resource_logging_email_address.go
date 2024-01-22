package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingEmailAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_email_address`: Set the recipients address\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingEmailAddressCreate,
		UpdateContext: resourceLoggingEmailAddressUpdate,
		ReadContext:   resourceLoggingEmailAddressRead,
		DeleteContext: resourceLoggingEmailAddressDelete,

		Schema: map[string]*schema.Schema{
			"email_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_address": {
							Type: schema.TypeString, Optional: true, Description: "Email address information of recipients",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingEmailAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingEmailAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingEmailAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingEmailAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceLoggingEmailAddressEmailList(d []interface{}) []edpt.LoggingEmailAddressEmailList {

	count1 := len(d)
	ret := make([]edpt.LoggingEmailAddressEmailList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingEmailAddressEmailList
		oi.EmailAddress = in["email_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingEmailAddress(d *schema.ResourceData) edpt.LoggingEmailAddress {
	var ret edpt.LoggingEmailAddress
	ret.Inst.EmailList = getSliceLoggingEmailAddressEmailList(d.Get("email_list").([]interface{}))
	//omit uuid
	return ret
}
