package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemRadiusServerDerivedAttributeUsergroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_radius_server_derived_attribute_usergroup`: Configure the derived attribute of user group.\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemRadiusServerDerivedAttributeUsergroupCreate,
		UpdateContext: resourceSystemRadiusServerDerivedAttributeUsergroupUpdate,
		ReadContext:   resourceSystemRadiusServerDerivedAttributeUsergroupRead,
		DeleteContext: resourceSystemRadiusServerDerivedAttributeUsergroupDelete,

		Schema: map[string]*schema.Schema{
			"attribute": {
				Type: schema.TypeString, Optional: true, Description: "'imei': Specify the IMEI attribute.; 'imsi': Specify the IMSI attribute.; 'msisdn': Specify the MSISDN attribute.; 'custom1': Specify the custom1 attribute.; 'custom2': Specify the custom2 attribute.; 'custom3': Specify the custom3 attribute.; 'custom4': Specify the custom4 attribute.; 'custom5': Specify the custom5 attribute.; 'custom6': Specify the custom6 attribute.;",
			},
			"regex": {
				Type: schema.TypeString, Optional: true, Description: "Specify the regular expression to parse the value from a RADIUS attribute.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemRadiusServerDerivedAttributeUsergroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerDerivedAttributeUsergroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServerDerivedAttributeUsergroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemRadiusServerDerivedAttributeUsergroupRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemRadiusServerDerivedAttributeUsergroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerDerivedAttributeUsergroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServerDerivedAttributeUsergroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemRadiusServerDerivedAttributeUsergroupRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemRadiusServerDerivedAttributeUsergroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerDerivedAttributeUsergroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServerDerivedAttributeUsergroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemRadiusServerDerivedAttributeUsergroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerDerivedAttributeUsergroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServerDerivedAttributeUsergroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemRadiusServerDerivedAttributeUsergroup(d *schema.ResourceData) edpt.SystemRadiusServerDerivedAttributeUsergroup {
	var ret edpt.SystemRadiusServerDerivedAttributeUsergroup
	ret.Inst.Attribute = d.Get("attribute").(string)
	ret.Inst.Regex = d.Get("regex").(string)
	//omit uuid
	return ret
}
