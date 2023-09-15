// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"fmt"
	"github.com/walkerus/go-wiremock"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSubjectModeSchemaWithEnhancedProviderBlock(t *testing.T) {
	ctx := context.Background()

	wiremockContainer, err := setupWiremock(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer wiremockContainer.Terminate(ctx)

	mockSchemaTestServerUrl := wiremockContainer.URI
	confluentCloudBaseUrl := ""
	wiremockClient := wiremock.NewClient(mockSchemaTestServerUrl)
	// nolint:errcheck
	defer wiremockClient.Reset()

	// nolint:errcheck
	defer wiremockClient.ResetAllScenarios()

	readCreatedSubjectModeResponse, _ := ioutil.ReadFile("../testdata/subject_mode/read_created_subject_mode.json")
	_ = wiremockClient.StubFor(wiremock.Get(wiremock.URLPathEqualTo(updateSubjectModePath)).
		InScenario(subjectModeDataSourceScenarioName).
		WhenScenarioStateIs(wiremock.ScenarioStateStarted).
		WillReturn(
			string(readCreatedSubjectModeResponse),
			contentTypeJSONHeader,
			http.StatusOK,
		))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		// https://www.terraform.io/docs/extend/testing/acceptance-tests/teststep.html
		// https://www.terraform.io/docs/extend/best-practices/testing.html#built-in-patterns
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSubjectModeDataSourceConfigWithEnhancedProviderBlock(confluentCloudBaseUrl, mockSchemaTestServerUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSchemaExists(fullSubjectModeDataSourceLabel),
					resource.TestCheckResourceAttr(fullSubjectModeDataSourceLabel, "id", fmt.Sprintf("%s/%s", testStreamGovernanceClusterId, testSubjectName)),
					resource.TestCheckResourceAttr(fullSubjectModeDataSourceLabel, "schema_registry_cluster.#", "0"),
					resource.TestCheckNoResourceAttr(fullSubjectModeDataSourceLabel, "schema_registry_cluster.0.id"),
					resource.TestCheckResourceAttr(fullSubjectModeDataSourceLabel, "subject_name", testSubjectName),
					resource.TestCheckResourceAttr(fullSubjectModeDataSourceLabel, "mode", testSubjectMode),
					resource.TestCheckResourceAttr(fullSubjectModeDataSourceLabel, "credentials.#", "0"),
					resource.TestCheckNoResourceAttr(fullSubjectModeDataSourceLabel, "credentials.0.key"),
					resource.TestCheckNoResourceAttr(fullSubjectModeDataSourceLabel, "credentials.0.secret"),
					resource.TestCheckNoResourceAttr(fullSubjectModeDataSourceLabel, "rest_endpoint"),
					resource.TestCheckResourceAttr(fullSubjectModeDataSourceLabel, "%", strconv.Itoa(testNumberOfSubjectModeDataSourceAttributes)),
				),
			},
		},
	})
}

func testAccCheckSubjectModeDataSourceConfigWithEnhancedProviderBlock(confluentCloudBaseUrl, mockServerUrl string) string {
	return fmt.Sprintf(`
	provider "confluent" {
      endpoint = "%s"
      schema_registry_rest_endpoint = "%s"
      schema_registry_api_key = "%s"
      schema_registry_api_secret = "%s"
      schema_registry_id = "%s"
    }
	data "confluent_subject_mode" "%s" {
	  subject_name = "%s"
	}
	`, confluentCloudBaseUrl, mockServerUrl, testSchemaRegistryKey, testSchemaRegistrySecret, testStreamGovernanceClusterId, testSchemaResourceLabel, testSubjectName)
}
