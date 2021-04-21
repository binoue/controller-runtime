module github.com/binoue/controller-runtime/examples/scratch-env

go 1.15

require (
	github.com/spf13/pflag v1.0.5
	k8s.io/client-go v0.19.2
	github.com/binoue/controller-runtime v0.0.0-00010101000000-000000000000
)

replace github.com/binoue/controller-runtime => ../..
