/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha2"
	. "github.com/cellery-io/mesh-controller/pkg/meta"
)

// func createGatewayLabels(gateway *v1alpha2.Gateway) map[string]string {
// 	labels := make(map[string]string, len(gateway.ObjectMeta.Labels)+2)
// 	labels[mesh.CellGatewayLabelKey] = gateway.Name
// 	labels[appLabelKey] = gateway.Name

// 	for k, v := range gateway.ObjectMeta.Labels {
// 		labels[k] = v
// 	}
// 	return labels
// }

func makeLabels(gateway *v1alpha2.Gateway) map[string]string {
	return UnionMaps(
		map[string]string{
			AppLabelKey:                  gateway.Name,
			VersionLabelKey:              "v1.0.0",
			ObservabilityGatewayLabelKey: gateway.Name,
		},
		gateway.Labels,
		map[string]string{
			GatewayLabelKey: gateway.Name,
		},
	)
}

func makeSelector(gateway *v1alpha2.Gateway) *metav1.LabelSelector {
	return &metav1.LabelSelector{MatchLabels: makeLabels(gateway)}
}

func makePodAnnotations(gateway *v1alpha2.Gateway) map[string]string {
	return UnionMaps(
		map[string]string{
			IstioSidecarInjectAnnotationKey: "false",
		},
		gateway.Annotations,
	)
}

func ServiceName(gateway *v1alpha2.Gateway) string {
	return gateway.Name + "-service"
}

func DeploymentName(gateway *v1alpha2.Gateway) string {
	return gateway.Name + "-deployment"
}

func JobName(gateway *v1alpha2.Gateway) string {
	return gateway.Name + "-job"
}

// func createGatewaySelector(gateway *v1alpha2.Gateway) *metav1.LabelSelector {
// 	return &metav1.LabelSelector{MatchLabels: createGatewayLabels(gateway)}
// }

func ApiPublisherConfigMap(gateway *v1alpha2.Gateway) string {
	return gateway.Name + "-config"
}

func IstioGatewayName(gateway *v1alpha2.Gateway) string {
	return gateway.Name
}

func IstioVirtualServiceName(gateway *v1alpha2.Gateway) string {
	return gateway.Name
}

func IstioIngressVirtualServiceName(gateway *v1alpha2.Gateway) string {
	return gateway.Name + "-ingress-virtual-service"
}

func GatewayFullK8sServiceName(gateway *v1alpha2.Gateway) string {
	return ServiceName(gateway) + "." + gateway.Namespace
}

func IstioDestinationRuleName(gateway *v1alpha2.Gateway) string {
	return gateway.Name + "-destination-rule"
}
