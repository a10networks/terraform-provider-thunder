package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodRadius() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_radius`: RADIUS type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodRadiusCreate,
		UpdateContext: resourceHealthMonitorMethodRadiusUpdate,
		ReadContext:   resourceHealthMonitorMethodRadiusRead,
		DeleteContext: resourceHealthMonitorMethodRadiusDelete,

		Schema: map[string]*schema.Schema{
			"radius": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RADIUS type",
			},
			"radius_expect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify what you expect from the response message",
			},
			"radius_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
			},
			"radius_port": {
				Type: schema.TypeInt, Optional: true, Default: 1812, Description: "Specify the RADIUS port, default is 1812 (Port number (default 1812))",
			},
			"radius_response_code": {
				Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 2,4-7) (Format is xx,xx-xx (xx between [1, 13]))",
			},
			"radius_secret": {
				Type: schema.TypeString, Optional: true, Description: "Configure shared secret of RADIUS server",
			},
			"radius_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the username",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodRadiusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRadiusCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRadius(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodRadiusRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodRadiusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRadiusUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRadius(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodRadiusRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodRadiusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRadiusDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRadius(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodRadiusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRadiusRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRadius(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodRadius(d *schema.ResourceData) edpt.HealthMonitorMethodRadius {
	var ret edpt.HealthMonitorMethodRadius
	ret.Inst.Radius = d.Get("radius").(int)
	//omit radius_encrypted
	ret.Inst.RadiusExpect = d.Get("radius_expect").(int)
	ret.Inst.RadiusPasswordString = d.Get("radius_password_string").(string)
	ret.Inst.RadiusPort = d.Get("radius_port").(int)
	ret.Inst.RadiusResponseCode = d.Get("radius_response_code").(string)
	ret.Inst.RadiusSecret = d.Get("radius_secret").(string)
	//omit radius_secret_encrypted
	ret.Inst.RadiusUsername = d.Get("radius_username").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
