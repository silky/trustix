// MIT License
//
// Copyright (c) 2020 Tweag IO
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/tweag/trustix/api"
	"github.com/tweag/trustix/correlator"
	pb "github.com/tweag/trustix/proto"
	"github.com/tweag/trustix/schema"
	"sync"
)

type TrustixCombinedRPCServer struct {
	pb.UnimplementedTrustixCombinedRPCServer
	logs       *TrustixCombinedRPCServerMap
	correlator correlator.LogCorrelator
}

func NewTrustixCombinedRPCServer(logs *TrustixCombinedRPCServerMap, correlator correlator.LogCorrelator) *TrustixCombinedRPCServer {
	return &TrustixCombinedRPCServer{
		logs:       logs,
		correlator: correlator,
	}
}

func (l *TrustixCombinedRPCServer) HashMap(ctx context.Context, in *pb.HashRequest) (*pb.HashMapResponse, error) {
	responses := make(map[string][]byte)

	var wg sync.WaitGroup
	var mux sync.Mutex

	hexInput := hex.EncodeToString(in.InputHash)
	log.WithField("inputHash", hexInput).Info("Received HashMap request")

	for name, l := range l.logs.Map() {
		// Create copies for goroutine
		name := name
		l := l

		wg.Add(1)

		go func() {
			defer wg.Done()

			log.WithFields(log.Fields{
				"inputHash": hexInput,
				"logName":   name,
			}).Info("Querying log")

			sth, err := l.GetSTH(&api.STHRequest{})
			if err != nil {
				fmt.Println(fmt.Errorf("could not get STH: %v", err))
				return
			}
			resp, err := l.GetMapValue(&api.GetMapValueRequest{
				Key:     in.InputHash,
				MapRoot: sth.MapRoot,
			})
			if err != nil {
				fmt.Println("could not query: %v", err)
				return
			}

			mux.Lock()
			responses[name] = resp.Value
			mux.Unlock()
		}()
	}

	wg.Wait()

	return &pb.HashMapResponse{
		Hashes: responses,
	}, nil

}

func (l *TrustixCombinedRPCServer) Decide(ctx context.Context, in *pb.HashRequest) (*pb.DecisionResponse, error) {

	hexInput := hex.EncodeToString(in.InputHash)
	log.WithField("inputHash", hexInput).Info("Received Decide request")

	var wg sync.WaitGroup
	var mux sync.Mutex
	var inputs []*correlator.LogCorrelatorInput
	var misses []string

	for name, l := range l.logs.Map() {
		// Create copies for goroutine
		name := name
		l := l

		wg.Add(1)

		go func() {
			defer wg.Done()

			log.WithFields(log.Fields{
				"inputHash": hexInput,
				"logName":   name,
			}).Info("Querying log")

			sth, err := l.GetSTH(&api.STHRequest{})
			if err != nil {
				fmt.Println(fmt.Errorf("could not get STH: %v", err))
				return
			}
			resp, err := l.GetMapValue(&api.GetMapValueRequest{
				Key:     in.InputHash,
				MapRoot: sth.MapRoot,
			})
			if err != nil {
				fmt.Println("could not query: %v", err)
				return
			}

			mapEntry := &schema.MapEntry{}
			err = proto.Unmarshal(resp.Value, mapEntry)
			if err != nil {
				fmt.Println("Could not unmarshal value")
				return
			}

			mux.Lock()
			defer mux.Unlock()

			if err != nil {
				fmt.Println(err)
				misses = append(misses, name)
				return
			}

			if len(resp.Value) == 0 {
				misses = append(misses, name)
				return
			}

			inputs = append(inputs, &correlator.LogCorrelatorInput{
				LogName:    name,
				OutputHash: hex.EncodeToString(mapEntry.Value),
			})

		}()
	}

	wg.Wait()

	resp := &pb.DecisionResponse{
		Misses: misses,
	}

	var decision *correlator.LogCorrelatorOutput
	if len(inputs) > 0 {

		var err error
		decision, err = l.correlator.Decide(inputs)
		if err != nil {
			return nil, err
		}

		outputHash, err := hex.DecodeString(decision.OutputHash)
		if err != nil {
			return nil, err
		}

		if len(decision.LogNames) > 0 {
			resp.Decision = &pb.OutputHashDecision{
				LogNames:   decision.LogNames,
				OutputHash: outputHash,
				Confidence: int32(decision.Confidence),
			}
		}
	}

	// inputMap := make(map[string][]byte)
	for _, input := range inputs {
		if input.OutputHash == decision.OutputHash {
			continue
		}

		h, err := hex.DecodeString(input.OutputHash)
		if err != nil {
			return nil, err
		}

		resp.Mismatches = append(resp.Mismatches, &pb.OutputHashResponse{
			LogName:    input.LogName,
			OutputHash: h,
		})
	}

	return resp, nil

}