
     ,-----.,--.                  ,--. ,---.   ,--.,------.  ,------.
    '  .--./|  | ,---. ,--.,--. ,-|  || o   \  |  ||  .-.  \ |  .---'
    |  |    |  || .-. ||  ||  |' .-. |`..'  |  |  ||  |  \  :|  `--, 
    '  '--'\|  |' '-' ''  ''  '\ `-' | .'  /   |  ||  '--'  /|  `---.
     `-----'`--' `---'  `----'  `---'  `--'    `--'`-------' `------'
    ----------------------------------------------------------------- 

https://ide.c9.io/vopijibib/gomon

### Load monitoring web application 
Create a simple web application that monitors load average on your machine:

Collect the machine load (using “uptime” for example)

Display in the application the key statistics as well as a history of load over the past 10 minutes in 10s intervals. We’d suggest a graphical representation using D3.js, but feel free to use another tool or representation if you prefer. Make it easy for the end-user to picture the situation!

Make sure a user can keep the web page open and monitor the load on their machine.

Whenever the load for the past 2 minutes exceeds 1 on average, add a message saying that “High load generated an alert - load = {value}, triggered at {time}”

Whenever the load average drops again below 1 on average for the past 2 minutes, Add another message explaining when the alert recovered.

Make sure all messages showing when alerting thresholds are crossed remain visible on the page for historical reasons.

Write a test for the alerting logic

Explain how you’d improve on this application design


### <del>HTTP log monitoring console program
<del>Create a simple console program that monitors HTTP traffic on your machine:

<del>Consume an actively written-to w3c-formatted HTTP access log

<del>Every 10s, display in the console the sections of the web site with the most hits (a section is defined as being what's before the second '/' in a URL. i.e. the section for "http://my.site.com/pages/create' is "http://my.site.com/pages"), as well as interesting summary statistics on the traffic as a whole.

<del>Make sure a user can keep the console app running and monitor traffic on their machine

<del>Whenever total traffic for the past 2 minutes exceeds a certain number on average, add a message saying that “High traffic generated an alert - hits = {value}, triggered at {time}”

<del>Whenever the total traffic drops again below that value on average for the past 2 minutes, add another message detailing when the alert recovered

<del>Make sure all messages showing when alerting thresholds are crossed remain visible on the page for historical reasons.

<del>Write a test for the alerting logic

<del>Explain how you’d improve on this application design
