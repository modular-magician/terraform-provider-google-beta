// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package sql

import (
	"testing"
)

func TestMaintenanceVersionDiffSuppress(t *testing.T) {
	cases := map[string]struct {
		Old, New       string
		ShouldSuppress bool
	}{
		"older configuration maintenance version than current version should suppress diff": {
			Old:            "MYSQL_8_0_26.R20220508.01_09",
			New:            "MYSQL_5_7_37.R20210508.01_03",
			ShouldSuppress: true,
		},
		"older configuration maintenance version than current version should suppress diff with lexicographically smaller database version": {
			Old:            "MYSQL_5_8_10.R20220508.01_09",
			New:            "MYSQL_5_8_7.R20210508.01_03",
			ShouldSuppress: true,
		},
		"newer configuration maintenance version than current version should not suppress diff": {
			Old:            "MYSQL_5_7_37.R20210508.01_03",
			New:            "MYSQL_8_0_26.R20220508.01_09",
			ShouldSuppress: false,
		},
	}

	for tn, tc := range cases {
		tc := tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()
			if maintenanceVersionDiffSuppress("version", tc.Old, tc.New, nil) != tc.ShouldSuppress {
				t.Fatalf("%q => %q expect DiffSuppress to return %t", tc.Old, tc.New, tc.ShouldSuppress)
			}
		})
	}
}

func TestDatabaseVersionDiffSuppress(t *testing.T) {
	cases := map[string]struct {
		Old, New       string
		ShouldSuppress bool
	}{
		"both old and new MySQL 8.0 versions should suppress diff": {
			Old:            "MYSQL_8_0_36.R20250213.00_00",
			New:            "MYSQL_8_0_39.R20250213.00_00",
			ShouldSuppress: true,
		},
		"old MySQL 5.x version with new MySQL 8.0 version should not suppress diff": {
			Old:            "MYSQL_5_7_21.R20250213.00_00",
			New:            "MYSQL_8_0_31.R20250213.00_00",
			ShouldSuppress: false,
		},
		"both old and new MySQL 5.x versions should not suppress diff": {
			Old:            "MYSQL_5_6_13.R20250213.00_00",
			New:            "MYSQL_5_7_23.R20250213.00_00",
			ShouldSuppress: false,
		},
	}

	for tn, tc := range cases {
		tc := tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()
			if databaseVersionDiffSuppress("version", tc.Old, tc.New, nil) != tc.ShouldSuppress {
				t.Fatalf("%q => %q expect DiffSuppress to return %t", tc.Old, tc.New, tc.ShouldSuppress)
			}
		})
	}
}
