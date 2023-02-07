package thunder

//Thunder resource FwServer

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceFwServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwServerCreate,
		UpdateContext: resourceFwServerUpdate,
		ReadContext:   resourceFwServerRead,
		DeleteContext: resourceFwServerDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_ipv6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fqdn_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"resolve_as": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"packet_capture_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FwServer (Inside resourceFwServerCreate) ")
		name1 := d.Get("name").(string)
		data := dataToFwServer(d)
		logger.Println("[INFO] received formatted data from method data to FwServer --")
		d.SetId(name1)
		err := go_thunder.PostFwServer(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwServerRead(ctx, d, meta)

	}
	return diags
}

func resourceFwServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwServer (Inside resourceFwServerRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetFwServer(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceFwServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying FwServer   (Inside resourceFwServerUpdate) ")
		data := dataToFwServer(d)
		logger.Println("[INFO] received formatted data from method data to FwServer ")
		err := go_thunder.PutFwServer(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwServerRead(ctx, d, meta)

	}
	return diags
}

func resourceFwServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceFwServerDelete) " + name1)
		err := go_thunder.DeleteFwServer(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToFwServer(d *schema.ResourceData) go_thunder.FwServer {
	var vc go_thunder.FwServer
	var c go_thunder.FwServerInstance
	c.FwServerInstanceName = d.Get("name").(string)
	c.FwServerInstanceServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	c.FwServerInstanceHost = d.Get("host").(string)
	c.FwServerInstanceFqdnName = d.Get("fqdn_name").(string)
	c.FwServerInstanceResolveAs = d.Get("resolve_as").(string)
	c.FwServerInstanceAction = d.Get("action").(string)
	c.FwServerInstanceHealthCheck = d.Get("health_check").(string)
	c.FwServerInstanceHealthCheckDisable = d.Get("health_check_disable").(int)
	c.FwServerInstanceUserTag = d.Get("user_tag").(string)

	FwServerInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.FwServerInstanceSamplingEnableCounters1 = make([]go_thunder.FwServerInstanceSamplingEnable, 0, FwServerInstanceSamplingEnableCount)

	for i := 0; i < FwServerInstanceSamplingEnableCount; i++ {
		var obj1 go_thunder.FwServerInstanceSamplingEnable
		prefix1 := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.FwServerInstanceSamplingEnableCounters1 = d.Get(prefix1 + "counters1").(string)
		c.FwServerInstanceSamplingEnableCounters1 = append(c.FwServerInstanceSamplingEnableCounters1, obj1)
	}

	FwServerInstancePortListCount := d.Get("port_list.#").(int)
	c.FwServerInstancePortListPortNumber = make([]go_thunder.FwServerInstancePortList, 0, FwServerInstancePortListCount)

	for i := 0; i < FwServerInstancePortListCount; i++ {
		var obj2 go_thunder.FwServerInstancePortList
		prefix2 := fmt.Sprintf("port_list.%d.", i)
		obj2.FwServerInstancePortListPortNumber = d.Get(prefix2 + "port_number").(int)
		obj2.FwServerInstancePortListProtocol = d.Get(prefix2 + "protocol").(string)
		obj2.FwServerInstancePortListAction = d.Get(prefix2 + "action").(string)
		obj2.FwServerInstancePortListHealthCheck = d.Get(prefix2 + "health_check").(string)
		obj2.FwServerInstancePortListHealthCheckDisable = d.Get(prefix2 + "health_check_disable").(int)
		obj2.FwServerInstancePortListUserTag = d.Get(prefix2 + "user_tag").(string)

		FwServerInstancePortListSamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
		obj2.FwServerInstancePortListSamplingEnableCounters1 = make([]go_thunder.FwServerInstancePortListSamplingEnable, 0, FwServerInstancePortListSamplingEnableCount)

		for i := 0; i < FwServerInstancePortListSamplingEnableCount; i++ {
			var obj2_1 go_thunder.FwServerInstancePortListSamplingEnable
			prefix2_1 := prefix2 + fmt.Sprintf("sampling_enable.%d.", i)
			obj2_1.FwServerInstancePortListSamplingEnableCounters1 = d.Get(prefix2_1 + "counters1").(string)
			obj2.FwServerInstancePortListSamplingEnableCounters1 = append(obj2.FwServerInstancePortListSamplingEnableCounters1, obj2_1)
		}

		obj2.FwServerInstancePortListPacketCaptureTemplate = d.Get(prefix2 + "packet_capture_template").(string)
		c.FwServerInstancePortListPortNumber = append(c.FwServerInstancePortListPortNumber, obj2)
	}

	vc.FwServerInstanceName = c
	return vc
}
