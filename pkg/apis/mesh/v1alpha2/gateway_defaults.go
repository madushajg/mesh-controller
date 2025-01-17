/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
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

package v1alpha2

func (g *Gateway) SetDefaults() {
	g.Spec.SetDefaults()
	g.Status.SetDefaults()
}

func (gs *GatewaySpec) SetDefaults() {
	// if cs.Type == "" {
	// 	cs.Type = ComponentTypeDeployment
	// }
	// cs.ScalingPolicy.SetDefaults()
	// for i, _ := range cs.Ports {
	// 	cs.Ports[i].SetDefaults()
	// }
}

// func (sp *ScalingPolicy) SetDefaults() {
// 	if sp.Replicas == nil && sp.Hpa == nil && sp.Kpa == nil {
// 		sp.Replicas = ptr.Int32(1)
// 	}
// 	if sp.Hpa != nil && sp.Hpa.MinReplicas == nil {
// 		sp.Hpa.MinReplicas = ptr.Int32(1)
// 	}

// 	if sp.Kpa != nil && sp.Kpa.MinReplicas == nil {
// 		sp.Kpa.MinReplicas = ptr.Int32(0)
// 	}
// }

// func (pm *PortMapping) SetDefaults() {
// 	if pm.Protocol == "" {
// 		pm.Protocol = ProtocolTCP
// 	}
// 	if pm.Name == "" {
// 		pm.Name = fmt.Sprintf("%d-%d", pm.Port, pm.TargetPort)
// 	}
// }

func (gs *GatewayStatus) SetDefaults() {
	// if cstat.Status == "" {
	// 	cstat.Status = ComponentCurrentStatusUnknown
	// }
}
