<!-- llm-readme-management spec=1 commit=634ec964adc19959a719f530a135e889a68015f5 template=terraform model=qwen3.8-27b-q4 digest=f68682b55584 generated=2026-10-12T07:47:19Z -->
<a href="https://hauke.cloud" target="_blank"><img src="https://img.shields.io/badge/home-hauke.cloud-brightgreen" alt="hauke.cloud" style="display: block;" /></a>
<a href="https://github.com/hauke-cloud" target="_blank"><img src="https://img.shields.io/badge/github-hauke.cloud-blue" alt="hauke.cloud Github Organisation" style="display: block;" /></a>
<a href="https://github.com/hauke-cloud/llm-readme-management" target="_blank"><img src="https://img.shields.io/badge/template-terraform-orange" alt="Repository type - terraform" style="display: block;" /></a>


# Template Repository


<img src="https://raw.githubusercontent.com/hauke-cloud/.github/main/resources/img/organisation-logo-small.png" alt="hauke.cloud logo" width="109" height="123" align="right">


<llm header hint="Say whether this is a reusable module or a root module that owns real state.">

This Go module provides the Kubernetes Custom Resource Definition type definitions for the hauke.cloud IoT platform, grouped under `iot.hauke.cloud/v1alpha1`. It is a reusable library you import into your own operator or service code to declare `Database`, `MQTTBridge`, `Device`, and `Schedule` resources. It is written for developers building the hauke-cloud operator projects.

</llm>


## :book: Description

<llm description>

This repository provides the Go type definitions for the Kubernetes Custom Resource Definitions (CRDs) used by the hauke.cloud IoT platform. If you are building a Kubernetes operator or service that reads, writes, or reconciles IoT resources, you import these types as a library rather than redefining them locally.

The module exposes a single package, `api/v1alpha1`, under the API group `iot.hauke.cloud/v1alpha1`. It contains struct definitions with kubebuilder validation markers, a shared `SecretReference` helper, scheme-registration boilerplate, and generated `DeepCopy` methods. You add it to your project with:

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

The four CRD types defined here are:

- **Database** – declarative PostgreSQL/TimescaleDB connection for storing IoT sensor data, with SSL, pooling, and batching parameters.
- **MQTTBridge** – MQTT broker connection with per-topic subscriptions, TLS, and device discovery.
- **Device** – an IoT device discovered on a bridge, carrying sensor type, alert conditions, and reachability status.
- **Schedule** – a cron-driven irrigation schedule for valve devices with optional execution conditions.

These types are consumed by the `hauke-cloud/database-manager`, `hauke-cloud/mqtt-sensor-exporter`, and `hauke-cloud/irrigator` operators, which install the corresponding CRDs into a cluster and act on the resources you create.

</llm>


## :clipboard: Requirements

<llm requirements hint="Give the Terraform version from .terraform-version and the provider constraints from versions.tf, plus the credentials the providers need.">

- **Go 1.25.3** — the toolchain version pinned in `go.mod`. You need it to build the module or run `go mod tidy`.
- **`controller-gen`** (from `sigs.k8s.io/controller-tools`) — required only if you modify the CRD type definitions and need to regenerate `zz_generated.deepcopy.go`.
- **`pre-commit`** — mandatory for contributors per `CONTRIBUTING.md`. Install the hooks with `pre-commit install` before pushing.

No Kubernetes cluster, cloud credentials, or Terraform/OpenTofu installation is required to use this repository. It is a Go type-definition library; the `.terraform-version` and `.opentofu-version` files at the repo root are vestigial and referenced by no code. Consumers who later deploy the CRDs into a cluster need that cluster and the consuming operators (`database-manager`, `mqtt-sensor-exporter`, `irrigator`), but that deployment is outside the scope of this repository.

</llm>


## 🚀 Getting started

<llm getting_started hint="terraform init, plan and apply, with the backend configuration the repository actually uses. Say plainly if apply touches real infrastructure.">

1. Clone the repository and enter the working directory.

```bash
git clone https://github.com/hauke-cloud/kubernetes-iot-api.git
cd kubernetes-iot-api
```

2. Resolve module dependencies.

```bash
go mod tidy
```

3. Verify the library compiles.

```bash
go build ./...
```

This repository is a Go type-definition library, not a Terraform project. The `.opentofu-version` and `.terraform-version` files at the root are vestigial; no `.tf` or `.hcl` files exist. Once the build succeeds, the module is ready to be imported by consumer operators (`hauke-cloud/database-manager`, `hauke-cloud/mqtt-sensor-exporter`, `hauke-cloud/irrigator`) via:

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

No cluster, credentials, or infrastructure provisioning is required to use the library. The CRD YAML manifests are not checked in here; the consuming operators are responsible for installing them into a target cluster.

</llm>


## :airplane: Usage

<llm usage hint="For a reusable module, the central example is a module block with source, version and the required variables filled in from variables.tf. For a root module, show the workflow instead.">

This repository is a Go type-definition library. You import it from operator projects (for example `hauke-cloud/database-manager`, `hauke-cloud/mqtt-sensor-exporter`, `hauke-cloud/irrigator`) to work with the `iot.hauke.cloud/v1alpha1` CRD types.

**Add the dependency**

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

**Register the CRD scheme in your operator**

Import the `v1alpha1` package and register it with your controller-runtime scheme:

```go
import (
    iotv1alpha1 "github.com/hauke-cloud/kubernetes-iot-api/api/v1alpha1"
)

iotv1alpha1.AddToScheme(scheme)
```

**Define CR manifests**

Once the consuming operator is running, you create custom resources in your cluster. For example, a `Database` CR that the `database-manager` operator reconciles:

```yaml
apiVersion: iot.hauke.cloud/v1alpha1
kind: Database
metadata:
  name: sensor-db
  namespace: iot
spec:
  host: timescale.internal
  port: 5432
  database: iot_sensors
  username: iot_writer
  passwordSecretRef:
    name: db-password
  sslMode: require
  supportedSensorTypes:
    - temperature
    - humidity
    - moisture
  maxConnections: 20
  batchSize: 200
  batchTimeout: 15
```

The referenced Secret must contain a key named `password`.

**Regenerate deepcopy code (contributors)**

After modifying any type in `api/v1alpha1/`, regenerate the DeepCopy methods:

```bash
controller-gen object:headerFile=hack/boilerplate.go.txt paths="./..."
```

</llm>


## :wrench: Configuration

<llm configuration hint="A table of the variables in variables.tf: name, type, default, required. Point at variables.tf for the full set and mention outputs.tf if it exists.">

This repository is a Go type-definition library, not a Terraform project. It exposes no Terraform variables, Helm values, environment variables, or CLI flags. The configuration surface is the CRD spec fields in `api/v1alpha1/`, which you set in Kubernetes Custom Resource manifests consumed by operators such as `hauke-cloud/database-manager` and `hauke-cloud/irrigator`.

The table below lists the required and defaulted fields for each CRD. The complete field set lives in `database_types.go`, `mqttbridge_types.go`, `device_types.go`, and `schedule_types.go` in that directory.

| CRD | Field | Type | Default | Required | Description |
|-----|-------|------|---------|----------|-------------|
| Database | host | string | — | yes | PostgreSQL/TimescaleDB host |
| Database | port | int | 5432 | yes | Database port |
| Database | database | string | — | yes | Database name |
| Database | username | string | — | yes | Database user |
| Database | sslMode | string | require | no | disable / require / verify-ca / verify-full |
| Database | maxConnections | int | 10 | no | Pool max (≤ 100) |
| Database | batchSize | int | 100 | no | Batch size (≤ 1000) |
| MQTTBridge | host | string | — | yes | Broker host |
| MQTTBridge | port | int | 1883 | no | Broker port |
| MQTTBridge | deviceType | string | tasmota | no | zigbee2mqtt / tasmota / generic |
| Device | bridgeRef.name | string | — | yes | Name of the MQTTBridge |
| Device | ieeeAddr | string | — | yes | IEEE address of the device |
| Schedule | cronExpression | string | — | yes | Cron expression |
| Schedule | durationSeconds | int | — | yes | Run duration (1–86400) |
| Schedule | timeZone | string | UTC | no | IANA time zone |

Optional fields such as `credentialsSecretRef`, `topics`, `corrections`, `alertCondition`, and `executionConditions` are defined in the same source files.

</llm>


## :hammer: Development

<llm development hint="Cover terraform fmt, validate, tflint and terraform-docs where the repository configures them.">

This repository contains no test files. The build check is:

```bash
go mod tidy
go build ./...
```

After modifying any type in `api/v1alpha1/`, regenerate the DeepCopy methods and commit the result:

```bash
controller-gen object:headerFile=hack/boilerplate.go.txt paths="./..."
```

The generated file is `api/v1alpha1/zz_generated.deepcopy.go`. Do not edit it by hand; a stale copy will cause a build failure in downstream consumers.

Install the pre-commit hooks before your first commit:

```bash
pre-commit install
pre-commit run --all-files
```

The configured hooks are `pre-commit-hooks v4.4.0` and `gitleaks v8.18.0`. A failing hook blocks the local commit.

CI enforces a conventional-commit title on every pull request via `amannn/action-semantic-pull-request`. A title such as `feat: add Schedule CRD` passes; a free-form title such as `added schedule` does not.

There is no Makefile, no `terraform fmt`, `terraform validate`, `tflint`, or `terraform-docs` step, and no linter beyond the pre-commit hooks. The `.opentofu-version` and `.terraform-version` files at the repository root are vestigial; no Terraform or OpenTofu code exists in this repository.

</llm>


## 📄 License

This Project is licensed under the GNU General Public License v3.0

- see the [LICENSE](LICENSE) file for details.


## :coffee: Contributing

To become a contributor, please check out the [CONTRIBUTING](CONTRIBUTING.md) file.


## :email: Contact

For any inquiries or support requests, please open an issue in this
repository or contact us at [contact@hauke.cloud](mailto:contact@hauke.cloud).
