/*
Copyright 2021 The CDI Authors.

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

package cluster

import (
	rbacv1 "k8s.io/api/rbac/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"kubevirt.io/containerized-data-importer/pkg/common"
	"kubevirt.io/containerized-data-importer/pkg/operator/resources/utils"
)

func createCronJobResources(args *FactoryArgs) []client.Object {
	return []client.Object{
		createCronJobClusterRole(),
		createCronJobClusterRoleBinding(args.Namespace),
	}
}

func getCronJobClusterPolicyRules() []rbacv1.PolicyRule {
	return []rbacv1.PolicyRule{
		{
			APIGroups: []string{
				"cdi.kubevirt.io",
			},
			Resources: []string{
				"dataimportcrons",
			},
			Verbs: []string{
				"get",
				"list",
				"update",
			},
		},
	}
}

func createCronJobClusterRoleBinding(namespace string) *rbacv1.ClusterRoleBinding {
	return utils.ResourceBuilder.CreateClusterRoleBinding(common.CDICronJobResourceName, common.CDICronJobResourceName, common.CDICronJobResourceName, namespace)
}

func createCronJobClusterRole() *rbacv1.ClusterRole {
	return utils.ResourceBuilder.CreateClusterRole(common.CDICronJobResourceName, getCronJobClusterPolicyRules())
}
