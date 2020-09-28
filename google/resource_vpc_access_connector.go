// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVPCAccessConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCAccessConnectorCreate,
		Read:   resourceVPCAccessConnectorRead,
		Delete: resourceVPCAccessConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVPCAccessConnectorImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ip_cidr_range": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the resource (Max 25 characters).`,
			},
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of a VPC network.`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Region where the VPC Access connector resides`,
			},
			"max_throughput": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(200, 1000),
				Description:  `Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 1000.`,
				Default:      1000,
			},
			"min_throughput": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(200, 1000),
				Description:  `Minimum throughput of the connector in Mbps. Default and min is 200.`,
				Default:      200,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully qualified name of this VPC connector`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the VPC access connector.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPCAccessConnectorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandVPCAccessConnectorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandVPCAccessConnectorNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	ipCidrRangeProp, err := expandVPCAccessConnectorIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !isEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	minThroughputProp, err := expandVPCAccessConnectorMinThroughput(d.Get("min_throughput"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_throughput"); !isEmptyValue(reflect.ValueOf(minThroughputProp)) && (ok || !reflect.DeepEqual(v, minThroughputProp)) {
		obj["minThroughput"] = minThroughputProp
	}
	maxThroughputProp, err := expandVPCAccessConnectorMaxThroughput(d.Get("max_throughput"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("max_throughput"); !isEmptyValue(reflect.ValueOf(maxThroughputProp)) && (ok || !reflect.DeepEqual(v, maxThroughputProp)) {
		obj["maxThroughput"] = maxThroughputProp
	}

	obj, err = resourceVPCAccessConnectorEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VPCAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors?connectorId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Connector: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Connector: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = vpcAccessOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Connector", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Connector: %s", err)
	}

	opRes, err = resourceVPCAccessConnectorDecoder(d, meta, opRes)
	if err != nil {
		return fmt.Errorf("Error decoding response from operation: %s", err)
	}
	if opRes == nil {
		return fmt.Errorf("Error decoding response from operation, could not find object")
	}

	if err := d.Set("name", flattenVPCAccessConnectorName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Connector %q: %#v", d.Id(), res)

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)

	return resourceVPCAccessConnectorRead(d, meta)
}

func resourceVPCAccessConnectorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VPCAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("VPCAccessConnector %q", d.Id()))
	}

	res, err = resourceVPCAccessConnectorDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing VPCAccessConnector because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}

	if err := d.Set("name", flattenVPCAccessConnectorName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("network", flattenVPCAccessConnectorNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenVPCAccessConnectorIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("state", flattenVPCAccessConnectorState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("min_throughput", flattenVPCAccessConnectorMinThroughput(res["minThroughput"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("max_throughput", flattenVPCAccessConnectorMaxThroughput(res["maxThroughput"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}

	return nil
}

func resourceVPCAccessConnectorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{VPCAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Connector %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Connector")
	}

	err = vpcAccessOperationWaitTime(
		config, res, project, "Deleting Connector", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Connector %q: %#v", d.Id(), res)
	return nil
}

func resourceVPCAccessConnectorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/connectors/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVPCAccessConnectorName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenVPCAccessConnectorNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVPCAccessConnectorIpCidrRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVPCAccessConnectorState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVPCAccessConnectorMinThroughput(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVPCAccessConnectorMaxThroughput(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandVPCAccessConnectorName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMinThroughput(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMaxThroughput(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceVPCAccessConnectorEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func resourceVPCAccessConnectorDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	if err := d.Set("self_link", res["name"].(string)); err != nil {
		return nil, fmt.Errorf("Error setting self_link: %s", err)
	}
	res["name"] = d.Get("name").(string)
	return res, nil
}
