/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"testing"

	"github.com/minio/dsync"
	xnet "github.com/manmandong/minio/pkg/net"
)

// Tests lock rpc client.
func TestLockRESTlient(t *testing.T) {
	host, err := xnet.ParseHost("localhost:9000")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	lkClient := newlockRESTClient(host)
	if lkClient.connected == false {
		t.Fatalf("unexpected error. connection failed")
	}

	// Attempt all calls.
	_, err = lkClient.RLock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Rlock to fail")
	}

	_, err = lkClient.Lock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Lock to fail")
	}

	_, err = lkClient.RUnlock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for RUnlock to fail")
	}

	_, err = lkClient.Unlock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Unlock to fail")
	}

	_, err = lkClient.ForceUnlock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for ForceUnlock to fail")
	}

	_, err = lkClient.Expired(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Expired to fail")
	}
}
