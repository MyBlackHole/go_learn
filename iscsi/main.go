package main

import (
    "os"
    "fmt"
    "strconv"
    "github.com/u-root/iscsinl"
)

func main() {
   ipc, err := iscsinl.ConnectNetlink()

   if err != nil {
       fmt.Printf("Error connecting to netlink: %s\n", err)
       return
   }

   sid, _ := strconv.Atoi(os.Args[1])

   err = ipc.StopConnection(uint32(sid), 0)

   if err != nil {
       fmt.Printf("Error stopping connection: %s\n", err)
   }

   err = ipc.DestroyConnection(uint32(sid), 0)

   if err != nil {
       fmt.Printf("Error destroying connection: %s\n", err)
   }

   err = ipc.DestroySession(uint32(sid))

   if err != nil {
       fmt.Printf("Error destroying session: %s\n", err)
   }
}
