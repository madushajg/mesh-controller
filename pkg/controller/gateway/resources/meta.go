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
	"github.com/cellery-io/mesh-controller/pkg/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cellery-io/mesh-controller/pkg/apis/mesh"
	"github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
)

func createGatewayLabels(gateway *v1alpha1.Gateway) map[string]string {
	labels := make(map[string]string, len(gateway.ObjectMeta.Labels)+2)
	labels[mesh.CellGatewayLabelKey] = gateway.Name
	labels[appLabelKey] = gateway.Name

	for k, v := range gateway.ObjectMeta.Labels {
		labels[k] = v
	}
	return labels
}

func createGatewaySelector(gateway *v1alpha1.Gateway) *metav1.LabelSelector {
	return &metav1.LabelSelector{MatchLabels: createGatewayLabels(gateway)}
}

func GatewayConfigMapName(gateway *v1alpha1.Gateway) string {
	return gateway.Name + "-config"
}

func GatewayDeploymentName(gateway *v1alpha1.Gateway) string {
	imageName, _, _ := extractImageInfo(gateway)
	if len(imageName) > 0 {
		return imageName + "-" + gateway.Name + "-deployment"
	}
	return gateway.Name + "-deployment"
}

func GatewayK8sServiceName(gateway *v1alpha1.Gateway) string {
	return gateway.Name + "-service"
}

func IstioGatewayName(gateway *v1alpha1.Gateway) string {
	return gateway.Name
}

func IstioVSName(gateway *v1alpha1.Gateway) string {
	return gateway.Name
}

func IstioIngressVirtualServiceName(gateway *v1alpha1.Gateway) string {
	return gateway.Name + "-ingress-virtual-service"
}

func GatewayFullK8sServiceName(gateway *v1alpha1.Gateway) string {
	return GatewayK8sServiceName(gateway) + "." + gateway.Namespace
}

func IstioDestinationRuleName(gateway *v1alpha1.Gateway) string {
	return gateway.Name + "-destination-rule"
}

func extractImageInfo(gateway *v1alpha1.Gateway) (string, string, string) {
	annotations := gateway.Annotations
	if annotations == nil || len(annotations) == 0 {
		// no annotations found
		return "", "", ""
	}
	return annotations[controller.CellImageNameAnnotation],
		annotations[controller.CellImageVersionAnnotation], annotations[controller.CellImageOrgAnnotation]
}
