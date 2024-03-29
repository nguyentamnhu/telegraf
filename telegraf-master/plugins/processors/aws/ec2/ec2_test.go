package ec2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/testutil"
)

func TestBasicStartup(t *testing.T) {
	p := newAwsEc2Processor()
	p.Log = &testutil.Logger{}
	p.ImdsTags = []string{"accountId", "instanceId"}
	acc := &testutil.Accumulator{}
	require.NoError(t, p.Init())

	require.Empty(t, acc.GetTelegrafMetrics())
	require.Empty(t, acc.Errors)
}

func TestBasicStartupWithEC2Tags(t *testing.T) {
	p := newAwsEc2Processor()
	p.Log = &testutil.Logger{}
	p.ImdsTags = []string{"accountId", "instanceId"}
	p.EC2Tags = []string{"Name"}
	acc := &testutil.Accumulator{}
	require.NoError(t, p.Init())

	require.Empty(t, acc.GetTelegrafMetrics())
	require.Empty(t, acc.Errors)
}

func TestBasicStartupWithCacheTTL(t *testing.T) {
	p := newAwsEc2Processor()
	p.Log = &testutil.Logger{}
	p.ImdsTags = []string{"accountId", "instanceId"}
	p.CacheTTL = config.Duration(12 * time.Hour)
	acc := &testutil.Accumulator{}
	require.NoError(t, p.Init())

	require.Empty(t, acc.GetTelegrafMetrics())
	require.Empty(t, acc.Errors)
}

func TestBasicStartupWithTagCacheSize(t *testing.T) {
	p := newAwsEc2Processor()
	p.Log = &testutil.Logger{}
	p.ImdsTags = []string{"accountId", "instanceId"}
	p.TagCacheSize = 100
	acc := &testutil.Accumulator{}
	require.NoError(t, p.Init())

	require.Empty(t, acc.GetTelegrafMetrics())
	require.Empty(t, acc.Errors)
}

func TestBasicInitNoTagsReturnAnError(t *testing.T) {
	p := newAwsEc2Processor()
	p.Log = &testutil.Logger{}
	p.ImdsTags = []string{}
	err := p.Init()
	require.Error(t, err)
}

func TestBasicInitInvalidTagsReturnAnError(t *testing.T) {
	p := newAwsEc2Processor()
	p.Log = &testutil.Logger{}
	p.ImdsTags = []string{"dummy", "qwerty"}
	err := p.Init()
	require.Error(t, err)
}
