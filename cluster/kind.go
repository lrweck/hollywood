package cluster

import "github.com/anthdm/hollywood/actor"

type KindOpts struct {
	autoActivate bool
	local        bool
}

type Kind struct {
	opts     KindOpts
	name     string
	producer actor.Producer
}

func NewKind(name string, p actor.Producer, opts KindOpts) *Kind {
	return &Kind{
		name:     name,
		opts:     opts,
		producer: p,
	}
}

// ActiveKind is a kind that is active somewhere on the cluster.
type ActiveKind struct {
	// pid of the activated kind
	pid *actor.PID
	// cid of the activated kind
	cid *CID
	// Wether the actor is activated on this cluster or not.
	isLocal bool
}

func (k ActiveKind) Equals(other ActiveKind) bool {
	return k.cid.ID == other.cid.ID &&
		k.cid.Kind == other.cid.Kind &&
		k.pid.ID == other.pid.ID &&
		k.pid.Address == other.pid.Address
}
