//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package topology_test

import (
	"path/filepath"
	"testing"

	"github.com/Zzorz/ghw/pkg/option"
	"github.com/Zzorz/ghw/pkg/topology"

	"github.com/Zzorz/ghw/testdata"
)

// nolint: gocyclo
func TestTopologyNUMADistances(t *testing.T) {
	testdataPath, err := testdata.SnapshotsDirectory()
	if err != nil {
		t.Fatalf("Expected nil err, but got %v", err)
	}

	multiNumaSnapshot := filepath.Join(testdataPath, "linux-amd64-intel-xeon-L5640.tar.gz")
	// from now on we use constants reflecting the content of the snapshot we requested,
	// which we reviewed beforehand. IOW, you need to know the content of the
	// snapshot to fully understand this test. Inspect it using
	// GHW_SNAPSHOT_PATH="/path/to/linux-amd64-intel-xeon-L5640.tar.gz" ghwc topology

	info, err := topology.New(option.WithSnapshot(option.SnapshotOptions{
		Path: multiNumaSnapshot,
	}))

	if err != nil {
		t.Fatalf("Expected nil err, but got %v", err)
	}
	if info == nil {
		t.Fatalf("Expected non-nil TopologyInfo, but got nil")
	}

	if len(info.Nodes) != 2 {
		t.Fatalf("Expected 2 nodes but got 0.")
	}

	for _, n := range info.Nodes {
		if len(n.Distances) != len(info.Nodes) {
			t.Fatalf("Expected distances to all known nodes")
		}
	}

	if info.Nodes[0].Distances[0] != info.Nodes[1].Distances[1] {
		t.Fatalf("Expected symmetric distance to self, got %v and %v", info.Nodes[0].Distances, info.Nodes[1].Distances)
	}

	if info.Nodes[0].Distances[1] != info.Nodes[1].Distances[0] {
		t.Fatalf("Expected symmetric distance to the other node, got %v and %v", info.Nodes[0].Distances, info.Nodes[1].Distances)
	}
}
