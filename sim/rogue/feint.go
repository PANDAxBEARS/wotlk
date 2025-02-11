package rogue

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/core/proto"
	"github.com/wowsims/wotlk/sim/core/stats"
)

func (rogue *Rogue) registerFeintSpell() {
	resourceType := stats.Energy
	baseCost := 20.0
	if rogue.HasMajorGlyph(proto.RogueMajorGlyph_GlyphOfFeint) {
		resourceType = 0
		baseCost = 0
	}
	rogue.Feint = rogue.RegisterSpell(core.SpellConfig{
		ActionID:     core.ActionID{SpellID: 48659},
		SpellSchool:  core.SpellSchoolPhysical,
		ProcMask:     core.ProcMaskMeleeMHSpecial,
		Flags:        core.SpellFlagMeleeMetrics,
		ResourceType: resourceType,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  time.Second,
			},
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: time.Second * 10,
			},
			IgnoreHaste: true,
		},

		DamageMultiplier: 0,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
		},
	})
	// Feint
	if rogue.Rotation.UseFeint {
		rogue.AddMajorCooldown(core.MajorCooldown{
			Spell:    rogue.Feint,
			Priority: core.CooldownPriorityDefault,
			Type:     core.CooldownTypeDPS,
		})
	}
}
