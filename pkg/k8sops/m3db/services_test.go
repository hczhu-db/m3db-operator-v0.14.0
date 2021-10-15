// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package m3db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	const svcName = "m3dbnode-m3db-cluster"
	fixture := getFixture("testM3DBCluster.yaml", t)
	svcCfg, err := GenerateM3DBService(fixture)
	require.NoError(t, err)
	k := newFakeK8sops(t)
	ctx := context.Background()
	svc, err := k.GetService(ctx, fixture, svcName)
	require.Nil(t, svc)
	require.NotNil(t, err)
	err = k.EnsureService(ctx, fixture, svcCfg)
	require.Nil(t, err)
	svc, err = k.GetService(ctx, fixture, svcName)
	require.NotNil(t, svc)
	require.Nil(t, err)
	err = k.EnsureService(ctx, fixture, svcCfg)
	require.Nil(t, err)
	err = k.DeleteService(ctx, fixture, svcName)
	require.Nil(t, err)
	err = k.DeleteService(ctx, fixture, svcName)
	require.NotNil(t, err)
}
