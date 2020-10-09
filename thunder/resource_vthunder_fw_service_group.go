package thunder

//Thunder resource FwServiceGroup

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFwServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwServiceGroupCreate,
		Update: resourceFwServiceGroupUpdate,
		Read:   resourceFwServiceGroupRead,
		Delete: resourceFwServiceGroupDelete,
		Schema: map[string]*schema.Schema{
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
			"member_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:        schema.TypeInt,
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
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
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

func resourceFwServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwServiceGroup (Inside resourceFwServiceGroupCreate) ")
		name := d.Get("name").(string)
		data := dataToFwServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to FwServiceGroup --")
		d.SetId(name)
		go_thunder.PostFwServiceGroup(client.Token, data, client.Host)

		return resourceFwServiceGroupRead(d, meta)

	}
	return nil
}

func resourceFwServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwServiceGroup (Inside resourceFwServiceGroupRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwServiceGroup(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying FwServiceGroup   (Inside resourceFwServiceGroupUpdate) ")
		name := d.Get("name").(string)
		data := dataToFwServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to FwServiceGroup ")
		d.SetId(name)
		go_thunder.PutFwServiceGroup(client.Token, name, data, client.Host)

		return resourceFwServiceGroupRead(d, meta)

	}
	return nil
}

func resourceFwServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceFwServiceGroupDelete) " + name)
		err := go_thunder.DeleteFwServiceGroup(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToFwServiceGroup(d *schema.ResourceData) go_thunder.FwServiceGroup {
	var vc go_thunder.FwServiceGroup
	var c go_thunder.FwServiceGroupInstance
	c.Protocol = d.Get("protocol").(string)
	c.UserTag = d.Get("user_tag").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwServiceSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwServiceSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	MemberListCount := d.Get("member_list.#").(int)
	c.Port = make([]go_thunder.FwServiceMemberList, 0, MemberListCount)

	for i := 0; i < MemberListCount; i++ {
		var obj1 go_thunder.FwServiceMemberList
		prefix := fmt.Sprintf("member_list.%d.", i)
		obj1.Port = d.Get(prefix + "port").(int)

		SamplingEnableCount := d.Get(prefix + "sampling_enable.#").(int)
		obj1.Counters1 = make([]go_thunder.FwServiceSamplingEnable, 0, SamplingEnableCount)

		for i := 0; i < SamplingEnableCount; i++ {
			var obj1_1 go_thunder.FwServiceSamplingEnable
			prefix1 := prefix + fmt.Sprintf(prefix+"sampling_enable.%d.", i)
			obj1_1.Counters1 = d.Get(prefix1 + "counters1").(string)
			obj1.Counters1 = append(obj1.Counters1, obj1_1)
		}

		obj1.UserTag = d.Get(prefix + "user_tag").(string)
		obj1.Name = d.Get(prefix + "name").(string)
		c.Port = append(c.Port, obj1)
	}

	c.HealthCheck = d.Get("health_check").(string)
	c.Name = d.Get("name").(string)

	vc.UUID = c
	return vc
}
