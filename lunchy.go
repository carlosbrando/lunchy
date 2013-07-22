package main

import (
  "flag"
  "fmt"
)

func showBanner() {
  fmt.Println(`Lunchy 1.0, the friendly launchctl wrapper
Usage: lunchy [start|stop|restart|ls|list|status|install|show|edit] [options]
    -F, --force                      Force start (disabled) agents
    -v, --verbose                    Show command executions
    -w, --write                      Persist command
    -l, --long                       Display absolute paths when listing agents

Supported commands:

 ls [-l] [pattern]       Show the list of installed agents, with optional [pattern] filter
 list [-l] [pattern]     Alias for 'ls'
 start [-wF] [pattern]   Start the first agent matching [pattern]
 stop [-w] [pattern]     Stop the first agent matching [pattern]
 restart [pattern]       Stop and start the first agent matching [pattern]
 status [pattern]        Show the PID and label for all agents, with optional [pattern] filter
 install [file]          Install [file] to ~/Library/LaunchAgents or /Library/LaunchAgents (whichever it finds first)
 show [pattern]          Show the contents of the launchctl daemon file
 edit [pattern]          Open the launchctl daemon file in the default editor (EDITOR environment variable)

-w will persist the start/stop command so the agent will load on startup or never load, respectively.
-l will display absolute paths of the launchctl daemon files when showing list of installed agents.

Example:
 lunchy ls
 lunchy ls -l nginx
 lunchy start -w redis
 lunchy stop mongo
 lunchy status mysql
 lunchy install /usr/local/Cellar/redis/2.2.2/io.redis.redis-server.plist
 lunchy show redis
 lunchy edit mongo

Note: if you run lunchy as root, you can manage daemons in /Library/LaunchDaemons also.`)
}

func checkOperation(command string, pattern string) {
  switch command {
  case "ls":
    ls(pattern)
  }
}

func ls(pattern string) {
  fmt.Println(pattern)
}

func main() {
  flag.Parse()

  args := flag.Args()
  if len(args) < 2 {
    showBanner()
    return
  }

  command := args[0]
  pattern := args[1]

  checkOperation(command, pattern)

  // fmt.Println(command)
  // fmt.Println(pattern)
}
