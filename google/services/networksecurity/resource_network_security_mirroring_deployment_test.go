// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/networksecurity/resource_network_security_mirroring_deployment_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package networksecurity_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccNetworkSecurityMirroringDeployment_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityMirroringDeployment_basic(context),
			},
			{
				ResourceName:            "google_network_security_mirroring_deployment.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
			{
				Config: testAccNetworkSecurityMirroringDeployment_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_network_security_mirroring_deployment.default", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_network_security_mirroring_deployment.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"update_time", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityMirroringDeployment_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  name                    = "tf-test-example-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork" {
  name          = "tf-test-example-subnet%{random_suffix}"
  region        = "us-central1"
  ip_cidr_range = "10.1.0.0/16"
  network       = google_compute_network.network.name
}

resource "google_compute_region_health_check" "health_check" {
  name     = "tf-test-example-hc%{random_suffix}"
  region   = "us-central1"
  http_health_check {
    port = 80
  }
}

resource "google_compute_region_backend_service" "backend_service" {
  name                  = "tf-test-example-bs%{random_suffix}"
  region                = "us-central1"
  health_checks         = [google_compute_region_health_check.health_check.id]
  protocol              = "UDP"
  load_balancing_scheme = "INTERNAL"
}

resource "google_compute_forwarding_rule" "forwarding_rule" {
  name                   = "tf-test-example-fwr%{random_suffix}"
  region                 = "us-central1"
  network                = google_compute_network.network.name
  subnetwork             = google_compute_subnetwork.subnetwork.name
  backend_service        = google_compute_region_backend_service.backend_service.id
  load_balancing_scheme  = "INTERNAL"
  ports                  = [6081]
  ip_protocol            = "UDP"
  is_mirroring_collector = true
}

resource "google_network_security_mirroring_deployment_group" "deployment_group" {
  mirroring_deployment_group_id = "tf-test-example-dg%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.network.id
}

resource "google_network_security_mirroring_deployment" "default" {
  mirroring_deployment_id    = "tf-test-example-deployment%{random_suffix}"
  location                   = "us-central1-a"
  forwarding_rule            = google_compute_forwarding_rule.forwarding_rule.id
  mirroring_deployment_group = google_network_security_mirroring_deployment_group.deployment_group.id
  description                = "initial description"
  labels = {
    foo = "bar"
  }
}
`, context)
}

func testAccNetworkSecurityMirroringDeployment_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  name                    = "tf-test-example-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork" {
  name          = "tf-test-example-subnet%{random_suffix}"
  region        = "us-central1"
  ip_cidr_range = "10.1.0.0/16"
  network       = google_compute_network.network.name
}

resource "google_compute_region_health_check" "health_check" {
  name     = "tf-test-example-hc%{random_suffix}"
  region   = "us-central1"
  http_health_check {
    port = 80
  }
}

resource "google_compute_region_backend_service" "backend_service" {
  name                  = "tf-test-example-bs%{random_suffix}"
  region                = "us-central1"
  health_checks         = [google_compute_region_health_check.health_check.id]
  protocol              = "UDP"
  load_balancing_scheme = "INTERNAL"
}

resource "google_compute_forwarding_rule" "forwarding_rule" {
  name                   = "tf-test-example-fwr%{random_suffix}"
  region                 = "us-central1"
  network                = google_compute_network.network.name
  subnetwork             = google_compute_subnetwork.subnetwork.name
  backend_service        = google_compute_region_backend_service.backend_service.id
  load_balancing_scheme  = "INTERNAL"
  ports                  = [6081]
  ip_protocol            = "UDP"
  is_mirroring_collector = true
}

resource "google_network_security_mirroring_deployment_group" "deployment_group" {
  mirroring_deployment_group_id = "tf-test-example-dg%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.network.id
}

resource "google_network_security_mirroring_deployment" "default" {
  mirroring_deployment_id    = "tf-test-example-deployment%{random_suffix}"
  location                   = "us-central1-a"
  forwarding_rule            = google_compute_forwarding_rule.forwarding_rule.id
  mirroring_deployment_group = google_network_security_mirroring_deployment_group.deployment_group.id
  description                = "updated description"
  labels = {
    foo = "goo"
  }
}
`, context)
}
