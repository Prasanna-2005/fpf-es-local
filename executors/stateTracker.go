package executors

import (
	"reflect"
	"sync"
)

// WorkPaperExporter is satisfied by any FEG entity (PrivateFund, a Curio
// deal, an InvesmentServices position, ...) that flattens its workpaper data
// into entity -> value via ExportWorkPaperData().
type WorkPaperExporter interface {
	ExportWorkPaperData() map[string]string
}

type FormDataExporter interface {
	ExportFormData() map[string]string
}

// StateChange is one entity whose value differs between two snapshots.
type StateChange struct {
	Entity   string
	Old, New string
	Change   bool
}

// StateTracker holds a snapshot of one entity instance's workpaper data and
// reports what changed against it later. Safe for concurrent use.
type StateTracker struct {
	mu       sync.Mutex
	baseline map[string]string
}

// NewStateTracker captures exp's current state as the baseline to diff
// future calls against.
func NewWPDataTracker(exp WorkPaperExporter) *StateTracker {
	return &StateTracker{baseline: exp.ExportWorkPaperData()}
}

func NewFormStateTracker(exp FormDataExporter) *StateTracker {
	return &StateTracker{baseline: exp.ExportFormData()}
}

// Snapshot re-captures the baseline from exp's current state.
func (s *StateTracker) Snapshot(exp WorkPaperExporter) {
	snap := exp.ExportWorkPaperData()
	s.mu.Lock()
	s.baseline = snap
	s.mu.Unlock()
}

// SnapshotForm re-captures the baseline from form data's current state.
func (s *StateTracker) SnapshotForm(exp FormDataExporter) {
	snap := exp.ExportFormData()
	s.mu.Lock()
	s.baseline = snap
	s.mu.Unlock()
}

// Diff compares exp's current state against the tracked baseline. Does not
// move the baseline forward -- call Snapshot separately if you want that.
func (s *StateTracker) Diff(exp WorkPaperExporter) []StateChange {
	s.mu.Lock()
	baseline := s.baseline
	s.mu.Unlock()
	return diffEntityMaps(baseline, exp.ExportWorkPaperData())
}

func (s *StateTracker) DiffForm(exp FormDataExporter) []StateChange {
	s.mu.Lock()
	baseline := s.baseline
	s.mu.Unlock()
	return diffEntityMaps(baseline, exp.ExportFormData())
}

// DiffWorkPaperData compares two exporters' current state directly --
// stateless, for when you already have both a "before" and "after" instance.
func DiffWorkPaperData(oldExp, newExp WorkPaperExporter) []StateChange {
	return diffEntityMaps(oldExp.ExportWorkPaperData(), newExp.ExportWorkPaperData())
}

func diffEntityMaps(old, new map[string]string) []StateChange {
	changes := make([]StateChange, 0)

	for entity, oldVal := range old {
		if newVal, ok := new[entity]; !ok || newVal != oldVal {
			changes = append(changes, StateChange{Entity: entity, Old: oldVal, New: new[entity], Change: true})
		} else {
			changes = append(changes, StateChange{Entity: entity, Old: oldVal, New: new[entity], Change: false})
		}
	}
	for entity, newVal := range new {
		if _, ok := old[entity]; !ok {
			changes = append(changes, StateChange{Entity: entity, Old: "", New: newVal, Change: true})
		}
	}
	return changes
}

func (f Form1bData) ExportFormData() (ansMapper map[string]string) {
	entityValueMapper := make(map[[2]string]string)
	ansMapper = make(map[string]string)
	CollectAnswers(reflect.ValueOf(f), entityValueMapper)
	for k, v := range entityValueMapper {
		ansMapper[k[0]+"_quantrium_"+k[1]] = v
	}
	return ansMapper
}
