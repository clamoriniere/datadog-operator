// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package utils

import (
	datadoghqv1alpha1 "github.com/DataDog/datadog-operator/pkg/apis/datadoghq/v1alpha1"
)

// GetAPIKeySecretName return the API key secret name
func GetAPIKeySecretName(dda *datadoghqv1alpha1.DatadogAgent) string {
	if dda.Spec.Credentials.APIKeyExistingSecret != "" {
		return dda.Spec.Credentials.APIKeyExistingSecret
	}
	return dda.Name
}

// GetAppKeySecretName return the APP key secret name
func GetAppKeySecretName(dda *datadoghqv1alpha1.DatadogAgent) string {
	if dda.Spec.Credentials.AppKeyExistingSecret != "" {
		return dda.Spec.Credentials.AppKeyExistingSecret
	}
	return dda.Name
}
