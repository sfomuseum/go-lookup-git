package git

import (
	"context"
	"testing"
	"github.com/sfomuseum/go-lookup"	
	"github.com/sfomuseum/go-lookup/catalog"
	"io"
)

func AppendFunc(ctx context.Context, c lookup.Catalog, fh io.ReadCloser) error {
	return nil
}

func TestGitLookerUpper(t *testing.T) {

	ctx := context.Background()
	
	lu, err := NewGitLookerUpper(ctx, "https://github.com/sfomuseum-data/sfomuseum-data-maps.git")

	if err != nil {
		t.Fatalf("Failed to create new GitLookerUpper, %v", err)
	}

	c, err := catalog.NewSyncMapCatalog(ctx, "syncmap://")

	if err != nil {
		t.Fatalf("Failed to create new SyncMapCatalog, %v", err)
	}
	
	err = lu.Append(ctx, c, AppendFunc)

	if err != nil {
		t.Fatalf("Failed to append lookup to catalog, %v", err)
	}
}
