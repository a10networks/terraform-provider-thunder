package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureDeletePacketCaptureFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_delete_packet_capture_file`: Delete the Packet capture pcapng file\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureDeletePacketCaptureFileCreate,
		UpdateContext: resourceVisibilityPacketCaptureDeletePacketCaptureFileUpdate,
		ReadContext:   resourceVisibilityPacketCaptureDeletePacketCaptureFileRead,
		DeleteContext: resourceVisibilityPacketCaptureDeletePacketCaptureFileDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete all files in this partition",
			},
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the name of the packet capture file to be deleted",
			},
		},
	}
}
func resourceVisibilityPacketCaptureDeletePacketCaptureFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureDeletePacketCaptureFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureDeletePacketCaptureFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureDeletePacketCaptureFileRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureDeletePacketCaptureFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureDeletePacketCaptureFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureDeletePacketCaptureFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureDeletePacketCaptureFileRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureDeletePacketCaptureFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureDeletePacketCaptureFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureDeletePacketCaptureFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureDeletePacketCaptureFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureDeletePacketCaptureFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureDeletePacketCaptureFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureDeletePacketCaptureFile(d *schema.ResourceData) edpt.VisibilityPacketCaptureDeletePacketCaptureFile {
	var ret edpt.VisibilityPacketCaptureDeletePacketCaptureFile
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
