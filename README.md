# Secure Supply Chain Go Demo

A demonstration repository showcasing secure software supply chain practices in Go, including build attestations and Software Bill of Materials (SBOM) generation using GitHub Actions.

## Features

- **Simple Go Application**: A basic "Hello World" style application written in Go
- **Automated Builds**: GitHub Actions workflow for building the application
- **SBOM Generation**: Automatic creation of Software Bill of Materials using Anchore's SBOM Action
- **Build Attestations**: Cryptographic attestations for both SBOM and build provenance using GitHub's attestation features

## Application

The application is a simple command-line tool that prints a greeting message:

```bash
# Default greeting
go run main.go
# Output: Hello, World! Welcome to the Secure Supply Chain Demo.

# Custom greeting
go run main.go "GitHub"
# Output: Hello, GitHub! Welcome to the Secure Supply Chain Demo.
```

## Building

To build the application locally:

```bash
go build -o secure-supply-chain-demo main.go
./secure-supply-chain-demo
```

## GitHub Actions Workflow

The workflow (`.github/workflows/build.yml`) performs the following steps:

1. **Checkout**: Retrieves the source code
2. **Setup Go**: Installs Go 1.21
3. **Build**: Compiles the binary
4. **Generate SBOM**: Creates a Software Bill of Materials in SPDX JSON format
5. **Attest SBOM**: Creates a cryptographic attestation for the SBOM
6. **Attest Build Provenance**: Creates a cryptographic attestation for the build artifact
7. **Upload Artifacts**: Stores the binary and SBOM as workflow artifacts

## Attestations

This demo uses GitHub's attestation features to provide:

- **SBOM Attestation**: Cryptographically signed record of all dependencies
- **Build Provenance Attestation**: Cryptographically signed record of how the artifact was built

These attestations can be verified using the GitHub CLI or API to ensure the integrity and provenance of the built artifacts.

## Requirements

- Go 1.21 or higher
- GitHub Actions (for automated builds and attestations)

## License

This is a demonstration project for educational purposes.