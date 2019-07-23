package shell

import (
	"io"

	// for types
	sh "github.com/ipfs/go-ipfs-api"
)

type Shell interface {
	Add(r io.Reader) (string, error)
	AddLink(target string) (string, error)

	Cat(ipfspath string) (io.ReadCloser, error)
	Get(ipfspath, outdir string) error
	ResolvePath(ipath string) (string, error)
	List(ipath string) ([]*sh.LsLink, error)

	NewObject(template string) (string, error)
	PatchLink(root, npath, childhash string, create bool) (string, error)
	Patch(root, action string, args ...string) (string, error)
}

type shellShim struct {
	*sh.Shell
}

//NOTE: upstream signature became (r, opts...)
func (ss *shellShim) Add(r io.Reader) (string, error) {
	return ss.Shell.Add(r)
}
