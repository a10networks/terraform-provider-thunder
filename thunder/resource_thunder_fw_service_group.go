package thunder

//Thunder resource FwServiceGroup

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceFwServiceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwServiceGroupCreate,
		UpdateContext: resourceFwServiceGroupUpdate,
		ReadContext:   resourceFwServiceGroupRead,
		DeleteContext: resourceFwServiceGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror_ip_repl": {
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
			"member_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
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

func resourceFwServiceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FwServiceGroup (Inside resourceFwServiceGroupCreate) ")
		name1 := d.Get("name").(string)
		data := dataToFwServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to FwServiceGroup --")
		d.SetId(name1)
		err := go_thunder.PostFwServiceGroup(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwServiceGroupRead(ctx, d, meta)

	}
	return diags
}

func resourceFwServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwServiceGroup (Inside resourceFwServiceGroupRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetFwServiceGroup(client.Token, name1, client.Host)
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

func resourceFwServiceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying FwServiceGroup   (Inside resourceFwServiceGroupUpdate) ")
		data := dataToFwServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to FwServiceGroup ")
		err := go_thunder.PutFwServiceGroup(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwServiceGroupRead(ctx, d, meta)

	}
	return diags
}

func resourceFwServiceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceFwServiceGroupDelete) " + name1)
		err := go_thunder.DeleteFwServiceGroup(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToFwServiceGroup(d *schema.ResourceData) go_thunder.FwServiceGroup {
	var vc go_thunder.FwServiceGroup
	var c go_thunder.FwServiceGroupInstance
	c.FwServiceGroupInstanceName = d.Get("name").(string)
	c.FwServiceGroupInstanceProtocol = d.Get("protocol").(string)
	c.FwServiceGroupInstanceHealthCheck = d.Get("health_check").(string)
	c.FwServiceGroupInstanceTrafficReplicationMirrorIPRepl = d.Get("traffic_replication_mirror_ip_repl").(int)
	c.FwServiceGroupInstanceUserTag = d.Get("user_tag").(string)

	FwServiceGroupInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.FwServiceGroupInstanceSamplingEnableCounters1 = make([]go_thunder.FwServiceGroupInstanceSamplingEnable, 0, FwServiceGroupInstanceSamplingEnableCount)

	for i := 0; i < FwServiceGroupInstanceSamplingEnableCount; i++ {
		var obj1 go_thunder.FwServiceGroupInstanceSamplingEnable
		prefix1 := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.FwServiceGroupInstanceSamplingEnableCounters1 = d.Get(prefix1 + "counters1").(string)
		c.FwServiceGroupInstanceSamplingEnableCounters1 = append(c.FwServiceGroupInstanceSamplingEnableCounters1, obj1)
	}

	c.FwServiceGroupInstancePacketCaptureTemplate = d.Get("packet_capture_template").(string)

	FwServiceGroupInstanceMemberListCount := d.Get("member_list.#").(int)
	c.FwServiceGroupInstanceMemberListPort = make([]go_thunder.FwServiceGroupInstanceMemberList, 0, FwServiceGroupInstanceMemberListCount)

	for i := 0; i < FwServiceGroupInstanceMemberListCount; i++ {
		var obj2 go_thunder.FwServiceGroupInstanceMemberList
		prefix2 := fmt.Sprintf("member_list.%d.", i)
		obj2.FwServiceGroupInstanceMemberListName = d.Get(prefix2 + "name").(string)
		obj2.FwServiceGroupInstanceMemberListPort = d.Get(prefix2 + "port").(int)
		obj2.FwServiceGroupInstanceMemberListUserTag = d.Get(prefix2 + "user_tag").(string)

		FwServiceGroupInstanceMemberListSamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
		obj2.FwServiceGroupInstanceMemberListSamplingEnableCounters1 = make([]go_thunder.FwServiceGroupInstanceMemberListSamplingEnable, 0, FwServiceGroupInstanceMemberListSamplingEnableCount)

		for i := 0; i < FwServiceGroupInstanceMemberListSamplingEnableCount; i++ {
			var obj2_1 go_thunder.FwServiceGroupInstanceMemberListSamplingEnable
			prefix2_1 := prefix2 + fmt.Sprintf("sampling_enable.%d.", i)
			obj2_1.FwServiceGroupInstanceMemberListSamplingEnableCounters1 = d.Get(prefix2_1 + "counters1").(string)
			obj2.FwServiceGroupInstanceMemberListSamplingEnableCounters1 = append(obj2.FwServiceGroupInstanceMemberListSamplingEnableCounters1, obj2_1)
		}

		obj2.FwServiceGroupInstanceMemberListPacketCaptureTemplate = d.Get(prefix2 + "packet_capture_template").(string)
		c.FwServiceGroupInstanceMemberListPort = append(c.FwServiceGroupInstanceMemberListPort, obj2)
	}

	vc.FwServiceGroupInstanceName = c
	return vc
}
