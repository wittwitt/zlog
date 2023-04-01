package zlog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

type dayWriter struct {
	dir        string
	filePrefix string
	f          *os.File
	sync.RWMutex
}

func newDayWriter(dir string, filePrefix string) (*dayWriter, error) {
	ld := &dayWriter{
		dir:        dir,
		filePrefix: filePrefix,
	}
	if err := ld.makeDateFile(time.Now()); err != nil {
		return nil, err
	}
	go ld.run()
	return ld, nil
}

func (p *dayWriter) Write(b []byte) (n int, err error) {
	p.RLock()
	defer p.RLocker().Unlock()
	return p.f.Write(b)
}

func (p *dayWriter) Sync() error {
	p.RLock()
	defer p.RLocker().Unlock()
	return p.f.Sync()
}

func (p *dayWriter) run() {
	logCron := cron.New()
	logCron.AddFunc("10 0 * * *", func() {
		now := time.Now()
		if err := p.makeDateFile(now); err != nil {
			log.Printf("makeDateFile: %v\n", err)
		}
	})
	logCron.Start()
}

func (p *dayWriter) makeDateFile(now time.Time) (err error) {
	defer func() {
		if rev := recover(); rev != nil {
			err = fmt.Errorf("recover: %v", rev)
		}
	}()

	if _, stat_err := os.Stat(p.dir); stat_err != nil {
		if errors.Is(stat_err, os.ErrNotExist) {
			if mk_err := os.MkdirAll(p.dir, os.ModePerm); mk_err != nil {
				return fmt.Errorf("make day dir: %v", mk_err)
			}
		} else {
			return fmt.Errorf("day dir stat: %v", stat_err)
		}
	}
	p.dir, err = filepath.Abs(p.dir)
	if err != nil {
		return err
	}

	fileName := filepath.Join(p.dir, p.filePrefix+"_"+now.Format("20060102")+".log")
	oldMask := syscall.Umask(0)
	newLogFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0640)
	syscall.Umask(oldMask)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %v", err)
	}

	oldLogFile := p.f
	p.Lock()
	p.f = newLogFile
	p.Unlock()

	if oldLogFile != nil {
		if err := oldLogFile.Sync(); err != nil {
			return fmt.Errorf("oldLogFile.Sync: %v", err)
		}
		if err := oldLogFile.Close(); err != nil {
			return fmt.Errorf("oldLogFile.Close: %v", err)
		}
	}
	return nil
}
