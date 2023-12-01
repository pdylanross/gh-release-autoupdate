package localfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

type LocalFileManager struct {
	basePath string
	appName  string
	tmp      string
}

func NewLocalFileManager(appName string) (*LocalFileManager, error) {
	basePath, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	basePath = path.Join(basePath, appName)
	return &LocalFileManager{basePath: basePath, appName: appName}, nil
}

func (lfm *LocalFileManager) LoadJSON(filename string, obj any) (bool, error) {
	fullpath := path.Join(lfm.basePath, filename)
	_, err := os.Stat(fullpath)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	contents, err := os.ReadFile(fullpath)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(contents, obj)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (lfm *LocalFileManager) SaveJSON(filename string, obj any) error {
	if err := lfm.EnsureBasePathExists(); err != nil {
		return err
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	fullpath := path.Join(lfm.basePath, filename)
	return os.WriteFile(fullpath, bytes, 0660)
}

func (lfm *LocalFileManager) EnsureBasePathExists() error {
	_, err := os.Stat(lfm.basePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(lfm.basePath, 0760)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (lfm *LocalFileManager) GetTemp() (string, error) {
	if lfm.tmp != "" {
		return lfm.tmp, nil
	}

	tmp, err := os.MkdirTemp("", fmt.Sprintf("%s-*", lfm.appName))
	if err != nil {
		return "", err
	}

	lfm.tmp = tmp
	return lfm.tmp, nil
}
