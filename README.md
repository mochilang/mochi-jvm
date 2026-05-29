# mochi-jvm

Mochi+Java bidirectional package bridge ([MEP-67](https://mochilang.org/docs/mep/mep-0067)).

[![CI](https://github.com/mochilang/mochi-jvm/actions/workflows/ci.yml/badge.svg)](https://github.com/mochilang/mochi-jvm/actions/workflows/ci.yml)

## Two directions

**Consume**: import Java libraries into Mochi.

```mochi
import java "com.google.guava:guava@33.4.8-jre" as guava
```

**Publish**: ship Mochi packages as Maven artifacts.

```
mochi pkg publish --to=maven-central
```

## Sub-packages

| Package | Description |
|---------|-------------|
| `errors` | Cross-cutting error types and skip-reason enum |
| `maven` | Maven Central HTTP client, coordinate parsing, JAR cache |
| `reflect` | Java reflection tool (runs a bundled JAR to extract class surfaces) |
| `typemap` | Java-to-Mochi type mapping table |
| `wrapper` | JNI wrapper class synthesiser and Java source emitter |
| `emit` | Mochi extern shim emitter and skip-report renderer |
| `jni` | JNI embedding (CGO, gated by build tag `java_jni`) |
| `build` | End-to-end build orchestration: Driver, Pipeline, javac, POM |
| `lock` | `mochi.lock` TOML integration for `[[java-package]]` tables |
| `publish` | Maven Central publish flow and Sigstore attestation |

## Requirements

- Go 1.22+
- JDK 17+ for the `jni` CGO path (gated by `-tags java_jni`; tests run without JDK)

## JNI embedding

The `jni` package requires CGO and JDK headers. Enable with:

```bash
go build -tags java_jni ./...
```

Without the tag, `jni.NewJVM` returns `jni.ErrJNINotAvailable` so CI runs clean without a JDK.

## Spec reference

See the [MEP-67 spec](https://mochilang.org/docs/mep/mep-0067) and [research bundle](https://mochilang.org/docs/research/0067/) for the full design.
