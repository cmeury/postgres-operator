/*
 Copyright 2017 Crunchy Data Solutions, Inc.
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

// Package tpr defines the ThirdPartyResources used within
// the crunchy operator, namely the PgDatabase and PgCluster
// types.
package tpr

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const POLICY_RESOURCE = "pgpolicies"

type PgPolicySpec struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Sql    string `json:"sql"`
	Status string `json:"status"`
}

type PgPolicy struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata"`

	Spec PgPolicySpec `json:"spec"`
}

type PgPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ListMeta `json:"metadata"`

	Items []PgPolicy `json:"items"`
}

func (e *PgPolicy) GetObjectKind() schema.ObjectKind {
	return &e.TypeMeta
}

func (e *PgPolicy) GetObjectMeta() metav1.Object {
	return &e.Metadata
}

func (el *PgPolicyList) GetObjectKind() schema.ObjectKind {
	return &el.TypeMeta
}

func (el *PgPolicyList) GetListMeta() metav1.List {
	return &el.Metadata
}

type PgPolicyListCopy PgPolicyList
type PgPolicyCopy PgPolicy

func (e *PgPolicy) UnmarshalJSON(data []byte) error {
	tmp := PgPolicyCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := PgPolicy(tmp)
	*e = tmp2
	return nil
}

func (el *PgPolicyList) UnmarshalJSON(data []byte) error {
	tmp := PgPolicyListCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := PgPolicyList(tmp)
	*el = tmp2
	return nil
}
