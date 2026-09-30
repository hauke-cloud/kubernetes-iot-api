<!-- llm-readme-management spec=1 commit=634ec964adc19959a719f530a135e889a68015f5 template=terraform model=qwen3.8-27b-q4 digest=f68682b55584 generated=2026-09-30T12:06:08Z -->
<a href="https://hauke.cloud" target="_blank"><img src="https://img.shields.io/badge/home-hauke.cloud-brightgreen" alt="hauke.cloud" style="display: block;" /></a>
<a href="https://github.com/hauke-cloud" target="_blank"><img src="https://img.shields.io/badge/github-hauke.cloud-blue" alt="hauke.cloud Github Organisation" style="display: block;" /></a>
<a href="https://github.com/hauke-cloud/llm-readme-management" target="_blank"><img src="https://img.shields.io/badge/template-terraform-orange" alt="Repository type - terraform" style="display: block;" /></a>


# Template Repository


<img src="https://raw.githubusercontent.com/hauke-cloud/.github/main/resources/img/organisation-logo-small.png" alt="hauke.cloud logo" width="109" height="123" align="right">


<llm header hint="Say whether this is a reusable module or a root module that owns real state.">

This Go library defines the shared Kubernetes CRD types (API group `iot.hauke.cloud/v1alpha1`) for the hauke.cloud IoT platform. It is a reusable module with no controllers or runnable binary, imported by the `database-manager`, `mqtt-sensor-exporter`, and `irrigator` operator repos. If you are writing a Go operator that consumes these CRDs, this is the package you import.

</llm>


## :book: Description

<llm description>

This repository is a Go library that defines the shared Kubernetes Custom Resource Definition (CRD) types for the hauke.cloud IoT platform, under the API group `iot.hauke.cloud/v1alpha1`. It exists so that multiple operator projects can import the same struct definitions, generated deepcopy methods, and scheme registration without duplicating them.

You consume this module when you are writing a Go operator that needs to read or write these CRDs. The library contains type definitions only — no controllers, no YAML manifests, and no runnable binary. You add it to your project with `go get github.com/hauke-cloud/kubernetes-iot-api` and import the `api/v1alpha1` package.

The four CRD types defined here are:

- **Database** — PostgreSQL/TimescaleDB connection settings for storing IoT sensor measurements, including SSL mode, pool sizing, and batching.
- **MQTTBridge** — MQTT broker connection for collecting sensor data, with TLS, per-topic subscriptions, and device-type selection.
- **Device** — a device discovered on an MQTT bridge, carrying sensor type, correction factors, and alert thresholds.
- **Schedule** — an irrigation schedule for valve devices, with cron timing, duration, timezone, and pre-execution conditions.

Within the `hauke-cloud` organisation, this module is shared by the `database-manager`, `mqtt-sensor-exporter`, and `irrigator` operator repositories.

</llm>


## :clipboard: Requirements

<llm requirements hint="Give the Terraform version from .terraform-version and the provider constraints from versions.tf, plus the credentials the providers need.">

- **Go 1.25.3** — pinned in `go.mod`; required to build and test the module.
- **`controller-gen`** (`sigs.k8s.io`) — used to regenerate `zz_generated.deepcopy.go` after type changes. No version is pinned in this repository.
- **`pre-commit`** — mandatory per `CONTRIBUTING.md`. The configured hooks are `pre-commit-hooks v4.4.0` and `gitleaks v8.18.0` (see `.pre-commit-config.yaml`).
- **Kubernetes cluster** (consumers only) — the CRDs defined here (`iot.hauke.cloud/v1alpha1`) must be installed in a cluster before any operator can use them. Per the README, the installing operators are `database-manager`, `mqtt-sensor-exporter`, and `irrigator`.

No cloud-provider credentials, Terraform/OpenTofu installation, or other runtime dependencies are required to work with this library. The `.terraform-version` (1.9) and `.opentofu-version` (1.8.0) files are inherited template scaffolding; no Terraform sources exist in this repository.

</llm>


## 🚀 Getting started

<llm getting_started hint="terraform init, plan and apply, with the backend configuration the repository actually uses. Say plainly if apply touches real infrastructure.">

1. Clone the repository.

```bash
git clone https://github.com/hauke-cloud/kubernetes-iot-api.git
cd kubernetes-iot-api
```

2. Sync the Go module dependencies.

```bash
go mod tidy
```

3. Compile the module to confirm it builds.

```bash
go build ./...
```

4. In a consuming operator project, add the library as a dependency.

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

The `v0.1.0` tag is documented in the README but could not be confirmed from the local clone; adjust the version if a newer tag has been published.

This repository contains no Terraform or OpenTofu source files. The `.terraform-version`, `.opentofu-version` pins and the `terraform-deploy.yml.template` workflow are inactive scaffolding inherited from the organisation template. There is nothing to `init`, `plan`, or `apply`, and no backend configuration to review.

</llm>


## :airplane: Usage

<llm usage hint="For a reusable module, the central example is a module block with source, version and the required variables filled in from variables.tf. For a root module, show the workflow instead.">

This repository is a Go library of Kubernetes CRD type definitions for the `iot.hauke.cloud/v1alpha1` API group. You consume it from operator projects such as `database-manager`, `mqtt-sensor-exporter`, or `irrigator`.

**Add the module to your operator project**

```bash
go get github.com/hauke-cloud/kubernetes-iot-api@v0.1.0
```

Import the `api/v1alpha1` package and call `AddToScheme` to register the `Database`, `MQTTBridge`, `Device`, and `Schedule` types with your controller-runtime scheme:

```go
import (
    iotv1alpha1 "github.com/hauke-cloud/kubernetes-iot-api/api/v1alpha1"
    "k8s.io/apimachinery/pkg/runtime"
)

func init() {
    iotv1alpha1.AddToScheme(runtime.NewScheme())
}
```

**Regenerate deepcopy code after editing types**

If you modify any type in `api/v1alpha1/`, regenerate `zz_generated.deepcopy.go` so the generated `DeepCopy` methods stay in sync:

```bash
controller-gen object:headerFile=hack/boilerplate.go.txt paths="./..."
```

**Verify the build**

```bash
go mod tidy
go build ./...
```

</llm>


## :wrench: Configuration

<llm configuration hint="A table of the variables in variables.tf: name, type, default, required. Point at variables.tf for the full set and mention outputs.tf if it exists.">

This repository exposes no Terraform variables, Helm values, environment variables, CLI flags, or configuration files. The configuration surface consists entirely of the CRD spec fields defined in `api/v1alpha1/`. The table below lists the most consequential fields; the type files hold the complete set.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| `Database.spec.host` | string | — | yes | PostgreSQL/TimescaleDB host |
| `Database.spec.port` | int32 | 5432 | yes | Database port |
| `Database.spec.sslMode` | enum | `require` | no | TLS mode: disable, require, verify-ca, verify-full |
| `Database.spec.maxConnections` | int | 10 | no | Pool maximum (max 100) |
| `Database.spec.batchSize` | int | 100 | no | Rows per batch insert |
| `MQTTBridge.spec.host` | string | — | yes | MQTT broker host |
| `MQTTBridge.spec.port` | int | 1883 | yes | MQTT broker port |
| `MQTTBridge.spec.deviceType` | enum | `tasmota` | no | zigbee2mqtt, tasmota, or generic |
| `MQTTBridge.spec.discoveryEnabled` | bool | true | no | Enable device discovery |
| `Device.spec.bridgeRef.name` | string | — | yes | Name of the MQTTBridge |
| `Device.spec.ieeeAddr` | string | — | yes | IEEE address of the device |
| `Device.spec.sensorType` | enum | — | no | One of nine sensor kinds |
| `Schedule.spec.cronExpression` | string | — | yes | Cron schedule |
| `Schedule.spec.durationSeconds` | int | — | yes | Valve open duration (1–86400) |
| `Schedule.spec.timeZone` | string | `UTC` | no | IANA timezone |

The full field set, including all enum values, nested structs, and kubebuilder validation annotations, lives in `api/v1alpha1/database_types.go`, `api/v1alpha1/mqttbridge_types.go`, `api/v1alpha1/device_types.go`, and `api/v1alpha1/schedule_types.go`.

</llm>


## :hammer: Development

<llm development hint="Cover terraform fmt, validate, tflint and terraform-docs where the repository configures them.">

Before pushing, install the pre-commit hooks and run them locally:

```bash
pre-commit install
pre-commit run --all-files
```

The configured hooks (`pre-commit-hooks v4.4.0`, `gitleaks v8.18.0`) check formatting and scan for leaked secrets. Run `pre-commit autoupdate` when you need to bump hook revisions.

CI does not build or test the Go code. The active workflows validate PR titles with `amannn/action-semantic-pull-request`, mark stale issues, and lock threads. Use a conventional-commit prefix (`feat:`, `fix:`, `chore:`, etc.) in your PR title or the title check will fail.

There are no test files in this repository. The closest thing to a test is compiling the module:

```bash
go mod tidy
go build ./...
```

After changing any type in `api/v1alpha1/`, regenerate the deepcopy methods and commit the result:

```bash
controller-gen object:headerFile=hack/boilerplate.go.txt paths="./..."
```

`zz_generated.deepcopy.go` is the only generated file; never edit it by hand.

You need Go 1.25.3 and `controller-gen` (sigs.k8s.io) installed. The Terraform/OpenTofu version pins (`.terraform-version`, `.opentofu-version`) and the `terraform-deploy.yml.template` workflow are template-repo leftovers; no Terraform sources exist in this repository, so `tofu fmt`, `tofu validate`, and `tflint` do not apply.

</llm>


## 📄 License

This Project is licensed under the GNU General Public License v3.0

- see the [LICENSE](LICENSE) file for details.


## :coffee: Contributing

To become a contributor, please check out the [CONTRIBUTING](CONTRIBUTING.md) file.


## :email: Contact

For any inquiries or support requests, please open an issue in this
repository or contact us at [contact@hauke.cloud](mailto:contact@hauke.cloud).
