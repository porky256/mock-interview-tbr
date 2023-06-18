package dal

type UserDatabaseProvider interface {
}

type MatchDatabaseProvider interface {
}

type SkillDatabaseProvider interface {
}

type GlobalDatabaseProvider interface {
	UserDatabaseProvider
	MatchDatabaseProvider
	SkillDatabaseProvider
}
