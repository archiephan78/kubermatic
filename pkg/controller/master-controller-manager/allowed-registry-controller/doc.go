/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

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

/*
Package allowedregistrycontroller contains a controller that is responsible for ensuring that the allowed registries
are being synced into corresponding Constraints and Constraint Templates, so that in the User clusters using OPA, users can
only deploy workloads that are from allowed registries

*/
package allowedregistrycontroller
