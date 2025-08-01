// Copyright 2025 openGemini Authors.
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

package memory

import "time"

const (
	defaultMaxMem = 64 << 30 // default memory value
	cacheInterval = 100 * time.Millisecond
)

type MemoryMonitor interface {
	MemUsedPct() float64    // memory usage percentage
	SysMem() (int64, int64) // get the total amount of memory and remaining capacity (total, available (kB))
}
