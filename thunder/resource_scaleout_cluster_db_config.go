package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterDbConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_db_config`: Configure zk prarameters\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterDbConfigCreate,
		UpdateContext: resourceScaleoutClusterDbConfigUpdate,
		ReadContext:   resourceScaleoutClusterDbConfigRead,
		DeleteContext: resourceScaleoutClusterDbConfigDelete,

		Schema: map[string]*schema.Schema{
			"broken_detect_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 12000, Description: "database connection broken detection timeout (mseconds) (12000 mseconds for default)",
			},
			"client_recv_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 13000, Description: "",
			},
			"clientport": {
				Type: schema.TypeInt, Optional: true, Description: "client session port",
			},
			"elect_conn_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1200, Description: "election connection timeout (mseconds) (1200 for default)",
			},
			"initlimit": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"loopback_intf_support": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "support loopback interface for scaleout database (enabled by default)",
			},
			"maxsessiontimeout": {
				Type: schema.TypeInt, Optional: true, Default: 30000, Description: "",
			},
			"minsessiontimeout": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "",
			},
			"more_election_packet": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "send more election packet in election period (enabled by default)",
			},
			"synclimit": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"ticktime": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterDbConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDbConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDbConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterDbConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterDbConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDbConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDbConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterDbConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterDbConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDbConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDbConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterDbConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDbConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDbConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterDbConfig(d *schema.ResourceData) edpt.ScaleoutClusterDbConfig {
	var ret edpt.ScaleoutClusterDbConfig
	ret.Inst.BrokenDetectTimeout = d.Get("broken_detect_timeout").(int)
	ret.Inst.ClientRecvTimeout = d.Get("client_recv_timeout").(int)
	ret.Inst.Clientport = d.Get("clientport").(int)
	ret.Inst.ElectConnTimeout = d.Get("elect_conn_timeout").(int)
	ret.Inst.Initlimit = d.Get("initlimit").(int)
	ret.Inst.LoopbackIntfSupport = d.Get("loopback_intf_support").(int)
	ret.Inst.Maxsessiontimeout = d.Get("maxsessiontimeout").(int)
	ret.Inst.Minsessiontimeout = d.Get("minsessiontimeout").(int)
	ret.Inst.MoreElectionPacket = d.Get("more_election_packet").(int)
	ret.Inst.Synclimit = d.Get("synclimit").(int)
	ret.Inst.Ticktime = d.Get("ticktime").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
