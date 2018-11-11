package predict

import (
	context "context"

	"github.com/rai-project/config"
	"github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework/agent"
	"github.com/rai-project/dlframework/framework/feature"
	"github.com/rai-project/dlframework/framework/options"
	common "github.com/rai-project/dlframework/framework/predict"
	mock "github.com/rai-project/mock-predictor"
)

type Predictor struct {
}

// New ...
func New(model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	return &Predictor{}
}

// Download ...
func (p *Predictor) Download(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) error {
	return nil
}

// Load ...
func (p *Predictor) Load(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	return &Predictor{}, nil
}

// GetPreprocessOptions ...
func (p *Predictor) GetPreprocessOptions(ctx context.Context) (common.PreprocessOptions, error) {

	return common.PreprocessOptions{
		Context: ctx,
	}, nil
}

// Predict ...
func (p *Predictor) Predict(ctx context.Context, data [][]float32, opts ...options.Option) error {
	return nil
}

// ReadPredictedFeatures ...
func (p *Predictor) ReadPredictedFeatures(ctx context.Context) ([]dlframework.Features, error) {

	batchSize := int(p.BatchSize())
	length := len(predictions) / batchSize
	output := make([]dlframework.Features, batchSize)

	for ii := 0; ii < batchSize; ii++ {
		rprobs := make([]*dlframework.Feature, length)
		for jj := 0; jj < length; jj++ {
			rprobs[jj] = feature.New(
				feature.ClassificationIndex(int32(jj)),
				feature.ClassificationName(p.features[jj]),
				feature.Probability(predictions[ii*length+jj].Probability),
			)
		}
		output[ii] = rprobs
	}
	return output, nil
}

// Reset ...
func (p *Predictor) Reset(ctx context.Context) error {
	return nil
}

// Close ...
func (p *Predictor) Close() error {
	return nil
}

func init() {
	config.AfterInit(func() {
		framework := mock.FrameworkManifest
		agent.AddPredictor(framework, &Predictor{
			Predictor: common.Predictor{
				Base: common.Base{
					Framework: framework,
				},
			},
		})
	})
}
