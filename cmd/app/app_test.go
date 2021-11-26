package app

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Integration(t *testing.T) {
	// ----- Setup ----- //
	var stdin bytes.Buffer
	var stdout bytes.Buffer

	// ----- Test cases ----- //
	type args struct {
		operations []byte
	}

	type want struct {
		results []string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Create account successfully",
			args: args{
				operations: []byte(`{"account": {"active-card": false, "available-limit": 750}}`),
			},
			want: want{
				results: []string{`{"account":{"active-card":false,"available-limit":750},"violations":[]}`},
			},
		},
		{
			name: "Creating an account that violates the Authorizer logic",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 175}}
									{"account": {"active-card": true, "available-limit": 350}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":175},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":175},"violations":["account-already-initialized"]}`,
				},
			},
		},
		{
			name: "Processing a transaction successfully",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 100}}
									{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T11:00:00.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":100},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":80},"violations":[]}`,
				},
			},
		},
		{
			name: "Processing a transaction which violates the account-not-initialized logic",
			args: args{
				operations: []byte(`{"transaction": {"merchant": "Uber Eats", "amount": 25, "time": "2020-12-01T11:07:00.000Z"}}
									{"account": {"active-card": true, "available-limit": 225}}
									{"transaction": {"merchant": "Uber Eats", "amount": 25, "time": "2020-12-01T11:07:00.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":null,"violations":["account-not-initialized"]}`,
					`{"account":{"active-card":true,"available-limit":225},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":200},"violations":[]}`,
				},
			},
		},
		{
			name: "Processing a transaction which violates card-not-active logic",
			args: args{
				operations: []byte(`{"account": {"active-card": false, "available-limit": 100}}
									{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T11:00:00.000Z"}}
									{"transaction": {"merchant": "Habbib's", "amount": 15, "time": "2019-02-13T11:15:00.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":false,"available-limit":100},"violations":[]}`,
					`{"account":{"active-card":false,"available-limit":100},"violations":["card-not-active"]}`,
					`{"account":{"active-card":false,"available-limit":100},"violations":["card-not-active"]}`,
				},
			},
		},
		{
			name: "Processing a transaction which violates insufficient-limit logic:\n",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 1000}}
									{"transaction": {"merchant": "Vivara", "amount": 1250, "time": "2019-02-13T11:00:00.000Z"}}
									{"transaction": {"merchant": "Samsung", "amount": 2500, "time": "2019-02-13T11:00:01.000Z"}}
									{"transaction": {"merchant": "Nike", "amount": 800, "time": "2019-02-13T11:01:01.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":1000},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":1000},"violations":["insufficient-limit"]}`,
					`{"account":{"active-card":true,"available-limit":1000},"violations":["insufficient-limit"]}`,
					`{"account":{"active-card":true,"available-limit":200},"violations":[]}`,
				},
			},
		},
		{
			name: "Processing a transaction which violates the high-frequency-small-interval logic",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 100}}
									{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T11:00:00.000Z"}}
 									{"transaction": {"merchant": "Habbib's", "amount": 20, "time": "2019-02-13T11:00:01.000Z"}}
 									{"transaction": {"merchant": "McDonald's", "amount": 20, "time": "2019-02-13T11:01:01.000Z"}}
									{"transaction": {"merchant": "Subway", "amount": 20, "time": "2019-02-13T11:01:31.000Z"}}
									{"transaction": {"merchant": "Burger King", "amount": 10, "time": "2019-02-13T12:00:00.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":100},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":80},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":60},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":40},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":40},"violations":["high-frequency-small-interval"]}`,
					`{"account":{"active-card":true,"available-limit":30},"violations":[]}`,
				},
			},
		},
		{
			name: "Processing a transaction which violates the doubled-transaction logic",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 100}}
 									{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T11:00:00.000Z"}}
									{"transaction": {"merchant": "McDonald's", "amount": 10, "time": "2019-02-13T11:00:01.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T11:00:02.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 15, "time": "2019-02-13T11:00:03.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":100},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":80},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":70},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":70},"violations":["doubled-transaction"]}`,
					`{"account":{"active-card":true,"available-limit":55},"violations":[]}`,
				},
			},
		},
		{
			name: "Processing transactions that violate multiple logics",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 100}}
 									{"transaction": {"merchant": "McDonald's", "amount": 10, "time": "2019-02-13T11:00:01.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T11:00:02.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 5, "time": "2019-02-13T11:00:07.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 5, "time": "2019-02-13T11:00:08.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 150, "time": "2019-02-13T11:00:18.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 190, "time": "2019-02-13T11:00:22.000Z"}}
 									{"transaction": {"merchant": "Burger King", "amount": 15, "time": "2019-02-13T12:00:27.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":100},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":90},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":70},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":65},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":65},"violations":["high-frequency-small-interval","doubled-transaction"]}`,
					`{"account":{"active-card":true,"available-limit":65},"violations":["insufficient-limit","high-frequency-small-interval"]}`,
					`{"account":{"active-card":true,"available-limit":65},"violations":["insufficient-limit","high-frequency-small-interval"]}`,
					`{"account":{"active-card":true,"available-limit":50},"violations":[]}`,
				},
			},
		},
		{
			name: "Violations should not be saved in the application's internal state",
			args: args{
				operations: []byte(`{"account": {"active-card": true, "available-limit": 1000}}
 									{"transaction": {"merchant": "Vivara", "amount": 1250, "time": "2019-02-13T11:00:00.000Z"}}
 									{"transaction": {"merchant": "Samsung", "amount": 2500, "time": "2019-02-13T11:00:01.000Z"}}
									{"transaction": {"merchant": "Nike", "amount": 800, "time": "2019-02-13T11:01:01.000Z"}}
 									{"transaction": {"merchant": "Uber", "amount": 80, "time": "2019-02-13T11:01:31.000Z"}}`),
			},
			want: want{
				results: []string{
					`{"account":{"active-card":true,"available-limit":1000},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":1000},"violations":["insufficient-limit"]}`,
					`{"account":{"active-card":true,"available-limit":1000},"violations":["insufficient-limit"]}`,
					`{"account":{"active-card":true,"available-limit":200},"violations":[]}`,
					`{"account":{"active-card":true,"available-limit":120},"violations":[]}`,
				},
			},
		},
	}

	// ----- Test runner ----- //
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			scanner := bufio.NewScanner(&stdout)
			defer stdin.Reset()
			defer stdout.Reset()

			stdin.Write(tt.args.operations)

			// Execute
			Start(&stdin, &stdout)

			// Verify
			index := 0
			for scanner.Scan() {
				assert.Equal(t, tt.want.results[index], scanner.Text())
				index += 1
			}
		})
	}
}
