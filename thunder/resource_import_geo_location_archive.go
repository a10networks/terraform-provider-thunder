package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportGeoLocationArchive() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_geo_location_archive`: Geo-location Archive File\n\n__PLACEHOLDER__",
		CreateContext: resourceImportGeoLocationArchiveCreate,
		UpdateContext: resourceImportGeoLocationArchiveUpdate,
		ReadContext:   resourceImportGeoLocationArchiveRead,
		DeleteContext: resourceImportGeoLocationArchiveDelete,

		Schema: map[string]*schema.Schema{
			"geo_location_archive_format": {
				Type: schema.TypeString, Optional: true, Description: "'GeoLite2-ASN': GeoLite2-ASN CSV Zipped File; 'GeoLite2-City': GeoLite2-City CSV Zipped File; 'GeoLite2-Country': GeoLite2-Country CSV Zipped File;",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceImportGeoLocationArchiveCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportGeoLocationArchiveCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportGeoLocationArchive(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportGeoLocationArchiveRead(ctx, d, meta)
	}
	return diags
}

func resourceImportGeoLocationArchiveUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportGeoLocationArchiveUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportGeoLocationArchive(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportGeoLocationArchiveRead(ctx, d, meta)
	}
	return diags
}
func resourceImportGeoLocationArchiveDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportGeoLocationArchiveDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportGeoLocationArchive(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportGeoLocationArchiveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportGeoLocationArchiveRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportGeoLocationArchive(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportGeoLocationArchive(d *schema.ResourceData) edpt.ImportGeoLocationArchive {
	var ret edpt.ImportGeoLocationArchive
	ret.Inst.GeoLocationArchiveFormat = d.Get("geo_location_archive_format").(string)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
