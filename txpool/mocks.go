// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package txpool

import (
	"sync"
)

// Ensure, that PoolMock does implement Pool.
// If this is not the case, regenerate this file with moq.
var _ Pool = &PoolMock{}

// PoolMock is a mock implementation of Pool.
//
// 	func TestSomethingThatUsesPool(t *testing.T) {
//
// 		// make and configure a mocked Pool
// 		mockedPool := &PoolMock{
// 			IdHashKnownFunc: func(hash []byte) bool {
// 				panic("mock out the IdHashKnown method")
// 			},
// 			NotifyNewPeerFunc: func(peerID PeerID)  {
// 				panic("mock out the NotifyNewPeer method")
// 			},
// 		}
//
// 		// use mockedPool in code that requires Pool
// 		// and then make assertions.
//
// 	}
type PoolMock struct {
	// IdHashKnownFunc mocks the IdHashKnown method.
	IdHashKnownFunc func(hash []byte) bool

	// NotifyNewPeerFunc mocks the NotifyNewPeer method.
	NotifyNewPeerFunc func(peerID PeerID)

	// calls tracks calls to the methods.
	calls struct {
		// IdHashKnown holds details about calls to the IdHashKnown method.
		IdHashKnown []struct {
			// Hash is the hash argument value.
			Hash []byte
		}
		// NotifyNewPeer holds details about calls to the NotifyNewPeer method.
		NotifyNewPeer []struct {
			// PeerID is the peerID argument value.
			PeerID PeerID
		}
	}
	lockIdHashKnown   sync.RWMutex
	lockNotifyNewPeer sync.RWMutex
}

// IdHashKnown calls IdHashKnownFunc.
func (mock *PoolMock) IdHashKnown(hash []byte) bool {
	callInfo := struct {
		Hash []byte
	}{
		Hash: hash,
	}
	mock.lockIdHashKnown.Lock()
	mock.calls.IdHashKnown = append(mock.calls.IdHashKnown, callInfo)
	mock.lockIdHashKnown.Unlock()
	if mock.IdHashKnownFunc == nil {
		var (
			bOut bool
		)
		return bOut
	}
	return mock.IdHashKnownFunc(hash)
}

// IdHashKnownCalls gets all the calls that were made to IdHashKnown.
// Check the length with:
//     len(mockedPool.IdHashKnownCalls())
func (mock *PoolMock) IdHashKnownCalls() []struct {
	Hash []byte
} {
	var calls []struct {
		Hash []byte
	}
	mock.lockIdHashKnown.RLock()
	calls = mock.calls.IdHashKnown
	mock.lockIdHashKnown.RUnlock()
	return calls
}

// NotifyNewPeer calls NotifyNewPeerFunc.
func (mock *PoolMock) NotifyNewPeer(peerID PeerID) {
	callInfo := struct {
		PeerID PeerID
	}{
		PeerID: peerID,
	}
	mock.lockNotifyNewPeer.Lock()
	mock.calls.NotifyNewPeer = append(mock.calls.NotifyNewPeer, callInfo)
	mock.lockNotifyNewPeer.Unlock()
	if mock.NotifyNewPeerFunc == nil {
		return
	}
	mock.NotifyNewPeerFunc(peerID)
}

// NotifyNewPeerCalls gets all the calls that were made to NotifyNewPeer.
// Check the length with:
//     len(mockedPool.NotifyNewPeerCalls())
func (mock *PoolMock) NotifyNewPeerCalls() []struct {
	PeerID PeerID
} {
	var calls []struct {
		PeerID PeerID
	}
	mock.lockNotifyNewPeer.RLock()
	calls = mock.calls.NotifyNewPeer
	mock.lockNotifyNewPeer.RUnlock()
	return calls
}
