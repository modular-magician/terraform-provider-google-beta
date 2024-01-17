// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package tpuv2_test

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccTpuV2Vm_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"runtime_version": envvar.GetTestTpuV2VmRuntimeVersionFromEnv(t),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckTpuV2VmDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTpuV2Vm_full(context),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "zone"},
			},
			{
				Config: testAccTpuV2Vm_update(context, true),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "zone"},
			},
			{
				Config: testAccTpuV2Vm_update(context, false),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "zone"},
			},
		},
	})
}

func testAccTpuV2Vm_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name        = "tf-test-tpu-%{random_suffix}"
  zone        = "us-central1-c"
  description = "Text description of the TPU."

  runtime_version  = "%{runtime_version}"
  accelerator_type = "v2-8"

  scheduling_config {
    preemptible = true
  }

  data_disks {
    source_disk = google_compute_disk.disk.id
    mode        = "READ_ONLY"
  }

  labels = {
    foo = "bar"
  }

  metadata = {
    foo = "bar"
  }

  tags = ["foo"]

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_compute_disk" "disk" {
  provider = google-beta

  name  = "tf-test-tpu-disk-%{random_suffix}"
  image = "debian-cloud/debian-11"
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-c"
}
`, context)
}

func testAccTpuV2Vm_update(context map[string]interface{}, preventDestroy bool) string {
	context["prevent_destroy"] = strconv.FormatBool(preventDestroy)

	return acctest.Nprintf(`
resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name        = "tf-test-tpu-%{random_suffix}"
  zone        = "us-central1-c"
  description = "Text description of the TPU updated."

  runtime_version  = "%{runtime_version}"
  accelerator_type = "v2-8"

  scheduling_config {
    preemptible = true
  }

  data_disks {
    source_disk = google_compute_disk.disk.id
    mode        = "READ_WRITE"
  }

  data_disks {
    source_disk = google_compute_disk.disk2.id
    mode        = "READ_ONLY"
  }

  labels = {
    baz = "bar"
  }

  metadata = {
    baz = "bar"
  }

  tags = ["baz"]

  lifecycle {
    prevent_destroy = %{prevent_destroy}
  }
}

resource "google_compute_disk" "disk" {
  provider = google-beta

  name  = "tf-test-tpu-disk-%{random_suffix}"
  image = "debian-cloud/debian-11"
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-c"
}

resource "google_compute_disk" "disk2" {
  provider = google-beta

  name  = "tf-test-tpu-disk2-%{random_suffix}"
  image = "debian-cloud/debian-11"
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-c"
}
`, context)
}
