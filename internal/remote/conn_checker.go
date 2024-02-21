package remote

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/loggerx"
	"github.com/sirupsen/logrus"
)

const (
	maxFailedAttempts = 3
)

// ConnChecker checks if Cloud connection is active.
type ConnChecker struct {
	failedAttempts atomic.Uint32
	deployClient   *DeploymentClient
	log            *logrus.Entry
}

// NewConnChecker returns new connection checker.
func NewConnChecker(cfg Config) *ConnChecker {
	return &ConnChecker{
		deployClient: NewDeploymentClient(cfg),
		log: loggerx.New(config.Logger{
			Level: "info",
		}).WithField("service", "auth-checker"),
	}
}

// IsConnectedWithCould returns error if Cloud connection is not active.
func (a *ConnChecker) IsConnectedWithCould() error {
	return a.deployClient.IsConnectedWithCould()
}

// AsyncSupervisor starts connection checker.
func (a *ConnChecker) AsyncSupervisor(onPermanentBlockage func()) {
	go func() {
		timer := time.NewTimer(a.nextCheck())
		defer timer.Stop()

		for {
			timer.Reset(a.nextCheck())
			<-timer.C
			err := a.deployClient.IsConnectedWithCould()
			if err != nil {
				a.failedAttempts.Add(1)
				a.log.WithField("attempt", a.failedAttempts.Load()).Error("Licence check failed")
				if a.ReachedPermanentBlockage() {
					onPermanentBlockage()
					return
				}
				continue
			}
			a.failedAttempts.Store(0)
		}
	}()
}

// ReachedPermanentBlockage returns true if reached max failed attempts.
func (a *ConnChecker) ReachedPermanentBlockage() bool {
	return a.failedAttempts.Load() > maxFailedAttempts
}

func (a *ConnChecker) nextCheck() time.Duration {
	//nolint: gosec
	i := rand.Intn(3) + 1 // [1,3]
	return time.Duration(i) * time.Hour
}
