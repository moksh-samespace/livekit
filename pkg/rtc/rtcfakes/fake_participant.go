// Code generated by counterfeiter. DO NOT EDIT.
package rtcfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc"
	"github.com/livekit/livekit-server/pkg/sfu"
	"github.com/livekit/livekit-server/proto/livekit"
	webrtc "github.com/pion/webrtc/v3"
)

type FakeParticipant struct {
	AddDownTrackStub        func(string, *sfu.DownTrack)
	addDownTrackMutex       sync.RWMutex
	addDownTrackArgsForCall []struct {
		arg1 string
		arg2 *sfu.DownTrack
	}
	AddICECandidateStub        func(webrtc.ICECandidateInit) error
	addICECandidateMutex       sync.RWMutex
	addICECandidateArgsForCall []struct {
		arg1 webrtc.ICECandidateInit
	}
	addICECandidateReturns struct {
		result1 error
	}
	addICECandidateReturnsOnCall map[int]struct {
		result1 error
	}
	AddSubscriberStub        func(rtc.Participant) error
	addSubscriberMutex       sync.RWMutex
	addSubscriberArgsForCall []struct {
		arg1 rtc.Participant
	}
	addSubscriberReturns struct {
		result1 error
	}
	addSubscriberReturnsOnCall map[int]struct {
		result1 error
	}
	AnswerStub        func(webrtc.SessionDescription) (webrtc.SessionDescription, error)
	answerMutex       sync.RWMutex
	answerArgsForCall []struct {
		arg1 webrtc.SessionDescription
	}
	answerReturns struct {
		result1 webrtc.SessionDescription
		result2 error
	}
	answerReturnsOnCall map[int]struct {
		result1 webrtc.SessionDescription
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	HandleNegotiateStub        func(webrtc.SessionDescription) error
	handleNegotiateMutex       sync.RWMutex
	handleNegotiateArgsForCall []struct {
		arg1 webrtc.SessionDescription
	}
	handleNegotiateReturns struct {
		result1 error
	}
	handleNegotiateReturnsOnCall map[int]struct {
		result1 error
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	OnCloseStub        func(func(rtc.Participant))
	onCloseMutex       sync.RWMutex
	onCloseArgsForCall []struct {
		arg1 func(rtc.Participant)
	}
	OnICECandidateStub        func(func(c *webrtc.ICECandidateInit))
	onICECandidateMutex       sync.RWMutex
	onICECandidateArgsForCall []struct {
		arg1 func(c *webrtc.ICECandidateInit)
	}
	OnOfferStub        func(func(webrtc.SessionDescription))
	onOfferMutex       sync.RWMutex
	onOfferArgsForCall []struct {
		arg1 func(webrtc.SessionDescription)
	}
	OnStateChangeStub        func(func(p rtc.Participant, oldState livekit.ParticipantInfo_State))
	onStateChangeMutex       sync.RWMutex
	onStateChangeArgsForCall []struct {
		arg1 func(p rtc.Participant, oldState livekit.ParticipantInfo_State)
	}
	OnTrackPublishedStub        func(func(rtc.Participant, rtc.PublishedTrack))
	onTrackPublishedMutex       sync.RWMutex
	onTrackPublishedArgsForCall []struct {
		arg1 func(rtc.Participant, rtc.PublishedTrack)
	}
	PeerConnectionStub        func() rtc.PeerConnection
	peerConnectionMutex       sync.RWMutex
	peerConnectionArgsForCall []struct {
	}
	peerConnectionReturns struct {
		result1 rtc.PeerConnection
	}
	peerConnectionReturnsOnCall map[int]struct {
		result1 rtc.PeerConnection
	}
	RemoveDownTrackStub        func(string, *sfu.DownTrack)
	removeDownTrackMutex       sync.RWMutex
	removeDownTrackArgsForCall []struct {
		arg1 string
		arg2 *sfu.DownTrack
	}
	RemoveSubscriberStub        func(string)
	removeSubscriberMutex       sync.RWMutex
	removeSubscriberArgsForCall []struct {
		arg1 string
	}
	SendJoinResponseStub        func([]rtc.Participant) error
	sendJoinResponseMutex       sync.RWMutex
	sendJoinResponseArgsForCall []struct {
		arg1 []rtc.Participant
	}
	sendJoinResponseReturns struct {
		result1 error
	}
	sendJoinResponseReturnsOnCall map[int]struct {
		result1 error
	}
	SendParticipantUpdateStub        func([]*livekit.ParticipantInfo) error
	sendParticipantUpdateMutex       sync.RWMutex
	sendParticipantUpdateArgsForCall []struct {
		arg1 []*livekit.ParticipantInfo
	}
	sendParticipantUpdateReturns struct {
		result1 error
	}
	sendParticipantUpdateReturnsOnCall map[int]struct {
		result1 error
	}
	SetRemoteDescriptionStub        func(webrtc.SessionDescription) error
	setRemoteDescriptionMutex       sync.RWMutex
	setRemoteDescriptionArgsForCall []struct {
		arg1 webrtc.SessionDescription
	}
	setRemoteDescriptionReturns struct {
		result1 error
	}
	setRemoteDescriptionReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func()
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	StateStub        func() livekit.ParticipantInfo_State
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
	}
	stateReturns struct {
		result1 livekit.ParticipantInfo_State
	}
	stateReturnsOnCall map[int]struct {
		result1 livekit.ParticipantInfo_State
	}
	ToProtoStub        func() *livekit.ParticipantInfo
	toProtoMutex       sync.RWMutex
	toProtoArgsForCall []struct {
	}
	toProtoReturns struct {
		result1 *livekit.ParticipantInfo
	}
	toProtoReturnsOnCall map[int]struct {
		result1 *livekit.ParticipantInfo
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParticipant) AddDownTrack(arg1 string, arg2 *sfu.DownTrack) {
	fake.addDownTrackMutex.Lock()
	fake.addDownTrackArgsForCall = append(fake.addDownTrackArgsForCall, struct {
		arg1 string
		arg2 *sfu.DownTrack
	}{arg1, arg2})
	stub := fake.AddDownTrackStub
	fake.recordInvocation("AddDownTrack", []interface{}{arg1, arg2})
	fake.addDownTrackMutex.Unlock()
	if stub != nil {
		fake.AddDownTrackStub(arg1, arg2)
	}
}

func (fake *FakeParticipant) AddDownTrackCallCount() int {
	fake.addDownTrackMutex.RLock()
	defer fake.addDownTrackMutex.RUnlock()
	return len(fake.addDownTrackArgsForCall)
}

func (fake *FakeParticipant) AddDownTrackCalls(stub func(string, *sfu.DownTrack)) {
	fake.addDownTrackMutex.Lock()
	defer fake.addDownTrackMutex.Unlock()
	fake.AddDownTrackStub = stub
}

func (fake *FakeParticipant) AddDownTrackArgsForCall(i int) (string, *sfu.DownTrack) {
	fake.addDownTrackMutex.RLock()
	defer fake.addDownTrackMutex.RUnlock()
	argsForCall := fake.addDownTrackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeParticipant) AddICECandidate(arg1 webrtc.ICECandidateInit) error {
	fake.addICECandidateMutex.Lock()
	ret, specificReturn := fake.addICECandidateReturnsOnCall[len(fake.addICECandidateArgsForCall)]
	fake.addICECandidateArgsForCall = append(fake.addICECandidateArgsForCall, struct {
		arg1 webrtc.ICECandidateInit
	}{arg1})
	stub := fake.AddICECandidateStub
	fakeReturns := fake.addICECandidateReturns
	fake.recordInvocation("AddICECandidate", []interface{}{arg1})
	fake.addICECandidateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) AddICECandidateCallCount() int {
	fake.addICECandidateMutex.RLock()
	defer fake.addICECandidateMutex.RUnlock()
	return len(fake.addICECandidateArgsForCall)
}

func (fake *FakeParticipant) AddICECandidateCalls(stub func(webrtc.ICECandidateInit) error) {
	fake.addICECandidateMutex.Lock()
	defer fake.addICECandidateMutex.Unlock()
	fake.AddICECandidateStub = stub
}

func (fake *FakeParticipant) AddICECandidateArgsForCall(i int) webrtc.ICECandidateInit {
	fake.addICECandidateMutex.RLock()
	defer fake.addICECandidateMutex.RUnlock()
	argsForCall := fake.addICECandidateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) AddICECandidateReturns(result1 error) {
	fake.addICECandidateMutex.Lock()
	defer fake.addICECandidateMutex.Unlock()
	fake.AddICECandidateStub = nil
	fake.addICECandidateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) AddICECandidateReturnsOnCall(i int, result1 error) {
	fake.addICECandidateMutex.Lock()
	defer fake.addICECandidateMutex.Unlock()
	fake.AddICECandidateStub = nil
	if fake.addICECandidateReturnsOnCall == nil {
		fake.addICECandidateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addICECandidateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) AddSubscriber(arg1 rtc.Participant) error {
	fake.addSubscriberMutex.Lock()
	ret, specificReturn := fake.addSubscriberReturnsOnCall[len(fake.addSubscriberArgsForCall)]
	fake.addSubscriberArgsForCall = append(fake.addSubscriberArgsForCall, struct {
		arg1 rtc.Participant
	}{arg1})
	stub := fake.AddSubscriberStub
	fakeReturns := fake.addSubscriberReturns
	fake.recordInvocation("AddSubscriber", []interface{}{arg1})
	fake.addSubscriberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) AddSubscriberCallCount() int {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	return len(fake.addSubscriberArgsForCall)
}

func (fake *FakeParticipant) AddSubscriberCalls(stub func(rtc.Participant) error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = stub
}

func (fake *FakeParticipant) AddSubscriberArgsForCall(i int) rtc.Participant {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	argsForCall := fake.addSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) AddSubscriberReturns(result1 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	fake.addSubscriberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) AddSubscriberReturnsOnCall(i int, result1 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	if fake.addSubscriberReturnsOnCall == nil {
		fake.addSubscriberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addSubscriberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) Answer(arg1 webrtc.SessionDescription) (webrtc.SessionDescription, error) {
	fake.answerMutex.Lock()
	ret, specificReturn := fake.answerReturnsOnCall[len(fake.answerArgsForCall)]
	fake.answerArgsForCall = append(fake.answerArgsForCall, struct {
		arg1 webrtc.SessionDescription
	}{arg1})
	stub := fake.AnswerStub
	fakeReturns := fake.answerReturns
	fake.recordInvocation("Answer", []interface{}{arg1})
	fake.answerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParticipant) AnswerCallCount() int {
	fake.answerMutex.RLock()
	defer fake.answerMutex.RUnlock()
	return len(fake.answerArgsForCall)
}

func (fake *FakeParticipant) AnswerCalls(stub func(webrtc.SessionDescription) (webrtc.SessionDescription, error)) {
	fake.answerMutex.Lock()
	defer fake.answerMutex.Unlock()
	fake.AnswerStub = stub
}

func (fake *FakeParticipant) AnswerArgsForCall(i int) webrtc.SessionDescription {
	fake.answerMutex.RLock()
	defer fake.answerMutex.RUnlock()
	argsForCall := fake.answerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) AnswerReturns(result1 webrtc.SessionDescription, result2 error) {
	fake.answerMutex.Lock()
	defer fake.answerMutex.Unlock()
	fake.AnswerStub = nil
	fake.answerReturns = struct {
		result1 webrtc.SessionDescription
		result2 error
	}{result1, result2}
}

func (fake *FakeParticipant) AnswerReturnsOnCall(i int, result1 webrtc.SessionDescription, result2 error) {
	fake.answerMutex.Lock()
	defer fake.answerMutex.Unlock()
	fake.AnswerStub = nil
	if fake.answerReturnsOnCall == nil {
		fake.answerReturnsOnCall = make(map[int]struct {
			result1 webrtc.SessionDescription
			result2 error
		})
	}
	fake.answerReturnsOnCall[i] = struct {
		result1 webrtc.SessionDescription
		result2 error
	}{result1, result2}
}

func (fake *FakeParticipant) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeParticipant) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeParticipant) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) HandleNegotiate(arg1 webrtc.SessionDescription) error {
	fake.handleNegotiateMutex.Lock()
	ret, specificReturn := fake.handleNegotiateReturnsOnCall[len(fake.handleNegotiateArgsForCall)]
	fake.handleNegotiateArgsForCall = append(fake.handleNegotiateArgsForCall, struct {
		arg1 webrtc.SessionDescription
	}{arg1})
	stub := fake.HandleNegotiateStub
	fakeReturns := fake.handleNegotiateReturns
	fake.recordInvocation("HandleNegotiate", []interface{}{arg1})
	fake.handleNegotiateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) HandleNegotiateCallCount() int {
	fake.handleNegotiateMutex.RLock()
	defer fake.handleNegotiateMutex.RUnlock()
	return len(fake.handleNegotiateArgsForCall)
}

func (fake *FakeParticipant) HandleNegotiateCalls(stub func(webrtc.SessionDescription) error) {
	fake.handleNegotiateMutex.Lock()
	defer fake.handleNegotiateMutex.Unlock()
	fake.HandleNegotiateStub = stub
}

func (fake *FakeParticipant) HandleNegotiateArgsForCall(i int) webrtc.SessionDescription {
	fake.handleNegotiateMutex.RLock()
	defer fake.handleNegotiateMutex.RUnlock()
	argsForCall := fake.handleNegotiateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) HandleNegotiateReturns(result1 error) {
	fake.handleNegotiateMutex.Lock()
	defer fake.handleNegotiateMutex.Unlock()
	fake.HandleNegotiateStub = nil
	fake.handleNegotiateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) HandleNegotiateReturnsOnCall(i int, result1 error) {
	fake.handleNegotiateMutex.Lock()
	defer fake.handleNegotiateMutex.Unlock()
	fake.HandleNegotiateStub = nil
	if fake.handleNegotiateReturnsOnCall == nil {
		fake.handleNegotiateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.handleNegotiateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeParticipant) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeParticipant) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeParticipant) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeParticipant) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeParticipant) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeParticipant) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeParticipant) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeParticipant) OnClose(arg1 func(rtc.Participant)) {
	fake.onCloseMutex.Lock()
	fake.onCloseArgsForCall = append(fake.onCloseArgsForCall, struct {
		arg1 func(rtc.Participant)
	}{arg1})
	stub := fake.OnCloseStub
	fake.recordInvocation("OnClose", []interface{}{arg1})
	fake.onCloseMutex.Unlock()
	if stub != nil {
		fake.OnCloseStub(arg1)
	}
}

func (fake *FakeParticipant) OnCloseCallCount() int {
	fake.onCloseMutex.RLock()
	defer fake.onCloseMutex.RUnlock()
	return len(fake.onCloseArgsForCall)
}

func (fake *FakeParticipant) OnCloseCalls(stub func(func(rtc.Participant))) {
	fake.onCloseMutex.Lock()
	defer fake.onCloseMutex.Unlock()
	fake.OnCloseStub = stub
}

func (fake *FakeParticipant) OnCloseArgsForCall(i int) func(rtc.Participant) {
	fake.onCloseMutex.RLock()
	defer fake.onCloseMutex.RUnlock()
	argsForCall := fake.onCloseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) OnICECandidate(arg1 func(c *webrtc.ICECandidateInit)) {
	fake.onICECandidateMutex.Lock()
	fake.onICECandidateArgsForCall = append(fake.onICECandidateArgsForCall, struct {
		arg1 func(c *webrtc.ICECandidateInit)
	}{arg1})
	stub := fake.OnICECandidateStub
	fake.recordInvocation("OnICECandidate", []interface{}{arg1})
	fake.onICECandidateMutex.Unlock()
	if stub != nil {
		fake.OnICECandidateStub(arg1)
	}
}

func (fake *FakeParticipant) OnICECandidateCallCount() int {
	fake.onICECandidateMutex.RLock()
	defer fake.onICECandidateMutex.RUnlock()
	return len(fake.onICECandidateArgsForCall)
}

func (fake *FakeParticipant) OnICECandidateCalls(stub func(func(c *webrtc.ICECandidateInit))) {
	fake.onICECandidateMutex.Lock()
	defer fake.onICECandidateMutex.Unlock()
	fake.OnICECandidateStub = stub
}

func (fake *FakeParticipant) OnICECandidateArgsForCall(i int) func(c *webrtc.ICECandidateInit) {
	fake.onICECandidateMutex.RLock()
	defer fake.onICECandidateMutex.RUnlock()
	argsForCall := fake.onICECandidateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) OnOffer(arg1 func(webrtc.SessionDescription)) {
	fake.onOfferMutex.Lock()
	fake.onOfferArgsForCall = append(fake.onOfferArgsForCall, struct {
		arg1 func(webrtc.SessionDescription)
	}{arg1})
	stub := fake.OnOfferStub
	fake.recordInvocation("OnOffer", []interface{}{arg1})
	fake.onOfferMutex.Unlock()
	if stub != nil {
		fake.OnOfferStub(arg1)
	}
}

func (fake *FakeParticipant) OnOfferCallCount() int {
	fake.onOfferMutex.RLock()
	defer fake.onOfferMutex.RUnlock()
	return len(fake.onOfferArgsForCall)
}

func (fake *FakeParticipant) OnOfferCalls(stub func(func(webrtc.SessionDescription))) {
	fake.onOfferMutex.Lock()
	defer fake.onOfferMutex.Unlock()
	fake.OnOfferStub = stub
}

func (fake *FakeParticipant) OnOfferArgsForCall(i int) func(webrtc.SessionDescription) {
	fake.onOfferMutex.RLock()
	defer fake.onOfferMutex.RUnlock()
	argsForCall := fake.onOfferArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) OnStateChange(arg1 func(p rtc.Participant, oldState livekit.ParticipantInfo_State)) {
	fake.onStateChangeMutex.Lock()
	fake.onStateChangeArgsForCall = append(fake.onStateChangeArgsForCall, struct {
		arg1 func(p rtc.Participant, oldState livekit.ParticipantInfo_State)
	}{arg1})
	stub := fake.OnStateChangeStub
	fake.recordInvocation("OnStateChange", []interface{}{arg1})
	fake.onStateChangeMutex.Unlock()
	if stub != nil {
		fake.OnStateChangeStub(arg1)
	}
}

func (fake *FakeParticipant) OnStateChangeCallCount() int {
	fake.onStateChangeMutex.RLock()
	defer fake.onStateChangeMutex.RUnlock()
	return len(fake.onStateChangeArgsForCall)
}

func (fake *FakeParticipant) OnStateChangeCalls(stub func(func(p rtc.Participant, oldState livekit.ParticipantInfo_State))) {
	fake.onStateChangeMutex.Lock()
	defer fake.onStateChangeMutex.Unlock()
	fake.OnStateChangeStub = stub
}

func (fake *FakeParticipant) OnStateChangeArgsForCall(i int) func(p rtc.Participant, oldState livekit.ParticipantInfo_State) {
	fake.onStateChangeMutex.RLock()
	defer fake.onStateChangeMutex.RUnlock()
	argsForCall := fake.onStateChangeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) OnTrackPublished(arg1 func(rtc.Participant, rtc.PublishedTrack)) {
	fake.onTrackPublishedMutex.Lock()
	fake.onTrackPublishedArgsForCall = append(fake.onTrackPublishedArgsForCall, struct {
		arg1 func(rtc.Participant, rtc.PublishedTrack)
	}{arg1})
	stub := fake.OnTrackPublishedStub
	fake.recordInvocation("OnTrackPublished", []interface{}{arg1})
	fake.onTrackPublishedMutex.Unlock()
	if stub != nil {
		fake.OnTrackPublishedStub(arg1)
	}
}

func (fake *FakeParticipant) OnTrackPublishedCallCount() int {
	fake.onTrackPublishedMutex.RLock()
	defer fake.onTrackPublishedMutex.RUnlock()
	return len(fake.onTrackPublishedArgsForCall)
}

func (fake *FakeParticipant) OnTrackPublishedCalls(stub func(func(rtc.Participant, rtc.PublishedTrack))) {
	fake.onTrackPublishedMutex.Lock()
	defer fake.onTrackPublishedMutex.Unlock()
	fake.OnTrackPublishedStub = stub
}

func (fake *FakeParticipant) OnTrackPublishedArgsForCall(i int) func(rtc.Participant, rtc.PublishedTrack) {
	fake.onTrackPublishedMutex.RLock()
	defer fake.onTrackPublishedMutex.RUnlock()
	argsForCall := fake.onTrackPublishedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) PeerConnection() rtc.PeerConnection {
	fake.peerConnectionMutex.Lock()
	ret, specificReturn := fake.peerConnectionReturnsOnCall[len(fake.peerConnectionArgsForCall)]
	fake.peerConnectionArgsForCall = append(fake.peerConnectionArgsForCall, struct {
	}{})
	stub := fake.PeerConnectionStub
	fakeReturns := fake.peerConnectionReturns
	fake.recordInvocation("PeerConnection", []interface{}{})
	fake.peerConnectionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) PeerConnectionCallCount() int {
	fake.peerConnectionMutex.RLock()
	defer fake.peerConnectionMutex.RUnlock()
	return len(fake.peerConnectionArgsForCall)
}

func (fake *FakeParticipant) PeerConnectionCalls(stub func() rtc.PeerConnection) {
	fake.peerConnectionMutex.Lock()
	defer fake.peerConnectionMutex.Unlock()
	fake.PeerConnectionStub = stub
}

func (fake *FakeParticipant) PeerConnectionReturns(result1 rtc.PeerConnection) {
	fake.peerConnectionMutex.Lock()
	defer fake.peerConnectionMutex.Unlock()
	fake.PeerConnectionStub = nil
	fake.peerConnectionReturns = struct {
		result1 rtc.PeerConnection
	}{result1}
}

func (fake *FakeParticipant) PeerConnectionReturnsOnCall(i int, result1 rtc.PeerConnection) {
	fake.peerConnectionMutex.Lock()
	defer fake.peerConnectionMutex.Unlock()
	fake.PeerConnectionStub = nil
	if fake.peerConnectionReturnsOnCall == nil {
		fake.peerConnectionReturnsOnCall = make(map[int]struct {
			result1 rtc.PeerConnection
		})
	}
	fake.peerConnectionReturnsOnCall[i] = struct {
		result1 rtc.PeerConnection
	}{result1}
}

func (fake *FakeParticipant) RemoveDownTrack(arg1 string, arg2 *sfu.DownTrack) {
	fake.removeDownTrackMutex.Lock()
	fake.removeDownTrackArgsForCall = append(fake.removeDownTrackArgsForCall, struct {
		arg1 string
		arg2 *sfu.DownTrack
	}{arg1, arg2})
	stub := fake.RemoveDownTrackStub
	fake.recordInvocation("RemoveDownTrack", []interface{}{arg1, arg2})
	fake.removeDownTrackMutex.Unlock()
	if stub != nil {
		fake.RemoveDownTrackStub(arg1, arg2)
	}
}

func (fake *FakeParticipant) RemoveDownTrackCallCount() int {
	fake.removeDownTrackMutex.RLock()
	defer fake.removeDownTrackMutex.RUnlock()
	return len(fake.removeDownTrackArgsForCall)
}

func (fake *FakeParticipant) RemoveDownTrackCalls(stub func(string, *sfu.DownTrack)) {
	fake.removeDownTrackMutex.Lock()
	defer fake.removeDownTrackMutex.Unlock()
	fake.RemoveDownTrackStub = stub
}

func (fake *FakeParticipant) RemoveDownTrackArgsForCall(i int) (string, *sfu.DownTrack) {
	fake.removeDownTrackMutex.RLock()
	defer fake.removeDownTrackMutex.RUnlock()
	argsForCall := fake.removeDownTrackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeParticipant) RemoveSubscriber(arg1 string) {
	fake.removeSubscriberMutex.Lock()
	fake.removeSubscriberArgsForCall = append(fake.removeSubscriberArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveSubscriberStub
	fake.recordInvocation("RemoveSubscriber", []interface{}{arg1})
	fake.removeSubscriberMutex.Unlock()
	if stub != nil {
		fake.RemoveSubscriberStub(arg1)
	}
}

func (fake *FakeParticipant) RemoveSubscriberCallCount() int {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	return len(fake.removeSubscriberArgsForCall)
}

func (fake *FakeParticipant) RemoveSubscriberCalls(stub func(string)) {
	fake.removeSubscriberMutex.Lock()
	defer fake.removeSubscriberMutex.Unlock()
	fake.RemoveSubscriberStub = stub
}

func (fake *FakeParticipant) RemoveSubscriberArgsForCall(i int) string {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	argsForCall := fake.removeSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) SendJoinResponse(arg1 []rtc.Participant) error {
	var arg1Copy []rtc.Participant
	if arg1 != nil {
		arg1Copy = make([]rtc.Participant, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.sendJoinResponseMutex.Lock()
	ret, specificReturn := fake.sendJoinResponseReturnsOnCall[len(fake.sendJoinResponseArgsForCall)]
	fake.sendJoinResponseArgsForCall = append(fake.sendJoinResponseArgsForCall, struct {
		arg1 []rtc.Participant
	}{arg1Copy})
	stub := fake.SendJoinResponseStub
	fakeReturns := fake.sendJoinResponseReturns
	fake.recordInvocation("SendJoinResponse", []interface{}{arg1Copy})
	fake.sendJoinResponseMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) SendJoinResponseCallCount() int {
	fake.sendJoinResponseMutex.RLock()
	defer fake.sendJoinResponseMutex.RUnlock()
	return len(fake.sendJoinResponseArgsForCall)
}

func (fake *FakeParticipant) SendJoinResponseCalls(stub func([]rtc.Participant) error) {
	fake.sendJoinResponseMutex.Lock()
	defer fake.sendJoinResponseMutex.Unlock()
	fake.SendJoinResponseStub = stub
}

func (fake *FakeParticipant) SendJoinResponseArgsForCall(i int) []rtc.Participant {
	fake.sendJoinResponseMutex.RLock()
	defer fake.sendJoinResponseMutex.RUnlock()
	argsForCall := fake.sendJoinResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) SendJoinResponseReturns(result1 error) {
	fake.sendJoinResponseMutex.Lock()
	defer fake.sendJoinResponseMutex.Unlock()
	fake.SendJoinResponseStub = nil
	fake.sendJoinResponseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) SendJoinResponseReturnsOnCall(i int, result1 error) {
	fake.sendJoinResponseMutex.Lock()
	defer fake.sendJoinResponseMutex.Unlock()
	fake.SendJoinResponseStub = nil
	if fake.sendJoinResponseReturnsOnCall == nil {
		fake.sendJoinResponseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendJoinResponseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) SendParticipantUpdate(arg1 []*livekit.ParticipantInfo) error {
	var arg1Copy []*livekit.ParticipantInfo
	if arg1 != nil {
		arg1Copy = make([]*livekit.ParticipantInfo, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.sendParticipantUpdateMutex.Lock()
	ret, specificReturn := fake.sendParticipantUpdateReturnsOnCall[len(fake.sendParticipantUpdateArgsForCall)]
	fake.sendParticipantUpdateArgsForCall = append(fake.sendParticipantUpdateArgsForCall, struct {
		arg1 []*livekit.ParticipantInfo
	}{arg1Copy})
	stub := fake.SendParticipantUpdateStub
	fakeReturns := fake.sendParticipantUpdateReturns
	fake.recordInvocation("SendParticipantUpdate", []interface{}{arg1Copy})
	fake.sendParticipantUpdateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) SendParticipantUpdateCallCount() int {
	fake.sendParticipantUpdateMutex.RLock()
	defer fake.sendParticipantUpdateMutex.RUnlock()
	return len(fake.sendParticipantUpdateArgsForCall)
}

func (fake *FakeParticipant) SendParticipantUpdateCalls(stub func([]*livekit.ParticipantInfo) error) {
	fake.sendParticipantUpdateMutex.Lock()
	defer fake.sendParticipantUpdateMutex.Unlock()
	fake.SendParticipantUpdateStub = stub
}

func (fake *FakeParticipant) SendParticipantUpdateArgsForCall(i int) []*livekit.ParticipantInfo {
	fake.sendParticipantUpdateMutex.RLock()
	defer fake.sendParticipantUpdateMutex.RUnlock()
	argsForCall := fake.sendParticipantUpdateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) SendParticipantUpdateReturns(result1 error) {
	fake.sendParticipantUpdateMutex.Lock()
	defer fake.sendParticipantUpdateMutex.Unlock()
	fake.SendParticipantUpdateStub = nil
	fake.sendParticipantUpdateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) SendParticipantUpdateReturnsOnCall(i int, result1 error) {
	fake.sendParticipantUpdateMutex.Lock()
	defer fake.sendParticipantUpdateMutex.Unlock()
	fake.SendParticipantUpdateStub = nil
	if fake.sendParticipantUpdateReturnsOnCall == nil {
		fake.sendParticipantUpdateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendParticipantUpdateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) SetRemoteDescription(arg1 webrtc.SessionDescription) error {
	fake.setRemoteDescriptionMutex.Lock()
	ret, specificReturn := fake.setRemoteDescriptionReturnsOnCall[len(fake.setRemoteDescriptionArgsForCall)]
	fake.setRemoteDescriptionArgsForCall = append(fake.setRemoteDescriptionArgsForCall, struct {
		arg1 webrtc.SessionDescription
	}{arg1})
	stub := fake.SetRemoteDescriptionStub
	fakeReturns := fake.setRemoteDescriptionReturns
	fake.recordInvocation("SetRemoteDescription", []interface{}{arg1})
	fake.setRemoteDescriptionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) SetRemoteDescriptionCallCount() int {
	fake.setRemoteDescriptionMutex.RLock()
	defer fake.setRemoteDescriptionMutex.RUnlock()
	return len(fake.setRemoteDescriptionArgsForCall)
}

func (fake *FakeParticipant) SetRemoteDescriptionCalls(stub func(webrtc.SessionDescription) error) {
	fake.setRemoteDescriptionMutex.Lock()
	defer fake.setRemoteDescriptionMutex.Unlock()
	fake.SetRemoteDescriptionStub = stub
}

func (fake *FakeParticipant) SetRemoteDescriptionArgsForCall(i int) webrtc.SessionDescription {
	fake.setRemoteDescriptionMutex.RLock()
	defer fake.setRemoteDescriptionMutex.RUnlock()
	argsForCall := fake.setRemoteDescriptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) SetRemoteDescriptionReturns(result1 error) {
	fake.setRemoteDescriptionMutex.Lock()
	defer fake.setRemoteDescriptionMutex.Unlock()
	fake.SetRemoteDescriptionStub = nil
	fake.setRemoteDescriptionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) SetRemoteDescriptionReturnsOnCall(i int, result1 error) {
	fake.setRemoteDescriptionMutex.Lock()
	defer fake.setRemoteDescriptionMutex.Unlock()
	fake.SetRemoteDescriptionStub = nil
	if fake.setRemoteDescriptionReturnsOnCall == nil {
		fake.setRemoteDescriptionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRemoteDescriptionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub()
	}
}

func (fake *FakeParticipant) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeParticipant) StartCalls(stub func()) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeParticipant) State() livekit.ParticipantInfo_State {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
	}{})
	stub := fake.StateStub
	fakeReturns := fake.stateReturns
	fake.recordInvocation("State", []interface{}{})
	fake.stateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeParticipant) StateCalls(stub func() livekit.ParticipantInfo_State) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = stub
}

func (fake *FakeParticipant) StateReturns(result1 livekit.ParticipantInfo_State) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 livekit.ParticipantInfo_State
	}{result1}
}

func (fake *FakeParticipant) StateReturnsOnCall(i int, result1 livekit.ParticipantInfo_State) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantInfo_State
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 livekit.ParticipantInfo_State
	}{result1}
}

func (fake *FakeParticipant) ToProto() *livekit.ParticipantInfo {
	fake.toProtoMutex.Lock()
	ret, specificReturn := fake.toProtoReturnsOnCall[len(fake.toProtoArgsForCall)]
	fake.toProtoArgsForCall = append(fake.toProtoArgsForCall, struct {
	}{})
	stub := fake.ToProtoStub
	fakeReturns := fake.toProtoReturns
	fake.recordInvocation("ToProto", []interface{}{})
	fake.toProtoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) ToProtoCallCount() int {
	fake.toProtoMutex.RLock()
	defer fake.toProtoMutex.RUnlock()
	return len(fake.toProtoArgsForCall)
}

func (fake *FakeParticipant) ToProtoCalls(stub func() *livekit.ParticipantInfo) {
	fake.toProtoMutex.Lock()
	defer fake.toProtoMutex.Unlock()
	fake.ToProtoStub = stub
}

func (fake *FakeParticipant) ToProtoReturns(result1 *livekit.ParticipantInfo) {
	fake.toProtoMutex.Lock()
	defer fake.toProtoMutex.Unlock()
	fake.ToProtoStub = nil
	fake.toProtoReturns = struct {
		result1 *livekit.ParticipantInfo
	}{result1}
}

func (fake *FakeParticipant) ToProtoReturnsOnCall(i int, result1 *livekit.ParticipantInfo) {
	fake.toProtoMutex.Lock()
	defer fake.toProtoMutex.Unlock()
	fake.ToProtoStub = nil
	if fake.toProtoReturnsOnCall == nil {
		fake.toProtoReturnsOnCall = make(map[int]struct {
			result1 *livekit.ParticipantInfo
		})
	}
	fake.toProtoReturnsOnCall[i] = struct {
		result1 *livekit.ParticipantInfo
	}{result1}
}

func (fake *FakeParticipant) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addDownTrackMutex.RLock()
	defer fake.addDownTrackMutex.RUnlock()
	fake.addICECandidateMutex.RLock()
	defer fake.addICECandidateMutex.RUnlock()
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	fake.answerMutex.RLock()
	defer fake.answerMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.handleNegotiateMutex.RLock()
	defer fake.handleNegotiateMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.onCloseMutex.RLock()
	defer fake.onCloseMutex.RUnlock()
	fake.onICECandidateMutex.RLock()
	defer fake.onICECandidateMutex.RUnlock()
	fake.onOfferMutex.RLock()
	defer fake.onOfferMutex.RUnlock()
	fake.onStateChangeMutex.RLock()
	defer fake.onStateChangeMutex.RUnlock()
	fake.onTrackPublishedMutex.RLock()
	defer fake.onTrackPublishedMutex.RUnlock()
	fake.peerConnectionMutex.RLock()
	defer fake.peerConnectionMutex.RUnlock()
	fake.removeDownTrackMutex.RLock()
	defer fake.removeDownTrackMutex.RUnlock()
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	fake.sendJoinResponseMutex.RLock()
	defer fake.sendJoinResponseMutex.RUnlock()
	fake.sendParticipantUpdateMutex.RLock()
	defer fake.sendParticipantUpdateMutex.RUnlock()
	fake.setRemoteDescriptionMutex.RLock()
	defer fake.setRemoteDescriptionMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.toProtoMutex.RLock()
	defer fake.toProtoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParticipant) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rtc.Participant = new(FakeParticipant)
