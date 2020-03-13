// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skipnumbers

import (
	"golang.org/x/exp/rand"
	"io"
	"log"
	"saientist.dev/spago/pkg/ml/ag"
	"saientist.dev/spago/pkg/ml/initializers"
	"saientist.dev/spago/pkg/ml/nn"
)

var _ nn.Model = &Model{}

// Model that makes predictions using the last state of the recurrent network.
type Model struct {
	RNN       nn.Model
	Predictor nn.Model
}

func NewModel(rnn, predictor nn.Model) *Model {
	return &Model{
		RNN:       rnn,
		Predictor: predictor,
	}
}

func (m *Model) Init() {
	src := rand.NewSource(1)
	m.RNN.ForEachParam(func(param *nn.Param) {
		if param.Type() == nn.Weights {
			// TODO: how to know the right gain for each param? Should the gain be a property of the param itself?
			initializers.XavierUniform(param.Value(), 1, src)
		}
	})
	m.Predictor.ForEachParam(func(param *nn.Param) {
		if param.Type() == nn.Weights {
			initializers.XavierUniform(param.Value(), initializers.Gain(ag.Softmax), src)
		}
	})
}

func (m *Model) ForEachParam(callback func(param *nn.Param)) {
	nn.ForEachParam(m, callback)
}

func (m *Model) Serialize(w io.Writer) (int, error) {
	return nn.Serialize(m, w)
}

func (m *Model) Deserialize(r io.Reader) (int, error) {
	return nn.Deserialize(m, r)
}

type Processor struct {
	opt       []interface{}
	model     *Model
	g         *ag.Graph
	RNN       nn.Processor
	Predictor nn.Processor
}

func (m *Model) NewProc(g *ag.Graph, opt ...interface{}) nn.Processor {
	p := &Processor{
		model:     m,
		opt:       opt,
		g:         g,
		RNN:       m.RNN.NewProc(g),
		Predictor: m.Predictor.NewProc(g),
	}
	p.init(opt)
	return p
}

func (p *Processor) init(opt []interface{}) {
	if len(opt) > 0 {
		log.Fatal("skipnumbers: invalid init options")
	}
}

func (p *Processor) Model() nn.Model       { return p.model }
func (p *Processor) Graph() *ag.Graph      { return p.g }
func (p *Processor) RequiresFullSeq() bool { return true }

func (p *Processor) Reset() {
	p.init(p.opt)
}

func (p *Processor) Forward(xs ...ag.Node) []ag.Node {
	return p.Predictor.Forward(p.RNN.Forward(xs...)[len(xs)-1])
}