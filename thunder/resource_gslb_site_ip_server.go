package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteIpServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site_ip_server`: Specify a real server for the GSLB site\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteIpServerCreate,
		UpdateContext: resourceGslbSiteIpServerUpdate,
		ReadContext:   resourceGslbSiteIpServerRead,
		DeleteContext: resourceGslbSiteIpServerDelete,

		Schema: map[string]*schema.Schema{
			"ip_server_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the real server name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
		},
	}
}
func resourceGslbSiteIpServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteIpServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteIpServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteIpServerRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteIpServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteIpServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteIpServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteIpServerRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteIpServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteIpServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteIpServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteIpServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteIpServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteIpServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbSiteIpServer(d *schema.ResourceData) edpt.GslbSiteIpServer {
	var ret edpt.GslbSiteIpServer
	ret.Inst.IpServerName = d.Get("ip_server_name").(string)
	//omit uuid
	ret.Inst.SiteName = d.Get("site_name").(string)
	return ret
}
