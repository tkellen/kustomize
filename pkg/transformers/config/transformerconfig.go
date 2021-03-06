/*
Copyright 2018 The Kubernetes Authors.

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

// Package config provides the functions to load default or user provided configurations
// for different transformers
package config

import (
	"log"
	"sort"

	"github.com/pkg/errors"

	"sigs.k8s.io/kustomize/v3/pkg/transformers/config/defaultconfig"
)

// TransformerConfig holds the data needed to perform transformations.
type TransformerConfig struct {
	NamePrefix        fsSlice  `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
	NameSuffix        fsSlice  `json:"nameSuffix,omitempty" yaml:"nameSuffix,omitempty"`
	NameSpace         fsSlice  `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	CommonLabels      fsSlice  `json:"commonLabels,omitempty" yaml:"commonLabels,omitempty"`
	CommonAnnotations fsSlice  `json:"commonAnnotations,omitempty" yaml:"commonAnnotations,omitempty"`
	NameReference     nbrSlice `json:"nameReference,omitempty" yaml:"nameReference,omitempty"`
	VarReference      fsSlice  `json:"varReference,omitempty" yaml:"varReference,omitempty"`
	Images            fsSlice  `json:"images,omitempty" yaml:"images,omitempty"`
	Replicas          fsSlice  `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}

// MakeEmptyConfig returns an empty TransformerConfig object
func MakeEmptyConfig() *TransformerConfig {
	return &TransformerConfig{}
}

// MakeDefaultConfig returns a default TransformerConfig.
func MakeDefaultConfig() *TransformerConfig {
	c, err := makeTransformerConfigFromBytes(
		defaultconfig.GetDefaultFieldSpecs())
	if err != nil {
		log.Fatalf("Unable to make default transformconfig: %v", err)
	}
	return c
}

// sortFields provides determinism in logging, tests, etc.
func (t *TransformerConfig) sortFields() {
	sort.Sort(t.NamePrefix)
	sort.Sort(t.NameSpace)
	sort.Sort(t.CommonLabels)
	sort.Sort(t.CommonAnnotations)
	sort.Sort(t.NameReference)
	sort.Sort(t.VarReference)
	sort.Sort(t.Images)
	sort.Sort(t.Replicas)
}

// AddPrefixFieldSpec adds a FieldSpec to NamePrefix
func (t *TransformerConfig) AddPrefixFieldSpec(fs FieldSpec) (err error) {
	t.NamePrefix, err = t.NamePrefix.mergeOne(FieldSpecConfig{FieldSpec: fs, Behavior: "add"})
	return err
}

// AddLabelFieldSpec adds a FieldSpec to CommonLabels
func (t *TransformerConfig) AddLabelFieldSpec(fs FieldSpec) (err error) {
	t.CommonLabels, err = t.CommonLabels.mergeOne(FieldSpecConfig{FieldSpec: fs, Behavior: "add"})
	return err
}

// AddAnnotationFieldSpec adds a FieldSpec to CommonAnnotations
func (t *TransformerConfig) AddAnnotationFieldSpec(fs FieldSpec) (err error) {
	t.CommonAnnotations, err = t.CommonAnnotations.mergeOne(FieldSpecConfig{FieldSpec: fs, Behavior: "add"})
	return err
}

// AddNamereferenceFieldSpec adds a NameBackReferences to NameReference
func (t *TransformerConfig) AddNamereferenceFieldSpec(
	nbrs NameBackReferences) (err error) {
	t.NameReference, err = t.NameReference.mergeOne(nbrs)
	return err
}

// Merge merges two TransformerConfigs objects into
// a new TransformerConfig object
func (t *TransformerConfig) Merge(input *TransformerConfig) (
	merged *TransformerConfig, err error) {
	if input == nil {
		return t, nil
	}
	merged = &TransformerConfig{}
	merged.NamePrefix, err = t.NamePrefix.mergeAll(input.NamePrefix)
	if err != nil {
		return nil, errors.Wrap(err, "NamePrefix")
	}
	merged.NameSpace, err = t.NameSpace.mergeAll(input.NameSpace)
	if err != nil {
		return nil, errors.Wrap(err, "NameSpace")
	}
	merged.CommonAnnotations, err = t.CommonAnnotations.mergeAll(
		input.CommonAnnotations)
	if err != nil {
		return nil, errors.Wrap(err, "CommonAnnotations")
	}
	merged.CommonLabels, err = t.CommonLabels.mergeAll(input.CommonLabels)
	if err != nil {
		return nil, errors.Wrap(err, "CommonLabels")
	}
	merged.VarReference, err = t.VarReference.mergeAll(input.VarReference)
	if err != nil {
		return nil, errors.Wrap(err, "VarReference")
	}
	merged.NameReference, err = t.NameReference.mergeAll(input.NameReference)
	if err != nil {
		return nil, errors.Wrap(err, "NameReference")
	}
	merged.Images, err = t.Images.mergeAll(input.Images)
	if err != nil {
		return nil, errors.Wrap(err, "Images")
	}
	merged.Replicas, err = t.Replicas.mergeAll(input.Replicas)
	if err != nil {
		return nil, errors.Wrap(err, "Replicas`")
	}
	merged.sortFields()
	return merged, nil
}

func (t *TransformerConfig) NamePrefixFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.NamePrefix)
}
func (t *TransformerConfig) NameSuffixFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.NameSuffix)
}
func (t *TransformerConfig) NameSpaceFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.NameSpace)
}
func (t *TransformerConfig) CommonLabelsFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.CommonLabels)
}
func (t *TransformerConfig) CommonAnnotationsFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.CommonAnnotations)
}
func (t *TransformerConfig) VarReferenceFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.VarReference)
}
func (t *TransformerConfig) ImagesFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.Images)
}
func (t *TransformerConfig) ReplicasFieldSpecs() FieldSpecs {
	return NewFieldSpecs(t.Replicas)
}
