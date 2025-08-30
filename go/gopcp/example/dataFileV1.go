package example

import (
	"os"
	"sync"
	"errors"
	"io"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

type myDataFile struct {
	f *os.File
	fmutex sync.RWMutex
	rcond *sync.Cond
	woffset int64
	roffset int64
	wmutex sync.Mutex
	rmutex sync.Mutex
	dataLen uint32
}

func NewDataFile(path string,dataLen uint32) (DataFile, error){
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile)Read()(rsn int64,d Data,err error){
	var offset int64
	df.rmutex.Lock()
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()
	rsn = offset/ int64(df.dataLen)
	bytes := make([]byte,df.dataLen)
	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.rcond.Wait()
				continue
			}
			df.fmutex.RUnlock()
			return  
		}
		d = bytes
		df.fmutex.RUnlock()
		return
	} 
}

func (df *myDataFile)Write(d Data) (wsn int64, err error ){
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	df.rcond.Signal()
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}