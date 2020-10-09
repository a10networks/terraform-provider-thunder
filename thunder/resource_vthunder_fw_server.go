package thunder

//Thunder resource FwServer

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFwServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwServerCreate,
		Update: resourceFwServerUpdate,
		Read:   resourceFwServerRead,
		Delete: resourceFwServerDelete,
		Schema: map[string]*schema.Schema{
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
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
						"port_number": {
							Type:        schema.TypeInt,
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
					},
				},
			},
			"uuid": {
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
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_ipv6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwServerCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwServer (Inside resourceFwServerCreate) ")
		name := d.Get("name").(string)
		data := dataToFwServer(d)
		logger.Println("[INFO] received formatted data from method data to FwServer --")
		d.SetId(name)
		go_thunder.PostFwServer(client.Token, data, client.Host)

		return resourceFwServerRead(d, meta)

	}
	return nil
}

func resourceFwServerRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwServer (Inside resourceFwServerRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwServer(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwServerUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying FwServer   (Inside resourceFwServerUpdate) ")
		name := d.Get("name").(string)
		data := dataToFwServer(d)
		logger.Println("[INFO] received formatted data from method data to FwServer ")
		d.SetId(name)
		go_thunder.PutFwServer(client.Token, name, data, client.Host)

		return resourceFwServerRead(d, meta)

	}
	return nil
}

func resourceFwServerDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceFwServerDelete) " + name)
		err := go_thunder.DeleteFwServer(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToFwServer(d *schema.ResourceData) go_thunder.FwServer {
	var vc go_thunder.FwServer
	var c go_thunder.FwServerInstance
	c.HealthCheckDisable = d.Get("health_check_disable").(int)

	PortListCount := d.Get("port_list.#").(int)
	c.Protocol = make([]go_thunder.FwServerPortList, 0, PortListCount)

	for i := 0; i < PortListCount; i++ {
		var obj1 go_thunder.FwServerPortList
		prefix := fmt.Sprintf("port_list.%d.", i)
		obj1.HealthCheckDisable = d.Get(prefix + "health_check_disable").(int)
		obj1.Protocol = d.Get(prefix + "protocol").(string)
		obj1.UserTag = d.Get(prefix + "user_tag").(string)

		SamplingEnableCount := d.Get(prefix + "sampling_enable.#").(int)
		obj1.Counters1 = make([]go_thunder.FwServerSamplingEnable, 0, SamplingEnableCount)

		for i := 0; i < SamplingEnableCount; i++ {
			var obj1_1 go_thunder.FwServerSamplingEnable
			prefix1 := prefix + fmt.Sprintf("sampling_enable.%d.", i)
			obj1_1.Counters1 = d.Get(prefix1 + "counters1").(string)
			obj1.Counters1 = append(obj1.Counters1, obj1_1)
		}

		obj1.PortNumber = d.Get(prefix + "port_number").(int)
		obj1.Action = d.Get(prefix + "action").(string)
		obj1.HealthCheck = d.Get(prefix + "health_check").(string)
		c.Protocol = append(c.Protocol, obj1)
	}

	c.FqdnName = d.Get("fqdn_name").(string)
	c.ResolveAs = d.Get("resolve_as").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwServerSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwServerSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.UserTag = d.Get("user_tag").(string)
	c.Host = d.Get("host").(string)
	c.Action = d.Get("action").(string)
	c.ServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	c.HealthCheck = d.Get("health_check").(string)
	c.Name = d.Get("name").(string)

	vc.UUID = c
	return vc
}
