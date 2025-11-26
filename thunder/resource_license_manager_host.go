package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseManagerHost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_license_manager_host`: Configure license manager host\n\n__PLACEHOLDER__",
		CreateContext: resourceLicenseManagerHostCreate,
		UpdateContext: resourceLicenseManagerHostUpdate,
		ReadContext:   resourceLicenseManagerHostRead,
		DeleteContext: resourceLicenseManagerHostDelete,

		Schema: map[string]*schema.Schema{
			"host_ipv4": {
				Type: schema.TypeString, Required: true, Description: "license server ip address (length:1-31)",
			},
			"host_ipv6": {
				Type: schema.TypeString, Required: true, Description: "Configure license manager server ipv6-address",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Configure the license manager port, default is 443",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLicenseManagerHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerHostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerHost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerHostRead(ctx, d, meta)
	}
	return diags
}

func resourceLicenseManagerHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerHostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerHost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerHostRead(ctx, d, meta)
	}
	return diags
}
func resourceLicenseManagerHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerHostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerHost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLicenseManagerHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerHostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerHost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLicenseManagerHost(d *schema.ResourceData) edpt.LicenseManagerHost {
	var ret edpt.LicenseManagerHost
	ret.Inst.HostIpv4 = d.Get("host_ipv4").(string)
	ret.Inst.HostIpv6 = d.Get("host_ipv6").(string)
	ret.Inst.Port = d.Get("port").(int)
	//omit uuid
	return ret
}
