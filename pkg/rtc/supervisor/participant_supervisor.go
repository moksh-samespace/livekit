package supervisor

import (
	"sync"
	"time"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"go.uber.org/atomic"
)

const (
	monitorInterval = 3 * time.Second
)

type ParticipantSupervisorParams struct {
	Logger logger.Logger
}

type trackMonitor struct {
	opMon types.OperationMonitor
	err   error
}

type ParticipantSupervisor struct {
	params ParticipantSupervisorParams

	lock                 sync.RWMutex
	isPublisherConnected bool
	publications         map[livekit.TrackID]*trackMonitor
	subscriptions        map[livekit.TrackID]*trackMonitor

	isStopped atomic.Bool

	onPublicationError  func(trackID livekit.TrackID)
	onSubscriptionError func(trackID livekit.TrackID)
}

func NewParticipantSupervisor(params ParticipantSupervisorParams) *ParticipantSupervisor {
	p := &ParticipantSupervisor{
		params:        params,
		publications:  make(map[livekit.TrackID]*trackMonitor),
		subscriptions: make(map[livekit.TrackID]*trackMonitor),
	}

	go p.checkState()

	return p
}

func (p *ParticipantSupervisor) Stop() {
	p.isStopped.Store(true)
}

func (p *ParticipantSupervisor) OnPublicationError(f func(trackID livekit.TrackID)) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.onPublicationError = f
}

func (p *ParticipantSupervisor) getOnPublicationError() func(trackID livekit.TrackID) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.onPublicationError
}

func (p *ParticipantSupervisor) OnSubscriptionError(f func(trackID livekit.TrackID)) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.onSubscriptionError = f
}

func (p *ParticipantSupervisor) getOnSubscriptionError() func(trackID livekit.TrackID) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.onSubscriptionError
}

func (p *ParticipantSupervisor) SetPublisherPeerConnectionConnected(isConnected bool) {
	p.lock.Lock()
	p.isPublisherConnected = isConnected

	for _, pm := range p.publications {
		pm.opMon.PostEvent(types.OperationMonitorEventPublisherPeerConnectionConnected, p.isPublisherConnected)
	}
	p.lock.Unlock()
}

func (p *ParticipantSupervisor) AddPublication(trackID livekit.TrackID) {
	p.lock.Lock()
	pm, ok := p.publications[trackID]
	if !ok {
		pm = &trackMonitor{
			opMon: NewPublicationMonitor(
				PublicationMonitorParams{
					TrackID:                   trackID,
					IsPeerConnectionConnected: p.isPublisherConnected,
					Logger:                    p.params.Logger,
				},
			),
		}
		p.publications[trackID] = pm
	}
	pm.opMon.PostEvent(types.OperationMonitorEventAddPendingPublication, nil)
	p.lock.Unlock()
}

func (p *ParticipantSupervisor) SetPublicationMute(trackID livekit.TrackID, isMuted bool) {
	p.lock.Lock()
	pm, ok := p.publications[trackID]
	if ok {
		pm.opMon.PostEvent(types.OperationMonitorEventSetPublicationMute, isMuted)
	}
	p.lock.Unlock()
}

func (p *ParticipantSupervisor) SetPublishedTrack(trackID livekit.TrackID, pubTrack types.LocalMediaTrack) {
	p.lock.RLock()
	pm, ok := p.publications[trackID]
	if ok {
		pm.opMon.PostEvent(types.OperationMonitorEventSetPublishedTrack, pubTrack)
	}
	p.lock.RUnlock()
}

func (p *ParticipantSupervisor) ClearPublishedTrack(trackID livekit.TrackID, pubTrack types.LocalMediaTrack) {
	p.lock.RLock()
	pm, ok := p.publications[trackID]
	if ok {
		pm.opMon.PostEvent(types.OperationMonitorEventClearPublishedTrack, pubTrack)
	}
	p.lock.RUnlock()
}

func (p *ParticipantSupervisor) UpdateSubscription(trackID livekit.TrackID, isSubscribe bool, sourceTrack types.MediaTrack) {
	p.lock.Lock()
	sm, ok := p.subscriptions[trackID]
	if !ok {
		sm = &trackMonitor{
			opMon: NewSubscriptionMonitor(SubscriptionMonitorParams{TrackID: trackID, Logger: p.params.Logger}),
		}
		p.subscriptions[trackID] = sm
	}
	sm.opMon.PostEvent(
		types.OperationMonitorEventUpdateSubscription,
		SubscriptionOpParams{
			SourceTrack: sourceTrack,
			IsSubscribe: isSubscribe,
		},
	)
	p.lock.Unlock()
}

func (p *ParticipantSupervisor) SetSubscribedTrack(trackID livekit.TrackID, subTrack types.SubscribedTrack, sourceTrack types.MediaTrack) {
	p.lock.RLock()
	sm, ok := p.subscriptions[trackID]
	if ok {
		sm.opMon.PostEvent(
			types.OperationMonitorEventSetSubscribedTrack,
			UpdateSubscribedTrackParams{
				SourceTrack:     sourceTrack,
				SubscribedTrack: subTrack,
			},
		)
	}
	p.lock.RUnlock()
}

func (p *ParticipantSupervisor) ClearSubscribedTrack(trackID livekit.TrackID, subTrack types.SubscribedTrack, sourceTrack types.MediaTrack) {
	p.lock.RLock()
	sm, ok := p.subscriptions[trackID]
	if ok {
		sm.opMon.PostEvent(
			types.OperationMonitorEventClearSubscribedTrack,
			UpdateSubscribedTrackParams{
				SourceTrack:     sourceTrack,
				SubscribedTrack: subTrack,
			},
		)
	}
	p.lock.RUnlock()
}

func (p *ParticipantSupervisor) checkState() {
	ticker := time.NewTicker(monitorInterval)
	defer ticker.Stop()

	for !p.isStopped.Load() {
		<-ticker.C

		p.checkPublications()
		p.checkSubscriptions()
	}
}

func (p *ParticipantSupervisor) checkPublications() {
	var erroredPublications []livekit.TrackID
	var removablePublications []livekit.TrackID
	p.lock.RLock()
	for trackID, pm := range p.publications {
		if err := pm.opMon.Check(); err != nil {
			if pm.err == nil {
				p.params.Logger.Errorw("supervisor error on publication", err, "trackID", trackID)
				pm.err = err
				erroredPublications = append(erroredPublications, trackID)
			}
		} else {
			if pm.err != nil {
				p.params.Logger.Infow("supervisor publication recovered", "trackID", trackID)
				pm.err = err
			}
			if pm.opMon.IsIdle() {
				removablePublications = append(removablePublications, trackID)
			}
		}
	}
	p.lock.RUnlock()

	p.lock.Lock()
	for _, trackID := range removablePublications {
		delete(p.publications, trackID)
	}
	p.lock.Unlock()

	if onPublicationError := p.getOnPublicationError(); onPublicationError != nil {
		for _, trackID := range erroredPublications {
			onPublicationError(trackID)
		}
	}
}

func (p *ParticipantSupervisor) checkSubscriptions() {
	var erroredSubscriptions []livekit.TrackID
	var removableSubscriptions []livekit.TrackID
	p.lock.RLock()
	for trackID, sm := range p.subscriptions {
		if err := sm.opMon.Check(); err != nil {
			if sm.err == nil {
				p.params.Logger.Errorw("supervisor error on subscription", err, "trackID", trackID)
				sm.err = err
				erroredSubscriptions = append(erroredSubscriptions, trackID)
			}
		} else {
			if sm.err != nil {
				p.params.Logger.Infow("supervisor subscription recovered", "trackID", trackID)
				sm.err = err
			}
			if sm.opMon.IsIdle() {
				removableSubscriptions = append(removableSubscriptions, trackID)
			}
		}
	}
	p.lock.RUnlock()

	p.lock.Lock()
	for _, trackID := range removableSubscriptions {
		delete(p.subscriptions, trackID)
	}
	p.lock.Unlock()

	if onSubscriptionError := p.getOnSubscriptionError(); onSubscriptionError != nil {
		for _, trackID := range erroredSubscriptions {
			onSubscriptionError(trackID)
		}
	}
}
