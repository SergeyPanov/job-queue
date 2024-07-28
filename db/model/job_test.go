package model

import (
	"log"
	"testing"
	"time"
)

func TestJob_Execute(t *testing.T) {
	type fields struct {
		startAt            time.Time
		executionTime      time.Duration
		SuccessProbability float64
		Results            chan Result
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				startAt:            time.Now(),
				executionTime:      time.Second,
				SuccessProbability: 0.0,
				Results:            make(chan Result),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				StartAt:            tt.fields.startAt,
				ExecutionTime:      tt.fields.executionTime,
				SuccessProbability: tt.fields.SuccessProbability,
				Results:            tt.fields.Results,
			}
			go j.Execute()

			res := <-j.Results

			if res.Err != nil {
				panic("fuck")
			}
		})
	}
}

func TestJob_Schedule(t *testing.T) {
	type fields struct {
		startAt            time.Time
		executionTime      time.Duration
		SuccessProbability float64
		Results            chan Result
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				startAt:            time.Now().Add(5 * time.Second),
				executionTime:      time.Second,
				SuccessProbability: 0.0,
				Results:            make(chan Result),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				StartAt:            tt.fields.startAt,
				ExecutionTime:      tt.fields.executionTime,
				SuccessProbability: tt.fields.SuccessProbability,
				Results:            tt.fields.Results,
			}
			go j.Schedule()

			res := <-j.Results

			if res.Err != nil {
				log.Fatalf(res.Err.Error())
			}
		})
	}
}
