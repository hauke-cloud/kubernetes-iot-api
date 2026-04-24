# Kubernetes IoT API

<img src="https://raw.githubusercontent.com/hauke-cloud/.github/main/resources/img/organisation-logo-small.png" alt="hauke.cloud logo" width="109" height="123" align="right">

Shared Kubernetes Custom Resource Definitions (CRDs) for the hauke.cloud IoT infrastructure.

## Overview

This repository contains Kubernetes API type definitions (`iot.hauke.cloud/v1alpha1`) that are shared across multiple hauke.cloud IoT services. By consolidating CRDs in a single repository, we ensure consistency and avoid duplication.

## API Group

All CRDs in this repository use the API group: **`iot.hauke.cloud/v1alpha1`**

## CRDs

All CRDs are in the `api/v1alpha1` package and share the `iot.hauke.cloud/v1alpha1` API group.

### Database

Represents a PostgreSQL/TimescaleDB database connection for storing IoT sensor data.

**Used by**:
- [database-manager](https://github.com/hauke-cloud/database-manager) - Manages database connections and lifecycle
- [mqtt-sensor-exporter](https://github.com/hauke-cloud/mqtt-sensor-exporter) - Stores MQTT sensor data
- Other IoT services that need database storage

**Example**:
```yaml
apiVersion: iot.hauke.cloud/v1alpha1
kind: Database
metadata:
  name: sensors-db
spec:
  host: postgres.default.svc.cluster.local
  port: 5432
  database: sensors
  username: app
  passwordSecretRef:
    name: db-credentials
  supportedSensorTypes:
    - moisture
    - temperature
```

### MQTTBridge

Represents an MQTT broker connection for collecting IoT sensor data.

**Used by**:
- [mqtt-sensor-exporter](https://github.com/hauke-cloud/mqtt-sensor-exporter) - MQTT data collector operator

**Example**:
```yaml
apiVersion: iot.hauke.cloud/v1alpha1
kind: MQTTBridge
metadata:
  name: tasmota-bridge
spec:
  host: mosquitto.mqtt.svc.cluster.local
  port: 1883
  deviceType: tasmota
  credentialsSecretRef:
    name: mqtt-credentials
```

### Device

Represents an IoT device discovered on an MQTT broker.

**Used by**:
- [mqtt-sensor-exporter](https://github.com/hauke-cloud/mqtt-sensor-exporter) - Tracks discovered devices

**Example**:
```yaml
apiVersion: iot.hauke.cloud/v1alpha1
kind: Device
metadata:
  name: sensor-001
spec:
  bridgeRef:
    name: tasmota-bridge
  friendlyName: "Living Room Sensor"
```

## Installation

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

## Usage

### Single Import for All CRDs

One import statement gives you access to all CRDs in the API group:

```go
import iotv1alpha1 "github.com/hauke-cloud/kubernetes-iot-api/api/v1alpha1"

// Use any CRD from the package:
db := &iotv1alpha1.Database{}
dbList := &iotv1alpha1.DatabaseList{}

// Future CRDs (when added):
// bridge := &iotv1alpha1.MQTTBridge{}
// device := &iotv1alpha1.Device{}
```

### Importing in Your Project

```go
import (
    iotv1alpha1 "github.com/hauke-cloud/kubernetes-iot-api/api/v1alpha1"
)

// Register with scheme
utilruntime.Must(iotv1alpha1.AddToScheme(scheme))
```

### Example: Watching Database Resources

```go
func (r *YourReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&yourv1alpha1.YourResource{}).
        Watches(
            &iotv1alpha1.Database{},
            handler.EnqueueRequestsFromMapFunc(r.findResourcesUsingDatabase),
        ).
        Complete(r)
}
```

### Example: Registering Service with Database

```go
import (
    iotv1alpha1 "github.com/hauke-cloud/kubernetes-iot-api/api/v1alpha1"
    "k8s.io/apimachinery/pkg/types"
)

// Register your service as connected to a database
err := iotv1alpha1.RegisterService(
    ctx,
    client,
    types.NamespacedName{Name: "sensors-db", Namespace: "default"},
    "my-service",
    "my-namespace",
)
```

## Development

### Generating Code

After modifying types, regenerate the deepcopy code:

```bash
controller-gen object:headerFile=hack/boilerplate.go.txt paths="./..."
```

### Testing

```bash
go mod tidy
go build ./...
```

## Versioning

This project follows semantic versioning. The current version is shown in git tags.

To use a specific version:

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

## Architecture

### Why a Shared API Repository?

1. **Consistency**: Single source of truth for shared types
2. **No Duplication**: Avoid copying CRD definitions across services
3. **Lightweight**: Import only type definitions, not controller logic
4. **Versioning**: Independent version lifecycle from operators
5. **Best Practice**: Follows Kubernetes ecosystem patterns (like `k8s.io/api`)

### Comparison to Other Projects

This pattern is used by:
- Kubernetes: `k8s.io/api` vs `k8s.io/kubernetes`
- cert-manager: Separate API package
- Istio: `istio.io/api` vs `istio.io/istio`

## Related Projects

- [database-manager](https://github.com/hauke-cloud/database-manager) - Database operator
- [mqtt-sensor-exporter](https://github.com/hauke-cloud/mqtt-sensor-exporter) - MQTT sensor data collector
- [irrigator](https://github.com/hauke-cloud/irrigator) - Irrigation controller

## Contributing

To become a contributor, please check out the [CONTRIBUTING](CONTRIBUTING.md) file.

## License

This Project is licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or support requests, please open an issue in this repository or contact us at [contact@hauke.cloud](mailto:contact@hauke.cloud).
