// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"easygoadmin/app/dao/internal"
)

// MemberLevelDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type MemberLevelDao struct {
	internal.MemberLevelDao
}

var (
	// MemberLevel is globally public accessible object for table ums_member_level operations.
	MemberLevel = MemberLevelDao{
		internal.MemberLevel,
	}
)

// Fill with you ideas below.