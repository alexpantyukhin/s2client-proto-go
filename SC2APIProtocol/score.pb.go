// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2clientprotocol/score.proto

package SC2APIProtocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Score_ScoreType int32

const (
	Score_Curriculum Score_ScoreType = 1
	Score_Melee      Score_ScoreType = 2
)

var Score_ScoreType_name = map[int32]string{
	1: "Curriculum",
	2: "Melee",
}
var Score_ScoreType_value = map[string]int32{
	"Curriculum": 1,
	"Melee":      2,
}

func (x Score_ScoreType) Enum() *Score_ScoreType {
	p := new(Score_ScoreType)
	*p = x
	return p
}
func (x Score_ScoreType) String() string {
	return proto.EnumName(Score_ScoreType_name, int32(x))
}
func (x *Score_ScoreType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Score_ScoreType_value, data, "Score_ScoreType")
	if err != nil {
		return err
	}
	*x = Score_ScoreType(value)
	return nil
}
func (Score_ScoreType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type Score struct {
	ScoreType        *Score_ScoreType `protobuf:"varint,6,opt,name=score_type,json=scoreType,enum=SC2APIProtocol.Score_ScoreType" json:"score_type,omitempty"`
	Score            *int32           `protobuf:"varint,7,opt,name=score" json:"score,omitempty"`
	ScoreDetails     *ScoreDetails    `protobuf:"bytes,8,opt,name=score_details,json=scoreDetails" json:"score_details,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Score) Reset()                    { *m = Score{} }
func (m *Score) String() string            { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()               {}
func (*Score) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Score) GetScoreType() Score_ScoreType {
	if m != nil && m.ScoreType != nil {
		return *m.ScoreType
	}
	return Score_Curriculum
}

func (m *Score) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Score) GetScoreDetails() *ScoreDetails {
	if m != nil {
		return m.ScoreDetails
	}
	return nil
}

type CategoryScoreDetails struct {
	None             *float32 `protobuf:"fixed32,1,opt,name=none" json:"none,omitempty"`
	Army             *float32 `protobuf:"fixed32,2,opt,name=army" json:"army,omitempty"`
	Economy          *float32 `protobuf:"fixed32,3,opt,name=economy" json:"economy,omitempty"`
	Technology       *float32 `protobuf:"fixed32,4,opt,name=technology" json:"technology,omitempty"`
	Upgrade          *float32 `protobuf:"fixed32,5,opt,name=upgrade" json:"upgrade,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CategoryScoreDetails) Reset()                    { *m = CategoryScoreDetails{} }
func (m *CategoryScoreDetails) String() string            { return proto.CompactTextString(m) }
func (*CategoryScoreDetails) ProtoMessage()               {}
func (*CategoryScoreDetails) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CategoryScoreDetails) GetNone() float32 {
	if m != nil && m.None != nil {
		return *m.None
	}
	return 0
}

func (m *CategoryScoreDetails) GetArmy() float32 {
	if m != nil && m.Army != nil {
		return *m.Army
	}
	return 0
}

func (m *CategoryScoreDetails) GetEconomy() float32 {
	if m != nil && m.Economy != nil {
		return *m.Economy
	}
	return 0
}

func (m *CategoryScoreDetails) GetTechnology() float32 {
	if m != nil && m.Technology != nil {
		return *m.Technology
	}
	return 0
}

func (m *CategoryScoreDetails) GetUpgrade() float32 {
	if m != nil && m.Upgrade != nil {
		return *m.Upgrade
	}
	return 0
}

type VitalScoreDetails struct {
	Life             *float32 `protobuf:"fixed32,1,opt,name=life" json:"life,omitempty"`
	Shields          *float32 `protobuf:"fixed32,2,opt,name=shields" json:"shields,omitempty"`
	Energy           *float32 `protobuf:"fixed32,3,opt,name=energy" json:"energy,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *VitalScoreDetails) Reset()                    { *m = VitalScoreDetails{} }
func (m *VitalScoreDetails) String() string            { return proto.CompactTextString(m) }
func (*VitalScoreDetails) ProtoMessage()               {}
func (*VitalScoreDetails) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *VitalScoreDetails) GetLife() float32 {
	if m != nil && m.Life != nil {
		return *m.Life
	}
	return 0
}

func (m *VitalScoreDetails) GetShields() float32 {
	if m != nil && m.Shields != nil {
		return *m.Shields
	}
	return 0
}

func (m *VitalScoreDetails) GetEnergy() float32 {
	if m != nil && m.Energy != nil {
		return *m.Energy
	}
	return 0
}

type ScoreDetails struct {
	// Sum of time any available structure able to produce a unit is not. The time stacks, as in, three idle barracks will increase idle_production_time three times quicker than just one.
	IdleProductionTime *float32 `protobuf:"fixed32,1,opt,name=idle_production_time,json=idleProductionTime" json:"idle_production_time,omitempty"`
	// Sum of time any worker is not mining. Note a worker building is not idle and three idle workers will increase this value three times quicker than just one.
	IdleWorkerTime *float32 `protobuf:"fixed32,2,opt,name=idle_worker_time,json=idleWorkerTime" json:"idle_worker_time,omitempty"`
	// Sum of minerals and vespene spent on completed units.
	TotalValueUnits *float32 `protobuf:"fixed32,3,opt,name=total_value_units,json=totalValueUnits" json:"total_value_units,omitempty"`
	// Sum of minerals and vespene spent on completed structures.
	TotalValueStructures *float32 `protobuf:"fixed32,4,opt,name=total_value_structures,json=totalValueStructures" json:"total_value_structures,omitempty"`
	// Sum of minerals and vespene of units, belonging to the opponent, that the player has destroyed.
	KilledValueUnits *float32 `protobuf:"fixed32,5,opt,name=killed_value_units,json=killedValueUnits" json:"killed_value_units,omitempty"`
	// Sum of minerals and vespene of structures, belonging to the opponent, that the player has destroyed.
	KilledValueStructures *float32 `protobuf:"fixed32,6,opt,name=killed_value_structures,json=killedValueStructures" json:"killed_value_structures,omitempty"`
	// Sum of minerals collected by the player.
	CollectedMinerals *float32 `protobuf:"fixed32,7,opt,name=collected_minerals,json=collectedMinerals" json:"collected_minerals,omitempty"`
	// Sum of vespene collected by the player.
	CollectedVespene *float32 `protobuf:"fixed32,8,opt,name=collected_vespene,json=collectedVespene" json:"collected_vespene,omitempty"`
	// Estimated income of minerals over the next minute based on the players current income. The unit is minerals per minute.
	CollectionRateMinerals *float32 `protobuf:"fixed32,9,opt,name=collection_rate_minerals,json=collectionRateMinerals" json:"collection_rate_minerals,omitempty"`
	// Estimated income of vespene over the next minute based on the players current income. The unit is vespene per minute.
	CollectionRateVespene *float32 `protobuf:"fixed32,10,opt,name=collection_rate_vespene,json=collectionRateVespene" json:"collection_rate_vespene,omitempty"`
	// Sum of spent minerals at the moment it is spent. For example, this number is incremented by 50 the moment an scv is queued in a command center.  It is decremented by 50 if that unit is canceled.
	SpentMinerals *float32 `protobuf:"fixed32,11,opt,name=spent_minerals,json=spentMinerals" json:"spent_minerals,omitempty"`
	// Sum of spent vespene at the moment it is spent. For example, this number is incremented by 50 when a reaper is queued but decremented by 50 if it is canceled.
	SpentVespene *float32 `protobuf:"fixed32,12,opt,name=spent_vespene,json=spentVespene" json:"spent_vespene,omitempty"`
	// Sum of food, or supply, utilized in the categories above.
	FoodUsed *CategoryScoreDetails `protobuf:"bytes,13,opt,name=food_used,json=foodUsed" json:"food_used,omitempty"`
	// Sum of enemies catagories destroyed in minerals.
	KilledMinerals *CategoryScoreDetails `protobuf:"bytes,14,opt,name=killed_minerals,json=killedMinerals" json:"killed_minerals,omitempty"`
	// Sum of enemies catagories destroyed in vespene.
	KilledVespene *CategoryScoreDetails `protobuf:"bytes,15,opt,name=killed_vespene,json=killedVespene" json:"killed_vespene,omitempty"`
	//  Sum of lost minerals for the player in each category.
	LostMinerals *CategoryScoreDetails `protobuf:"bytes,16,opt,name=lost_minerals,json=lostMinerals" json:"lost_minerals,omitempty"`
	// Sum of lost vespene for the player in each category.
	LostVespene *CategoryScoreDetails `protobuf:"bytes,17,opt,name=lost_vespene,json=lostVespene" json:"lost_vespene,omitempty"`
	// Sum of the lost minerals via destroying the players own units/buildings.
	FriendlyFireMinerals *CategoryScoreDetails `protobuf:"bytes,18,opt,name=friendly_fire_minerals,json=friendlyFireMinerals" json:"friendly_fire_minerals,omitempty"`
	// Sum of the lost vespene via destroying the players own units/buildings.
	FriendlyFireVespene *CategoryScoreDetails `protobuf:"bytes,19,opt,name=friendly_fire_vespene,json=friendlyFireVespene" json:"friendly_fire_vespene,omitempty"`
	// Sum of used minerals for the player in each category for each existing unit or upgrade. Therefore if a unit died worth 50 mierals this number will be decremented by 50.
	UsedMinerals *CategoryScoreDetails `protobuf:"bytes,20,opt,name=used_minerals,json=usedMinerals" json:"used_minerals,omitempty"`
	// Sum of used vespene for the player in each category. Therefore if a unit died worth 50 vespene this number will be decremented by 50.
	UsedVespene *CategoryScoreDetails `protobuf:"bytes,21,opt,name=used_vespene,json=usedVespene" json:"used_vespene,omitempty"`
	// Sum of used minerals throughout the entire game for each category. Unliked used_minerals, this value is never decremented.
	TotalUsedMinerals *CategoryScoreDetails `protobuf:"bytes,22,opt,name=total_used_minerals,json=totalUsedMinerals" json:"total_used_minerals,omitempty"`
	// Sum of used vespene throughout the entire game for each category. Unliked used_vespene, this value is never decremented.
	TotalUsedVespene *CategoryScoreDetails `protobuf:"bytes,23,opt,name=total_used_vespene,json=totalUsedVespene" json:"total_used_vespene,omitempty"`
	// Sum of damage dealt to the player's opponent for each category.
	TotalDamageDealt *VitalScoreDetails `protobuf:"bytes,24,opt,name=total_damage_dealt,json=totalDamageDealt" json:"total_damage_dealt,omitempty"`
	// Sum of damage taken by the player for each category.
	TotalDamageTaken *VitalScoreDetails `protobuf:"bytes,25,opt,name=total_damage_taken,json=totalDamageTaken" json:"total_damage_taken,omitempty"`
	// Sum of health healed by the player. Note that technology can be healed (by queens) or repaired (by scvs).
	TotalHealed      *VitalScoreDetails `protobuf:"bytes,26,opt,name=total_healed,json=totalHealed" json:"total_healed,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ScoreDetails) Reset()                    { *m = ScoreDetails{} }
func (m *ScoreDetails) String() string            { return proto.CompactTextString(m) }
func (*ScoreDetails) ProtoMessage()               {}
func (*ScoreDetails) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ScoreDetails) GetIdleProductionTime() float32 {
	if m != nil && m.IdleProductionTime != nil {
		return *m.IdleProductionTime
	}
	return 0
}

func (m *ScoreDetails) GetIdleWorkerTime() float32 {
	if m != nil && m.IdleWorkerTime != nil {
		return *m.IdleWorkerTime
	}
	return 0
}

func (m *ScoreDetails) GetTotalValueUnits() float32 {
	if m != nil && m.TotalValueUnits != nil {
		return *m.TotalValueUnits
	}
	return 0
}

func (m *ScoreDetails) GetTotalValueStructures() float32 {
	if m != nil && m.TotalValueStructures != nil {
		return *m.TotalValueStructures
	}
	return 0
}

func (m *ScoreDetails) GetKilledValueUnits() float32 {
	if m != nil && m.KilledValueUnits != nil {
		return *m.KilledValueUnits
	}
	return 0
}

func (m *ScoreDetails) GetKilledValueStructures() float32 {
	if m != nil && m.KilledValueStructures != nil {
		return *m.KilledValueStructures
	}
	return 0
}

func (m *ScoreDetails) GetCollectedMinerals() float32 {
	if m != nil && m.CollectedMinerals != nil {
		return *m.CollectedMinerals
	}
	return 0
}

func (m *ScoreDetails) GetCollectedVespene() float32 {
	if m != nil && m.CollectedVespene != nil {
		return *m.CollectedVespene
	}
	return 0
}

func (m *ScoreDetails) GetCollectionRateMinerals() float32 {
	if m != nil && m.CollectionRateMinerals != nil {
		return *m.CollectionRateMinerals
	}
	return 0
}

func (m *ScoreDetails) GetCollectionRateVespene() float32 {
	if m != nil && m.CollectionRateVespene != nil {
		return *m.CollectionRateVespene
	}
	return 0
}

func (m *ScoreDetails) GetSpentMinerals() float32 {
	if m != nil && m.SpentMinerals != nil {
		return *m.SpentMinerals
	}
	return 0
}

func (m *ScoreDetails) GetSpentVespene() float32 {
	if m != nil && m.SpentVespene != nil {
		return *m.SpentVespene
	}
	return 0
}

func (m *ScoreDetails) GetFoodUsed() *CategoryScoreDetails {
	if m != nil {
		return m.FoodUsed
	}
	return nil
}

func (m *ScoreDetails) GetKilledMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.KilledMinerals
	}
	return nil
}

func (m *ScoreDetails) GetKilledVespene() *CategoryScoreDetails {
	if m != nil {
		return m.KilledVespene
	}
	return nil
}

func (m *ScoreDetails) GetLostMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.LostMinerals
	}
	return nil
}

func (m *ScoreDetails) GetLostVespene() *CategoryScoreDetails {
	if m != nil {
		return m.LostVespene
	}
	return nil
}

func (m *ScoreDetails) GetFriendlyFireMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.FriendlyFireMinerals
	}
	return nil
}

func (m *ScoreDetails) GetFriendlyFireVespene() *CategoryScoreDetails {
	if m != nil {
		return m.FriendlyFireVespene
	}
	return nil
}

func (m *ScoreDetails) GetUsedMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.UsedMinerals
	}
	return nil
}

func (m *ScoreDetails) GetUsedVespene() *CategoryScoreDetails {
	if m != nil {
		return m.UsedVespene
	}
	return nil
}

func (m *ScoreDetails) GetTotalUsedMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.TotalUsedMinerals
	}
	return nil
}

func (m *ScoreDetails) GetTotalUsedVespene() *CategoryScoreDetails {
	if m != nil {
		return m.TotalUsedVespene
	}
	return nil
}

func (m *ScoreDetails) GetTotalDamageDealt() *VitalScoreDetails {
	if m != nil {
		return m.TotalDamageDealt
	}
	return nil
}

func (m *ScoreDetails) GetTotalDamageTaken() *VitalScoreDetails {
	if m != nil {
		return m.TotalDamageTaken
	}
	return nil
}

func (m *ScoreDetails) GetTotalHealed() *VitalScoreDetails {
	if m != nil {
		return m.TotalHealed
	}
	return nil
}

func init() {
	proto.RegisterType((*Score)(nil), "SC2APIProtocol.Score")
	proto.RegisterType((*CategoryScoreDetails)(nil), "SC2APIProtocol.CategoryScoreDetails")
	proto.RegisterType((*VitalScoreDetails)(nil), "SC2APIProtocol.VitalScoreDetails")
	proto.RegisterType((*ScoreDetails)(nil), "SC2APIProtocol.ScoreDetails")
	proto.RegisterEnum("SC2APIProtocol.Score_ScoreType", Score_ScoreType_name, Score_ScoreType_value)
}

func init() { proto.RegisterFile("s2clientprotocol/score.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x5b, 0x6f, 0xe2, 0x46,
	0x14, 0xc7, 0x65, 0x1a, 0x92, 0x70, 0x00, 0x07, 0x26, 0x84, 0xb8, 0x55, 0xd4, 0x52, 0x7a, 0x11,
	0xea, 0x25, 0xad, 0x50, 0x15, 0xf5, 0xa9, 0x52, 0x14, 0xd4, 0x36, 0x5a, 0x45, 0x1b, 0x39, 0x24,
	0x7b, 0x79, 0xb1, 0x2c, 0xfb, 0x40, 0x46, 0x19, 0x3c, 0x68, 0x3c, 0xce, 0x8a, 0xaf, 0xb1, 0xaf,
	0xfb, 0x85, 0xf6, 0x63, 0xad, 0x66, 0xc6, 0x63, 0x0c, 0xc9, 0xc3, 0x7a, 0x5f, 0xd0, 0x9c, 0xcb,
	0xff, 0x77, 0xfe, 0x1c, 0xc6, 0x18, 0x4e, 0xd2, 0x71, 0xc4, 0x28, 0x26, 0x72, 0x29, 0xb8, 0xe4,
	0x11, 0x67, 0x7f, 0xa4, 0x11, 0x17, 0x78, 0xaa, 0x43, 0xe2, 0xde, 0x5c, 0x8c, 0xcf, 0xaf, 0x2f,
	0xaf, 0xf3, 0xda, 0xf0, 0xa3, 0x03, 0xf5, 0x1b, 0x55, 0x27, 0xff, 0x00, 0xe8, 0xc6, 0x40, 0xae,
	0x96, 0xe8, 0xed, 0x0e, 0x9c, 0x91, 0x3b, 0xfe, 0xee, 0x74, 0xb3, 0xfd, 0x54, 0xb7, 0x9a, 0xcf,
	0xe9, 0x6a, 0x89, 0x7e, 0x23, 0xb5, 0x47, 0xd2, 0x83, 0xba, 0x0e, 0xbc, 0xbd, 0x81, 0x33, 0xaa,
	0xfb, 0x26, 0x20, 0xe7, 0xd0, 0x36, 0xd4, 0x18, 0x65, 0x48, 0x59, 0xea, 0xed, 0x0f, 0x9c, 0x51,
	0x73, 0x7c, 0xf2, 0x2c, 0x78, 0x62, 0x7a, 0xfc, 0x56, 0x5a, 0x8a, 0x86, 0x3f, 0x43, 0xa3, 0x18,
	0x48, 0x5c, 0x80, 0x8b, 0x4c, 0x08, 0x1a, 0x65, 0x2c, 0x5b, 0x74, 0x1c, 0xd2, 0x80, 0xfa, 0x15,
	0x32, 0xc4, 0x4e, 0x6d, 0xf8, 0xde, 0x81, 0xde, 0x45, 0x28, 0x71, 0xce, 0xc5, 0xaa, 0x8c, 0x23,
	0x04, 0x76, 0x12, 0x9e, 0xa0, 0xe7, 0x0c, 0x9c, 0x51, 0xcd, 0xd7, 0x67, 0x95, 0x0b, 0xc5, 0x62,
	0xe5, 0xd5, 0x4c, 0x4e, 0x9d, 0x89, 0x07, 0x7b, 0x18, 0xf1, 0x84, 0x2f, 0x56, 0xde, 0x57, 0x3a,
	0x6d, 0x43, 0xf2, 0x2d, 0x80, 0xc4, 0xe8, 0x3e, 0xe1, 0x8c, 0xcf, 0x57, 0xde, 0x8e, 0x2e, 0x96,
	0x32, 0x4a, 0x99, 0x2d, 0xe7, 0x22, 0x8c, 0xd1, 0xab, 0x1b, 0x65, 0x1e, 0x0e, 0xdf, 0x40, 0xf7,
	0x8e, 0xca, 0x90, 0x6d, 0x1b, 0x62, 0x74, 0x56, 0x18, 0x52, 0x67, 0x85, 0x48, 0xef, 0x29, 0xb2,
	0x38, 0xcd, 0x3d, 0xd9, 0x90, 0xf4, 0x61, 0x17, 0x13, 0x14, 0x73, 0xeb, 0x2a, 0x8f, 0x86, 0x1f,
	0xda, 0xd0, 0xda, 0xc0, 0xfe, 0x09, 0x3d, 0x1a, 0x33, 0x0c, 0x96, 0x82, 0xc7, 0x59, 0x24, 0x29,
	0x4f, 0x02, 0x49, 0x17, 0x76, 0x0c, 0x51, 0xb5, 0xeb, 0xa2, 0x34, 0xa5, 0x0b, 0x24, 0x23, 0xe8,
	0x68, 0xc5, 0x3b, 0x2e, 0x1e, 0x50, 0x98, 0x6e, 0x33, 0xdd, 0x55, 0xf9, 0x57, 0x3a, 0xad, 0x3b,
	0x7f, 0x81, 0xae, 0xe4, 0x32, 0x64, 0xc1, 0x63, 0xc8, 0x32, 0x0c, 0xb2, 0x84, 0xca, 0x34, 0xf7,
	0x73, 0xa0, 0x0b, 0x77, 0x2a, 0x7f, 0xab, 0xd2, 0xe4, 0x2f, 0xe8, 0x97, 0x7b, 0x53, 0x29, 0xb2,
	0x48, 0x66, 0x02, 0xd3, 0x7c, 0x73, 0xbd, 0xb5, 0xe0, 0xa6, 0xa8, 0x91, 0xdf, 0x80, 0x3c, 0x50,
	0xc6, 0x30, 0xde, 0x18, 0x61, 0xd6, 0xd9, 0x31, 0x95, 0xd2, 0x8c, 0x33, 0x38, 0xde, 0xe8, 0x2e,
	0x0d, 0xd9, 0xd5, 0x92, 0xa3, 0x92, 0xa4, 0x34, 0xe5, 0x77, 0x20, 0x11, 0x67, 0x0c, 0x23, 0x89,
	0x71, 0xb0, 0xa0, 0x09, 0x8a, 0x90, 0xa5, 0xfa, 0xca, 0xd6, 0xfc, 0x6e, 0x51, 0xb9, 0xca, 0x0b,
	0xe4, 0x57, 0x58, 0x27, 0x83, 0x47, 0x4c, 0x97, 0x98, 0xa0, 0xbe, 0xc2, 0x35, 0xbf, 0x53, 0x14,
	0xee, 0x4c, 0x9e, 0xfc, 0x0d, 0x5e, 0x9e, 0x53, 0xab, 0x17, 0xa1, 0xc4, 0xf5, 0x84, 0x86, 0xd6,
	0xf4, 0xd7, 0x75, 0x3f, 0x94, 0x58, 0x8c, 0x39, 0x83, 0xe3, 0x6d, 0xa5, 0x1d, 0x06, 0xe6, 0xdb,
	0x6c, 0x0a, 0xed, 0xc4, 0x9f, 0xc0, 0x55, 0x07, 0xb9, 0x9e, 0xd3, 0xd4, 0xed, 0x6d, 0x9d, 0x2d,
	0xf0, 0x3f, 0x80, 0x49, 0x14, 0xd0, 0x96, 0xee, 0x6a, 0xe9, 0xa4, 0x65, 0x9d, 0x43, 0x63, 0xc6,
	0x79, 0x1c, 0x64, 0x29, 0xc6, 0x5e, 0x5b, 0x3f, 0xa5, 0x3f, 0x6e, 0x3f, 0xa5, 0xcf, 0x3d, 0x5e,
	0xfe, 0xbe, 0x92, 0xdd, 0xa6, 0x18, 0x93, 0x2b, 0x38, 0xc8, 0x7f, 0x94, 0xc2, 0x8f, 0x5b, 0x01,
	0xe4, 0x1a, 0x71, 0x61, 0xfb, 0x05, 0xb8, 0xf6, 0x37, 0xce, 0x7d, 0x1f, 0x54, 0xa0, 0xb5, 0xf3,
	0x0b, 0x90, 0x7f, 0xbd, 0x4b, 0x68, 0x33, 0x9e, 0x96, 0x36, 0xd5, 0xa9, 0xc0, 0x6a, 0x29, 0x69,
	0xe1, 0xeb, 0x3f, 0xd0, 0x71, 0xe1, 0xaa, 0x5b, 0x81, 0xd4, 0x54, 0x4a, 0xeb, 0xe9, 0x2d, 0xf4,
	0x67, 0x82, 0x62, 0x12, 0xb3, 0x55, 0x30, 0xa3, 0xa2, 0x74, 0x5d, 0x48, 0x05, 0x64, 0xcf, 0x32,
	0xfe, 0xa5, 0x62, 0x7d, 0xa5, 0x5e, 0xc3, 0xd1, 0x26, 0xdb, 0xba, 0x3d, 0xac, 0x80, 0x3e, 0x2c,
	0xa3, 0x4b, 0x9b, 0x54, 0x77, 0x64, 0x6d, 0xb6, 0x57, 0x65, 0x93, 0x4a, 0x5a, 0xde, 0xa4, 0x46,
	0x59, 0x6f, 0x47, 0x55, 0x36, 0xa9, 0x94, 0xd6, 0xd3, 0x14, 0x0e, 0xcd, 0x5f, 0xce, 0xa6, 0xb3,
	0x7e, 0x05, 0x9e, 0xf9, 0x7f, 0xbb, 0x2d, 0xdb, 0xf3, 0x81, 0x94, 0xa8, 0xd6, 0xe4, 0x71, 0x05,
	0x68, 0xa7, 0x80, 0x5a, 0xa7, 0x2f, 0x2d, 0x33, 0x0e, 0x17, 0xe1, 0x5c, 0xbd, 0x17, 0x43, 0x26,
	0x3d, 0x4f, 0x33, 0xbf, 0xdf, 0x66, 0x3e, 0x79, 0x75, 0xe4, 0xc0, 0x89, 0xd6, 0x4e, 0x94, 0xf4,
	0x09, 0x50, 0x86, 0x0f, 0x98, 0x78, 0x5f, 0x7f, 0x09, 0x70, 0xaa, 0xa4, 0x64, 0x02, 0x2d, 0x03,
	0xbc, 0xc7, 0x90, 0x61, 0xec, 0x7d, 0xf3, 0xb9, 0xa8, 0xa6, 0x96, 0xfd, 0xaf, 0x55, 0x9f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xfd, 0xb5, 0x39, 0xbe, 0x87, 0x08, 0x00, 0x00,
}
