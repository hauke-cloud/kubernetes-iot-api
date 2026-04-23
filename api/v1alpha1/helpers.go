/*
Copyright 2026 hauke.cloud.

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
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RegisterService registers a service as connected to the database
// This updates the status.connectedServices field without triggering reconciliation
// in the database-manager controller (which uses generation-based skipping)
func RegisterService(ctx context.Context, c client.Client, databaseName types.NamespacedName, serviceName, serviceNamespace string) error {
	db := &Database{}
	if err := c.Get(ctx, databaseName, db); err != nil {
		return err
	}

	// Check if service is already registered
	now := metav1.NewTime(time.Now())
	found := false
	for i := range db.Status.ConnectedServices {
		if db.Status.ConnectedServices[i].Name == serviceName &&
			db.Status.ConnectedServices[i].Namespace == serviceNamespace {
			// Update last seen time
			db.Status.ConnectedServices[i].LastSeenTime = &now
			found = true
			break
		}
	}

	if !found {
		// Add new service
		db.Status.ConnectedServices = append(db.Status.ConnectedServices, ConnectedService{
			Name:         serviceName,
			Namespace:    serviceNamespace,
			LastSeenTime: &now,
		})
	}

	// Update status - this will NOT trigger reconciliation in database-manager
	// because generation doesn't change
	return c.Status().Update(ctx, db)
}

// UnregisterService removes a service from the connected services list
func UnregisterService(ctx context.Context, c client.Client, databaseName types.NamespacedName, serviceName, serviceNamespace string) error {
	db := &Database{}
	if err := c.Get(ctx, databaseName, db); err != nil {
		return err
	}

	// Remove service from list
	filtered := make([]ConnectedService, 0, len(db.Status.ConnectedServices))
	for _, svc := range db.Status.ConnectedServices {
		if svc.Name != serviceName || svc.Namespace != serviceNamespace {
			filtered = append(filtered, svc)
		}
	}
	db.Status.ConnectedServices = filtered

	return c.Status().Update(ctx, db)
}
