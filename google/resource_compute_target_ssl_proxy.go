// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceComputeTargetSslProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetSslProxyCreate,
		Read:   resourceComputeTargetSslProxyRead,
		Update: resourceComputeTargetSslProxyUpdate,
		Delete: resourceComputeTargetSslProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetSslProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backend_service": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to the BackendService resource.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"certificate_map": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A reference to the CertificateMap resource uri that identifies a certificate map
associated with the given target proxy. This field can only be set for global target proxies.
Accepted format is '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'.`,
				ExactlyOneOf: []string{"ssl_certificates", "certificate_map"},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"proxy_header": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"NONE", "PROXY_V1", ""}),
				Description: `Specifies the type of proxy header to append before sending data to
the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]`,
				Default: "NONE",
			},
			"ssl_certificates": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of SslCertificate resources that are used to authenticate
connections between users and the load balancer. At least one
SSL certificate must be specified.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
				ExactlyOneOf: []string{"ssl_certificates", "certificate_map"},
			},
			"ssl_policy": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `A reference to the SslPolicy resource that will be associated with
the TargetSslProxy resource. If not set, the TargetSslProxy
resource will not have any SSL policy configured.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"proxy_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeTargetSslProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeTargetSslProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeTargetSslProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	proxyHeaderProp, err := expandComputeTargetSslProxyProxyHeader(d.Get("proxy_header"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("proxy_header"); !isEmptyValue(reflect.ValueOf(proxyHeaderProp)) && (ok || !reflect.DeepEqual(v, proxyHeaderProp)) {
		obj["proxyHeader"] = proxyHeaderProp
	}
	serviceProp, err := expandComputeTargetSslProxyBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(serviceProp)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["service"] = serviceProp
	}
	sslCertificatesProp, err := expandComputeTargetSslProxySslCertificates(d.Get("ssl_certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_certificates"); !isEmptyValue(reflect.ValueOf(sslCertificatesProp)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
		obj["sslCertificates"] = sslCertificatesProp
	}
	certificateMapProp, err := expandComputeTargetSslProxyCertificateMap(d.Get("certificate_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificate_map"); !isEmptyValue(reflect.ValueOf(certificateMapProp)) && (ok || !reflect.DeepEqual(v, certificateMapProp)) {
		obj["certificateMap"] = certificateMapProp
	}
	sslPolicyProp, err := expandComputeTargetSslProxySslPolicy(d.Get("ssl_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_policy"); !isEmptyValue(reflect.ValueOf(sslPolicyProp)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
		obj["sslPolicy"] = sslPolicyProp
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetSslProxy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetSslProxy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TargetSslProxy: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/targetSslProxies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating TargetSslProxy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TargetSslProxy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TargetSslProxy %q: %#v", d.Id(), res)

	return resourceComputeTargetSslProxyRead(d, meta)
}

func resourceComputeTargetSslProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetSslProxy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeTargetSslProxy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeTargetSslProxyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeTargetSslProxyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("proxy_id", flattenComputeTargetSslProxyProxyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeTargetSslProxyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("proxy_header", flattenComputeTargetSslProxyProxyHeader(res["proxyHeader"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("backend_service", flattenComputeTargetSslProxyBackendService(res["service"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("ssl_certificates", flattenComputeTargetSslProxySslCertificates(res["sslCertificates"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("certificate_map", flattenComputeTargetSslProxyCertificateMap(res["certificateMap"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("ssl_policy", flattenComputeTargetSslProxySslPolicy(res["sslPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading TargetSslProxy: %s", err)
	}

	return nil
}

func resourceComputeTargetSslProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetSslProxy: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("proxy_header") {
		obj := make(map[string]interface{})

		proxyHeaderProp, err := expandComputeTargetSslProxyProxyHeader(d.Get("proxy_header"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("proxy_header"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, proxyHeaderProp)) {
			obj["proxyHeader"] = proxyHeaderProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}/setProxyHeader")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetSslProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetSslProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("backend_service") {
		obj := make(map[string]interface{})

		serviceProp, err := expandComputeTargetSslProxyBackendService(d.Get("backend_service"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
			obj["service"] = serviceProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}/setBackendService")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetSslProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetSslProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("ssl_certificates") {
		obj := make(map[string]interface{})

		sslCertificatesProp, err := expandComputeTargetSslProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_certificates"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
			obj["sslCertificates"] = sslCertificatesProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}/setSslCertificates")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetSslProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetSslProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("certificate_map") {
		obj := make(map[string]interface{})

		certificateMapProp, err := expandComputeTargetSslProxyCertificateMap(d.Get("certificate_map"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("certificate_map"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, certificateMapProp)) {
			obj["certificateMap"] = certificateMapProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}/setCertificateMap")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetSslProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetSslProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("ssl_policy") {
		obj := make(map[string]interface{})

		sslPolicyProp, err := expandComputeTargetSslProxySslPolicy(d.Get("ssl_policy"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
			obj["sslPolicy"] = sslPolicyProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}/setSslPolicy")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetSslProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetSslProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeTargetSslProxyRead(d, meta)
}

func resourceComputeTargetSslProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetSslProxy: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetSslProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TargetSslProxy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TargetSslProxy")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting TargetSslProxy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TargetSslProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeTargetSslProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/global/targetSslProxies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/targetSslProxies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeTargetSslProxyCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetSslProxyDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetSslProxyProxyId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenComputeTargetSslProxyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetSslProxyProxyHeader(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetSslProxyBackendService(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeTargetSslProxySslCertificates(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeTargetSslProxyCertificateMap(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetSslProxySslPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeTargetSslProxyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxyProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxyBackendService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("backendServices", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeTargetSslProxySslCertificates(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: nil")
		}
		f, err := parseGlobalFieldValue("sslCertificates", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeTargetSslProxyCertificateMap(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxySslPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("sslPolicies", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for ssl_policy: %s", err)
	}
	return f.RelativeLink(), nil
}
