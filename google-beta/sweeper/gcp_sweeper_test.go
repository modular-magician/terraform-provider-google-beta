// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package sweeper_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/accessapproval"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/accesscontextmanager"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/activedirectory"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/alloydb"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/apigateway"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/apigee"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/appengine"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/artifactregistry"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/backupdr"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/beyondcorp"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/biglake"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigquery"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryanalyticshub"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryconnection"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigquerydatapolicy"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigquerydatatransfer"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryreservation"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigtable"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/billing"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/binaryauthorization"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/certificatemanager"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudasset"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuild"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuildv2"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudfunctions"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudfunctions2"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudidentity"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudids"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudiot"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudrun"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudrunv2"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudscheduler"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudtasks"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containeranalysis"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containerattached"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/corebilling"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/databasemigrationservice"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datacatalog"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataform"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datafusion"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datalossprevention"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataplex"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataproc"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataprocmetastore"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datastore"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datastream"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/deploymentmanager"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dialogflow"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dialogflowcx"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dns"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/documentai"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/documentaiwarehouse"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/essentialcontacts"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/filestore"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebase"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebasedatabase"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebaseextensions"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebasehosting"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebasestorage"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firestore"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gameservices"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkebackup"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkehub"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkehub2"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkeonprem"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/healthcare"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iam2"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iambeta"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iamworkforcepool"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iap"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/identityplatform"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/kms"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/logging"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/looker"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/memcache"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/mlengine"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/monitoring"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkconnectivity"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkmanagement"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networksecurity"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkservices"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/notebooks"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/orgpolicy"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/osconfig"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/oslogin"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/privateca"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/publicca"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsub"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsublite"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/redis"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/resourcemanager"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/runtimeconfig"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/secretmanager"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/securitycenter"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/securityscanner"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/servicedirectory"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/servicemanagement"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/serviceusage"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/sourcerepo"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/spanner"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/sql"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/storage"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/storageinsights"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/storagetransfer"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/tags"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/tpu"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/vertexai"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/vmwareengine"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/vpcaccess"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/workflows"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/workstations"

	// Manually add the services for DCL resource and handwritten resource sweepers if they are not in the above list
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/apikeys"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/clouddeploy"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/composer"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/container"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containeraws"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containerazure"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/eventarc"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebase"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebaserules"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkconnectivity"
	_ "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/recaptchaenterprise"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}
