package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExportGeoLocationArchive() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_export_geo_location_archive`: Geo-location Archive File\n\n__PLACEHOLDER__",
		CreateContext: resourceExportGeoLocationArchiveCreate,
		UpdateContext: resourceExportGeoLocationArchiveUpdate,
		ReadContext:   resourceExportGeoLocationArchiveRead,
		DeleteContext: resourceExportGeoLocationArchiveDelete,

		Schema: map[string]*schema.Schema{
			"geo_location_archive_name": {
				Type: schema.TypeString, Optional: true, Description: "'GeoLite2-ASN-Archive': GeoLite2-ASN CSV Zipped File; 'GeoLite2-City-Archive': GeoLite2-City CSV Zipped File; 'GeoLite2-Country-Archive': GeoLite2-Country CSV Zipped File;",
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
func resourceExportGeoLocationArchiveCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportGeoLocationArchiveCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportGeoLocationArchive(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceExportGeoLocationArchiveRead(ctx, d, meta)
	}
	return diags
}

func resourceExportGeoLocationArchiveUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportGeoLocationArchiveUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportGeoLocationArchive(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceExportGeoLocationArchiveRead(ctx, d, meta)
	}
	return diags
}
func resourceExportGeoLocationArchiveDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportGeoLocationArchiveDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportGeoLocationArchive(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceExportGeoLocationArchiveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportGeoLocationArchiveRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportGeoLocationArchive(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointExportGeoLocationArchive(d *schema.ResourceData) edpt.ExportGeoLocationArchive {
	var ret edpt.ExportGeoLocationArchive
	ret.Inst.GeoLocationArchiveName = d.Get("geo_location_archive_name").(string)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
