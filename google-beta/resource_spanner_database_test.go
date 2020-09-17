package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSpannerDatabase_basic(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	rnd := randString(t, 10)
	instanceName := fmt.Sprintf("my-instance-%s", rnd)
	databaseName := fmt.Sprintf("mydb_%s", rnd)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpannerDatabaseDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerDatabase_basic(instanceName, databaseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("google_spanner_database.basic", "state"),
				),
			},
			{
				// Test import with default Terraform ID
				ResourceName:            "google_spanner_database.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ddl"},
			},
			{
				Config: testAccSpannerDatabase_basicUpdate(instanceName, databaseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("google_spanner_database.basic", "state"),
				),
			},
			{
				// Test import with default Terraform ID
				ResourceName:            "google_spanner_database.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ddl"},
			},
			{
				ResourceName:            "google_spanner_database.basic",
				ImportStateId:           fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instanceName, databaseName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ddl"},
			},
			{
				ResourceName:            "google_spanner_database.basic",
				ImportStateId:           fmt.Sprintf("instances/%s/databases/%s", instanceName, databaseName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ddl"},
			},
			{
				ResourceName:            "google_spanner_database.basic",
				ImportStateId:           fmt.Sprintf("%s/%s", instanceName, databaseName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ddl"},
			},
		},
	})
}

func testAccSpannerDatabase_basic(instanceName, databaseName string) string {
	return fmt.Sprintf(`
resource "google_spanner_instance" "basic" {
  name         = "%s"
  config       = "regional-us-central1"
  display_name = "display-%s"
  num_nodes    = 1
}

resource "google_spanner_database" "basic" {
  instance = google_spanner_instance.basic.name
  name     = "%s"
  ddl = [
	"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
	"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
  ]
}
`, instanceName, instanceName, databaseName)
}

func testAccSpannerDatabase_basicUpdate(instanceName, databaseName string) string {
	return fmt.Sprintf(`
resource "google_spanner_instance" "basic" {
  name         = "%s"
  config       = "regional-us-central1"
  display_name = "display-%s"
  num_nodes    = 1
}

resource "google_spanner_database" "basic" {
  instance = google_spanner_instance.basic.name
  name     = "%s"
  ddl = [
	"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
	"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
	"CREATE TABLE t3 (t3 INT64 NOT NULL,) PRIMARY KEY(t3)",
  ]
}
`, instanceName, instanceName, databaseName)
}

// Unit Tests for type spannerDatabaseId
func TestDatabaseNameForApi(t *testing.T) {
	id := spannerDatabaseId{
		Project:  "project123",
		Instance: "instance456",
		Database: "db789",
	}
	actual := id.databaseUri()
	expected := "projects/project123/instances/instance456/databases/db789"
	expectEquals(t, expected, actual)
}

// Unit Tests for ForceNew when the change in ddl
func TestSpannerDatabase_resourceSpannerDBDdlCustomDiffFuncForceNew(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		before   interface{}
		after    interface{}
		forcenew bool
	}{
		"remove_old_statements": {
			before: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)"},
			after: []interface{}{
				"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)"},
			forcenew: true,
		},
		"append_new_statements": {
			before: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)"},
			after: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
				"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
			},
			forcenew: false,
		},
		"no_change": {
			before: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)"},
			after: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)"},
			forcenew: false,
		},
		"order_of_statments_change": {
			before: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
				"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
				"CREATE TABLE t3 (t3 INT64 NOT NULL,) PRIMARY KEY(t3)",
			},
			after: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
				"CREATE TABLE t3 (t3 INT64 NOT NULL,) PRIMARY KEY(t3)",
				"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
			},
			forcenew: true,
		},
		"missing_an_old_statement": {
			before: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
				"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
				"CREATE TABLE t3 (t3 INT64 NOT NULL,) PRIMARY KEY(t3)",
			},
			after: []interface{}{
				"CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
				"CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
			},
			forcenew: true,
		},
	}

	for tn, tc := range cases {
		d := &ResourceDiffMock{
			Before: map[string]interface{}{
				"ddl": tc.before,
			},
			After: map[string]interface{}{
				"ddl": tc.after,
			},
		}
		err := resourceSpannerDBDdlCustomDiffFunc(d)
		if err != nil {
			t.Errorf("failed, expected no error but received - %s for the condition %s", err, tn)
		}
		if d.IsForceNew != tc.forcenew {
			t.Errorf("ForceNew not setup correctly for the condition-'%s', expected:%v;actual:%v", tn, tc.forcenew, d.IsForceNew)
		}
	}
}
