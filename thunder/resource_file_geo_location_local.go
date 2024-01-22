package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileGeoLocationLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_geo_location_local`: Display geo-location files\n\n__PLACEHOLDER__",
		CreateContext: resourceFileGeoLocationLocalCreate,
		UpdateContext: resourceFileGeoLocationLocalUpdate,
		ReadContext:   resourceFileGeoLocationLocalRead,
		DeleteContext: resourceFileGeoLocationLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
			},
			"dst_file": {
				Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "geo-location database local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
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
func resourceFileGeoLocationLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGeoLocationLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileGeoLocationLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGeoLocationLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileGeoLocationLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileGeoLocationLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileGeoLocationLocal(d *schema.ResourceData) edpt.FileGeoLocationLocal {
	var ret edpt.FileGeoLocationLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
