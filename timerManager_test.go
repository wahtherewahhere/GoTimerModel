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

import(
     "testing"
     "time"
)

func Test_CreateTimerManuallyAndRegisteToManagerList_Success (t *testing.T) {
	timer := new(Timer);
	
	timerManager := GetTimerManager();
	timerManager.RegisteTimer(timer);
}

/*
func Test_CreateTimerManuallyAndRegisteToManagerList_Failed (t *testing.T) {
	timerManager := GetTimerManager();
	timerManager.RegisteTimer(nil);
	
	defer func(){
		if r := recover(); r != nil {
			t.Log("success");
		} else {
			t.Error("Test Error");
		}
	}();
}
*/

func Test_CreateTimerByTimerManage_Success (t *testing.T) {
	// 
	timerManager := GetTimerManager();
	// create a timer which interval is 1s
	// it'll return a timer's pointer
	timer := timerManager.CreateAndRegisteTimer(1000);
	
	if timer == nil { t.Error("Test Error"); }
}

func Test_TimerWithoutStarted (t *testing.T) {
	// 
	timerManager := GetTimerManager();
	//
	timer := timerManager.CreateAndRegisteTimer(1000);
	
	result := timer.IsTimeout();
	
	if result == false { t.Error("Test Error"); }
}

func Test_TimerHasBeStarted (t *testing.T) {
	
	timerManager := GetTimerManager();
	
	timer := timerManager.CreateAndRegisteTimer(1000);
	timer.Start();
	
	result := timer.IsTimeout();
	
	if result == true { t.Error("Test Error"); }
}

func Test_TimerHasBeChangedIntervalAndStarted (t *testing.T) {
	
	timerManager := GetTimerManager();
	
	timer := timerManager.CreateAndRegisteTimer(1000);
	timer.SetIntervalAndStart(3000);
	
	result := timer.IsTimeout();
	
	if result == true { t.Error("Test Error"); }
}

func Test_TimerHasBeStartedAndWillNotBeTimeout (t *testing.T) {
	
	timerManager := GetTimerManager();
	
	timer := timerManager.CreateAndRegisteTimer(1000);
	timer.SetIntervalAndStart(3000);
	time.Sleep(1 * time.Second)
	
	result := timer.IsTimeout();
	if result == true { t.Error("Test Error"); }
}

func Test_TimerHasBeStartedAndWillBeTimeout (t *testing.T) {
	
	timerManager := GetTimerManager();
	
	timer := timerManager.CreateAndRegisteTimer(1000);
	timer.SetIntervalAndStart(3000);
	time.Sleep(4 * time.Second)
	
	result := timer.IsTimeout();
	if result == false { t.Error("Test Error"); }
}

