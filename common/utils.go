package common

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// IDoc2VecWord is the interface used to construct a word for the Doc2Vec algorithm. Anything
// meeting this interface can supply data to Doc2Vec
type IDoc2VecWord interface {
	Make() string
}

//======================================================================

type IStringProvider interface {
	More() bool
	Next() string
}

type ArrayStringProvider struct {
	Data []string
	Cur  int
}

func NewArrayStringProvider(data []string) *ArrayStringProvider {
	res := &ArrayStringProvider{
		Data: data,
	}
	var _ IStringProvider = res
	return res
}

func (a *ArrayStringProvider) More() bool {
	return a.Cur < len(a.Data)
}

func (a *ArrayStringProvider) inc() {
	a.Cur++
}

func (a *ArrayStringProvider) Next() string {
	defer a.inc()
	return a.Data[a.Cur]
}

//======================================================================

type ArrayDoc2VecStringProvider struct {
	Data []IDoc2VecWord
	Cur  int
}

func NewDoc2VecArrayStringProvider(data []IDoc2VecWord) *ArrayDoc2VecStringProvider {
	res := &ArrayDoc2VecStringProvider{
		Data: data,
	}
	var _ IStringProvider = res
	return res
}

func (a *ArrayDoc2VecStringProvider) More() bool {
	return a.Cur < len(a.Data)
}

func (a *ArrayDoc2VecStringProvider) inc() {
	a.Cur++
}

func (a *ArrayDoc2VecStringProvider) Next() string {
	defer a.inc()
	return a.Data[a.Cur].Make()
}

//======================================================================

type IEndpointData interface {
	Name() string
	Words() IStringProvider
}

type IModelDataProvider interface {
	More() bool
	Next() IEndpointData
	Reset()
}

//======================================================================

type EndpointData struct {
	TheName string
	//TheWords *ArrayStringProvider
	TheWords IStringProvider
}

func (e *EndpointData) Name() string {
	return e.TheName
}

func (e *EndpointData) Words() IStringProvider {
	return e.TheWords
}

//======================================================================

type IResettableReader interface {
	io.Reader
	io.Seeker
}

type ModelFileDataProvider struct {
	reader    IResettableReader
	done      bool
	dataready bool
	scanner   *bufio.Scanner
	curEP     *EndpointData
	//curWords *ArrayStringProvider
}

//func NewModelFileDataProvider(reader io.Reader) *ModelFileDataProvider {
func NewModelFileDataProvider(reader IResettableReader) *ModelFileDataProvider {
	res := &ModelFileDataProvider{
		//filename: filename,
		reader: reader,
		curEP:  &EndpointData{},
	}

	// file, err := os.Open(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	res.scanner = bufio.NewScanner(reader)
	res.scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*100)
	//res.tryToAdvance()

	return res
}

func (m *ModelFileDataProvider) tryToAdvance() {
	res := m.scanner.Scan()
	if res {
		line := m.scanner.Text()
		items := strings.Split(line, "\t")
		if len(items) < 2 {
			panic(errors.New(fmt.Sprintf("Strange input line: %s", line)))
		}
		m.curEP.TheName = items[0]
		m.curEP.TheWords = NewArrayStringProvider(strings.Split(items[1], " "))
	}

	m.dataready = true
	m.done = !res
}

func (m *ModelFileDataProvider) Reset() {
	//m2 := NewModelFileDataProvider(m.file)
	//*m = *m2
	m.reader.Seek(0, io.SeekStart)

	m.dataready = false
	m.scanner = bufio.NewScanner(m.reader)
	m.scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*100)
}

func (m *ModelFileDataProvider) More() bool {
	if !m.dataready {
		m.tryToAdvance()
	}
	return !m.done
}

func (m *ModelFileDataProvider) Next() IEndpointData {
	//defer m.tryToAdvance()
	//res := m.curEP
	//return res
	m.dataready = false // rescan next time
	return m.curEP
}

//======================================================================
// Local Variables:
// indent-tabs-mode: nil
// tab-width: 4
// fill-column: 120
// End:
