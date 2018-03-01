package main

import(
    "log"
    "log/syslog"
    "fmt"
    "github.com/hpcloud/tail"
    flag "github.com/ogier/pflag"
)

// Flags
var (
    syslogProtocol string
    syslogServer string
    syslogPort string
    sigsciAgentLog string
)


func main() {
    flag.Parse()

    // Create the syslog string
    syslogUrl := syslogServer + ":" + syslogPort
    // Configure logger to write to the syslog. You could do this in init(), too.
    logwriter, e := syslog.Dial(syslogProtocol,syslogUrl, syslog.LOG_NOTICE, "sigsci-agent")
    if e == nil {
        log.SetOutput(logwriter)
    } else {
        log.Fatal("Failed to dial syslog")
    }

    // Now from anywhere else in your program, you can use this:
    log.Print("Starting sigsci-agent monitor script")

    t, err := tail.TailFile(sigsciAgentLog, tail.Config{
    Follow: true,
    ReOpen: true})
    for line := range t.Lines {
        fmt.Println(line.Text)
        log.Print(line.Text)
    }

    if err != nil {
            log.Print(err)
    }
}


func init() {
 flag.StringVarP(&syslogProtocol, "protocol", "p", "udp", "Syslog protocol (tcp/udp)")
 flag.StringVarP(&syslogServer, "server", "s", "localhost", "Syslog server fqdn (eg localhost, 127.0.001)")
 flag.StringVarP(&syslogPort, "port", "P", "514", "Syslog port to connect to, default 514")
 flag.StringVarP(&sigsciAgentLog, "sigsci-log", "l", "/var/log/sigsci.log", "Path to the SigSci Agent log")
}
