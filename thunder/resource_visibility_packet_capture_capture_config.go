package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureCaptureConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_capture_config`: Packet Capture-Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureCaptureConfigCreate,
		UpdateContext: resourceVisibilityPacketCaptureCaptureConfigUpdate,
		ReadContext:   resourceVisibilityPacketCaptureCaptureConfigRead,
		DeleteContext: resourceVisibilityPacketCaptureCaptureConfigDelete,

		Schema: map[string]*schema.Schema{
			"concurrent_captures": {
				Type: schema.TypeInt, Optional: true, Description: "Enable and specify maximum concurrent 3 tuple filter based captures in seperate pcaps.",
			},
			"concurrent_captures_age": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify the time in minutes upto which a 3 tuple filter based capture will be kept active(default 1)",
			},
			"concurrent_conn_per_capture": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Specify maximum number of concurrent connnections(5 tuple matches) to be captured within in a 3 tuple based capture. (default 1",
			},
			"concurrent_conn_tag": {
				Type: schema.TypeInt, Optional: true, Description: "Enable and specify maximum concurrent connnections(only 5 tuple based) to be captured in common pcaps.",
			},
			"create_pcap_files_now": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Operational command to force create temporary pcapng files before completion (for global/non 3 tuple based captures)",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable packet capture (default enabled)",
			},
			"disable_auto_merge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable auto merging per CPU pcapng files(default enabled)",
			},
			"enable_continuous_global_capture": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable continuous capture of packets for the global capture(non 3 tuple based capture) overriding size limits",
			},
			"file_count": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the number of continuous pcapng files that can be created for capturing packets (default 10)",
			},
			"file_size": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify pcapng filesize in MB, Will be distributed per CPU (default 1)",
			},
			"keep_pcap_files_after_merge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep original per CPU pcapng files after auto merging pcapng files(default disabled)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify the name of the capture-config",
			},
			"number_of_packets_per_capture": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Maximum number of packets per global or dynamic capture (default 0 unlimited)",
			},
			"number_of_packets_per_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify maximum number of packets to be captured in a 5 tuple based connection (default 0 unlimited).",
			},
			"number_of_packets_total": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Maximum number of packets for all captures (default 0 unlimited)",
			},
			"packet_length": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Packet length in Bytes to capture (Default 128)",
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
func resourceVisibilityPacketCaptureCaptureConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureCaptureConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureCaptureConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureCaptureConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureCaptureConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureCaptureConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureCaptureConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureCaptureConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureCaptureConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureCaptureConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureCaptureConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureCaptureConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureCaptureConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureCaptureConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureCaptureConfig(d *schema.ResourceData) edpt.VisibilityPacketCaptureCaptureConfig {
	var ret edpt.VisibilityPacketCaptureCaptureConfig
	ret.Inst.ConcurrentCaptures = d.Get("concurrent_captures").(int)
	ret.Inst.ConcurrentCapturesAge = d.Get("concurrent_captures_age").(int)
	ret.Inst.ConcurrentConnPerCapture = d.Get("concurrent_conn_per_capture").(int)
	ret.Inst.ConcurrentConnTag = d.Get("concurrent_conn_tag").(int)
	ret.Inst.CreatePcapFilesNow = d.Get("create_pcap_files_now").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DisableAutoMerge = d.Get("disable_auto_merge").(int)
	ret.Inst.EnableContinuousGlobalCapture = d.Get("enable_continuous_global_capture").(int)
	ret.Inst.FileCount = d.Get("file_count").(int)
	ret.Inst.FileSize = d.Get("file_size").(int)
	ret.Inst.KeepPcapFilesAfterMerge = d.Get("keep_pcap_files_after_merge").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NumberOfPacketsPerCapture = d.Get("number_of_packets_per_capture").(int)
	ret.Inst.NumberOfPacketsPerConn = d.Get("number_of_packets_per_conn").(int)
	ret.Inst.NumberOfPacketsTotal = d.Get("number_of_packets_total").(int)
	ret.Inst.PacketLength = d.Get("packet_length").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
