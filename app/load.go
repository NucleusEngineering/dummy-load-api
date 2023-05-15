// Copyright 2023 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"log"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

func GenerateLoad(coresCount int, timeMilliSeconds int, percentage int, memoryMegabytes int) {

	runtime.GOMAXPROCS(coresCount)
	var wg sync.WaitGroup

	/*
		Every loop consists of running the cycle + sleep
		Running cycle ~= 10ms = 10000 microseconds
		Unit of work = Running cycle / 100 (%) = 100
	*/

	// Unit of work in microsecnods
	unitOfWork := 100

	// Run cycle duration
	runDuration := unitOfWork * percentage

	// Total cycle duration minus run duration
	sleepDuration := unitOfWork*100 - runDuration

	log.Printf("Running %d%% load on %d core(s) for %d milliseconds.", percentage, coresCount, timeMilliSeconds)

	wg.Add(coresCount)

	for i := 0; i < coresCount; i++ {
		go func() {
			start := time.Now()

			defer func() {
				log.Printf("Done faking load. Existing go routine after %d ms", time.Since(start).Milliseconds())
				wg.Done()
			}()

			runtime.LockOSThread()

			perRoutineMemory := int(memoryMegabytes / coresCount)

			// Allocate and fill memory. Using C extern to avoid any possible interaction with GC and go's memory management
			cptr := C.malloc(1024 * 1024 * C.ulong(perRoutineMemory))
			C.memset(cptr, 0, 1024*1024*C.ulong(perRoutineMemory))

			log.Printf("Allocated memory: %d Mb\n", perRoutineMemory)

			for {
				begin := time.Now()
				for {
					// Did the run cycle end? If so, break out of it and sleep for the rest
					if time.Since(begin) > time.Duration(runDuration)*time.Microsecond {
						break
					}
				}

				// Did the total work time end? If yes, exit go routine.
				if time.Since(start) > time.Duration(timeMilliSeconds)*time.Millisecond {
					break
				}

				time.Sleep(time.Duration(sleepDuration) * time.Microsecond)
			}

			defer func() {
				log.Printf("De-allocating memory")
				C.free(unsafe.Pointer(cptr))
			}()

		}()
	}

	wg.Wait()
}
