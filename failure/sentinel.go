/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package failure

// Sentinel [constant-errors]
//
// [constant-errors]: https://dave.cheney.net/2016/04/07/constant-errors
type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	NoError        = Sentinel("no error")
	NotFound       = Sentinel("not found")
	NotImplemented = Sentinel("not implemented")
)

/*
####### END ############################################################################################################
*/
