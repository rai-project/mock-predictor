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
	"github.com/rai-project/tracer"
)

type Predictor struct {
	Framework dlframework.FrameworkManifest
	Model     dlframework.ModelManifest
	options   *options.Options
}

// New
// creates a new predictor given a model and a set of options
func New(model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	framework, err := model.ResolveFramework()
	if err != nil {
		return nil, err
	}

	return &Predictor{
		Framework: framework,
		Model:     model,
		options:   options.New(opts...),
	}, nil
}

// Download
// downloads models: there are none really
func (p *Predictor) Download(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) error {
	return nil
}

// Loads a model into memory
// loads a new predictor given a model and a set of options
func (p *Predictor) Load(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	framework, err := model.ResolveFramework()
	if err != nil {
		return nil, err
	}

	return &Predictor{
		Framework: framework,
		Model:     model,
		options:   options.New(opts...),
	}, nil
}

// GetPreprocessOptions
// get the processing options needed to run the model
func (p *Predictor) GetPreprocessOptions(ctx context.Context) (common.PreprocessOptions, error) {
	return common.PreprocessOptions{
		Context: ctx,
	}, nil
}

// Predict ...
// perform the prediction. since this a mock framework it does not need to perform anything
func (p *Predictor) Predict(ctx context.Context, data [][]float32, opts ...options.Option) error {
	return nil
}

// gets the options used for the predictor
func (p Predictor) GetPredictionOptions(ctx context.Context) (*options.Options, error) {
	return p.options, nil
}

// get information about the predictor .. i.e. framework and model
func (p Predictor) Info() (dlframework.FrameworkManifest, dlframework.ModelManifest, error) {
	return p.Framework, p.Model, nil
}

// ReadPredictedFeatures
// reads the predicted features. for this mock framework the prediction results in mock
// as the label with 100% probability
func (p *Predictor) ReadPredictedFeatures(ctx context.Context) ([]dlframework.Features, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, "read_predicted_features")
	defer span.Finish()

	batchSize := int(p.options.BatchSize())
	output := make([]dlframework.Features, batchSize)

	for ii := 0; ii < batchSize; ii++ {
		output[ii] = []*dlframework.Feature{
			feature.New(
				feature.ClassificationIndex(0),
				feature.ClassificationLabel("mock"),
				feature.Probability(1),
			),
		}
	}
	return output, nil
}

// Reset
// resets the predictor state
func (p *Predictor) Reset(ctx context.Context) error {
	return nil
}

// Close
// cloeses the predictor
func (p *Predictor) Close() error {
	return nil
}

// register the predictor as an agent
func init() {
	config.AfterInit(func() {
		framework := mock.FrameworkManifest
		agent.AddPredictor(framework, &Predictor{
			Framework: framework,
		})
	})
}
