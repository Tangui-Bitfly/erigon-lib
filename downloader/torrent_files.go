package downloader

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/ledgerwatch/erigon-lib/common/dir"
	"golang.org/x/exp/slices"
)

// AtomicTorrentFS - does provide thread-safe CRUD operations on .torrent files
type AtomicTorrentFS struct {
	lock sync.Mutex
	dir  string
}

func NewAtomicTorrentFS(dir string) *AtomicTorrentFS {
	return &AtomicTorrentFS{dir: dir}
}

func (tf *AtomicTorrentFS) Exists(name string) bool {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.exists(name)
}

func (tf *AtomicTorrentFS) exists(name string) bool {
	if !strings.HasSuffix(name, ".torrent") {
		name += ".torrent"
	}
	return dir.FileExist(filepath.Join(tf.dir, name))
}
func (tf *AtomicTorrentFS) Delete(name string) error {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.delete(name)
}

func (tf *AtomicTorrentFS) delete(name string) error {
	if !strings.HasSuffix(name, ".torrent") {
		name += ".torrent"
	}
	return os.Remove(filepath.Join(tf.dir, name))
}

func (tf *AtomicTorrentFS) Create(name string, res []byte) (ts *torrent.TorrentSpec, created bool, err error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()

	if !tf.exists(name) {
		err = tf.create(name, res)
		if err != nil {
			return nil, false, err
		}
	}

	ts, err = tf.load(filepath.Join(tf.dir, name))
	if err != nil {
		return nil, false, err
	}
	return ts, false, nil
}

func (tf *AtomicTorrentFS) create(name string, res []byte) error {
	if !strings.HasSuffix(name, ".torrent") {
		name += ".torrent"
	}
	if len(res) == 0 {
		return fmt.Errorf("try to write 0 bytes to file: %s", name)
	}

	fPath := filepath.Join(tf.dir, name)
	f, err := os.Create(fPath + ".tmp")
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(res); err != nil {
		return err
	}
	if err = f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(fPath+".tmp", fPath); err != nil {
		return err
	}

	return nil
}

func (tf *AtomicTorrentFS) createFromMetaInfo(fPath string, mi *metainfo.MetaInfo) error {
	file, err := os.Create(fPath + ".tmp")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := mi.Write(file); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(fPath+".tmp", fPath); err != nil {
		return err
	}
	return nil
}

func (tf *AtomicTorrentFS) CreateWithMetaInfo(info *metainfo.Info, additionalMetaInfo *metainfo.MetaInfo) (created bool, err error) {
	name := info.Name
	if !strings.HasSuffix(name, ".torrent") {
		name += ".torrent"
	}
	mi, err := CreateMetaInfo(info, additionalMetaInfo)
	if err != nil {
		return false, err
	}

	tf.lock.Lock()
	defer tf.lock.Unlock()

	if tf.exists(name) {
		return false, nil
	}
	if err = tf.createFromMetaInfo(filepath.Join(tf.dir, name), mi); err != nil {
		return false, err
	}
	return true, nil
}

func (tf *AtomicTorrentFS) LoadByName(name string) (*torrent.TorrentSpec, error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.load(filepath.Join(tf.dir, name))
}

func (tf *AtomicTorrentFS) LoadByPath(fPath string) (*torrent.TorrentSpec, error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.load(fPath)
}

func (tf *AtomicTorrentFS) load(fPath string) (*torrent.TorrentSpec, error) {
	if !strings.HasSuffix(fPath, ".torrent") {
		fPath += ".torrent"
	}
	mi, err := metainfo.LoadFromFile(fPath)
	if err != nil {
		return nil, fmt.Errorf("LoadFromFile: %w, file=%s", err, fPath)
	}
	mi.AnnounceList = Trackers
	return torrent.TorrentSpecFromMetaInfoErr(mi)
}

const ProhibitNewDownloadsFileName = "prohibit_new_downloads.lock"

// Erigon "download once" - means restart/upgrade/downgrade will not download files (and will be fast)
// After "download once" - Erigon will produce and seed new files
// Downloader will able: seed new files (already existing on FS), download uncomplete parts of existing files (if Verify found some bad parts)
func (tf *AtomicTorrentFS) ProhibitNewDownloads(whitelistAdd, whitelistRemove []string) (whitelist []string, err error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.prohibitNewDownloads(whitelistAdd, whitelistRemove)
}

func (tf *AtomicTorrentFS) prohibitNewDownloads(whitelistAdd, whitelistRemove []string) (whitelist []string, err error) {
	fPath := filepath.Join(tf.dir, ProhibitNewDownloadsFileName)
	exist := dir.FileExist(fPath)

	var _currentWhiteList []string
	if exist {
		torrentListJsonBytes, err := os.ReadFile(fPath)
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}
		if len(torrentListJsonBytes) > 0 {
			if err := json.Unmarshal(torrentListJsonBytes, &_currentWhiteList); err != nil {
				return nil, fmt.Errorf("unmarshal: %w", err)
			}
		}
	}

	whiteList := make([]string, 0, len(_currentWhiteList))
	for _, it := range _currentWhiteList {
		if slices.Contains(whitelistRemove, it) {
			continue
		}
		whiteList = append(whiteList, it)
	}

	for _, it := range whitelistAdd {
		if slices.Contains(whiteList, it) {
			whiteList = append(whiteList, it)
			continue
		}
	}
	slices.Sort(whiteList)

	whiteListBytes, err := json.Marshal(whiteList)
	if err != nil {
		return _currentWhiteList, fmt.Errorf("marshal: %w", err)
	}
	if err := dir.WriteFileWithFsync(fPath, whiteListBytes, 0644); err != nil {
		return _currentWhiteList, fmt.Errorf("write: %w", err)
	}
	return whiteList, nil
}

func (tf *AtomicTorrentFS) NewDownloadsAreProhibited(name string) (prohibited bool, err error) {
	tf.lock.Lock()
	defer tf.lock.Unlock()
	return tf.newDownloadsAreProhibited(name)
}

func (tf *AtomicTorrentFS) newDownloadsAreProhibited(name string) (prohibited bool, err error) {
	fPath := filepath.Join(tf.dir, ProhibitNewDownloadsFileName)
	exists := dir.FileExist(fPath)
	if !exists { // no .lock - means all allowed
		return false, nil
	}

	var whiteList []string
	whiteListBytes, err := os.ReadFile(fPath)
	if err != nil {
		return false, fmt.Errorf("NewDownloadsAreProhibited: read file: %w", err)
	}
	if len(whiteListBytes) > 0 {
		if err := json.Unmarshal(whiteListBytes, &whiteList); err != nil {
			return false, fmt.Errorf("NewDownloadsAreProhibited: unmarshal: %w", err)
		}
	}

	for _, whiteListedItem := range whiteList {
		if strings.Contains(name, whiteListedItem) {
			return false, nil
		}
	}
	return true, nil
}
