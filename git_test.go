package git

import (
	"context"
	"encoding/json"
	"github.com/sfomuseum/go-lookup/catalog"
	"github.com/sfomuseum/go-lookup/iterator"
	"io"
	"testing"
)

type properties struct {
	Id   int64  `json:"wof:id"`
	Name string `json:"wof:name"`
}

type feature struct {
	Properties properties `json:"properties"`
}

func AppendFunc(ctx context.Context, c catalog.Catalog, fh io.ReadCloser) error {

	var f *feature

	dec := json.NewDecoder(fh)
	err := dec.Decode(&f)

	if err != nil {
		return err
	}

	props := f.Properties
	c.LoadOrStore(props.Name, props.Id)

	return nil
}

func TestGitLookerUpper(t *testing.T) {

	ctx := context.Background()

	lu, err := iterator.NewIterator(ctx, "https://github.com/sfomuseum-data/sfomuseum-data-maps.git")

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

	key := "SFO (2014)"
	expected := int64(1477881751)

	v, ok := c.Load(key)

	if !ok {
		t.Fatalf("Failed to locate key '%s'", key)
	}

	if v.(int64) != expected {
		t.Fatalf("Unexpected value. Wanted '%d' but got '%d'", expected, v.(int64))
	}

}
