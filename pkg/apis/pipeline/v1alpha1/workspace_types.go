/*
Copyright 2019 The Tekton Authors

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

package v1alpha1

import (
	"path/filepath"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	corev1 "k8s.io/api/core/v1"
)

// WorkspaceDeclaration is a declaration of a volume that a Task requires.
type WorkspaceDeclaration struct {
	// Name is the name by which you can bind the volume at runtime.
	Name string `json:"name"`
	// Description is an optional human readable description of this volume.
	// +optional
	Description string `json:"description,omitempty"`
	// MountPath overrides the directory that the volume will be made available at.
	// +optional
	MountPath string `json:"mountPath,omitempty"`
}

// GetMountPath returns the mountPath for w which is the MountPath if provided or the
// default if not.
func (w *WorkspaceDeclaration) GetMountPath() string {
	if w.MountPath != "" {
		return w.MountPath
	}
	return filepath.Join(pipeline.WorkspaceDir, w.Name)
}

// WorkspaceBinding maps a Task's declared workspace to a Volume.
// Currently we only support PersistentVolumeClaims and EmptyDir.
type WorkspaceBinding struct {
	// Name is the name of the workspace populated by the volume.
	Name string `json:"name"`
	// SubPath is optionally a directory on the volume which should be used
	// for this binding (i.e. the volume will be mounted at this sub directory).
	// +optional
	SubPath string `json:"subPath,omitempty"`
	// PersistentVolumeClaimVolumeSource represents a reference to a
	// PersistentVolumeClaim in the same namespace. Either this OR EmptyDir can be used.
	// +optional
	PersistentVolumeClaim *corev1.PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
	// EmptyDir represents a temporary directory that shares a Task's lifetime.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// Either this OR PersistentVolumeClaim can be used.
	// +optional
	EmptyDir *corev1.EmptyDirVolumeSource `json:"emptyDir,omitempty"`
}
