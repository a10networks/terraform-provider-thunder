package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseManagerConnect() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_license_manager_connect`: Connect to license manager to activate\n\n__PLACEHOLDER__",
		CreateContext: resourceLicenseManagerConnectCreate,
		UpdateContext: resourceLicenseManagerConnectUpdate,
		ReadContext:   resourceLicenseManagerConnectRead,
		DeleteContext: resourceLicenseManagerConnectDelete,

		Schema: map[string]*schema.Schema{
			"connect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Connect to license manager to activate",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLicenseManagerConnectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerConnectCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerConnect(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerConnectRead(ctx, d, meta)
	}
	return diags
}

func resourceLicenseManagerConnectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerConnectUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerConnect(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerConnectRead(ctx, d, meta)
	}
	return diags
}
func resourceLicenseManagerConnectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerConnectDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerConnect(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLicenseManagerConnectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerConnectRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerConnect(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLicenseManagerConnect(d *schema.ResourceData) edpt.LicenseManagerConnect {
	var ret edpt.LicenseManagerConnect
	ret.Inst.Connect = d.Get("connect").(int)
	//omit uuid
	return ret
}
