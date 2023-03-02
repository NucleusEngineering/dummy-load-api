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

import (
	"crypto/rand"
	"log"
	"runtime"
	"time"
)

func GenerateCPUAndMemoryLoad(coresCount int, timeMilliSeconds int, percentage int, memoryMegabytes int) {
	runtime.GOMAXPROCS(coresCount)

	// Every loop (1 unit of work) consists of running the cycle + sleep
	// 1 unit of work ~= 100ms

	// 100ms in microserconds = 100000 microseconds
	unitOfWork := 1000

	// Run duration
	runMicrosecond := unitOfWork * percentage

	// Sleep duration
	sleepMicrosecond := unitOfWork*100 - runMicrosecond

	log.Printf("Running %d%% load on %d core(s) for %d milliseconds.", percentage, coresCount, timeMilliSeconds)

	for i := 0; i < coresCount; i++ {
		go func() {
			runtime.LockOSThread()

			// Allocate evenly per thread
			alloc(int(memoryMegabytes / coresCount))
			start := time.Now()

			for {
				begin := time.Now()
				for {
					if time.Since(begin) > time.Duration(runMicrosecond)*time.Microsecond {
						break
					}
				}

				if time.Since(start) > time.Duration(timeMilliSeconds)*time.Millisecond {
					log.Printf("Done faking load. Existing go routine")
					break
				}

				time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
			}

			// Force memory cleanup
			runtime.GC()
		}()
	}

	time.Sleep(time.Duration(timeMilliSeconds) * time.Millisecond)
}

func alloc(mb int) {
	token := make([]byte, 1024*1024*mb)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}

	log.Printf("Allocated memory: %d bytes\n", len(token))
}
