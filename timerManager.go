// (The MIT License)
//
// Copyright (C) 2013 wahtherewahhere@github.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and 
// associated documentation files (the "Software"), 
// to deal in the Software without restriction, including without limitation the rights to 
// use, copy, modify, merge, publish, distribute, sublicense, 
// and/or sell copies of the Software, 
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be 
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. 
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, 
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, 
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
//

package GoTimerModel

import (
	"container/list"
)

type TimerManagerInstance struct {
	timerList *list.List  // pointer to timer
	counter int64
	owner string
}

func GetTimerManager() (*TimerManagerInstance) {
	instance := new(TimerManagerInstance);
	instance.timerList = new(list.List);
	
	return instance;
}

func (self *TimerManagerInstance) SetOwnerName (name string) {
	self.owner = name;
}

// 在外面產生一個timer，
// 調用此function，
// 該timer便會進入timerList，被manager服務
func (self *TimerManagerInstance) RegisteTimer (t *Timer) {
	
	if t == nil {
		panic("Insert a nil pointer of timer into RegisteTimer()");
	}
	
	self.timerList.PushBack(t);
	self.counter++;
}

// 幫忙產生一個 timer並註冊到 list 裡面
func (self *TimerManagerInstance) CreateAndRegisteTimer (interval int64) (t *Timer){
	// 產生 timer
	timer := new(Timer);
	timer.Interval = interval;
	timer.OnCounting = false;
	// 註冊 timer
	self.RegisteTimer(timer);
	return timer;
}

