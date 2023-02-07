package thunder

import (
	"context"

	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileGeoLocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_geo_location`: Display geo-location files\n\n__PLACEHOLDER__",
		CreateContext: resourceFileGeoLocationCreate,
		UpdateContext: resourceFileGeoLocationUpdate,
		ReadContext:   resourceFileGeoLocationRead,
		DeleteContext: resourceFileGeoLocationDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
			},
			"dst_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "destination file name for copy and rename action",
			},
			"file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "geo-location database local file name",
			},
			"file_handle": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "full path of the uploaded file",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customized tag",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "uuid of the object",
			},
		},
	}
}
func resourceFileGeoLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceFileGeoLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}
func resourceFileGeoLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocation(d)
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceFileGeoLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}
func dataToEndpointFileGeoLocation(d *schema.ResourceData) edpt.FileGeoLocation {
	var ret edpt.FileGeoLocation
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	return ret
}
