package cache

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/moutend/clickupctl/internal/cache/sqllite"
)

//go:embed schema.sql
var ddl string

type FileCache struct {
	queries *sqllite.Queries
}

func (f *FileCache) Load(ctx context.Context, req *http.Request, now time.Time) (*http.Response, error) {
	if req.Method != http.MethodGet {
		return nil, nil
	}

	response, err := f.queries.GetResponse(ctx, req.URL.Path)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if now.Unix() > response.ExpiredAt {
		return nil, nil
	}

	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(response.Body)),
	}

	return res, nil
}

func (f *FileCache) Save(ctx context.Context, req *http.Request, res *http.Response) error {
	if req.Method != http.MethodGet || res.StatusCode != http.StatusOK {
		return nil
	}

	defer res.Body.Close()

	buffer := &bytes.Buffer{}

	body, err := io.ReadAll(io.TeeReader(res.Body, buffer))

	if err != nil {
		return err
	}

	res.Body = io.NopCloser(buffer)

	_, err = f.queries.GetResponse(ctx, req.URL.Path)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	now := time.Now().UTC()
	cachedAt := now.Unix()
	expiredAt := now.Add(3 * time.Minute).Unix()

	if err == sql.ErrNoRows {
		_, err = f.queries.CreateResponse(ctx, sqllite.CreateResponseParams{
			Path:      req.URL.Path,
			Body:      string(body),
			CachedAt:  cachedAt,
			ExpiredAt: expiredAt,
		})
	} else {
		_, err = f.queries.UpdateResponse(ctx, sqllite.UpdateResponseParams{
			Path:      req.URL.Path,
			Body:      string(body),
			CachedAt:  cachedAt,
			ExpiredAt: expiredAt,
		})
	}

	return err
}

func NewFileCache(ctx context.Context) (*FileCache, error) {
	homeDirPath, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	clickupctlPath := filepath.Join(homeDirPath, ".clickupctl")

	db, err := sql.Open("sqlite3", "file://"+clickupctlPath)

	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(clickupctlPath); os.IsNotExist(err) {
		if _, err := db.ExecContext(ctx, ddl); err != nil {
			return nil, err
		}
	}

	cache := &FileCache{
		queries: sqllite.New(db),
	}

	return cache, nil
}
