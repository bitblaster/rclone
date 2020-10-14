// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	//_ "github.com/rclone/rclone/backend/all" // import all backends
	_ "github.com/rclone/rclone/backend/amazonclouddrive"
	_ "github.com/rclone/rclone/backend/googlecloudstorage"
	_ "github.com/rclone/rclone/backend/local"
	"github.com/rclone/rclone/cmd"
	//_ "github.com/rclone/rclone/cmd/all"    // import all commands
	_ "github.com/rclone/rclone/cmd/cat"
	_ "github.com/rclone/rclone/cmd/copy"
	_ "github.com/rclone/rclone/cmd/copyto"
	_ "github.com/rclone/rclone/cmd/copyurl"
	_ "github.com/rclone/rclone/cmd/delete"
	_ "github.com/rclone/rclone/cmd/deletefile"
	_ "github.com/rclone/rclone/cmd/info"
	_ "github.com/rclone/rclone/cmd/ls"
	_ "github.com/rclone/rclone/cmd/lsl"
	_ "github.com/rclone/rclone/cmd/mkdir"
	_ "github.com/rclone/rclone/cmd/move"
	_ "github.com/rclone/rclone/cmd/moveto"
	_ "github.com/rclone/rclone/cmd/rmdir"
	_ "github.com/rclone/rclone/cmd/rmdirs"
	_ "github.com/rclone/rclone/cmd/sync"
	_ "github.com/rclone/rclone/cmd/version"

	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
