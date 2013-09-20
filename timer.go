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
	"time"
)

type Timer struct {
	Interval int64
	StartTime int64
	OnCounting bool
}

func (self *Timer) Start () {
	self.StartTime = getTimeMillionSecond();
	self.OnCounting = true;
}

func (self *Timer) Stop () {
	self.OnCounting = false;
}

func (self *Timer) SetInterval (interval int64) {
	self.Interval = interval;
}

func (self *Timer) SetIntervalAndStart (interval int64) {
	self.Interval = interval;
	self.Start();
}

func (self *Timer) IsCounting () (bool) {
	return self.OnCounting;
}

// 是否已經到達 timeout 標準
func (self *Timer) IsTimeout () (bool) {
	if getTimeMillionSecond() - self.StartTime >= self.Interval {
		return true
	}
	return false
}

// 複製
func (self *Timer) Clone() (*Timer) {
	newTimer := new(Timer)
	newTimer.Interval = self.Interval
	newTimer.StartTime = self.StartTime
	newTimer.OnCounting = false
	return newTimer
}

//
func getTimeMillionSecond () int64 {
	return (time.Now().UnixNano() / 1e6);
}