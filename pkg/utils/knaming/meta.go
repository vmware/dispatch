///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package knaming

import (
	"encoding/json"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	dapi "github.com/vmware/dispatch/pkg/api/v1"
)

const (
	NameLabel    = "dispatchframework.io/name"
	ProjectLabel = "dispatchframework.io/project"
	KnTypeLabel  = "knative.dev/type"

	DefaultProject = "default"

	InitialObjectAnnotation = "dispatchframework.io/initialObject"
)

func ToJSONString(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		// TODO the right thing
		panic(errors.Wrapf(err, "could not JSON-marshal this: %+v", obj))
	}
	return string(bs)
}

func ToObjectMeta(meta dapi.Meta, initialObject interface{}) *v1.ObjectMeta {
	if meta.Project == "" {
		meta.Project = DefaultProject
	}
	name := meta.Project + "-" + meta.Name
	labels := map[string]string{NameLabel: meta.Name, ProjectLabel: meta.Project}

	switch initialObject.(type) {
	case dapi.Function:
		name = "d-fn-" + name
		labels[KnTypeLabel] = "function"
	}

	return &v1.ObjectMeta{
		Name:        name,
		Labels:      labels,
		Annotations: map[string]string{InitialObjectAnnotation: ToJSONString(initialObject)},
	}
}

func SetLabel(meta *v1.ObjectMeta, key, value string) *v1.ObjectMeta {
	meta.Labels[key] = value
	return meta
}