package emulator

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name:  "server",
	Usage: "start object storage server",
	// Flags:  append(ServerFlags, GlobalFlags...),
	// 执行此命令参数是执行 serverMain
	Action: serverMain,
	Flags:  ServerFlags,
	CustomHelpTemplate: `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} {{if .VisibleFlags}}[FLAGS] {{end}}DIR1 [DIR2..]
  {{.HelpName}} {{if .VisibleFlags}}[FLAGS] {{end}}DIR{1...64}
  {{.HelpName}} {{if .VisibleFlags}}[FLAGS] {{end}}DIR{1...64} DIR{65...128}

DIR:
  DIR points to a directory on a filesystem. When you want to combine
  multiple drives into a single large system, pass one directory per
  filesystem separated by space. You may also use a '...' convention
  to abbreviate the directory arguments. Remote directories in a
  distributed setup are encoded as HTTP(s) URIs.
{{if .VisibleFlags}}
FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}
EXAMPLES:
  1. Start MinIO server on "/home/shared" directory.
     {{.Prompt}} {{.HelpName}} /home/shared

  2. Start single node server with 64 local drives "/mnt/data1" to "/mnt/data64".
     {{.Prompt}} {{.HelpName}} /mnt/data{1...64}

  3. Start distributed MinIO server on an 32 node setup with 32 drives each, run following command on all the nodes
     {{.Prompt}} {{.HelpName}} http://node{1...32}.example.com/mnt/export{1...32}

  4. Start distributed MinIO server in an expanded setup, run the following command on all the nodes
     {{.Prompt}} {{.HelpName}} http://node{1...16}.example.com/mnt/export{1...32} \
            http://node{17...64}.example.com/mnt/export{1...64}

  5. Start distributed MinIO server, with FTP and SFTP servers on all interfaces via port 8021, 8022 respectively
     {{.Prompt}} {{.HelpName}} http://node{1...4}.example.com/mnt/export{1...4} \
           --ftp="address=:8021" --ftp="passive-port-range=30000-40000" \
           --sftp="address=:8022" --sftp="ssh-private-key=${HOME}/.ssh/id_rsa"
`,
}

func serverMain(ctx *cli.Context) error {
	fmt.Println("start server main")

	err := buildServerCtxt(ctx, &globalServerCtxt)

	if err != nil {
		fmt.Errorf("buildServerCtxt error: %s\n", err)
	}

	handler, err := configureServerHandler(globalServerCtxt)
	if err != nil {
		fmt.Errorf("configureServerHandler error: %s\n", err)
	}

	_, err = newObjectLayer(GlobalContext, globalServerCtxt)
	if err != nil {
		fmt.Errorf("configureServerHandler error: %s\n", err)
	}
	return listenAndServe(globalServerCtxt, handler)
}

func newObjectLayer(ctx context.Context, server serverCtxt) (newObject ObjectLayer, err error) {
	object := &Objects{}
	setObjectLayer(object)
    disk, err := newStorageAPI(server.Disk)
    if err != nil {
        return
    }
    object.disk = disk
    globalLocalDrive = disk
	return
}
