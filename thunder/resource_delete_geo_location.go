package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteGeoLocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_geo_location`: Delete geo-location file\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteGeoLocationCreate,
		UpdateContext: resourceDeleteGeoLocationUpdate,
		ReadContext:   resourceDeleteGeoLocationRead,
		DeleteContext: resourceDeleteGeoLocationDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete all files",
			},
			"filename": {
				Type: schema.TypeString, Optional: true, Description: "Specify file to be deleted",
			},
		},
	}
}
func resourceDeleteGeoLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGeoLocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGeoLocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteGeoLocationRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteGeoLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGeoLocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGeoLocation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteGeoLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGeoLocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGeoLocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteGeoLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGeoLocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGeoLocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteGeoLocation(d *schema.ResourceData) edpt.DeleteGeoLocation {
	var ret edpt.DeleteGeoLocation
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Filename = d.Get("filename").(string)
	return ret
}
