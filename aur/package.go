package aur

type Package struct {
	//URL to remote Git repository.
	//Basically it is a https://aur.archlinux.org/<Package.Name>.git
	Remote string
	//LocalPath is a path to local git repository.
	LocalPath      string
	Description    string      `json:"Description"`
	FirstSubmitted int         `json:"FirstSubmitted"`
	ID             int         `json:"ID"`
	LastModified   int         `json:"LastModified"`
	Maintainer     string      `json:"Maintainer"`
	Name           string      `json:"Name"`
	NumVotes       int         `json:"NumVotes"`
	OutOfDate      interface{} `json:"OutOfDate"`
	PackageBase    string      `json:"PackageBase"`
	PackageBaseID  int         `json:"PackageBaseID"`
	Popularity     float64     `json:"Popularity"`
	URL            string      `json:"URL"`
	URLPath        string      `json:"URLPath"`
	Version        string      `json:"Version"`

	//Info specific
	Depends     []string `json:"Depends,omitempty"`
	Keywords    []string `json:"Keywords,omitempty"`
	License     []string `json:"License,omitempty"`
	MakeDepends []string `json:"MakeDepends,omitempty"`
	OptDepends  []string `json:"OptDepends,omitempty"`
	Submitter   string   `json:"Submitter,omitempty"`
}

//Depends        []string `json:"Depends,omitempty"`
//Description    string   `json:"Description,omitempty"`
//FirstSubmitted int      `json:"FirstSubmitted,omitempty"`
//ID             int      `json:"ID,omitempty"`
//Keywords       []string `json:"Keywords,omitempty"`
//LastModified   int      `json:"LastModified,omitempty"`
//License        []string `json:"License,omitempty"`
//Maintainer     string   `json:"Maintainer,omitempty"`
//MakeDepends    []string `json:"MakeDepends,omitempty"`
//Name           string   `json:"Name,omitempty"`
//NumVotes       int      `json:"NumVotes,omitempty"`
//OptDepends     []string `json:"OptDepends,omitempty"`
//OutOfDate      any      `json:"OutOfDate,omitempty"`
//PackageBase    string   `json:"PackageBase,omitempty"`
//PackageBaseID  int      `json:"PackageBaseID,omitempty"`
//Popularity     float64  `json:"Popularity,omitempty"`
//Submitter      string   `json:"Submitter,omitempty"`
//URL            string   `json:"URL,omitempty"`
//URLPath        string   `json:"URLPath,omitempty"`
//Version        string   `json:"Version,omitempty"`
