package main

import "github.com/marcusolsson/tui-go"

// The sign at the register view
var logo = `
     ________  _____________________  ______________
    /__  _/ / / / ___/_  __/ ____/ /_/ / __  /_  __/
      / // / / /\__ \ / / / /   / __  / /_/ / / /
  ___/ // /_/ /___/ // / / /___/ / / / / / / / /
 /____//_____/_____//_/ /_____/_/ /_/_/ /_/ /_/
`

var ldb *LocalDatabase

// The globUsername when user logs in
var globUsername string

// The user id
var globUserid int

/* Didn't name it globHistory because it
only affects the chats page*/
var history *tui.Box
