package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHostname() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_hostname`: Set system's network name\n\n__PLACEHOLDER__",
		CreateContext: resourceHostnameCreate,
		UpdateContext: resourceHostnameUpdate,
		ReadContext:   resourceHostnameRead,
		DeleteContext: resourceHostnameDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Optional: true, Description: "This System's Network Name",
			},
		},
	}
}
func resourceHostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHostnameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHostname(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHostnameRead(ctx, d, meta)
	}
	return diags
}

func resourceHostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHostnameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHostname(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHostnameRead(ctx, d, meta)
	}
	return diags
}
func resourceHostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHostnameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHostname(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHostnameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHostname(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHostname(d *schema.ResourceData) edpt.Hostname {
	var ret edpt.Hostname
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	return ret
}
