package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRadiusProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_radius_profile`: Configure LSN RADIUS Profile\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnRadiusProfileCreate,
		UpdateContext: resourceCgnv6LsnRadiusProfileUpdate,
		ReadContext:   resourceCgnv6LsnRadiusProfileRead,
		DeleteContext: resourceCgnv6LsnRadiusProfileDelete,

		Schema: map[string]*schema.Schema{
			"lid_profile_index": {
				Type: schema.TypeInt, Required: true, Description: "LSN RADIUS Profile Index",
			},
			"radius": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": {
							Type: schema.TypeString, Optional: true, Description: "'custom1': Configure RADIUS Attribute Custom 1; 'custom2': Configure RADIUS Attribute Custom 2; 'custom3': Configure RADIUS Attribute Custom 3; 'custom4': Configure RADIUS Attribute Custom 4; 'custom5': Configure RADIUS Attribute Custom 5; 'custom6': Configure RADIUS Attribute Custom 6; 'imei': Configure RADIUS Attribute IMEI; 'imsi': Configure RADIUS Attribute IMSI; 'msisdn': Configure RADIUS Attribute MSISDN; 'default': Configure default;",
						},
						"starts_with": {
							Type: schema.TypeString, Optional: true, Description: "Value of the attribute",
						},
						"starts_with_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
						},
						"exact_value": {
							Type: schema.TypeString, Optional: true, Description: "Value of the attribute",
						},
						"exact_value_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
						},
						"default_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnRadiusProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRadiusProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRadiusProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRadiusProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnRadiusProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRadiusProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRadiusProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRadiusProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnRadiusProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRadiusProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRadiusProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnRadiusProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRadiusProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRadiusProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnRadiusProfileRadius(d []interface{}) []edpt.Cgnv6LsnRadiusProfileRadius {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRadiusProfileRadius, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRadiusProfileRadius
		oi.Attribute = in["attribute"].(string)
		oi.StartsWith = in["starts_with"].(string)
		oi.StartsWithLsnLid = in["starts_with_lsn_lid"].(int)
		oi.ExactValue = in["exact_value"].(string)
		oi.ExactValueLsnLid = in["exact_value_lsn_lid"].(int)
		oi.DefaultLsnLid = in["default_lsn_lid"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRadiusProfile(d *schema.ResourceData) edpt.Cgnv6LsnRadiusProfile {
	var ret edpt.Cgnv6LsnRadiusProfile
	ret.Inst.LidProfileIndex = d.Get("lid_profile_index").(int)
	ret.Inst.Radius = getSliceCgnv6LsnRadiusProfileRadius(d.Get("radius").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
