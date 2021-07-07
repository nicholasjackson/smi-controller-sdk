# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace


# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# download kubebuilder and generate
RUN curl -L https://github.com/kubernetes-sigs/kubebuilder/releases/download/v3.1.0/kubebuilder_linux_${TARGETARCH} \
     -o /tmp/kubebuilder && \
    mv /tmp/kubebuilder /usr/local/bin/kubebuilder && \
    make generate

# Build
RUN CGO_ENABLED=0 GO111MODULE=on go build -a -o smi-controller /workspace

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/smi-controller .
USER nonroot:nonroot

ENTRYPOINT ["/smi-controller"]
