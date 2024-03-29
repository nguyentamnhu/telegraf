package reverse_dns

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/influxdata/telegraf/metric"
	"github.com/influxdata/telegraf/testutil"
)

func TestSimpleReverseLookupIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	now := time.Now()
	m := metric.New("name", map[string]string{
		"dest_ip": "1.1.1.1",
	}, map[string]interface{}{
		"source_ip": "127.0.0.1",
	}, now)

	dns := newReverseDNS()
	dns.Log = &testutil.Logger{}
	dns.Lookups = []lookupEntry{
		{
			Field: "source_ip",
			Dest:  "source_name",
		},
		{
			Tag:  "dest_ip",
			Dest: "dest_name",
		},
	}
	acc := &testutil.Accumulator{}
	err := dns.Start(acc)
	require.NoError(t, err)
	err = dns.Add(m, acc)
	require.NoError(t, err)
	dns.Stop()
	// should be processed now.

	require.Len(t, acc.GetTelegrafMetrics(), 1)
	processedMetric := acc.GetTelegrafMetrics()[0]
	_, ok := processedMetric.GetField("source_name")
	require.True(t, ok)
	tag, ok := processedMetric.GetTag("dest_name")
	require.True(t, ok)
	require.EqualValues(t, "one.one.one.one.", tag)
}
