package roster

import "../common"

type Roster interface {
	ListContacts(token string) (*[]common.Profile, error)
}
