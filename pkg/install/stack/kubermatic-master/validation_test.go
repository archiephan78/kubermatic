/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubermaticmaster

import (
	"testing"

	"github.com/Masterminds/semver/v3"

	operatorv1alpha1 "k8c.io/kubermatic/v2/pkg/crd/operator/v1alpha1"

	"k8s.io/utils/pointer"
)

func TestClusterVersionIsConfigured(t *testing.T) {
	testcases := []struct {
		name       string
		version    *semver.Version
		versions   operatorv1alpha1.KubermaticVersioningConfiguration
		configured bool
	}{
		{
			name:    "version is directly supported",
			version: semver.MustParse("1.0.0"),
			versions: operatorv1alpha1.KubermaticVersioningConfiguration{
				Versions: []*semver.Version{
					semver.MustParse("1.0.0"),
				},
			},
			configured: true,
		},
		{
			name:    "version is not configured",
			version: semver.MustParse("1.0.0"),
			versions: operatorv1alpha1.KubermaticVersioningConfiguration{
				Versions: []*semver.Version{
					semver.MustParse("2.0.0"),
				},
			},
			configured: false,
		},
		{
			name:    "update constraint matches because it's auto update",
			version: semver.MustParse("1.0.0"),
			versions: operatorv1alpha1.KubermaticVersioningConfiguration{
				Updates: []operatorv1alpha1.Update{
					{
						From:      "1.0.0",
						To:        "2.0.0",
						Automatic: pointer.Bool(true),
					},
				},
			},
			configured: true,
		},
		{
			name:    "constraint expression matches",
			version: semver.MustParse("1.2.3"),
			versions: operatorv1alpha1.KubermaticVersioningConfiguration{
				Updates: []operatorv1alpha1.Update{
					{
						From:      "1.2.*",
						To:        "2.0.0",
						Automatic: pointer.Bool(true),
					},
				},
			},
			configured: true,
		},
		{
			name:    "update constraint does not match because it's no auto update",
			version: semver.MustParse("1.0.0"),
			versions: operatorv1alpha1.KubermaticVersioningConfiguration{
				Updates: []operatorv1alpha1.Update{
					{
						From:      "1.0.0",
						To:        "2.0.0",
						Automatic: pointer.Bool(false),
					},
				},
			},
			configured: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			config := operatorv1alpha1.KubermaticConfiguration{
				Spec: operatorv1alpha1.KubermaticConfigurationSpec{
					Versions: operatorv1alpha1.KubermaticVersionsConfiguration{
						Kubernetes: testcase.versions,
					},
				},
			}

			constraints, err := getAutoUpdateConstraints(&config)
			if err != nil {
				t.Fatalf("Failed to determine update constraints: %v", err)
			}

			configured := clusterVersionIsConfigured(testcase.version, &config, constraints)
			if configured != testcase.configured {
				if testcase.configured {
					t.Fatalf("Expected %q to be supported, but it is not.", testcase.version)
				} else {
					t.Fatalf("Expected %q to not be supported, but it is.", testcase.version)
				}
			}
		})
	}
}
