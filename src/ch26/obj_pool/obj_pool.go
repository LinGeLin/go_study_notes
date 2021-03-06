package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	// 创建 channel
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	// channel 中添加对象
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <- p.bufChan:
		return ret, nil
	case <- time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default: // 放不进去时走该分支
		return errors.New("overflow")
	}
}