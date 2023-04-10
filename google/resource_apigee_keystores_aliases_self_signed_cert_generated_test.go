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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccApigeeKeystoresAliasesSelfSignedCert_apigeeEnvKeystoreAliasSelfSignedCertExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          GetTestOrgFromEnv(t),
		"billing_account": GetTestBillingAccountFromEnv(t),
		"random_suffix":   RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApigeeKeystoresAliasesSelfSignedCertDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeKeystoresAliasesSelfSignedCert_apigeeEnvKeystoreAliasSelfSignedCertExample(context),
			},
			{
				ResourceName:            "google_apigee_keystores_aliases_self_signed_cert.apigee_environment_keystore_ss_alias",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id", "environment", "keystore", "key_size", "sig_alg", "subject", "cert_validity_in_days"},
			},
		},
	})
}

func testAccApigeeKeystoresAliasesSelfSignedCert_apigeeEnvKeystoreAliasSelfSignedCertExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
  depends_on = [google_project_service.apigee]
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [google_project_service.servicenetworking]
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_environment" "apigee_environment_keystore_ss_alias" {
  org_id       = google_apigee_organization.apigee_org.id
  name         = "tf-test%{random_suffix}"
  description  = "Apigee Environment"
  display_name = "environment-1"
}

resource "google_apigee_env_keystore" "apigee_environment_keystore_alias" {
  name       = "tf-test-keystore%{random_suffix}"
  env_id     = google_apigee_environment.apigee_environment_keystore_ss_alias.id
}

resource "google_apigee_keystores_aliases_self_signed_cert" "apigee_environment_keystore_ss_alias" {
  environment			      = google_apigee_environment.apigee_environment_keystore_ss_alias.name
  org_id				        = google_apigee_organization.apigee_org.name
  keystore				      = google_apigee_env_keystore.apigee_environment_keystore_alias.name
  alias                 = "tf test-alias%{random_suffix}"
  key_size              = 1024
  sig_alg               = "SHA512withRSA"
  cert_validity_in_days = 4
  subject {
      common_name = "selfsigned_example"
      country_code = "US"
      locality = "TX"
      org = "CCE"
      org_unit = "PSO"
  }   
}
`, context)
}

func testAccCheckApigeeKeystoresAliasesSelfSignedCertDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_apigee_keystores_aliases_self_signed_cert" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ApigeeBasePath}}organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ApigeeKeystoresAliasesSelfSignedCert still exists at %s", url)
			}
		}

		return nil
	}
}
