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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDocumentAIProcessorDefaultVersion_documentaiDefaultVersionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIProcessorDefaultVersion_documentaiDefaultVersionExample(context),
			},
			{
				ResourceName:            "google_document_ai_processor_default_version.processor",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"processor"},
			},
		},
	})
}

func testAccDocumentAIProcessorDefaultVersion_documentaiDefaultVersionExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_processor" "processor" {
  location = "us"
  display_name = "tf-test-test-processor%{random_suffix}"
  type = "OCR_PROCESSOR"
}

resource "google_document_ai_processor_default_version" "processor" {
  processor = google_document_ai_processor.processor.id
  version = "${google_document_ai_processor.processor.id}/processorVersions/stable"

  lifecycle {
    ignore_changes = [
      # Using "stable" or "rc" will return a specific version from the API; suppressing the diff.
      version,
    ]
  }
}
`, context)
}
