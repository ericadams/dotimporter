package inner

type inner struct {
	member string
}
type Inner struct {
	PublicMember  string
	privateMember string
}

func newPublicMember() Inner {
	return Inner{
		PublicMember:  "public",
		privateMember: "private",
	}
}
