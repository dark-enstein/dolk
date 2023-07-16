package awspile

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/dark-enstein/dolk/internal"
	"gopkg.in/yaml.v3"
)

const (
	SageMaker = "sagemaker"
)

var (
	ErrUnmarshalingYaml = fmt.Sprintf("unable to unmarshal %v options")
)

type SageMakerOpts struct {
	Opts NoteBookOpts `yaml:"notebook-opt"`
}

type NoteBookOpts struct {
	public          bool     `yaml:"public"`
	sourceRepo      RepoCred `yaml:"source-repo"`
	sourceFs        string   `yaml:"source-fs"`
	volumeInfo      string   `yaml:"volume_info"`
	acceleratorType string   `yaml:"accelerator_type"`
}

func (s *SageMakerOpts) String() string {
	return fmt.Sprintf("%v", *s)
}

type SageMakerBuilder struct {
	UUID     string
	Provider string
	Version  string
	Name     string
	Tags     []string
	Options  string
	Stack    *internal.ContextStack
}

type RepoCred struct {
	url   string `yaml:"url"`
	token string `yaml:"token"`
}

func (b *SageMakerBuilder) Deploy() ([]byte, error) {
	trace, _ := b.Stack.LogInit()
	client := b.NewClient()
	trace.Info().Msgf("created %v client: %v", SageMaker, client)
	return []byte("SageMakerDeployed"), nil
}

func (b *SageMakerBuilder) NewClient() *sagemaker.Client {
	trace, _ := b.Stack.LogInit()
	//createSession()
	trace.Info().Msgf("starting %v session", SageMaker)
	options := b._setupClientOptions()
	return sagemaker.New(*options, nil)
}

func (b *SageMakerBuilder) _setupClientOptions() *sagemaker.Options {
	return &sagemaker.Options{
		APIOptions:               nil,
		AppID:                    "",
		ClientLogMode:            0,
		Credentials:              nil,
		DefaultsMode:             "",
		EndpointOptions:          sagemaker.EndpointResolverOptions{},
		EndpointResolver:         nil,
		HTTPSignerV4:             nil,
		IdempotencyTokenProvider: nil,
		Logger:                   nil,
		Region:                   "",
		RetryMaxAttempts:         0,
		RetryMode:                "",
		Retryer:                  nil,
		RuntimeEnvironment:       aws.RuntimeEnvironment{},
		HTTPClient:               nil,
	}
}

func (b *SageMakerBuilder) extractOpts() (name, version, uuid string, tags []string,
	options *SageMakerOpts) {
	return b.Name, b.Version, b.UUID, b.Tags, b.unfurlYamlOpts(b.Options)
}

func (b *SageMakerBuilder) unfurlYamlOpts(src string) *SageMakerOpts {
	trace, log := b.Stack.LogInit()
	opts := &SageMakerOpts{}
	err := yaml.Unmarshal([]byte(src), opts)
	if err != nil {
		log.Error().Msgf(ErrUnmarshalingYaml, SageMaker)
		return nil
	}
	trace.Info().Msgf("initalized %v options", SageMaker)
	trace.Trace().Msg(opts.String())

	return opts
}
