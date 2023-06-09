/*
Copyright 2023. projectsveltos.io. All rights reserved.

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

package controllers

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/projectsveltos/libsveltos/lib/set"
)

var (
	GetCurrentReferences               = (*AddonConstraintReconciler).getCurrentReferences
	RequeueAddonConstraintForReference = (*AddonConstraintReconciler).requeueAddonConstraintForReference
	GetMatchingClusters                = (*AddonConstraintReconciler).getMatchingClusters
	UpdateReferenceMap                 = (*AddonConstraintReconciler).updateReferenceMap
	UpdateClusterMap                   = (*AddonConstraintReconciler).updateClusterMap
	CleanMaps                          = (*AddonConstraintReconciler).cleanMaps
	CollectContentOfConfigMap          = (*AddonConstraintReconciler).collectContentOfConfigMap
	CollectContentOfSecret             = (*AddonConstraintReconciler).collectContentOfSecret
	CollectOpenapiValidations          = (*AddonConstraintReconciler).collectOpenapiValidations
)

var (
	ShouldAddClusterEntry = shouldAddClusterEntry
	AddClusterEntry       = addClusterEntry
	AnnotateCluster       = annotateCluster
)

var (
	AddTypeInformationToObject = addTypeInformationToObject
	WalkDir                    = walkDir
)

func (m *manager) GetMap() *map[corev1.ObjectReference]*set.Set {
	return &m.addonConstraints
}

func Reset() {
	managerInstance = nil
}
